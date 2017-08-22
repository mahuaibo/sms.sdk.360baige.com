package send

import (
	. "sms.sdk.360baige.com/client"
	"fmt"
)

const (
	conf_accessKey_id     = "LTAIMKJ3UerrTf0q"
	conf_accessKey_secret = "xumzZFxwNIOwj9h7VybBRBhlqJ2BYl"
)

func MessageCode(companyName, code, recNum string) error {
	client := NewClient(&Credentials{conf_accessKey_id, conf_accessKey_secret})
	client.SetDebug(false)
	paramString := ` { "companyName" : "` + companyName + `" , "code" : "` + code + `" } `;
	fmt.Println(recNum,paramString)
	request := NewSMSRequest("SingleSendSms", "百鸽互联", "SMS_35020179", recNum, paramString)
	response := &ErrorResponse{}
	//3
	err := client.Query(request, response)
	return err
}
