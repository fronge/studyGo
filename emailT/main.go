package main

import (
	"fmt"
	net_mail "net/mail"
	"studyGo/emailT/tools"
	"time"

	"github.com/emersion/go-imap"
	id "github.com/emersion/go-imap-id"
	"github.com/emersion/go-imap/client"
	"github.com/emersion/go-message"
)

// 登录函数
func loginEmail(Eserver, UserName, Password string) (*client.Client, error) {
	c, err := client.DialTLS(Eserver, nil)
	if err != nil {
		return nil, err
	}
	//登陆
	if err = c.Login(UserName, Password); err != nil {
		return nil, err
	}
	return c, nil
}

// 邮件接收

func emailList(Eserver, UserName, Password string) (err error) {
	c, err := loginEmail(Eserver, UserName, Password)
	if err != nil {
		fmt.Println(err)
		return
	}
	idClient := id.NewClient(c)
	idClient.ID(
		id.ID{
			id.FieldName:    "IMAPClient",
			id.FieldVersion: "3.1.0",
		},
	)

	defer c.Close()
	// 选择收件箱
	mbox, err := c.Select("INBOX", false)
	if err != nil {

		fmt.Println("select inbox err: ", err)
		return
	}
	if mbox.Messages == 0 {
		return
	}
	// 选择收取邮件的时间段
	criteria := imap.NewSearchCriteria()
	// 收取7天之内的邮件
	t1, err := time.Parse("2006-01-02 15:04:05", "2020-03-02 15:04:05")
	criteria.Since = t1
	// 按条件查询邮件
	ids, err := c.Search(criteria)
	if err != nil {
		fmt.Println(err)
	}
	if len(ids) == 0 {
		return
	}
	seqset := new(imap.SeqSet)
	seqset.AddNum(ids...)
	sect := &imap.BodySectionName{}
	messages := make(chan *imap.Message, 100)
	done := make(chan error, 1)
	go func() {
		done <- c.Fetch(seqset, []imap.FetchItem{sect.FetchItem()}, messages)
	}()
	for msg := range messages {
		r := msg.GetBody(sect)
		m, err := message.Read(r)
		if err != nil {
			fmt.Println(err)
			// return err
		}
		header := m.Header
		emailDate, _ := net_mail.ParseDate(header.Get("Date"))
		subject := tools.GetSubject(header)
		from := tools.GetFrom(header)
		// 读取邮件内容
		// body, _ := tools.ParseBody(m.Body)
		fmt.Printf("%s 在时间为:%v 发送了主题为:%s \n", from, emailDate, subject)
	}
	return
}

func main() {
	// emailList("imap.163.com:993", "z_frange@163.com", "WAPHMMRZZWDXJUOZ")
	emailList("imap.163.com:993", "yike123120@163.com", "750059?")
}
