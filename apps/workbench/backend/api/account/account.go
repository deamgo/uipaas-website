package account

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"

	"github.com/deamgo/workbench/auth/jwt"
	"github.com/deamgo/workbench/context"
	"github.com/deamgo/workbench/dao"
	developerDO "github.com/deamgo/workbench/dao/developer"
	"github.com/deamgo/workbench/db"
	"github.com/deamgo/workbench/pkg/e"
	"github.com/deamgo/workbench/pkg/encryption"
	"github.com/deamgo/workbench/pkg/logger"
	"github.com/deamgo/workbench/pkg/types"
	"github.com/deamgo/workbench/service/developer"
	workspace "github.com/deamgo/workbench/service/workspace"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type Resp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type UserPostReq struct {
	*developer.Developer
}

type SignUpSuccessResp struct {
	CodeKey string `json:"code_key"`
}
type SendMailResp struct {
	CodeKey string `json:"code_key"`
}
type VerifyReq struct {
	*developer.Developer
	CodeKey string `json:"code_key" validate:"required"`
	Code    int    `json:"code" validate:"required"`
}

type ForgotReq struct {
	Email string `json:"email" validate:"email"`
}

type LSR struct {
	Token string `json:"token"`
}

func SignUp(ctx context.ApplicationContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			req UserPostReq
			err error
		)
		// get The Parameters
		err = c.ShouldBind(&req.Developer)
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
		err = validate.Struct(req.Developer)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, types.NewValidResponse(&Resp{
				Code: e.Failed,
				Msg:  "The parameters do not match",
				Data: nil,
			}))
			return
		}

		var u *developerDO.DeveloperDO
		// check Whether The Mailbox Is Occupied
		u, err = ctx.UserService.DeveloperGetByEmail(c, req.Developer)
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
		pwd := encryption.EncryptPwd(req.Password)
		req.Password = pwd
		// add
		var codeHash string
		codeHash, err = ctx.UserService.DeveloperAdd(c, req.Developer)
		if err != nil {
			switch err {
			case dao.DBError:
				c.AbortWithStatusJSON(http.StatusInternalServerError, types.NewErrorResponse("failed to add developer"))
			default:
				c.AbortWithStatusJSON(http.StatusBadRequest, types.NewValidResponse(&Resp{
					Code: e.Failed,
					Msg:  err.Error(),
					Data: nil,
				}))
			}
			return
		}
		//ctx.WorkspaceService.WorkspaceCreate()
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
		// Delete the verification code that has already been used
		db.RedisDB.HDel(req.CodeKey, "code")
		// Modify developer deactivate
		err = ctx.UserService.DeveloperStatusModifyByEmail(c, req.Developer)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, types.NewValidResponse(&Resp{
				Code: e.Failed,
				Msg:  e.SignUpError,
				Data: nil,
			}))
			return
		}

		dpl, _ := ctx.UserService.DeveloperGetByEmail(c, req.Developer)

		//PasswordEncryption
		password := encryption.EncryptPwd(req.Password)

		if password != dpl.Password {
			c.AbortWithStatusJSON(http.StatusBadRequest, types.NewValidResponse(&Resp{
				Code: e.Failed,
				Msg:  "The account or password is incorrect",
				Data: nil,
			}))
			return
		}
		id, err := strconv.ParseInt(dpl.ID, 10, 64)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
		}
		// create default workspace
		var workspace = &workspace.Workspace{
			Name:        dpl.Username + "'s Workspace",
			Logo:        "http://121.41.78.218:80/images/9ccb00dbd1.jpg",
			Label:       "default workspace",
			Description: "default workspace",
			CreatedBy:   uint64(id),
			UpdatedBy:   uint64(id),
		}
		_, err = ctx.WorkspaceService.WorkspaceCreate(c, workspace)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, types.NewValidResponse(&Resp{
				Code: e.Failed,
				Msg:  "The workspace failed to be created",
				Data: nil,
			}))
			return
		}
		//	generate A Token And Return It
		var t string
		t, _ = jwt.GenToken(dpl.ID)
		c.AbortWithStatusJSON(http.StatusOK, types.NewValidResponse(&Resp{
			Code: e.Success,
			Msg:  "Registration Successful, signing in...",
			Data: LSR{Token: t},
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
		findUser, err := ctx.UserService.DeveloperGetByEmail(c, req.Developer)
		if err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				c.AbortWithStatusJSON(http.StatusInternalServerError, types.NewValidResponse(&Resp{
					Code: e.Failed,
					Msg:  err.Error(),
					Data: nil,
				}))
				logger.Error(err)
				return
			}
		}
		if findUser == nil {
			//	doesNotExist
			c.AbortWithStatusJSON(http.StatusInternalServerError, types.NewValidResponse(&Resp{
				Code: e.Failed,
				Msg:  "The account or password is incorrect",
				Data: nil,
			}))
			return
		}
		//	decrypt The Password

		//PasswordEncryption
		password := encryption.EncryptPwd(req.Password)

		if password != findUser.Password {
			c.AbortWithStatusJSON(http.StatusBadRequest, types.NewValidResponse(&Resp{
				Code: e.Failed,
				Msg:  "The account or password is incorrect",
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

func Logout() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		// remove token
		jwt.RevokeToken(authHeader)
		c.AbortWithStatusJSON(http.StatusOK, types.NewValidResponse(&Resp{
			Code: e.Success,
			Msg:  "Logout successfully",
			Data: nil,
		}))
	}
}
func SendForgotVerify(ctx context.ApplicationContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		SendMail(c, ctx)
	}
}

func SendMail(c *gin.Context, ctx context.ApplicationContext) {
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

	var u *developerDO.DeveloperDO
	u, err = ctx.UserService.DeveloperGetByEmail(c, &developer.Developer{Email: req.Email})
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusInternalServerError, types.NewValidResponse(&Resp{
				Code: e.Failed,
				Msg:  err.Error(),
				Data: nil,
			}))
			logger.Error(err)
			return
		}
	}

	if u == nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, types.NewValidResponse(&Resp{
			Code: e.Failed,
			Msg:  "The developer does not exist",
			Data: nil,
		}))
		return
	}

	var codeHash string
	codeHash, err = ctx.UserService.ForgotVerifySend(c, &developer.Developer{Email: req.Email})
	if err != nil {
		logger.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.AbortWithStatusJSON(http.StatusOK, types.NewValidResponse(&Resp{
		Code: e.Success,
		Msg:  "The email has been sent, please pay attention to check",
		Data: SendMailResp{CodeKey: codeHash},
	}))
}

