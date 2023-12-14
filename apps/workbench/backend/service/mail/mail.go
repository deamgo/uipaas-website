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

var (
	SUBJECT_LINE = "UIPaaS Email verification code"
	ADDRESSER    = "uipaas@tests.run"
	CLOSING      = "If you didn't try to sign in or change email,please ignore this email."
)

type MailService interface {
	SendMail(ctx context.Context, emailStr string) int
}

type mailService struct {
}

func NewMailService() MailService {
	return &mailService{}
}

func (us mailService) SendMail(ctx context.Context, emailStr string) int {

	// Set up a random seed
	rand.NewSource(time.Now().UnixNano())
	// Generate a four-digit random number
	randomNumber := rand.Intn(9000) + 1000
	e := email.NewEmail()
	// Set the sender's mailbox
	e.From = fmt.Sprintf("UIPaaS <%v>", ADDRESSER)
	// Set up the recipient's mailbox
	e.To = []string{emailStr}
	// Set up a subject
	e.Subject = SUBJECT_LINE
	// Set the content of the file to be sent
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
	// Set the server-related configurations
	err := e.Send("smtp.feishu.cn:25", smtp.PlainAuth("", "uipaas@tests.run", "rR9rJvSiXkfAm44h", "smtp.feishu.cn"))
	if err != nil {
		log.Fatal(err)
	}
	return randomNumber

}
