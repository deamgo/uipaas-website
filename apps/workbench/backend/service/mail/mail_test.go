package mail

import (
	"context"
	"fmt"
	"testing"
)

func TestMailService_SendMail(t *testing.T) {
	ctx := context.Background()
	ms := mailService{}
	email := "zhaokang0610@gmail.com"
	number := ms.SendMail(ctx, email)
	fmt.Println(number)
}
