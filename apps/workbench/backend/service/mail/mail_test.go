package mail

import (
	"context"
	"fmt"
	"testing"
)

func TestMailService_SendMail(t *testing.T) {
	ctx := context.Background()
	ms := mailService{}
	email := "2734170020@qq.com"
	number := ms.SendMail(ctx, email)
	fmt.Println(number)
}
