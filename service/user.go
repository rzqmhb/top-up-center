package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/rzqmhb/top-up-center/models"
	repo "github.com/rzqmhb/top-up-center/repository"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService interface {
	Login(user *models.User) (string, error)
	Register(user *models.User) error
	Logout(user *models.User) error
	GetByUsername(username string) (*models.User, error)
}

type userService struct {
	userRepo repo.UserRepository
	sessionRepo repo.SessionRepository
}

func NewUserService(userRepo repo.UserRepository, sessionRepo repo.SessionRepository) UserService {
	return &userService{
		userRepo: userRepo,
		sessionRepo: sessionRepo,
	}
}

func (u *userService) Login(user *models.User) (string, error) {
	var (
		dbUser *models.User
		err error
	)

	if len(user.Name) == 0 {
		dbUser, err = u.userRepo.GetByEmail(user.Email)
	} else {
		dbUser, err = u.userRepo.GetByUsername(user.Name)
	}
	if err != nil {
		return "", fmt.Errorf("error while getting user data: %s", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return "", errors.New("password incorrect")
		}
		return "", fmt.Errorf("error while comparing password: %s", err)
	}
	
	var expirationTime time.Time = time.Now().Add(time.Minute * 30)
	var newClaims *models.Claims = &models.Claims{
		Username: dbUser.Name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	var token *jwt.Token = jwt.NewWithClaims(jwt.SigningMethodHS256, newClaims)
	tokenString, err := token.SignedString(models.JWTKey)
	if err != nil {
		return "", fmt.Errorf("error while signing token: %s", err)
	}

	var newSession = &models.Session{
		UserName: dbUser.Name,
		Token: tokenString,
		Expiry: expirationTime,
	}

	_, err = u.sessionRepo.GetByUsername(dbUser.Name)
	if err != nil {
		err = u.sessionRepo.Store(newSession)
	} else {
		err = u.sessionRepo.Update(user.Name, newSession)
	}

	return tokenString, nil
}

func (u *userService) Register(user *models.User) error {
	user.CreatedAt, user.UpdatedAT = time.Now(), time.Now()
	passwordInBytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	if err != nil {
		return fmt.Errorf("error while hashing password: %s", err)
	}
	user.Password = string(passwordInBytes)

	err = u.userRepo.Store(user)
	if err != nil {
		if err == gorm.ErrDuplicatedKey {
			return errors.New("email or username already exists")
		}
		return err
	}
	return nil
}

func (u *userService) Logout(user *models.User) error {
	currentSession, err := u.sessionRepo.GetByUsername(user.Name)
	if err != nil {
		return err
	}
	err = u.sessionRepo.Delete(currentSession.Token)
	if err != nil {
		return err
	}
	return nil
}

func (u *userService) GetByUsername(username string) (*models.User, error) {
	return u.userRepo.GetByUsername(username)
}