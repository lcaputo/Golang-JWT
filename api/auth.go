package api

import (
	"elprogramador.co/go/echo/database"
	"elprogramador.co/go/echo/helpers"
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"strings"
	"time"

	"elprogramador.co/go/echo/models"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func Register(c echo.Context) error {
	u := new(models.User)

	if err := helpers.BindAndValidate(c, u); err != nil {
		return err
	}

	var exist models.User

	database.DB.First(&exist, "email = ?", u.Email)

	fmt.Println(u.ID)
	if exist.ID == 0 {
		hash, _ := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
		u.Password = string(hash)
		database.DB.Create(u)
	} else {
		return SendResponse(c, http.StatusBadRequest, "User already exists", nil)
	}

	//token, _ := user.GenerateToken()

	var cookie http.Cookie

	cookie.Name = "token"
	//cookie.Value = token
	cookie.Expires = time.Now().Add(7 * 24 * time.Hour)

	c.SetCookie(&cookie)

	return c.JSON(http.StatusOK, u)
}

// función para generar un token JWT
func generateToken(user *models.User) (string, error) {
	// definir los claims del token
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
	})

	// generar el token firmado con una clave secreta
	token, err := claims.SignedString([]byte("mi_secreto"))
	if err != nil {
		return "", err
	}
	return token, nil
}

// función para verificar un email y password
func verifyUser(email, password string) (*models.User, error) {
	// buscar al usuario en la base de datos
	user, err := getUserByEmail(email)
	if err != nil {
		return nil, err
	}

	// verificar la contraseña
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, err
	}

	return user, nil
}

// función para obtener un usuario por email
func getUserByEmail(email string) (*models.User, error) {
	// consulta a la base de datos para buscar al usuario
	// ...
	u := new(models.User)
	// TODO FIX
	//database.DB.First(&u, "email = ?", email)

	if u.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}
	// simular un usuario encontrado
	user := &models.User{
		Email:    email,
		Password: u.Password, // contraseña: secret
	}
	return user, nil
}

func Login(c echo.Context) error {
	u := new(models.User)
	if err := helpers.BindAndValidate(c, u); err != nil {
		return SendResponse(c, http.StatusBadRequest, "Invalid form", err)
	}
	// verificar el usuario
	//user, err := verifyUser(u.Email, u.Password)
	//if err != nil {
	//	return err
	//}
	database.DB.First(&u, "email = ?", u.Email)

	keyBase64 := os.Getenv("PRIVATE_KEY")
	keyData, _ := base64.StdEncoding.DecodeString(keyBase64)
	key, _ := jwt.ParseRSAPrivateKeyFromPEM(keyData)

	token2 := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"id":        u.ID,
		"email":     u.Email,
		"iss":       "elprogramador.co",
		"timestamp": time.Now().Unix(),
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	})
	signed, _ := token2.SignedString(key)
	return SendResponse(c, http.StatusOK, "Login successful", signed)
}

func SignedLoginToken(u *models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": u.Email,
	})

	return token.SignedString([]byte("mi_secreto"))
}

func TokenInfo(c echo.Context) error {
	// obtener el token del request
	tokenString := c.Request().Header.Get("Authorization")
	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

	fmt.Println(tokenString)
	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Get the key from the public key file
		keyBase64 := os.Getenv("PUBLIC_KEY")
		keyData, _ := base64.StdEncoding.DecodeString(keyBase64)
		publicKey, _ := jwt.ParseRSAPublicKeyFromPEM(keyData)

		return publicKey, nil
	})

	// Check if there was an error parsing the token
	if err != nil {
		fmt.Printf("error parsing token: %v\n", err)
		return err
	}

	// Check if the token is valid
	if !token.Valid {
		fmt.Printf("token is not valid\n")
		return err
	}

	// Get the claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		fmt.Printf("error getting claims\n")
		return err
	}

	return SendResponse(c, http.StatusOK, "Token info", claims)

}

// // Login : Login Router
// func (AuthRouter) Login(c echo.Context) error {
// 	type RequestBody struct {
// 		Username string `json:"username" validate:"required"`
// 		Password string `json:"password" validate:"required"`
// 	}

// 	var body RequestBody

// 	if err := c.Bind(&body); err != nil {
// 		return c.NoContent(http.StatusNotFound)
// 	}
// 	if err := c.Validate(&body); err != nil {
// 		return err
// 	}

// 	db, _ := c.Get("db").(*gorm.DB)

// 	var user models.User

// 	if err := db.Where("username = ?", body.Username).First(&user).Error; err != nil {
// 		return c.NoContent(http.StatusConflict)
// 	}

// 	if bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(body.Password)) != nil {
// 		return c.NoContent(http.StatusInternalServerError)
// 	}

// 	token, _ := user.GenerateToken()

// 	var cookie http.Cookie

// 	cookie.Name = "token"
// 	cookie.Value = token
// 	cookie.Expires = time.Now().Add(7 * 24 * time.Hour)

// 	c.SetCookie(&cookie)

// 	return c.JSON(http.StatusOK, echo.Map{
// 		"token": token,
// 		"user":  user,
// 	})
// }

// // Logout : Logout Router
// func (AuthRouter) Logout(c echo.Context) error {
// 	tokenCookie, _ := c.Get("tokenCookie").(*http.Cookie)

// 	tokenCookie.Value = ""
// 	tokenCookie.Expires = time.Unix(0, 0)

// 	c.SetCookie(tokenCookie)

// 	return c.NoContent(200)
// }
