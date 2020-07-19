package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strings"

	"github.com/axgle/mahonia"
	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
	"github.com/emersion/go-message/mail"
)

const (
	server   = "imap.exmail.qq.com:993"
	username = "mailbackups05@aimsen.com"
	password = "Ruyu500"

	// server   = "imap.exmail.qq.com:993"
	// username = "zhangfengguang@hua-yong.com"
	// password = "admin123ZFG"
)

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

func readMessage(msg *imap.Message, section *imap.BodySectionName) {
	fmt.Println("======into readMessage=======")
	r := msg.GetBody(section)
	if r == nil {
		log.Fatal("Server didn't returned message body")
	}

	// Create a new mail reader
	mr, err := mail.CreateReader(r)
	if err != nil {
		// log.Fatal(err)
		fmt.Println(err, "++++++++++++++++")
	}
	// Print some info about the message
	header := mr.Header
	if date, err := header.Date(); err == nil {
		log.Println("Date:", date)
	}
	if _, err := header.AddressList("From"); err == nil {
		from := header.Get("From")
		fmt.Printf("ID:%v", header.MessageID)
		if strings.Contains(from, "b1.service@zhaopinmail.com") {
			for {
				p, err := mr.NextPart()
				if err == io.EOF {
					break
				} else if err != nil {
					fmt.Printf("NextPart error%v", err)
				}

				switch h := p.Header.(type) {
				case *mail.InlineHeader:
					// This is the message's text (can be plain-text or HTML)
					fmt.Println("======================line========================")

					b, _ := ioutil.ReadAll(p.Body)
					var emailText string
					isGbk := isGBK(b)
					if isGbk {
						emailText = convertToString(string(b), "gbk", "utf-8")
						fmt.Printf("gbk:%v\n", emailText)
					} else {
						emailText = string(b)
						fmt.Printf("UTF8:%v\n", emailText)
					}
				case *mail.AttachmentHeader:
					// This is an attachment
					filename, _ := h.Filename()
					fmt.Printf("Got attachment: %v", filename)
				default:
					fmt.Println("=======defalut=========")
				}
			}
		}
	}
	// if to, err := header.AddressList("To"); err == nil {
	// 	log.Println("To:", to)
	// }
	// if subject, err := header.Subject(); err == nil {
	// 	fmt.Println(subject)
	// } else {
	// 	fmt.Println("SUBJECT ERROR:", ConvertToString(subject, "GBK", "utf-8"))
	// }
	// Process each message's part

}

// func connection(server string) (c *client.Client, err error) {
// 	c, err = client.DialTLS(server, nil)
// 	return c, err
// }

// func login(c *client.Client, username, password string) (err error) {
// 	//登陆
// 	err = c.Login(username, password)
// 	return err
// }

func main() {
	// c, err := connection(server)
	// if err != nil {
	// 	fmt.Println("connect error: %v", err)
	// }
	// c, err = login(c, username, password)
	// if err != nil {
	// 	fmt.Println("connect error: %v", err)
	// }

	var c *client.Client
	var err error
	c, err = client.DialTLS(server, nil)
	if err != nil {
		log.Fatal(err)
	}
	//登陆
	if err := c.Login(username, password); err != nil {
		log.Fatal(err)
	}
	// mailboxes := make(chan *imap.MailboxInfo, 20)

	// go func() {
	// 	c.List("", "*", mailboxes)
	// }()
	//列取邮件夹

	mbox, err := c.Select("INBOX", false)
	if err != nil {
		fmt.Printf("error:%v", err)
	}
	if mbox.Messages == 0 {
		fmt.Println("邮件量为0")
	}
	to := mbox.Messages
	fmt.Printf("%v", to)
	seqSet := new(imap.SeqSet)
	// seqSet.AddNum(to)
	seqSet.AddRange(mbox.Messages-2, mbox.Messages)
	section := &imap.BodySectionName{}
	items := []imap.FetchItem{section.FetchItem()}
	messages := make(chan *imap.Message, 2)
	go func() {
		if err := c.Fetch(seqSet, items, messages); err != nil {
			fmt.Println("list error:", err)
		}
	}()

	for msg := range messages {
		if msg == nil {
			log.Fatal("Server didn't returned message")
		}
		readMessage(msg, section)
	}
}
