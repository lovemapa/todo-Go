package todocontroller

import (
	configuration "TODO/api/Configuration"
	constants "TODO/api/Constant"
	helper "TODO/api/Helpers"
	models "TODO/api/Models/Todo"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//CreateTodo godoc
// @Summary                    Create Todo
// @Description                Add a new Todo
// @Tags                       todo
// @Accept                     json
// @Produce                    json
// @Param                      todo body     models.Todo true "Todo Data"
// @Success                    200  {object} models.Todo
// @securityDefinitions.apiKey token
// @in                         header
// @name                       Authorization
// @Security                   JWT
// @Router                     /todo/create [post]
func CreateTodo(c *gin.Context) {

	var Todo models.Todo
	val := reflect.ValueOf(c.Keys["user_id"])

	Todo.ID = bson.NewObjectId()
	Todo.User = bson.ObjectIdHex(val.String())
	Todo.Date = time.Now()
	Todo.Status = false
	
	jsonErr:=c.BindJSON(&Todo)

	if jsonErr!=nil{
		helper.RespondWithError(c, http.StatusBadRequest,jsonErr)
		return
	 }

	if Todo.Name == "" {
		helper.RespondWithError(c, http.StatusBadRequest, "Please Provide valid name")
		return
	}


	mongoSession := configuration.ConnectDb(constants.Database)
	defer mongoSession.Close()

	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()

	getCollection := sessionCopy.DB(constants.Database).C("todo")
	err := getCollection.Insert(Todo)

	if err != nil {
		helper.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	helper.RespondWithSuccess(c, http.StatusOK, constants.TodoCreatedSuccess, Todo)

}

//GetAllTodos godoc
// @Summary                    Get Todos
// @Description                Get All TODOs
// @Tags                       todo
// @Produce                    json
// @securityDefinitions.apiKey token
// @in                         header
// @name                       Authorization
// @Security                   JWT
// @Router                     /todo/getTodos [get]
func GetTodos(c *gin.Context) {

	val := reflect.ValueOf(c.Keys["user_id"])

	resp := []bson.M{}
	mongoSession := configuration.ConnectDb(constants.Database)
	defer mongoSession.Close()

	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()

	getCollection := sessionCopy.DB(constants.Database).C("todo")
	// err := getCollection.Find(bson.M{}).All(&resp)

	err := getCollection.Find(bson.M{"user": bson.ObjectIdHex(val.String())}).All(&resp)

	if err != nil {
		helper.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	helper.RespondWithSuccess(c, http.StatusOK, constants.ListFetchedSuccess, resp)

}

//GetTodoByID godoc
// @Summary                    Get Todo
// @Description                Get a TODO by ID
// @Tags                       todo
// @Produce                    json
// @Param                      todoId path string true "Todo ID"
// // @Success      200  {object}  models.Todo
// @securityDefinitions.apiKey token
// @in                         header
// @name                       Authorization
// @Security                   JWT
// @Router                     /todo/getTodo/{todoId} [get]
func GetTodo(c *gin.Context) {
	todoID := c.Param("todoId")
	_id := bson.ObjectIdHex(todoID)
	resp := bson.M{}
	mongoSession := configuration.ConnectDb(constants.Database)
	defer mongoSession.Close()

	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()

	getCollection := sessionCopy.DB(constants.Database).C("todo")

	err := getCollection.FindId(_id).One(&resp)

	if err != nil {
		helper.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}
	helper.RespondWithSuccess(c, http.StatusOK, constants.ListFetchedSuccess, resp)

}


//UpdateTodoByID godoc
// @Summary                    Update Todo
// @Description                Update TODO by ID
// @Tags                       todo
// @Produce                    json
// @Param                      todoId path string      true "Todo ID"
// @Param                      todo   body models.Todo true "Todo Data"
// @securityDefinitions.apiKey token
// @in                         header
// @name                       Authorization
// @Security                   JWT
// @Router                     /todo/updateTodo/{todoId} [patch]
func UpdateTodo(c *gin.Context) {
	todoID := c.Param("todoId")
	var Todo models.Todo

	jsonErr:=c.BindJSON(&Todo)
	if jsonErr!=nil{
		fmt.Print(jsonErr)
		helper.RespondWithError(c, http.StatusBadRequest,jsonErr)
		return
	 }

	_id := bson.ObjectIdHex(todoID)
	if Todo.Name == "" {
		helper.RespondWithError(c, http.StatusBadRequest, "Please Provide valid name")
		return
	}
	resp := bson.M{}
	mongoSession := configuration.ConnectDb(constants.Database)
	defer mongoSession.Close()

	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()

	getCollection := sessionCopy.DB(constants.Database).C("todo")

	change := mgo.Change{
		Update:    bson.M{"$set": bson.M{"name": Todo.Name}},
		ReturnNew: true,
	}

	info, err := getCollection.Find(bson.M{"_id": _id}).Apply(change, &resp)
	// err := getCollection.UpdateId(_id, bson.M{"$set": bson.M{
	// 	"name": Todo.Name,
	// }})

	if err != nil {
		helper.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}
	fmt.Print(info)
	helper.RespondWithSuccess(c, http.StatusOK, constants.ListFetchedSuccess, resp)
}


//DeleteTodoByID godoc
// @Summary                    Delete Todo
// @Description                Delete TODO by ID
// @Tags                       todo
// @Produce                    json
// @Param                      todoId path string true "Todo ID"
// @securityDefinitions.apiKey token
// @in                         header
// @name                       Authorization
// @Security                   JWT
// @Router                     /todo/deleteTodo/{todoId} [delete]
func DeleteTodo(c *gin.Context) {
	todoID := c.Param("todoId")
	_id := bson.ObjectIdHex(todoID)
	resp := bson.M{}
	mongoSession := configuration.ConnectDb(constants.Database)
	defer mongoSession.Close()

	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()

	getCollection := sessionCopy.DB(constants.Database).C("todo")

	err := getCollection.RemoveId(_id)

	if err != nil {
		helper.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	helper.RespondWithSuccess(c, http.StatusOK, constants.DeletedSuccessfully, resp)
}
