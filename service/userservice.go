package service

import (
	"MyGram/models"
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type Userinterf interface {
	Register(user *models.User) (*models.User, error)
	Login(user *models.User, tempPassword string) (*models.User, error)
	Update(user *models.User) (*models.User, error)
	GetToken(id uint, email string, password string) string
	CheckToken(compareToken string, id uint, email string, password string) error
	VerifToken(tempToken string) (float64, error)
}

type UserService struct {
}

func NewUserService() Userinterf {
	return &UserService{}
}

func (servicuser *UserService) Register(user *models.User) (*models.User, error) {
	if user.Username == "" {
		return nil, errors.New("Username cannot be empty")
	}
	if user.Email == "" {
		return nil, errors.New("Email cannot be empty")
	}
	if len(user.Password) < 6 {
		return nil, errors.New("password must be minimum 6 characters")
	}
	if user.Age < 8 {
		return nil, errors.New("age must be greater than 8")
	}
	return user, nil
}

func (servicuser *UserService) Login(user *models.User, tempPassword string) (*models.User, error) {
	if user.Email == "" {
		return nil, errors.New("email cannot be empty")
	}
	password := []byte(tempPassword)
	//check password salah
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), password); err != nil {
		return nil, errors.New("password salah")
	}
	return user, nil
}

func (servicuser *UserService) Update(user *models.User) (*models.User, error) {
	if user.Email == "" {
		return nil, errors.New("email cannot be empty")
	}
	if user.Username == "" {
		return nil, errors.New("username cannot be empty")
	}
	return user, nil
}

func (servicuser *UserService) GetToken(id uint, email string, password string) string {
	claims := jwt.MapClaims{
		"id":    id,
		"email": email,
	}

	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, _ := parseToken.SignedString([]byte(password))

	return signedToken
}

func (servicuser *UserService) CheckToken(compareToken string, id uint, email string, password string) error {
	token := servicuser.GetToken(id, email, password)
	if compareToken == token {
		fmt.Println("berhasil")
		return nil
	} else {
		fmt.Println("tidak berhasil")
		return errors.New("username cannot be empty")
	}
	//compare
}

func (servicuser *UserService) VerifToken(TempToken string) (float64, error) {
	tokenString := TempToken
	claimstoken := jwt.MapClaims{}
	var Verifykey []byte
	token, _ := jwt.ParseWithClaims(tokenString, claimstoken, func(token *jwt.Token) (interface{}, error) {
		return Verifykey, nil
	})

	payload := token.Claims.(jwt.MapClaims)
	id := payload["id"].(float64)
	if !isIntegral(id) {
		return 0, errors.New("invalid token")
	}
	return id, nil
}
func isIntegral(val float64) bool {
	return val == float64(int(val))
}
