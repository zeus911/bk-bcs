/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.,
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under,
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

package variable

import (
	"context"
	"errors"
	"fmt"

	"github.com/spf13/viper"

	pbcommon "bk-bscp/internal/protocol/common"
	pbdatamanager "bk-bscp/internal/protocol/datamanager"
	pb "bk-bscp/internal/protocol/templateserver"
	"bk-bscp/pkg/common"
	"bk-bscp/pkg/logger"
)

// CreateAction create a variable object
type CreateAction struct {
	viper      *viper.Viper
	dataMgrCli pbdatamanager.DataManagerClient

	req  *pb.CreateVariableReq
	resp *pb.CreateVariableResp

	newVid string
}

// NewCreateAction creates new CreateAction
func NewCreateAction(viper *viper.Viper, dataMgrCli pbdatamanager.DataManagerClient,
	req *pb.CreateVariableReq, resp *pb.CreateVariableResp) *CreateAction {
	action := &CreateAction{viper: viper, dataMgrCli: dataMgrCli, req: req, resp: resp}

	action.resp.Seq = req.Seq
	action.resp.ErrCode = pbcommon.ErrCode_E_OK
	action.resp.ErrMsg = "OK"

	return action
}

// Err setup error code message in response and return the error.
func (act *CreateAction) Err(errCode pbcommon.ErrCode, errMsg string) error {
	act.resp.ErrCode = errCode
	act.resp.ErrMsg = errMsg
	return errors.New(errMsg)
}

// Input handles the input messages.
func (act *CreateAction) Input() error {
	if err := act.verify(); err != nil {
		return act.Err(pbcommon.ErrCode_E_TPL_PARAMS_INVALID, err.Error())
	}
	return nil
}

// Output handles the output messages.
func (act *CreateAction) Output() error {
	// do nothing.
	return nil
}

func (act *CreateAction) verify() error {
	if err := common.VerifyID(act.req.Bid, "bid"); err != nil {
		return err
	}

	switch pbcommon.VariableType(act.req.Type) {
	case pbcommon.VariableType_VT_CLUSTER:
		if err := common.VerifyNormalName(act.req.Cluster, "cluster"); err != nil {
			return err
		}
		if err := common.VerifyClusterLabels(act.req.ClusterLabels); err != nil {
			return err
		}
	case pbcommon.VariableType_VT_ZONE:
		if err := common.VerifyNormalName(act.req.Cluster, "cluster"); err != nil {
			return err
		}
		if err := common.VerifyClusterLabels(act.req.ClusterLabels); err != nil {
			return err
		}
		if err := common.VerifyNormalName(act.req.Zone, "zone"); err != nil {
			return err
		}
	}

	if err := common.VerifyVarKey(act.req.Key); err != nil {
		return err
	}

	if err := common.VerifyVarValue(act.req.Value); err != nil {
		return err
	}

	if err := common.VerifyMemo(act.req.Memo); err != nil {
		return err
	}

	if err := common.VerifyNormalName(act.req.Creator, "creator"); err != nil {
		return err
	}

	return nil
}

func (act *CreateAction) genVariableID() error {
	id, err := common.GenVariableid()
	if err != nil {
		return err
	}
	act.newVid = id
	return nil
}

func (act *CreateAction) createVariable() (pbcommon.ErrCode, string) {
	ctx, cancel := context.WithTimeout(context.Background(), act.viper.GetDuration("datamanager.calltimeout"))
	defer cancel()

	req := &pbdatamanager.CreateVariableReq{
		Seq:           act.req.Seq,
		Bid:           act.req.Bid,
		Cluster:       act.req.Cluster,
		ClusterLabels: act.req.ClusterLabels,
		Zone:          act.req.Zone,
		Vid:           act.newVid,
		Type:          act.req.Type,
		Key:           act.req.Key,
		Value:         act.req.Value,
		Memo:          act.req.Memo,
		Creator:       act.req.Creator,
	}

	logger.V(2).Infof("CreateVariable[%d]| request to datamanager CreateVariable, %+v", req.Seq, req)

	resp, err := act.dataMgrCli.CreateVariable(ctx, req)
	if err != nil {
		return pbcommon.ErrCode_E_TPL_SYSTEM_UNKONW, fmt.Sprintf("request to datamanager CreateVariable, %+v", err)
	}
	// when already exist, return the existed vid
	act.resp.Vid = resp.Vid

	if resp.ErrCode != pbcommon.ErrCode_E_OK {
		return resp.ErrCode, resp.ErrMsg
	}

	return pbcommon.ErrCode_E_OK, "OK"
}

// Do do action
func (act *CreateAction) Do() error {

	// generate variable id
	if err := act.genVariableID(); err != nil {
		return act.Err(pbcommon.ErrCode_E_TPL_SYSTEM_UNKONW, err.Error())
	}

	// create variable
	if errCode, errMsg := act.createVariable(); errCode != pbcommon.ErrCode_E_OK {
		return act.Err(errCode, errMsg)
	}
	return nil
}
