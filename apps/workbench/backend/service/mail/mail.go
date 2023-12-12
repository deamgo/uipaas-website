package mail

import (
	"context"
	"fmt"
	"github.com/jordan-wright/email"
	"log"
	"math/rand"
	"net/smtp"
	"time"
)

// 1.邮件说明
// Subject line(主题)：UIPaaS Email verification code
// Addresser（发件人）：notify@mail.uipaas.com
// Receiver（收件人）：-
// Salutation（问候语）：-
// Main point（正文）：UIPaaS Email verification code
// Closing（收尾）：If you didn't try to sign in or change email,please ignore this email.
// Sign off（签名）：logo+website
var (
	SUBJECT_LINE = "UIPaaS Email verification code"
	ADDRESSER    = "uipaas@tests.run"
	CLOSING      = "If you didn't try to sign in or change email,please ignore this email."
)

func SendMail(ctx context.Context, emailStr string) int {

	// 设置随机种子
	rand.Seed(time.Now().UnixNano())
	// 生成四位随机数
	randomNumber := rand.Intn(9000) + 1000
	e := email.NewEmail()
	//设置发送方的邮箱
	e.From = fmt.Sprintf("UIPaaS <%v>", ADDRESSER)
	// 设置接收方的邮箱
	e.To = []string{emailStr}
	//设置主题
	e.Subject = SUBJECT_LINE
	//设置文件发送的内容
	html := fmt.Sprintf(`
		<h1>UIPaaS Email verification code</h1>
		<br/>
		<h1 style= "bakcgroundcolor">%d</h1>
		<p>If you didn't try to sign in or change email,please ignore this email.</p>
		<br/>
		<h1>UIPaas</h1>
		<p>AI-base, future-oriented,drag-and-drop development platform.</p>
		<a href ="https://uipaas.com">https://www.uipaas.com</a>
	`, randomNumber)
	e.HTML = []byte(html)
	//设置服务器相关的配置
	err := e.Send("smtp.feishu.cn:25", smtp.PlainAuth("", "uipaas@tests.run", "rR9rJvSiXkfAm44h", "smtp.feishu.cn"))
	if err != nil {
		log.Fatal(err)
	}
	return randomNumber

}
