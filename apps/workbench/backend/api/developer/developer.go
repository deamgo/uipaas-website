package developer

import (
	"errors"
	"fmt"
	"github.com/deamgo/workbench/auth/jwt"
	"net/http"
	"regexp"
	"time"

	"github.com/deamgo/workbench/context"
	"github.com/deamgo/workbench/db"
	"github.com/deamgo/workbench/pkg/e"
	"github.com/deamgo/workbench/pkg/encryption"
	"github.com/deamgo/workbench/pkg/logger"
	"github.com/deamgo/workbench/pkg/types"
	"github.com/deamgo/workbench/service/developer"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Resp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func DeveloperGetByID(ctx context.ApplicationContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")

		id, err := jwt.ExtractIDFromToken(authHeader)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
		}
		dlp, err := ctx.UserService.DeveloperGetByID(c, id)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.AbortWithStatusJSON(http.StatusOK, types.NewValidResponse(&Resp{
					Code: e.Failed,
					Msg:  "The user does not exist.",
					Data: nil,
				}))
				return
			} else {
				c.AbortWithStatus(http.StatusInternalServerError)
				return
			}
		}
		c.AbortWithStatusJSON(http.StatusOK, types.NewValidResponse(&Resp{
			Code: e.Success,
			Msg:  "The user information was successfully obtained.",
			Data: dlp,
		}))
	}
}
func DeveloperNameModify(ctx context.ApplicationContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			ID       string `json:"id" validate:"required"`
			UserName string `json:"username" validate:"required,min=4,max=20"`
		}
		err := c.ShouldBind(&req)
		if err != nil {
			fmt.Println(err)
			return
		}
		req.ID = c.Param("id")

		// Create an authenticator
		validate := validator.New()

		// Perform validation
		if err := validate.Struct(req); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, types.NewValidResponse(&Resp{
				Code: e.Failed,
				Msg:  "The parameters are not formatted correctly",
				Data: nil,
			}))
			return
		}
		dlp := &developer.Developer{
			ID:       req.ID,
			Username: req.UserName,
		}

		err = ctx.UserService.DeveloperNameModifyByID(c, dlp)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, types.NewValidResponse(&Resp{
				Code: e.Failed,
				Msg:  err.Error(),
				Data: nil,
			}))
			return
		}
		c.AbortWithStatusJSON(http.StatusOK, types.NewValidResponse(&Resp{
			Code: e.Success,
			Msg:  "The username has been modified successfully",
			Data: nil,
		}))
	}
}
func VerifyEmailAndPwd(ctx context.ApplicationContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			Email    string `json:"email" validate:"required,email"`
			Password string `json:"password" validate:"required,verifyPwd"`
		}
		err := c.ShouldBind(&req)
		if err != nil {
			fmt.Println(err)
			return
		}
		validate := validator.New()
		err = validate.RegisterValidation("verifyPwd", verifyPwd)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, types.NewValidResponse(&Resp{
				Code: e.Failed,
				Msg:  err.Error(),
				Data: nil,
			}))
			return
		}
		err = validate.Struct(req)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, types.NewValidResponse(&Resp{
				Code: e.Failed,
				Msg:  "The parameters are not formatted correctly",
				Data: nil,
			}))
			return
		}
		dlp := &developer.Developer{
			Email:    req.Email,
			Password: req.Password,
		}
		dlp.Password = encryption.EncryptPwd(dlp.Password)
		findDlp, err := ctx.UserService.DeveloperGetByEmailAndPwd(c, dlp)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.AbortWithStatusJSON(http.StatusOK, types.NewValidResponse(&Resp{
					Code: e.Failed,
					Msg:  "Incorrect password.",
					Data: nil,
				}))
				logger.Error(err)
				return
			} else {
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}
		if findDlp != nil {
			// The first step in modifying the mailbox logic is to store the state in Redis
			db.RedisDB.HSet(developer.GetEmailHashStr(findDlp.Email), "mod-email-step", 1)

			// Set the expiration time of the verification code to 5 minutes
			db.RedisDB.Expire(developer.GetEmailHashStr(findDlp.Email), 60*5*time.Second)
			c.AbortWithStatusJSON(http.StatusOK, types.NewValidResponse(&Resp{
				Code: e.Success,
				Msg:  "The email and password are correct.",
				Data: nil,
			}))
		}
	}
}
func SendModifyEmailVerify(ctx context.ApplicationContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var email struct {
			Email    string `json:"email" validate:"required,email"`
			OldEmail string `json:"old_email" validate:"required,email"`
		}
		err := c.ShouldBind(&email)
		if err != nil {
			fmt.Println(err)
			return
		}
		validate := validator.New()
		err = validate.Struct(email)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, types.NewValidResponse(&Resp{
				Code: e.Failed,
				Msg:  "The parameters are not formatted correctly",
				Data: nil,
			}))
			return
		}
		// Verify that the password check and email check have passed
		step, _ := db.RedisDB.HGet(developer.GetEmailHashStr(email.OldEmail), "mod-email-step").Int()
		if step != 1 {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		codeKey, err := ctx.UserService.SendModifyEmailVerify(c, &developer.Developer{Email: email.Email})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, types.NewValidResponse(&Resp{
				Code: e.Failed,
				Msg:  err.Error(),
				Data: nil,
			}))
			return
		}
		db.RedisDB.HSet(developer.GetEmailHashStr(email.OldEmail), "mod-email-step", 2)

		c.AbortWithStatusJSON(http.StatusOK, types.NewValidResponse(&Resp{
			Code: e.Success,
			Msg:  "The email has been sent, please pay attention to check",
			Data: struct {
				CodeKey string `json:"code_key"`
			}{codeKey},
		}))
	}
}
func VerifyEmailVerificationCode(ctx context.ApplicationContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			OldEmail string `json:"old_email" validate:"required,email"`
			Email    string `json:"email" validate:"required,email" `
			CodeKey  string `json:"code_key" validate:"required"`
			Code     int    `json:"code" validate:"required"`
		}

		err := c.ShouldBind(&req)
		if err != nil {
			fmt.Println(err)
			return
		}
		validate := validator.New()
		err = validate.Struct(req)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, types.NewValidResponse(&Resp{
				Code: e.Failed,
				Msg:  "The parameters are not formatted correctly",
				Data: nil,
			}))
			return
		}
		isExists := db.RedisDB.HExists(req.CodeKey, "code")
		if !isExists.Val() {
			c.AbortWithStatusJSON(http.StatusBadRequest, types.NewValidResponse(&Resp{
				Code: e.Failed,
				Msg:  e.VerifyCodeExpired,
				Data: nil,
			}))
			return
		}

		getCode, _ := db.RedisDB.HGet(req.CodeKey, "code").Int()
		if getCode != req.Code {
			c.AbortWithStatusJSON(http.StatusBadRequest, types.NewValidResponse(&Resp{
				Code: e.Failed,
				Msg:  e.InvalidVerifyCode,
				Data: nil,
			}))
			return
		}
		// Verify that the password check and email check have passed
		step, _ := db.RedisDB.HGet(developer.GetEmailHashStr(req.OldEmail), "mod-email-step").Int()
		if step != 2 {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		// Delete the verification code that has already been used
		db.RedisDB.HDel(req.CodeKey, "code")

		// Modify the mail
		dlp := &developer.Developer{
			Email: req.Email,
		}
		err = ctx.UserService.DeveloperEmailModifyByEmail(c, req.OldEmail, dlp)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, types.NewValidResponse(&Resp{
				Code: e.Failed,
				Msg:  err.Error(),
				Data: nil,
			}))
			return
		}
		c.AbortWithStatusJSON(http.StatusOK, types.NewValidResponse(&Resp{
			Code: e.Success,
			Msg:  "The email address is modified",
			Data: nil,
		}))

	}
}
func SendModifyPwdVerify(ctx context.ApplicationContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			Email string `json:"email" validate:"required,email"`
		}
		err := c.ShouldBind(&req)
		if err != nil {
			fmt.Println(err)
			return
		}
		validate := validator.New()
		err = validate.Struct(req)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, types.NewValidResponse(&Resp{
				Code: e.Failed,
				Msg:  "The parameters are not formatted correctly",
				Data: nil,
			}))
			return
		}
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, types.NewValidResponse(&Resp{
				Code: e.Failed,
				Msg:  err.Error(),
				Data: nil,
			}))
			return
		}
		codeKey, err := ctx.UserService.ForgotVerifySend(c, &developer.Developer{Email: req.Email})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, types.NewValidResponse(&Resp{
				Code: e.Failed,
				Msg:  err.Error(),
				Data: nil,
			}))
			return
		}

		c.AbortWithStatusJSON(http.StatusOK, types.NewValidResponse(&Resp{
			Code: e.Success,
			Msg:  "The email has been sent, please pay attention to check",
			Data: struct {
				CodeKey string `json:"code_key"`
			}{codeKey},
		}))
	}
}
func VerifyPwdVerificationCode(ctx context.ApplicationContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			Email   string `json:"email" validate:"required,email"`
			CodeKey string `json:"code_key" validate:"required"`
			Code    int    `json:"code" validate:"required"`
		}
		err := c.ShouldBind(&req)
		if err != nil {
			fmt.Println(err)
			return
		}
		validate := validator.New()
		if err = validate.Struct(req); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, types.NewValidResponse(&Resp{
				Code: e.Failed,
				Msg:  "The parameters are not formatted correctly",
				Data: nil,
			}))
			return
		}
		isExists := db.RedisDB.HExists(req.CodeKey, "code")
		if !isExists.Val() {
			c.AbortWithStatusJSON(http.StatusBadRequest, types.NewValidResponse(&Resp{
				Code: e.Failed,
				Msg:  e.VerifyCodeExpired,
				Data: nil,
			}))
			return
		}

		getCode, _ := db.RedisDB.HGet(req.CodeKey, "code").Int()
		if getCode != req.Code {
			c.AbortWithStatusJSON(http.StatusBadRequest, types.NewValidResponse(&Resp{
				Code: e.Failed,
				Msg:  e.InvalidVerifyCode,
				Data: nil,
			}))
			return
		}
		db.RedisDB.HSet(developer.GetEmailHashStr(req.Email), "mod-pwd-step", 1)

		// Delete the verification code that has already been used
		db.RedisDB.HDel(req.CodeKey, "code")

		c.AbortWithStatusJSON(http.StatusOK, types.NewValidResponse(&Resp{
			Code: e.Success,
			Msg:  "Success",
			Data: nil,
		}))
	}
}

