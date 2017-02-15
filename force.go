package main

import (
	"encoding/base64"
)

type ForceClient struct {
	portType    *MetadataPortType
	loginResult *LoginResult
}

func NewForceClient(endpoint string, apiversion string) *ForceClient {
	portType := NewMetadataPortType("https://"+endpoint+"/services/Soap/u/"+apiversion, true, nil)
	return &ForceClient{portType: portType}
}

func (client *ForceClient) Login(username string, password string) error {
	loginRequest := LoginRequest{Username: username, Password: password}
	loginResponse, err := client.portType.Login(&loginRequest)
	if err != nil {
		return err
	}
	client.loginResult = &loginResponse.LoginResult
	return nil
}

func (client *ForceClient) Deploy(buf []byte) (*DeployResponse, error) {
	request := Deploy{
		ZipFile:       base64.StdEncoding.EncodeToString(buf),
		DeployOptions: nil,
	}
	sessionHeader := SessionHeader{
		SessionId: client.loginResult.SessionId,
	}
	client.portType.SetHeader(&sessionHeader)
	client.portType.SetServerUrl(client.loginResult.MetadataServerUrl)

	return client.portType.Deploy(&request)
}

func (client *ForceClient) CheckDeployStatus(resultId *ID) (*CheckDeployStatusResponse, error) {
	check_request := CheckDeployStatus{AsyncProcessId: resultId, IncludeDetails: true}
	return client.portType.CheckDeployStatus(&check_request)
}

func (client *ForceClient) Retrieve() (*RetrieveResponse, error) {
	request := Retrieve{
		RetrieveRequest: &RetrieveRequest{
			ApiVersion: 38,
			Unpackaged: &Package{
				Types: []*PackageTypeMembers{
					&PackageTypeMembers{
						Name:    "ApexClass",
						Members: []string{"HelloSpm_Dep"},
					},
				},
			},
		},
	}
	sessionHeader := SessionHeader{
		SessionId: client.loginResult.SessionId,
	}
	client.portType.SetHeader(&sessionHeader)
	client.portType.SetServerUrl(client.loginResult.MetadataServerUrl)
	return client.portType.Retrieve(&request)
}

func (client *ForceClient) CheckRetrieveStatus(id *ID) (*CheckRetrieveStatusResponse, error) {
	request := CheckRetrieveStatus{
		AsyncProcessId: id,
	}
	sessionHeader := SessionHeader{
		SessionId: client.loginResult.SessionId,
	}
	client.portType.SetHeader(&sessionHeader)
	client.portType.SetServerUrl(client.loginResult.MetadataServerUrl)
	return client.portType.CheckRetrieveStatus(&request)
}
