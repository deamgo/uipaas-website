package mail

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/smtp"
	"os"
	"path"
	"runtime"
	"time"

	"github.com/Boostport/mjml-go"
	"github.com/jordan-wright/email"
)

var (
	SUBJECT_LINE = "UIPaaS Email verification code"
	ADDRESSER    = "uipaas@tests.run"
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
	htmlStr, err := parseMJMLFile(getCurrentAbPathByCaller()+"/mjml/verification_code.mjml", ctx)
	if err != nil {
		log.Fatal(err)
	}
	// 定义包含变量的结构
	data := struct {
		Code int
	}{
		Code: randomNumber,
	}
	tmpl, _ := template.New("mjml").Parse(htmlStr)
	var rendered bytes.Buffer
	_ = tmpl.Execute(&rendered, data)
	e.HTML = rendered.Bytes()
	// Set the server-related configurations
	err = e.Send("smtp.feishu.cn:25", smtp.PlainAuth("", "uipaas@tests.run", "rR9rJvSiXkfAm44h", "smtp.feishu.cn"))
	if err != nil {
		log.Fatal(err)
	}
	return randomNumber

}

func parseMJMLFile(filePath string, ctx context.Context) (string, error) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	html, err := mjml.ToHTML(ctx, string(file), mjml.WithMinify(true))
	if err != nil {
		return "", err
	}
	return html, nil
}

func getCurrentAbPathByCaller() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(filename)
	}
	return abPath
}
