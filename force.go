package main

import (
  "log"
  "encoding/base64"
  "time"
)

type ForceClient struct {
  portType *MetadataPortType
  loginResult *LoginResult
}

func NewForceClient (endpoint string, apiversion string) (*ForceClient) {
  portType := NewMetadataPortType("https://" + endpoint + "/services/Soap/u/" + apiversion, true, nil)
  return &ForceClient{portType: portType}
}

func (client *ForceClient) Login(username string, password string) (error) {
  loginRequest := LoginRequest{Username: username, Password: password }
  loginResponse, err := client.portType.Login(&loginRequest)
  if err != nil {
    return err
  }
  client.loginResult = &loginResponse.LoginResult
  return nil
}

func (client *ForceClient) DeployAndCheckResult(buf []byte, pollseconds int) (error) {
  request := Deploy{
    ZipFile: base64.StdEncoding.EncodeToString(buf),
    DeployOptions: nil,
  }
  sessionHeader := SessionHeader{
    SessionId: client.loginResult.SessionId,
  }
  client.portType.SetHeader(&sessionHeader)
  client.portType.SetServerUrl(client.loginResult.MetadataServerUrl)

  response, err := client.portType.Deploy(&request)
  if err != nil {
    panic(err)
  }
  log.Println("Deploying...")
  for {
    time.Sleep(time.Duration(pollseconds) * time.Second)
    log.Println("Check Deploy Result...")
    check_request := CheckDeployStatus{AsyncProcessId: response.Result.Id, IncludeDetails: true}
    check_response, err := client.portType.CheckDeployStatus(&check_request)
    if err != nil {
      panic(err)
    }
    if check_response.Result.Done {
      log.Println("Deploy is successful")
      return nil
    }
  }
  return nil
}

