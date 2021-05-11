package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func doResumeSDK() error {
	// 线上需要换成内网 ip : 172.23.201.53
	ResumeSDKURL := "http://101.37.18.181:2015/api/ResumeParser"

	b, _ := ioutil.ReadFile("./aa.pdf")
	encodeString := base64.StdEncoding.EncodeToString(b)
	resume := make(map[string]interface{})
	resume["base_cont"] = encodeString
	resume["fname"] = "aaa.pdf"
	resume["uid"] = "1811191"
	resume["pwd"] = "538610"

	bytesData, err := json.Marshal(resume)
	req, err := http.NewRequest("POST", ResumeSDKURL, bytes.NewBuffer(bytesData))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", `application/json`)
	req.Header.Set("Authentication", `Basic username="admin",password="2015"`)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	fmt.Println(string(body))

	return nil
}

func main() {
	err := doResumeSDK()
	fmt.Println(err)
}