func ResetPassword(ctx context.ApplicationContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req VerifyReq
		err := c.ShouldBind(&req)
		if err != nil {
			fmt.Println(err)
			return
		}
		validate := validator.New()
		err = validate.Var(req.CodeKey, "required")
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
		}
		err = validate.Var(req.Code, "required")
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
		}

		var u *developerDO.DeveloperDO
		// check Whether The Mailbox Is exist
		u, err = ctx.UserService.DeveloperGetByEmail(c, req.Developer)
		if err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				c.AbortWithStatusJSON(http.StatusInternalServerError, types.NewValidResponse(&Resp{
					Code: e.Failed,
					Msg:  err.Error(),
					Data: nil,
				}))
				logger.Error(err)
			}
		}
		if u == nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, types.NewValidResponse(&Resp{
				Code: e.Failed,
				Msg:  "The developer does not exist",
				Data: nil,
			}))
			return
		}
		// Verify that the mailbox has not been maliciously altered
		emailHashStr := developer.GetEmailHashStr(u.Email)
		if req.CodeKey != emailHashStr {
			c.AbortWithStatusJSON(http.StatusBadRequest, types.NewValidResponse(&Resp{
				Code: e.Failed,
				Msg:  "The email address is incorrect",
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
				Msg:  "The password is not formatted correctly",
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

		// Delete the verification code that has already been used
		db.RedisDB.HDel(req.CodeKey, "code")
		// Password encryption
		password := encryption.EncryptPwd(req.Password)
		if err != nil {
			logger.Error(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		req.Password = string(password)
		err = ctx.UserService.DeveloperPasswordModifyByEmail(c, req.Developer)
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
			Msg:  "Password successfully reset, redirecting to login page.",
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
