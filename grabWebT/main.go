package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// 添加账号
func AddEmail() {
	emailData := make(map[string]string)
	emailData["password"] = "password"
	emailData["user_name"] = "user_nameeeee"
	// emailData["server"] = "imap.exmail.qq.com:993"
	emailData["last_time"] = "2020-01-02T15:04:05Z"
	// emailData["is_deleted"] = "1"

	b, _ := json.Marshal(emailData)
	req, _ := http.NewRequest("POST", "http://127.0.0.1:63334/emails/Add", bytes.NewBuffer(b))
	client := &http.Client{}
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

// 修改账号
func saveEmail() {
	emailData := make(map[string]string)
	emailData["id"] = "46"
	emailData["password"] = "passwordddd"
	emailData["user_name"] = "user_nameeeee1"
	emailData["server"] = "imap.exmail.qq.com:993"
	emailData["last_time"] = "2020-01-02T15:04:05Z"
	emailData["is_deleted"] = "0"

	b, _ := json.Marshal(emailData)
	req, _ := http.NewRequest("POST", "http://127.0.0.1:63334/emails/Save", bytes.NewBuffer(b))
	client := &http.Client{}
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func main() {
	saveEmail()
	// AddEmail()
}
