/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package command

import (
	"encoding/json"

	"github.com/hyperledger/aries-framework-go/pkg/controller/command"
	cmdvdr "github.com/hyperledger/aries-framework-go/pkg/controller/command/vdr"

	"github.com/trustbloc/agent-sdk/cmd/agent-mobile/pkg/wrappers/models"
)

// VDR contains necessary fields for each of its operations.
type VDR struct {
	handlers map[string]command.Exec
}

// ResolveDID resolve did.
func (v *VDR) ResolveDID(request *models.RequestEnvelope) *models.ResponseEnvelope {
	args := cmdvdr.IDArg{}

	if err := json.Unmarshal(request.Payload, &args); err != nil {
		return &models.ResponseEnvelope{Error: &models.CommandError{Message: err.Error()}}
	}

	response, cmdErr := exec(v.handlers[cmdvdr.ResolveDIDCommandMethod], args)
	if cmdErr != nil {
		return &models.ResponseEnvelope{Error: cmdErr}
	}

	return &models.ResponseEnvelope{Payload: response}
}

// SaveDID saves the did doc to the store.
func (v *VDR) SaveDID(request *models.RequestEnvelope) *models.ResponseEnvelope {
	args := cmdvdr.DIDArgs{}

	if err := json.Unmarshal(request.Payload, &args); err != nil {
		return &models.ResponseEnvelope{Error: &models.CommandError{Message: err.Error()}}
	}

	response, cmdErr := exec(v.handlers[cmdvdr.SaveDIDCommandMethod], args)
	if cmdErr != nil {
		return &models.ResponseEnvelope{Error: cmdErr}
	}

	return &models.ResponseEnvelope{Payload: response}
}

// GetDID retrieves the did from the store.
func (v *VDR) GetDID(request *models.RequestEnvelope) *models.ResponseEnvelope {
	args := cmdvdr.IDArg{}

	if err := json.Unmarshal(request.Payload, &args); err != nil {
		return &models.ResponseEnvelope{Error: &models.CommandError{Message: err.Error()}}
	}

	response, cmdErr := exec(v.handlers[cmdvdr.GetDIDCommandMethod], args)
	if cmdErr != nil {
		return &models.ResponseEnvelope{Error: cmdErr}
	}

	return &models.ResponseEnvelope{Payload: response}
}

// GetDIDRecords retrieves the did doc containing name and didID.
func (v *VDR) GetDIDRecords(request *models.RequestEnvelope) *models.ResponseEnvelope {
	response, cmdErr := exec(v.handlers[cmdvdr.GetDIDsCommandMethod], request.Payload)
	if cmdErr != nil {
		return &models.ResponseEnvelope{Error: cmdErr}
	}

	return &models.ResponseEnvelope{Payload: response}
}