func ModifyPwd(ctx context.ApplicationContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			Email    string `json:"email" validate:"required,email"`
			Password string `json:"password" validate:"required,verifyPwd"`
		}
		err := c.ShouldBind(&req)
		if err != nil {
			fmt.Println(err)
			return
		}
		validate := validator.New()
		if err = validate.RegisterValidation("verifyPwd", verifyPwd); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, types.NewValidResponse(&Resp{
				Code: e.Failed,
				Msg:  err.Error(),
				Data: nil,
			}))
			return
		}
		if err = validate.Struct(req); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, types.NewValidResponse(&Resp{
				Code: e.Failed,
				Msg:  "The parameters are not formatted correctly",
				Data: nil,
			}))
			return
		}
		// Verify that the password check and email check have passed
		step, _ := db.RedisDB.HGet(developer.GetEmailHashStr(req.Email), "mod-pwd-step").Int()
		if step != 1 {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		var dlp = &developer.Developer{
			Email:    req.Email,
			Password: req.Password,
		}
		dlp.Password = encryption.EncryptPwd(dlp.Password)
		err = ctx.UserService.DeveloperPasswordModifyByEmail(c, dlp)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, types.NewValidResponse(&Resp{
				Code: e.Failed,
				Msg:  err.Error(),
				Data: nil,
			}))
			return
		}
		c.AbortWithStatusJSON(http.StatusOK, types.NewValidResponse(&Resp{
			Code: e.Success,
			Msg:  "The password was successfully changed.",
			Data: nil,
		}))
	}
}

// validator password
func verifyPwd(f validator.FieldLevel) bool {
	val := f.Field().String()
	if len(val) < 8 || len(val) > 20 { // length需要通过验证
		fmt.Println("pwd length error")
		return false
	}

	pwdPattern := `^[a-zA-Z0-9]{8,20}$`
	reg, err := regexp.Compile(pwdPattern) // filter exclude chars
	if err != nil {
		fmt.Println(err)
		return false
	}

	match := reg.MatchString(val)
	if !match {
		fmt.Println("not match error.")
		return false
	}
	return true
}
