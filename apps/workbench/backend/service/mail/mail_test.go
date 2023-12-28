package mail

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMailService_SendMail(t *testing.T) {
	ctx := context.Background()
	ms := mailService{}
	email := "zhaokang0610@gmail.com"
	number := ms.SendVerificationCodeMail(ctx, email)
	fmt.Println(number)
}

func TestMailService_SendWorkspaceInviteMail(t *testing.T) {
	ctx := context.Background()
	ms := mailService{}
	email := "zhaokang0610@gmail.com"
	number := ms.SendWorkspaceInviteMail(ctx, email, "test22")
	fmt.Println(number)
}

func TestParseMJMLFile_Success(t *testing.T) {
	ctx := context.Background()
	filePath := getCurrentAbPathByCaller() + "/mjml/verification_code.mjml"
	html, err := parseMJMLFile(filePath, ctx)
	assert.Nil(t, err)
	assert.NotEmpty(t, html)
}

func TestParseMJMLFile_FileNotFound(t *testing.T) {
	ctx := context.Background()
	filePath := "nonexistent.mjml"
	html, err := parseMJMLFile(filePath, ctx)
	assert.NotNil(t, err)
	assert.Empty(t, html)
}

func TestGetCurrentAbPathByCaller_Success(t *testing.T) {
	path := getCurrentAbPathByCaller()
	assert.NotEmpty(t, path)
}
