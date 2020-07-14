package helper

import (
	modelTodo "TODO/api/Models/Todo"
	modelUser "TODO/api/Models/User"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtKey = []byte("my_secret_key")

type claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// CreateToken fore creating token
func CreateToken(_id string) (string, int64, error) {

	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = _id
	atClaims["exp"] = time.Now().Add(time.Minute*15).Unix() * 1000
	expires := atClaims["exp"].(int64)
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte("my_secret_key"))
	if err != nil {
		return "", expires, err
	}
	return token, expires, nil

}

//RespondWithError for sending errors
func RespondWithError(c *gin.Context, code int, message interface{}) {

	c.AbortWithStatusJSON(code, gin.H{
		"error":  message,
		"status": false,
	})
}

//RespondWithSuccess for success response
func RespondWithSuccess(c *gin.Context, code int, message string, data interface{}) {
	c.AbortWithStatusJSON(code, gin.H{
		"data":    data,
		"message": message,
		"status":  true,
	})
}

//ValidateSignupInput for validting user Input
func ValidateSignupInput(user modelUser.User) string {
	var errMsg string

	// Check first_name if valid.
	if user.FirstName == "" {
		errMsg = "Please enter a valid first name."
		return errMsg
	}

	// Check last_name if valid.
	if user.LastName == "" {
		errMsg = "Please enter a valid last name."
		return errMsg

	}

	// Check email_id if valid.
	if user.Email == "" {
		errMsg = "Please enter a valid email id."
		return errMsg
	}

	//check password if valid
	if user.Password == "" {
		errMsg = "Please enter a valid password"
		return errMsg
	}

	//check password if valid
	if user.Mobile == "" {
		errMsg = "Please enter a valid mobile"
		return errMsg
	}
	return ""
}

// ValidateLoginInput for login Input
func ValidateLoginInput(user modelUser.UserLogin) string {
	var errMsg string

	if user.Email == "" {
		errMsg = "Please provide email"
		return errMsg
	}

	if user.Password == "" {
		errMsg = "Please provide password"
		return errMsg
	}

	return ""
}

// ValidateTodoInput for login Input
func ValidateTodoInput(todo modelTodo.Todo) string {
	var errMsg string

	if todo.Name == "" {
		errMsg = "Please provide valid Name"
		return errMsg
	}

	return ""
}
