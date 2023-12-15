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
		<!DOCTYPE html>
<html lang="en">

<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>UIPaaS Email verification</title>
</head>

<body>
  <div style="position: relative; margin-left: auto; width: 730px; height: 362px;">
    <h1
      style="position: absolute; top: 43px; left: 53px; width: max-content; height: 28px; font-size: 18px; font-weight: bold; line-height: 28px; color: #3D3D3D;">
      UIPaaS Email verification code
    </h1>
    <div
      style="position: absolute; left: 53px; top: 137px; width: 113px; height: 44px; background-color: rgba(216, 216, 216, 0.4392);">
      <span
        style="position: absolute; top: 8px; left: 21px; width: max-content; height: 28px; color: #3D3D3D; font-size: 18px; font-weight: normal; line-height: 28px; letter-spacing: 10px;">
        %d
      </span>
    </div>
    <p
      style="position: absolute; left: 54px; top: 200px; margin: 0; width: 432; height: 28px; font-size: 14px; font-weight: normal; line-height: 28px;color: #3D3D3D;">
      If you didn't try to sign in or change email,please ignore this email.</p>
    <div
      style="position: absolute; left: 54px; top: 263px; width: 120px; height: 18.14px; background-image: url('data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHhtbG5zOnhsaW5rPSJodHRwOi8vd3d3LnczLm9yZy8xOTk5L3hsaW5rIiBmaWxsPSJub25lIiB2ZXJzaW9uPSIxLjEiIHdpZHRoPSIxMjAiIGhlaWdodD0iMTguMTM5NTI2MzY3MTg3NSIgdmlld0JveD0iMCAwIDEyMCAxOC4xMzk1MjYzNjcxODc1Ij48Zz48Zz48Zz48cGF0aCBkPSJNMjAuMzY3OTgwNDY4MTM5NjUsMC4wMDEyMTUwOUwyMC41Mjc3MzA0NjgxMzk2NDgsMEwyMC41Mjc3MzA0NjgxMzk2NDgsMTguMTM5NUwxMC44Njc2MjA0NjgxMzk2NDgsMTguMTM5NUwxMC44Njc2MjA0NjgxMzk2NDgsOS4wNjk3N0MxMC44Njc2MjA0NjgxMzk2NDgsNC4xNjAzNSwxNS4wMjIxODA0NjgxMzk2NDgsMC4xNjE5OSwyMC4yMTA0NDA0NjgxMzk2NDgsMC4wMDQ3OTk2N0wyMC4zNjc5ODA0NjgxMzk2NSwwLjAwMTIxNTA5WiIgZmlsbC1ydWxlPSJldmVub2RkIiBmaWxsPSIjMDg3MUYwIiBmaWxsLW9wYWNpdHk9IjEiLz48L2c+PGc+PHBhdGggZD0iTTAuMTU5NzQ4LDAuMDAxMjE1MDlMMCwwTDAsOS4wNjk3N0MwLDE0LjAyODgsNC4yMzg5MSwxOC4wNTgyLDkuNTAwMzYsMTguMTM4M0w5LjY2MDExLDE4LjEzOTVMOS42NjAxMSw5LjA2OTc3QzkuNjYwMTEsNC4xNjAzNSw1LjUwNTU1LDAuMTYxOTksMC4zMTcyODEsMC4wMDQ3OTk2N0wwLjE1OTc0OCwwLjAwMTIxNTA5WiIgZmlsbC1ydWxlPSJldmVub2RkIiBmaWxsPSIjMDg3MUYwIiBmaWxsLW9wYWNpdHk9IjEiLz48L2c+PGc+PHBhdGggZD0iTTIxLjczNTI0MDkzNjI3OTI5NywxLjM1NTMxZS0zMUwyMy41NDY1MTA5MzYyNzkyOTcsMS4zNTUzMWUtMzFDMjcuODgxMzAwOTM2Mjc5Mjk1LC03Ljk3NDY4ZS0xNiwzMS4zOTUzNTA5MzYyNzkyOTYsMy41MTkyNSwzMS4zOTUzNTA5MzYyNzkyOTYsNy44NjA0NkMzMS4zOTUzNTA5MzYyNzkyOTYsMTIuMjAxNywyNy44ODEzMDA5MzYyNzkyOTUsMTUuNzIwOSwyMy41NDY1MTA5MzYyNzkyOTcsMTUuNzIwOUwyMS43MzUyNDA5MzYyNzkyOTcsMTUuNzIwOUwyMS43MzUyNDA5MzYyNzkyOTcsMS4zNTUzMWUtMzFaIiBmaWxsLXJ1bGU9ImV2ZW5vZGQiIGZpbGw9IiMwODcxRjAiIGZpbGwtb3BhY2l0eT0iMSIvPjwvZz48L2c+PGc+PHBhdGggZD0iTTM5Ljc2NzQ0MDc5NTg5ODQ0LDEwLjk5NjM3NDM5MzY1Mzg3QzM5Ljc2NzQ0MDc5NTg5ODQ0LDE1LjI5MzM3NDM5MzY1Mzg3LDQyLjcxNTYyMDc5NTg5ODQ0LDE3LjM5NTE3NDM5MzY1Mzg3LDQ2LjI5NTU1MDc5NTg5ODQzNiwxNy4zOTUxNzQzOTM2NTM4N0M0OS44NzU0NDA3OTU4OTg0NCwxNy4zOTUxNzQzOTM2NTM4Nyw1Mi45NjQwNDA3OTU4OTg0NCwxNS4yOTMzNzQzOTM2NTM4Nyw1Mi45NjQwNDA3OTU4OTg0NCwxMC45OTYzNzQzOTM2NTM4N0w1Mi45NjQwNDA3OTU4OTg0NCwwLjkzMTIwNTM5MzY1Mzg2OTZMNDkuNjY0ODkwNzk1ODk4NDQsMC45MzEyMDUzOTM2NTM4Njk2TDQ5LjY2NDg5MDc5NTg5ODQ0LDExLjAxOTc3NDM5MzY1Mzg3QzQ5LjY2NDg5MDc5NTg5ODQ0LDEzLjI2MTY3NDM5MzY1Mzg3LDQ4LjQ0ODE5MDc5NTg5ODQ0LDE0LjQyOTI3NDM5MzY1Mzg3LDQ2LjM0MjM0MDc5NTg5ODQ0LDE0LjQyOTI3NDM5MzY1Mzg3QzQ0LjI1OTkwMDc5NTg5ODQ0LDE0LjQyOTI3NDM5MzY1Mzg3LDQzLjA0MzE5MDc5NTg5ODQ0LDEzLjI2MTY3NDM5MzY1Mzg3LDQzLjA0MzE5MDc5NTg5ODQ0LDExLjAxOTc3NDM5MzY1Mzg3TDQzLjA0MzE5MDc5NTg5ODQ0LDAuOTMxMjA1MzkzNjUzODY5NkwzOS43Njc0NDA3OTU4OTg0NCwwLjkzMTIwNTM5MzY1Mzg2OTZMMzkuNzY3NDQwNzk1ODk4NDQsMTAuOTk2Mzc0MzkzNjUzODdaTTU2LjEyMjg0MDc5NTg5ODQ0LDE3LjIzMTY3NDM5MzY1Mzg3TDU5LjM5ODU0MDc5NTg5ODQ0LDE3LjIzMTY3NDM5MzY1Mzg3TDU5LjM5ODU0MDc5NTg5ODQ0LDAuOTMxMjA1MzkzNjUzODY5Nkw1Ni4xMjI4NDA3OTU4OTg0NCwwLjkzMTIwNTM5MzY1Mzg2OTZMNTYuMTIyODQwNzk1ODk4NDQsMTcuMjMxNjc0MzkzNjUzODdaTTY1LjkwMzI0MDc5NTg5ODQ0LDguMzM0MTQ0MzkzNjUzODY5TDY1LjkwMzI0MDc5NTg5ODQ0LDMuNTkzNDY0MzkzNjUzODY5NUw2OC41MDA0NDA3OTU4OTg0NCwzLjU5MzQ2NDM5MzY1Mzg2OTVDNzAuMzQ4OTQwNzk1ODk4NDQsMy41OTM0NjQzOTM2NTM4Njk1LDcxLjE2Nzg0MDc5NTg5ODQ0LDQuNDgwODg0MzkzNjUzODcsNzEuMTY3ODQwNzk1ODk4NDQsNS45NzU0ODQzOTM2NTM4NjlDNzEuMTY3ODQwNzk1ODk4NDQsNy40MjMzNzQzOTM2NTM4NjksNzAuMzQ4OTQwNzk1ODk4NDQsOC4zMzQxNDQzOTM2NTM4NjksNjguNTAwNDQwNzk1ODk4NDQsOC4zMzQxNDQzOTM2NTM4NjlMNjUuOTAzMjQwNzk1ODk4NDQsOC4zMzQxNDQzOTM2NTM4NjlaTTc0LjUzNzI0MDc5NTg5ODQ0LDUuOTc1NDg0MzkzNjUzODY5Qzc0LjUzNzI0MDc5NTg5ODQ0LDMuMTQ5NzU0MzkzNjUzODY5Nyw3Mi41NzE3NDA3OTU4OTg0NCwwLjkzMTIwNTM5MzY1Mzg2OTYsNjguNjQwODQwNzk1ODk4NDQsMC45MzEyMDUzOTM2NTM4Njk2TDYyLjYyNzU0MDc5NTg5ODQ0LDAuOTMxMjA1MzkzNjUzODY5Nkw2Mi42Mjc1NDA3OTU4OTg0NCwxNy4yMzE2NzQzOTM2NTM4N0w2NS45MDMyNDA3OTU4OTg0NCwxNy4yMzE2NzQzOTM2NTM4N0w2NS45MDMyNDA3OTU4OTg0NCwxMC45NzMwNzQzOTM2NTM4NjlMNjguNjQwODQwNzk1ODk4NDQsMTAuOTczMDc0MzkzNjUzODY5QzcyLjg1MjU0MDc5NTg5ODQzLDEwLjk3MzA3NDM5MzY1Mzg2OSw3NC41MzcyNDA3OTU4OTg0NCw4LjQ1MDkxNDM5MzY1Mzg3LDc0LjUzNzI0MDc5NTg5ODQ0LDUuOTc1NDg0MzkzNjUzODY5Wk03Ni4wMTEzNDA3OTU4OTg0MywxMC43MTYxNzQzOTM2NTM4N0M3Ni4wMTEzNDA3OTU4OTg0MywxNC43NTYyNzQzOTM2NTM4Nyw3OC42MzE5NDA3OTU4OTg0MywxNy40NDE4NzQzOTM2NTM4Nyw4MS45MDc2NDA3OTU4OTg0MywxNy40NDE4NzQzOTM2NTM4N0M4My45NjY3NDA3OTU4OTg0MywxNy40NDE4NzQzOTM2NTM4Nyw4NS40NDA4NDA3OTU4OTg0NCwxNi40NjEwNzQzOTM2NTM4Nyw4Ni4yMTI5NDA3OTU4OTg0NSwxNS4zNDAwNzQzOTM2NTM4N0w4Ni4yMTI5NDA3OTU4OTg0NSwxNy4yMzE2NzQzOTM2NTM4N0w4OS41MTIwNDA3OTU4OTg0MywxNy4yMzE2NzQzOTM2NTM4N0w4OS41MTIwNDA3OTU4OTg0Myw0LjI5NDA1NDM5MzY1Mzg2OTVMODYuMjEyOTQwNzk1ODk4NDUsNC4yOTQwNTQzOTM2NTM4Njk1TDg2LjIxMjk0MDc5NTg5ODQ1LDYuMTM4OTU0MzkzNjUzODY5NUM4NS40NDA4NDA3OTU4OTg0NCw1LjA2NDcwNDM5MzY1Mzg2OSw4NC4wMTM1NDA3OTU4OTg0NCw0LjA4Mzg3NDM5MzY1Mzg3LDgxLjkzMTA0MDc5NTg5ODQ0LDQuMDgzODc0MzkzNjUzODdDNzguNjMxOTQwNzk1ODk4NDMsNC4wODM4NzQzOTM2NTM4Nyw3Ni4wMTEzNDA3OTU4OTg0Myw2LjY3NjA3NDM5MzY1Mzg2OSw3Ni4wMTEzNDA3OTU4OTg0MywxMC43MTYxNzQzOTM2NTM4N1pNODYuMjEyOTQwNzk1ODk4NDUsMTAuNzYyODc0MzkzNjUzODdDODYuMjEyOTQwNzk1ODk4NDUsMTMuMjE0OTc0MzkzNjUzODcsODQuNTc1MDQwNzk1ODk4NDUsMTQuNTY5NDc0MzkzNjUzODcsODIuNzczNDQwNzk1ODk4NDQsMTQuNTY5NDc0MzkzNjUzODdDODEuMDE4NTQwNzk1ODk4NDUsMTQuNTY5NDc0MzkzNjUzODcsNzkuMzU3MjQwNzk1ODk4NDMsMTMuMTY4Mjc0MzkzNjUzODY5LDc5LjM1NzI0MDc5NTg5ODQzLDEwLjcxNjE3NDM5MzY1Mzg3Qzc5LjM1NzI0MDc5NTg5ODQzLDguMjY0MDg0MzkzNjUzODcsODEuMDE4NTQwNzk1ODk4NDUsNi45NTYzMTQzOTM2NTM4NjksODIuNzczNDQwNzk1ODk4NDQsNi45NTYzMTQzOTM2NTM4NjlDODQuNTc1MDQwNzk1ODk4NDUsNi45NTYzMTQzOTM2NTM4NjksODYuMjEyOTQwNzk1ODk4NDUsOC4zMTA3OTQzOTM2NTM4Nyw4Ni4yMTI5NDA3OTU4OTg0NSwxMC43NjI4NzQzOTM2NTM4N1pNOTEuODc1MzQwNzk1ODk4NDQsMTAuNzE2MTc0MzkzNjUzODdDOTEuODc1MzQwNzk1ODk4NDQsMTQuNzU2Mjc0MzkzNjUzODcsOTQuNDk1OTQwNzk1ODk4NDMsMTcuNDQxODc0MzkzNjUzODcsOTcuNzcxNjQwNzk1ODk4NDMsMTcuNDQxODc0MzkzNjUzODdDOTkuODMwNzQwNzk1ODk4NDQsMTcuNDQxODc0MzkzNjUzODcsMTAxLjMwNDg0MDc5NTg5ODQzLDE2LjQ2MTA3NDM5MzY1Mzg3LDEwMi4wNzY5NDA3OTU4OTg0NCwxNS4zNDAwNzQzOTM2NTM4N0wxMDIuMDc2OTQwNzk1ODk4NDQsMTcuMjMxNjc0MzkzNjUzODdMMTA1LjM3NjA0MDc5NTg5ODQzLDE3LjIzMTY3NDM5MzY1Mzg3TDEwNS4zNzYwNDA3OTU4OTg0Myw0LjI5NDA1NDM5MzY1Mzg2OTVMMTAyLjA3Njk0MDc5NTg5ODQ0LDQuMjk0MDU0MzkzNjUzODY5NUwxMDIuMDc2OTQwNzk1ODk4NDQsNi4xMzg5NTQzOTM2NTM4Njk1QzEwMS4zMDQ4NDA3OTU4OTg0Myw1LjA2NDcwNDM5MzY1Mzg2OSw5OS44Nzc1NDA3OTU4OTg0NCw0LjA4Mzg3NDM5MzY1Mzg3LDk3Ljc5NTA0MDc5NTg5ODQ0LDQuMDgzODc0MzkzNjUzODdDOTQuNDk1OTQwNzk1ODk4NDMsNC4wODM4NzQzOTM2NTM4Nyw5MS44NzUzNDA3OTU4OTg0NCw2LjY3NjA3NDM5MzY1Mzg2OSw5MS44NzUzNDA3OTU4OTg0NCwxMC43MTYxNzQzOTM2NTM4N1pNMTAyLjA3Njk0MDc5NTg5ODQ0LDEwLjc2Mjg3NDM5MzY1Mzg3QzEwMi4wNzY5NDA3OTU4OTg0NCwxMy4yMTQ5NzQzOTM2NTM4NywxMDAuNDM5MDQwNzk1ODk4NDQsMTQuNTY5NDc0MzkzNjUzODcsOTguNjM3NDQwNzk1ODk4NDQsMTQuNTY5NDc0MzkzNjUzODdDOTYuODgyNTQwNzk1ODk4NDQsMTQuNTY5NDc0MzkzNjUzODcsOTUuMjIxMjQwNzk1ODk4NDQsMTMuMTY4Mjc0MzkzNjUzODY5LDk1LjIyMTI0MDc5NTg5ODQ0LDEwLjcxNjE3NDM5MzY1Mzg3Qzk1LjIyMTI0MDc5NTg5ODQ0LDguMjY0MDg0MzkzNjUzODcsOTYuODgyNTQwNzk1ODk4NDQsNi45NTYzMTQzOTM2NTM4NjksOTguNjM3NDQwNzk1ODk4NDQsNi45NTYzMTQzOTM2NTM4NjlDMTAwLjQzOTA0MDc5NTg5ODQ0LDYuOTU2MzE0MzkzNjUzODY5LDEwMi4wNzY5NDA3OTU4OTg0NCw4LjMxMDc5NDM5MzY1Mzg3LDEwMi4wNzY5NDA3OTU4OTg0NCwxMC43NjI4NzQzOTM2NTM4N1pNMTIwLjAwMDA0MDc5NTg5ODQ0LDEyLjUxNDM3NDM5MzY1Mzg3QzEyMC4wMDAwNDA3OTU4OTg0NCw2Ljc2OTQ4NDM5MzY1Mzg3LDExMS42MjM0NDA3OTU4OTg0Myw4LjU5MTAyNDM5MzY1Mzg3LDExMS42MjM0NDA3OTU4OTg0Myw1LjMyMTU5NDM5MzY1Mzg3QzExMS42MjM0NDA3OTU4OTg0MywzLjk5MDQ2NDM5MzY1Mzg2OTcsMTEyLjU4Mjc0MDc5NTg5ODQzLDMuMzU5OTM0MzkzNjUzODY5NSwxMTMuODY5NjQwNzk1ODk4NDMsMy40MDY2MzQzOTM2NTM4Njk0QzExNS4yNzM1NDA3OTU4OTg0NCwzLjQyOTk5NDM5MzY1Mzg2OTcsMTE2LjEzOTI0MDc5NTg5ODQzLDQuMjcwNzA0MzkzNjUzODcsMTE2LjIwOTQ0MDc5NTg5ODQzLDUuMzQ0OTQ0MzkzNjUzODY5TDExOS44MTI4NDA3OTU4OTg0NCw1LjM0NDk0NDM5MzY1Mzg2OUMxMTkuNjAyMjQwNzk1ODk4NDQsMi40MDI0NTQzOTM2NTM4Njk2LDExNy4zMDkyNDA3OTU4OTg0MywwLjY5NzY3NDM5MzY1Mzg2OTYsMTEzLjk2MzI0MDc5NTg5ODQ0LDAuNjk3Njc0MzkzNjUzODY5NkMxMTAuNTkzOTQwNzk1ODk4NDMsMC42OTc2NzQzOTM2NTM4Njk2LDEwOC4yMDcyNDA3OTU4OTg0NCwyLjQ3MjUxNDM5MzY1Mzg3LDEwOC4yMDcyNDA3OTU4OTg0NCw1LjQxNTAwNDM5MzY1Mzg2OUMxMDguMTgzODQwNzk1ODk4NDMsMTEuMzIzMzc0MzkzNjUzODcsMTE2LjYwNzI0MDc5NTg5ODQzLDkuMTk4MjE0MzkzNjUzODcsMTE2LjYwNzI0MDc5NTg5ODQzLDEyLjcyNDU3NDM5MzY1Mzg2OUMxMTYuNjA3MjQwNzk1ODk4NDMsMTMuODkyMTc0MzkzNjUzODcsMTE1LjY5NDc0MDc5NTg5ODQ0LDE0LjcwOTU3NDM5MzY1Mzg3LDExNC4xNTA0NDA3OTU4OTg0MywxNC43MDk1NzQzOTM2NTM4N0MxMTIuNjI5NTQwNzk1ODk4NDQsMTQuNzA5NTc0MzkzNjUzODcsMTExLjc2Mzg0MDc5NTg5ODQzLDEzLjg0NTQ3NDM5MzY1Mzg3LDExMS42NzAyNDA3OTU4OTg0NCwxMi41Mzc2NzQzOTM2NTM4N0wxMDguMTYwNDQwNzk1ODk4NDQsMTIuNTM3Njc0MzkzNjUzODdDMTA4LjIwNzI0MDc5NTg5ODQ0LDE1LjYyMDI3NDM5MzY1Mzg2OSwxMTAuODI3ODQwNzk1ODk4NDQsMTcuMzk1MTc0MzkzNjUzODcsMTE0LjIyMDY0MDc5NTg5ODQzLDE3LjM5NTE3NDM5MzY1Mzg3QzExNy45NDA5NDA3OTU4OTg0NCwxNy4zOTUxNzQzOTM2NTM4NywxMjAuMDAwMDQwNzk1ODk4NDQsMTUuMDgzMTc0MzkzNjUzODcsMTIwLjAwMDA0MDc5NTg5ODQ0LDEyLjUxNDM3NDM5MzY1Mzg3WiIgZmlsbD0iIzA4MDgwOCIgZmlsbC1vcGFjaXR5PSIxIi8+PC9nPjwvZz48L3N2Zz4=');">
    </div>
    <div
      style="position: absolute; left: 54px; top: 288px; width: 503px; height: 56px; display: flex; flex-direction: column; justify-content: space-between;">
      <p style="margin: 0; font-size: 14px; font-weight: normal; line-height: 28px;">AI-base,
        future-oriented,drag-and-drop
        development platform.</p>
      <a href="https://www.uipaas.com"
        style="font-size: 14px; font-weight: normal; line-height: 28px; text-decoration: none; color: #1677FF;">https://www.uipaas.com</a>
    </div>
  </div>
</body>

</html>
	`, randomNumber)
	e.HTML = []byte(html)
	// Set the server-related configurations
	err := e.Send("smtp.feishu.cn:25", smtp.PlainAuth("", "uipaas@tests.run", "rR9rJvSiXkfAm44h", "smtp.feishu.cn"))
	if err != nil {
		log.Fatal(err)
	}
	return randomNumber

}
