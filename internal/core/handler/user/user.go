package user

import (
	"encoding/json"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/protect-we-network/server/internal/packages/config"
	"github.com/protect-we-network/server/internal/packages/logger"
	"github.com/protect-we-network/server/internal/packages/user"
)

type User struct{}

type RegisterRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (*User) Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		logger.Error(err)
		return
	}

	requestUser := &user.User{
		Email:    req.Email,
		Username: req.Username,
		Password: req.Password,
	}

	err := user.CreateUser(c.Request.Context(), requestUser)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create user"})
		logger.Error(err)
		return
	}

	queryUserParams := &user.QueryUserParams{Username: &requestUser.Username}
	queryUser, err := user.QueryUser(c.Request.Context(), queryUserParams)
	if err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		logger.Error(err)
		return
	}

	maxAge := int(time.Hour * 24 * 7)
	userString, _ := json.Marshal(queryUser)
	c.SetCookie(config.AuthenticationHeader, jwt.EncodeSegment(userString), maxAge, "/", "", false, true)

	c.JSON(200, gin.H{"message": "User created successfully"})
}

type LoginRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (*User) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		logger.Error(err)
		return
	}

	queryUserParams := &user.QueryUserParams{Username: &req.Username}
	user, err := user.QueryUser(c.Request.Context(), queryUserParams)
	if err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	if user.Password != req.Password {
		c.JSON(403, gin.H{"error": "Invalid password"})
		return
	}

	maxAge := int(time.Hour * 24 * 7)
	userString, _ := json.Marshal(user)
	token := jwt.EncodeSegment(userString)
	c.SetCookie(config.AuthenticationHeader, token, maxAge, "/", "", false, true)

	c.JSON(200, gin.H{"message": "Login successful"})
}

func (*User) Logout(c *gin.Context) {
	c.SetCookie(config.AuthenticationHeader, "", -1, "/", "", false, true)
	c.JSON(200, gin.H{"message": "Login successful"})
}

type UserManage struct{}
