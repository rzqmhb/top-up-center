package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rzqmhb/top-up-center/models"
	"github.com/rzqmhb/top-up-center/service"
)

type UserAPI interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	Logout(c *gin.Context)
}

type userAPI struct {
	userService service.UserService
}

func NewUserAPI(userService service.UserService) UserAPI {
	return &userAPI{userService: userService}
}

func (u *userAPI) Register(c *gin.Context) {
	var userFromBody models.User

	if err := c.ShouldBindJSON(&userFromBody); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error":"failed decode json"})
		return
	}

	if userFromBody.Email == "" || userFromBody.Name == "" || userFromBody.Password == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error":"registration data is empty"})
		return
	}

	var userToRegister = &models.User{
		Name: userFromBody.Name,
		Email: userFromBody.Email,
		Password: userFromBody.Password,
	}

	err := u.userService.Register(userToRegister)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error":"internal server error", "cause":err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message":"register success"})
}

func (u *userAPI) Login(c *gin.Context) {
	var userFromBody, userTologin models.User

	if err := c.ShouldBindJSON(&userFromBody); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error":"failed decode json"})
		return
	}

	if (userFromBody.Email == "" && userFromBody.Name == "") || userFromBody.Password == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error":"authentication data is empty"})
		return
	}

	if userFromBody.Email == "" {
		userTologin = models.User{
			Name: userFromBody.Name,
			Password: userFromBody.Password,
		}
	} else {
		userTologin = models.User{
			Email: userFromBody.Email,
			Password: userFromBody.Password,
		}
	}

	token, err := u.userService.Login(&userTologin)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error":"internal server error","cause":err.Error()})
		return
	}
	expitationTime := time.Now().Add(30 * time.Minute)
	c.SetCookie("session_token", token, expitationTime.Second(), "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "login success",})
}

func (u *userAPI) Logout(c *gin.Context) {
	username, _ := c.Get("username")
	user, err := u.userService.GetByUsername(username.(string))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error":"internal server error", "cause":err.Error()})
		return
	}
	u.userService.Logout(user)
	c.SetCookie("session_token", "", -1, "/", "", false, true)
	c.Redirect(http.StatusSeeOther, "/login")
}