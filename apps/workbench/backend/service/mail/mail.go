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
	"time"

	"github.com/deamgo/workbench/pkg/consts"

	"github.com/Boostport/mjml-go"
	"github.com/jordan-wright/email"
)

type MailService interface {
	SendVerificationCodeMail(ctx context.Context, emailStr string) int
	SendWorkspaceInviteMail(ctx context.Context, emailStr, workspaceID string) error
}

type mailService struct {
}

func NewMailService() MailService {
	return &mailService{}
}

func (us mailService) SendVerificationCodeMail(ctx context.Context, emailStr string) int {

	// Set up a random seed
	rand.NewSource(time.Now().UnixNano())
	// Generate a four-digit random number
	randomNumber := rand.Intn(9000) + 1000
	e := email.NewEmail()
	// Set the sender's mailbox
	e.From = fmt.Sprintf("UIPaaS <%v>", consts.ADDRESSER)
	// Set up the recipient's mailbox
	e.To = []string{emailStr}
	// Set up a subject
	e.Subject = consts.SUBJECT_LINE
	// Set the content of the file to be sent
	htmlStr, err := parseMJMLFile("./mjml/verification_code.mjml", ctx)
	if err != nil {
		log.Fatal(err)
	}
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

func (us mailService) SendWorkspaceInviteMail(ctx context.Context, emailStr, workspaceID string) error {

	url := "https://uipaas.com/" + workspaceID
	e := email.NewEmail()
	// Set the sender's mailbox
	e.From = fmt.Sprintf("UIPaaS <%v>", consts.ADDRESSER)
	// Set up the recipient's mailbox
	e.To = []string{emailStr}
	// Set up a subject
	e.Subject = consts.SUBJECT_LINE
	// Set the content of the file to be sent
	htmlStr, err := parseMJMLFile("./mjml/workspace_invite.mjml", ctx)
	if err != nil {
		return err
	}
	data := struct {
		URL string
	}{
		URL: url,
	}
	tmpl, _ := template.New("mjml").Parse(htmlStr)
	var rendered bytes.Buffer
	_ = tmpl.Execute(&rendered, data)
	e.HTML = rendered.Bytes()
	// Set the server-related configurations
	err = e.Send("smtp.feishu.cn:25", smtp.PlainAuth("", "uipaas@tests.run", "rR9rJvSiXkfAm44h", "smtp.feishu.cn"))
	if err != nil {
		return err
	}
	return err

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
