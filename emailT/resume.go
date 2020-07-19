package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"time"

	"github.com/axgle/mahonia"
	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
)

const (
	server   = "imap.exmail.qq.com:993"
	username = ""
	password = ""
)

func ConvertToString(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}

// func readMessage(msg *imap.Message, section *imap.BodySectionName) (time.Time, error) {
// 	r := msg.GetBody(section)
// 	if r == nil {
// 		log.Fatal("Server didn't returned message body")
// 	}

// 	// Create a new mail reader
// 	mr, err := mail.CreateReader(r)
// 	if err != nil {

// 	}
// 	var emailDate time.Time
// 	header := mr.Header
// 	if date, err := header.Date(); err == nil {
// 		emailDate = date
// 	}
// 	if from, err := header.AddressList("From"); err == nil {
// 		log.Println("From:", from)
// 	}

// 	for {
// 		p, err := mr.NextPart()
// 		if err == io.EOF {
// 			break
// 		} else if err != nil {
// 			fmt.Printf("===NextPart error:%v", err)
// 		}

// switch h := p.Header.(type) {
// case *message_mail.InlineHeader:
// 	// This is the message's text (can be plain-text or HTML)
// 	b, _ := ioutil.ReadAll(p.Body)
// 	fmt.Printf("Got text: %v", string(b))
// case *message_mail.AttachmentHeader:
// 	// This is an attachment
// 	filename, _ := h.Filename()
// 	fmt.Printf("Got attachment: %v", filename)
// }
// }
// }

func main() {
	var c *client.Client
	var err error
	log.Println("Connecting to server...")
	c, err = client.DialTLS(server, nil)
	//连接失败报错
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected")
	//登陆
	if err := c.Login(username, password); err != nil {
		log.Fatal(err)
	}
	log.Println("Logged in")
	mailboxes := make(chan *imap.MailboxInfo, 20)
	go func() {
		c.List("", "*", mailboxes)
	}()
	//列取邮件夹
	for m := range mailboxes {

		mbox, err := c.Select(m.Name, false)
		fmt.Printf("%v\n", m.Name)
		if err != nil {
			log.Fatal(err)
		}
		if mbox.Messages == 0 {
			// log.Fatal("No message in mailbox")
			fmt.Println("邮件量为0")
			continue
		}
		to := mbox.Messages
		seqSet := new(imap.SeqSet)
		// seqSet.AddNum(mbox.Messages)
		seqSet.AddRange(mbox.Messages-10, mbox.Messages)
		section := &imap.BodySectionName{}
		items := []imap.FetchItem{section.FetchItem()}
		messages := make(chan *imap.Message, 100)
		go func() {
			if err := c.Fetch(seqSet, items, messages); err != nil {
				log.Fatal(err)
			}
		}()

		log.Printf("%s : %d", m.Name, to)
	}
}

func isGBK(data []byte) bool {
	length := len(data)
	var i int = 0
	for i < length {
		//fmt.Printf("for %x\n", data[i])
		if data[i] <= 0xff {
			//编码小于等于127,只有一个字节的编码，兼容ASCII吗
			i++
			continue
		} else {
			//大于127的使用双字节编码
			if data[i] >= 0x81 &&
				data[i] <= 0xfe &&
				data[i+1] >= 0x40 &&
				data[i+1] <= 0xfe &&
				data[i+1] != 0xf7 {
				i += 2
				continue
			} else {
				return false
			}
		}
	}
	return true
}

func convertToString(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}

func readMessage(msg *imap.Message, section *imap.BodySectionName) (time.Time, error) {
	r := msg.GetBody(section)

	if r == nil {
		log.Fatal("Server didn't returned message body")
	}
	var emailDate time.Time
	// Create a new mail reader
	mr, err := message_mail.CreateReader(r)
	if err == nil {
		header := mr.Header
		if date, err := header.Date(); err == nil {
			emailDate = date
		} else {
			fmt.Println("==第2层:", err)
		}
		if from, err := header.AddressList("From"); err == nil {
			log.Println("From:", from)
		} else {
			fmt.Println("==第3层:", err)
		}

		for {
			fmt.Println("======================line========================")
			p, err := mr.NextPart()
			if err == io.EOF {
				break
			} else if err != nil {
				fmt.Printf("NextPart error%v", err)
			}

			switch h := p.Header.(type) {
			case *message_mail.InlineHeader:
				var emailText string
				b, _ := ioutil.ReadAll(p.Body)
				if isGbk := isGBK(b); isGbk {
					emailText = convertToString(string(b), "gbk", "utf-8")
					fmt.Printf("GBK:%v\n", emailText)
				} else {
					emailText = string(b)
					fmt.Printf("UTF8:%v\n", emailText)
				}
				fmt.Printf("====EmailTxt:%v", emailText)
			case *message_mail.AttachmentHeader:
				// This is an attachment
				filename, _ := h.Filename()
				fmt.Printf("Got attachment: %v", filename)
			default:
				fmt.Println("=======defalut=========")
			}
		}
	} else {
		fmt.Println("==第一层:", err)
	}
	return emailDate, err
}
