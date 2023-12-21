package mail

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMailService_SendMail(t *testing.T) {
	ctx := context.Background()
	ms := mailService{}
	email := "zhaokang0610@gmail.com"
	number := ms.SendMail(ctx, email)
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
