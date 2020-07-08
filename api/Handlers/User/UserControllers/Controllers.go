package usercontrollers

import (
	configuration "TODO/api/Configuration"
	CONSTANTS "TODO/api/Constant"
	helper "TODO/api/Helpers"
	models "TODO/api/Models/User"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//Register advertisment
func Register(c *gin.Context) {

	var userData models.User

	// userErr := json.NewDecoder(c.Request.Body).Decode(&userData)
	params := struct {
		ProfilePhoto string `json:"ProfilePhoto,default=pawan"`
	}{}

	c.Bind(&params)

	email := c.PostForm("email")
	firstName := c.PostForm("firstName")
	lastName := c.PostForm("lastName")
	password := c.PostForm("password")
	mobile := c.PostForm("mobile")

	userData.Email = email
	userData.FirstName = firstName
	userData.LastName = lastName
	userData.Password = password
	userData.Mobile = mobile

	// if userErr != nil {

	// 	helper.RespondWithError(c, http.StatusBadRequest, CONSTANTS.EmptyReqBody)
	// 	return
	// }
	validateInputErr := helper.ValidateSignupInput(userData)
	if validateInputErr != "" {
		helper.RespondWithError(c, http.StatusBadRequest, validateInputErr)
		return
	}
	hashedPassword, hashError := bcrypt.GenerateFromPassword([]byte(userData.Password), bcrypt.DefaultCost)
	if hashError != nil {
		helper.RespondWithError(c, http.StatusBadRequest, hashError)
		return
	}
	userData.Password = string(hashedPassword)
	userData.Date = time.Now().Round(time.Millisecond).UnixNano() / (int64(time.Millisecond) / int64(time.Nanosecond))

	mongoSession := configuration.ConnectDb(CONSTANTS.Database)
	defer mongoSession.Close()

	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()

	getCollection := sessionCopy.DB(CONSTANTS.Database).C("user")

	index := mgo.Index{
		Key:        []string{"email"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}

	indexerr := getCollection.EnsureIndex(index)

	if indexerr != nil {
		helper.RespondWithError(c, http.StatusBadRequest, indexerr)
		return
	}

	userData.ID = bson.NewObjectId()

	token, _ := helper.CreateToken(userData.ID.Hex())
	userData.Token = token

	// file, FileErr := c.FormFile("file")
	// if file == nil {
	// 	helper.RespondWithError(c, http.StatusBadRequest, CONSTANTS.FileMissing)
	// 	return
	// }
	// fmt.Print(file.Filename)

	// if FileErr != nil {

	// 	helper.RespondWithError(c, http.StatusBadRequest, FileErr)
	// 	return

	// }

	// fileName := "Static/" + time.Now().Format("2006-01-02 15:04:05.000000") + file.Filename

	// userData.ProfilePhoto = "/" + fileName
	// uploadErr := c.SaveUploadedFile(file, fileName)
	// if uploadErr != nil {
	// 	helper.RespondWithError(c, http.StatusBadRequest, uploadErr)
	// 	return
	// }
	err := getCollection.Insert(userData)

	if err != nil {
		if mgo.IsDup(err) == true {

			helper.RespondWithError(c, http.StatusBadRequest, CONSTANTS.AccountAlreadyExists)
			return
		}
		helper.RespondWithError(c, http.StatusBadRequest, err)
		return
	}

	helper.RespondWithSuccess(c, http.StatusOK, CONSTANTS.CreatedSuccssfully, userData)

}

// Login for advertiser login
func Login(c *gin.Context) {
	var Login models.UserLogin
	var userData models.User
	userErr := json.NewDecoder(c.Request.Body).Decode(&Login)
	LoginErr := helper.ValidateLoginInput(Login)
	if LoginErr != "" {
		helper.RespondWithError(c, http.StatusBadRequest, LoginErr)
		return
	}
	if userErr != nil {

		helper.RespondWithError(c, http.StatusBadRequest, userErr)
		return
	}

	mongoSession := configuration.ConnectDb(CONSTANTS.Database)
	defer mongoSession.Close()

	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()

	getCollection := sessionCopy.DB(CONSTANTS.Database).C("user")

	err := getCollection.Find(bson.M{"email": Login.Email}).One(&userData)
	if err != nil {
		helper.RespondWithError(c, http.StatusBadRequest, CONSTANTS.AccountNotExists)
		return
	}

	PassErr := bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(Login.Password))
	if PassErr != nil {
		helper.RespondWithError(c, http.StatusBadRequest, CONSTANTS.IncorrectPassword)
		return
	}
	token, _ := helper.CreateToken(userData.ID.Hex())
	userData.Token = token
	err = getCollection.UpdateId(bson.ObjectIdHex(userData.ID.Hex()), userData)
	helper.RespondWithSuccess(c, http.StatusOK, CONSTANTS.LoggedInSuccess, userData)

}
