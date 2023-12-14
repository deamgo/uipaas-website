package user

import (
	"errors"
	"fmt"
	"github.com/deamgo/workbench/dao/user"
	"github.com/deamgo/workbench/pkg/logger"
	"gorm.io/gorm"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"

	"github.com/deamgo/workbench/auth/jwt"
	"github.com/deamgo/workbench/context"
	"github.com/deamgo/workbench/dao"
	"github.com/deamgo/workbench/db"
	"github.com/deamgo/workbench/pkg/e"
	"github.com/deamgo/workbench/pkg/types"
	"github.com/deamgo/workbench/service/users"
)

type Resp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type UserPostReq struct {
	*users.User
}

type SignUpSuccessResp struct {
	CodeKey string
}
type VerifyReq struct {
	*users.User
	CodeKey string `json:"code_key"`
	Code    int    `json:"code"`
}

type ForgotReq struct {
	Email string `json:"email" validate:"email"`
}

type LSR struct {
	Token string
}

func SignUp(ctx context.ApplicationContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			req UserPostReq
			err error
		)
		// get The Parameters
		err = c.ShouldBind(&req.User)
		if err != nil {
			fmt.Println(err)
			return
		}
		// verify The Parameter Format
		validate := validator.New()
		err = validate.RegisterValidation("verifyPwd", verifyPwd)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, types.NewValidResponse(&Resp{
				Code: e.Failed,
				Msg:  err.Error(),
				Data: nil,
			}))
			return
		}
		err = validate.Struct(req.User)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, types.NewValidResponse(&Resp{
				Code: e.Failed,
				Msg:  "theParametersAreNotFormattedCorrectly",
				Data: nil,
			}))
			return
		}

		var u *user.DeveloperDO
		// check Whether The Mailbox Is Occupied
		u, err = ctx.UserService.UserGetByEmail(c, req.User)
		if err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				c.AbortWithStatus(http.StatusInternalServerError)
				return
			}
		}
		if u != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, types.NewValidResponse(&Resp{
				Code: e.Failed,
				Msg:  e.EmailHasBeenOccupied,
				Data: nil,
			}))
			return
		}

		//PasswordEncryption
		password, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.MinCost)
		req.Password = string(password)
		// add
		var codeHash string
		codeHash, err = ctx.UserService.UserAdd(c, req.User)
		if err != nil {
			switch err {
			case dao.DBError:
				c.AbortWithStatusJSON(http.StatusInternalServerError, types.NewErrorResponse("failed to add users"))
			default:
				c.AbortWithStatusJSON(http.StatusBadRequest, types.NewValidResponse(&Resp{
					Code: e.Failed,
					Msg:  err.Error(),
					Data: nil,
				}))
			}
			return
		}
		c.AbortWithStatusJSON(http.StatusOK, types.NewValidResponse(&Resp{
			Code: e.Success,
			Msg:  e.AddMsgSuccess,
			Data: SignUpSuccessResp{CodeKey: codeHash},
		}))
	}
}

// SignUpVerify Verify the email verification code at the time of registration
func SignUpVerify(ctx context.ApplicationContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req VerifyReq
		err := c.ShouldBind(&req)
		if err != nil {
			fmt.Println(err)
			return
		}
		isExists := db.RedisDB.HExists(req.CodeKey, "code").Val()
		if !isExists {
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
		// Modify users deactivate
		err = ctx.UserService.UserStatusModifyByEmail(c, req.User)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, types.NewValidResponse(&Resp{
				Code: e.Failed,
				Msg:  e.SignUpError,
				Data: nil,
			}))
			return
		}

		c.AbortWithStatusJSON(http.StatusOK, types.NewValidResponse(&Resp{
			Code: e.Success,
			Msg:  "Success",
			Data: nil,
		}))
	}
}

func SignIn(ctx context.ApplicationContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			req UserPostReq
		)
		err := c.ShouldBind(&req)
		if err != nil {
			fmt.Println(err)
			return
		}
		//	Check if the username exists
		findUser, err := ctx.UserService.UserGetByEmail(c, req.User)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, types.NewValidResponse(&Resp{
				Code: e.Failed,
				Msg:  err.Error(),
				Data: nil,
			}))
			logger.Error(err)
			return
		}
		if findUser == nil {
			//	doesNotExist
			c.AbortWithStatusJSON(http.StatusInternalServerError, types.NewValidResponse(&Resp{
				Code: e.Failed,
				Msg:  "The user does not exist",
				Data: nil,
			}))
			return
		}
		//	decrypt The Password
		err = bcrypt.CompareHashAndPassword([]byte(req.Password), []byte(findUser.Password))

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, types.NewValidResponse(&Resp{
				Code: e.Failed,
				Msg:  "Wrong password",
				Data: nil,
			}))
			return
		}
		//	generate A Token And Return It
		var t string
		t, _ = jwt.GenToken(findUser.ID)
		fmt.Println(t)
		c.AbortWithStatusJSON(http.StatusOK, types.NewValidResponse(&Resp{
			Code: e.Success,
			Msg:  e.LoginSuccess,
			Data: LSR{Token: t},
		}))
	}
}

func ForgotVerifySend(ctx context.ApplicationContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			req ForgotReq
			err error
		)
		// get The Parameters
		err = c.ShouldBind(&req)
		if err != nil {
			fmt.Println(err)
			return
		}
		// verify The Parameter Format
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

		var u *user.DeveloperDO
		u, err = ctx.UserService.UserGetByEmail(c, &users.User{Email: req.Email})
		if err != nil {
			logger.Error(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		if u == nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, types.NewValidResponse(&Resp{
				Code: e.Failed,
				Msg:  "The users does not exist",
				Data: nil,
			}))
			return
		}

		var codeHash string
		codeHash, err = ctx.UserService.ForgotVerifySend(c, &users.User{Email: req.Email})
		if err != nil {
			logger.Error(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		c.AbortWithStatusJSON(http.StatusOK, types.NewValidResponse(&Resp{
			Code: e.Success,
			Msg:  "The email has been sent, please pay attention to check",
			Data: codeHash,
		}))
	}
}

func ResetPassword(ctx context.ApplicationContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req VerifyReq
		err := c.ShouldBind(&req)
		if err != nil {
			fmt.Println(err)
			return
		}

		var u *user.DeveloperDO
		// check Whether The Mailbox Is Occupied
		u, err = ctx.UserService.UserGetByEmail(c, req.User)
		if err != nil {
			logger.Error(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		if u != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, types.NewValidResponse(&Resp{
				Code: e.Failed,
				Msg:  e.EmailHasBeenOccupied,
				Data: nil,
			}))
			return
		}
		pwdPattern := `^[a-zA-Z0-9]{8,20}$`
		reg, err := regexp.Compile(pwdPattern) // filter exclude chars
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, types.NewValidResponse(&Resp{
				Code: e.Failed,
				Msg:  err.Error(),
				Data: nil,
			}))
			return
		}
		match := reg.MatchString(req.Password)
		if !match {
			c.AbortWithStatusJSON(http.StatusBadRequest, types.NewValidResponse(&Resp{
				Code: e.Failed,
				Msg:  "theParametersAreNotFormattedCorrectly",
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
		// Password encryption
		password, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.MinCost)
		if err != nil {
			logger.Error(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		req.Password = string(password)
		err = ctx.UserService.UserPasswordModifyByEmail(c, req.User)
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
			Msg:  "Success",
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
