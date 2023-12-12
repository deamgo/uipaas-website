package user

import (
	"fmt"
	"log"
	"math/rand"
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
	"github.com/deamgo/workbench/service/user"
)

type Resp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type UserPostReq struct {
	*user.User
}

type SignUpSuccessResp struct {
	CodeKey string
}
type VerifyReq struct {
	*user.User
	CodeKey string
	Code    int
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
		err = validate.Struct(req.User)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, types.NewValidResponse(&Resp{
				Code: e.Failed,
				Msg:  "theParametersAreNotFormattedCorrectly",
				Data: nil,
			}))
			return
		}

		var u *user.User
		// verifyWhether The InvitationCode Exists
		u, err = ctx.UserService.UserGetByInvitationCode(c, req.User)
		if u == nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, types.NewValidResponse(&Resp{
				Code: e.Failed,
				Msg:  e.InvalidInvitationCode,
				Data: nil,
			}))
			return
		}
		// check Whether The Mailbox Is Occupied
		u, err = ctx.UserService.UserGetByEmail(c, req.User)
		if u != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, types.NewValidResponse(&Resp{
				Code: e.Failed,
				Msg:  e.EmailHasBeenOccupied,
				Data: nil,
			}))
			return
		}

		// generate An Invitation Code
		req.InvitationCode = GenerateInviteCode()
		//PasswordEncryption
		password, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.MinCost)
		req.Password = string(password)
		// add
		var codeHash string
		codeHash, err = ctx.UserService.UserAdd(c, req.User)
		if err != nil {
			switch err {
			case dao.DBError:
				c.AbortWithStatusJSON(http.StatusInternalServerError, types.NewErrorResponse("failed to add user"))
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

// verify
func SignUpVerify(ctx context.ApplicationContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req VerifyReq
		err := c.ShouldBind(&req)
		if err != nil {
			fmt.Println(err)
			return
		}
		isExists := db.RedisDB.HExists(req.CodeKey, "code").Val()
		if isExists {
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
		// Modify user deactivate
		err = ctx.UserService.UserDeactivateModifyByEmail(c, req.User)
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
		if findUser == nil {
			//	doesNotExist
			log.Fatal(err)
		}
		//	decrypt The Password
		_ = bcrypt.CompareHashAndPassword([]byte(req.Password), []byte(findUser.Password))
		//	generate A Token And Return It
		var t string
		t, err = jwt.GenToken(findUser.Username)
		fmt.Println(t)
		c.AbortWithStatusJSON(http.StatusOK, types.NewValidResponse(&Resp{
			Code: e.Success,
			Msg:  e.LoginSuccess,
			Data: LSR{Token: t},
		}))
	}
}

func GenerateInviteCode() string {
	// define The Character Set
	alphabets := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")

	// generate Random Strings
	inviteCode := make([]rune, 6)
	for i := 0; i < len(inviteCode); i++ {
		inviteCode[i] = alphabets[rand.Intn(len(alphabets))]
	}

	// return The Invitation Code
	return string(inviteCode)
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
