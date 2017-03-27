package client

import (
	"testing"
	"fmt"
)

const (
	conf_accessKey_id = "" //"LTAIMKJ3UerrTf0q"
	conf_accessKey_secret = "" //"xumzZFxwNIOwj9h7VybBRBhlqJ2BYl"
)

func TestClient(t *testing.T) {
	//1
	client := NewClient(&Credentials{conf_accessKey_id, conf_accessKey_secret})
	client.SetDebug(true)
	//2
	paramString := ` { "companyName" : "百鸽小学" , "code" : "666666" } `;
	request := NewSMSRequest("SingleSendSms", "百鸽互联", "SMS_35020179", "18911545460", paramString)
	response := &ErrorResponse{}
	//3
	err := client.Query(request,response)
	if err == nil {
		fmt.Println("发送成功")
	}else{
		fmt.Println(err)
	}
}
