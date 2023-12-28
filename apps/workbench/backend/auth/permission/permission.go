package auth

import (
	"encoding/json"
	"errors"
	"log"
	"strings"

	"github.com/deamgo/workbench/auth/jwt"
	"github.com/deamgo/workbench/db"
	"github.com/deamgo/workbench/util"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

type DevWorkspace struct {
	DeveloperID string `json:"developer_id"`
	WorkspaceID string `json:"workspace_id"`
	Role        string `json:"role"`
}

func (dw *DevWorkspace) TableName() string {
	return "workspace_developer_relation"
}

func AuthenticationPermissions(c *gin.Context) {
	// Determine whether the user role has permission
	pathRoles := GetCurrentPathRole(c)
	devRole := GetCurrentDeveloperRole(c)
	for i := range pathRoles {
		if pathRoles[i] == devRole {
			c.Next()
			return
		}
	}
	c.JSON(403, gin.H{
		"message": "No access",
	})
	c.Abort()
}

func GetCurrentPathRole(c *gin.Context) []string {
	path := c.Request.URL.Path
	path = parsePath(path)
	method := c.Request.Method
	roles := util.GetRoles()
	for i := range roles {
		role := roles[i]
		if role.Path == path && role.Method == method {
			return util.TransStrToArr(role.Role)
		}
	}
	return nil
}

func GetCurrentDeveloperRole(c *gin.Context) string {
	userId, _ := jwt.ExtractIDFromToken(c.Request.Header.Get("Authorization"))
	workSpaceId := c.Param("workspace_id")

	// Example Query the role of a user
	var devWorkspace *DevWorkspace
	// Check cache first
	devWorkspaceJSON, err := db.RedisDB.Get(userId).Result()
	if err != nil {
		if !errors.Is(err, redis.Nil) {
			log.Println(err)
		}
	}
	if err = json.Unmarshal([]byte(devWorkspaceJSON), &devWorkspace); err != nil {
		log.Println(err)
	}

	if devWorkspace == nil {
		db.DB.Model(&devWorkspace).
			Where("developer_id= ? and workspace_id= ?", userId, workSpaceId).
			Find(&devWorkspace)
		devWorkspaceJSON, err := json.Marshal(devWorkspace)
		if err != nil {
			log.Println(err)
		}
		db.RedisDB.Set(userId, devWorkspaceJSON, 0)
	}
	return devWorkspace.Role
}

func parsePath(path string) string {
	segments := strings.Split(path, "/")
	for i, segment := range segments {
		if segment == "workspace" && i+1 < len(segments) {
			segments[i+1] = ":workspace_id"
		}
	}
	return strings.Join(segments, "/")
}
