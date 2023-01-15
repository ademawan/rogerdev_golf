package middlewares

import (
	"errors"
	config "rogerdev_golf/configs"
	"rogerdev_golf/entities"

	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func GenerateToken(u entities.User) (string, error) {
	if u.UserId == "" {
		return "cannot Generate token", errors.New("user_id == null")
	}

	codes := jwt.MapClaims{
		"user_id": u.UserId,
		// "email":    u.Email,
		// "password": u.Password,
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
		"auth": true,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, codes)
	// fmt.Println(token)
	return token.SignedString([]byte(config.JWT_SECRET))
}
func ExtractTokenUserUid(e echo.Context) string {
	user := e.Get("user").(*jwt.Token) //convert to jwt token from interface
	if user.Valid {
		codes := user.Claims.(jwt.MapClaims)
		id := codes["user_id"].(string)
		return id
	}
	return ""
}

// func ExtractRoles(e echo.Context) bool {
// 	user := e.Get("user").(*jwt.Token) //convert to jwt token from interface
// 	if user.Valid {
// 		codes := user.Claims.(jwt.MapClaims)
// 		id := codes["roles"].(bool)
// 		return id
// 	}
// 	return false
// }
