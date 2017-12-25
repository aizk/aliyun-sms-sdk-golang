package main

import (
	"fmt"
	"github.com/aliyun-sms-go-sdk/sms"
	"os"
)

// modify it to yours
const (
	ACCESSID  = "$ACCESSID"
	ACCESSKEY = "$ACCESSKEY"
)

func main() {
	sms.HttpDebugEnable = true
	c := sms.New(ACCESSID, ACCESSKEY)
	c.SetEndPoint("https://dysmsapi.aliyuncs.com")

	// send to one person
	e, err := c.SendOne("1861234****", "多协云", "SMS_22175101", `{"code":"123456"}`)
	if err != nil {
		fmt.Println("send sms failed", err, e.Error())
		os.Exit(0)
	}

	//send to more than one person
	e, err = c.SendMulti([]string{"1375821****", "1835718****"}, "多协云", "SMS_22175101", `{"company":"duoxieyun"}`)
	if err != nil {
		fmt.Println("send sms failed", err, e.Error())
		os.Exit(0)
	}
	fmt.Println("send sms succeed", e.GetRequestId())
}
