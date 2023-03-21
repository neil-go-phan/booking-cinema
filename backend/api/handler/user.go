package handler

import (
	"booking-cinema-backend/api/presenter"
	"booking-cinema-backend/helper"
	userservice "booking-cinema-backend/services/user"
	"errors"
	"net/http"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/gin-gonic/gin"
)

var ADMIN_ROLE = "Super admin"

type UserHandler struct {
	handler  userservice.UserServices
	ESClient *elasticsearch.Client
}

func NewUserHandler(handler userservice.UserServices) *UserHandler {
	userHandler := &UserHandler{
		handler:  handler,
		// ESClient: ESClient,
	}
	return userHandler
}

func (userHandler *UserHandler) CheckLogin(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Logged"})
}

func (userHandler *UserHandler) Token(c *gin.Context) {
	username, _ := c.Get("username")
	role, _ := c.Get("role")
	accessToken, err := generateAccessToken(username.(string), role.(string))
	if err != nil {
		c.Error(errors.New(helper.ERROR_GENERATE_TOKEN_FAIL.ErrorName)) //nolint:errcheck
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Successful token reissue", "accessToken": accessToken})
}

func (userHandler *UserHandler) GetUsers(c *gin.Context) {
	role, _ := c.Get("role")
	if role != ADMIN_ROLE {
		c.Error(errors.New(helper.ERROR_NO_PERMISSION.ErrorName)) //nolint:errcheck
		c.Abort()
		return
	}
	users, err := userHandler.handler.ListUsers()
	if err != nil {
		c.Error(errors.New(helper.ERROR_BAD_REQUEST.ErrorName)) //nolint:errcheck
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, users)
}

func (userHandler *UserHandler) SignIn(c *gin.Context) {
	var inputUser presenter.User
	err := c.BindJSON(&inputUser)
	if err != nil {
		c.Error(errors.New(helper.ERROR_WHEN_PARSE_RESPONSE_BODY.ErrorName)) //nolint:errcheck
		c.Abort()
		return
	}
	err = validateUsernameAndPassword(&inputUser)
	if err != nil {
		c.Error(errors.New(helper.ERROR_INPUT_INVALID.ErrorName)) //nolint:errcheck
		c.Abort()
		return
	}
	user := newServicesUser(&inputUser)
	// verify user
	ok, err := userHandler.handler.VerifyUser(user.Username, *user)
	if err != nil || !ok {
		c.Error(errors.New(helper.ERROR_SIGNIN_INCORRECT.ErrorName)) //nolint:errcheck
		c.Abort()
		return
	}
	// generate tokens
	fullInfoUser, _ := userHandler.handler.GetUser(user.Username)
	accessToken, err := generateAccessToken(fullInfoUser.Username, fullInfoUser.RoleName)
	if err != nil {
		c.Error(errors.New(helper.ERROR_GENERATE_TOKEN_FAIL.ErrorName)) //nolint:errcheck
		c.Abort()
		return
	}
	refreshToken, err := GenerateRefreshToken(fullInfoUser.Username, fullInfoUser.RoleName)
	if err != nil {
		c.Error(errors.New(helper.ERROR_GENERATE_TOKEN_FAIL.ErrorName)) //nolint:errcheck
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Sign in success", "accessToken": accessToken, "refreshToken": refreshToken})
}

func (userHandler *UserHandler) SignUp(c *gin.Context) {
	var inputUser presenter.User
	err := c.BindJSON(&inputUser)
	if err != nil {
		c.Error(errors.New(helper.ERROR_WHEN_PARSE_RESPONSE_BODY.ErrorName)) //nolint:errcheck
		c.Abort()
		return
	}
	err = validateSignUp(&inputUser)
	if err != nil {
		c.Error(errors.New(helper.ERROR_INPUT_INVALID.ErrorName)) //nolint:errcheck
		c.Abort()
		return
	}

	user := newServicesUser(&inputUser)

	checkUser, err := userHandler.handler.GetUser(user.Username)
	if err != nil {
		c.Error(errors.New(helper.ERROR_BAD_REQUEST.ErrorName)) //nolint:errcheck
		c.Abort()
		return
	}
	if checkUser.Username != "" {
		c.Error(errors.New(helper.ERROR_USERNAME_TAKEN.ErrorName)) //nolint:errcheck
		c.Abort()
		return
	}
	_, err = userHandler.handler.CreateUser(user)
	if err != nil {
		c.Error(errors.New(helper.ERROR_SERVER.ErrorName)) //nolint:errcheck
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Sign up success"})
}

func (userHandler *UserHandler) DeleteUser(c *gin.Context) {
	role, _ := c.Get("role") 
	if role != ADMIN_ROLE {
		c.Error(errors.New(helper.ERROR_NO_PERMISSION.ErrorName)) //nolint:errcheck
		c.Abort()
		return
	}
	username := c.Param("username")
	err := userHandler.handler.DeleteUser(username)
	if err != nil {
		c.Error(errors.New(helper.ERROR_DELETE_FAIL.ErrorName)) //nolint:errcheck
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Delete user success"})
}



// func (userHandler *UserHandler) UpdateUser(c *gin.Context) {
// 	roleFromRequest, _ := c.Get("role")
// 	if roleFromRequest != ADMIN_ROLE {
// 		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No permission granted"})
// 		return
// 	}

// 	var inputUser presenter.User
// 	c.BindJSON(&inputUser)
// 	user := newServicesUser(&inputUser)
// 	user.Role = inputUser.RoleName
// 	err := userHandler.handler.UpdateUser(user)

// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Fail to update"})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Update user success"})
// }

// func (userHandler *UserHandler) SearchUser(c *gin.Context) {
// 	roleFromRequest, _ := c.Get("role")
// 	if roleFromRequest != ADMIN_ROLE {
// 		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No permission granted"})
// 		return
// 	}

// 	var query string
// 	if query, _ = c.GetQuery("q"); query == "" {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "no search query present"})
// 		return
// 	}
// 	//body := `{"query" : { "match_all" : {}" }}`
// 	body := fmt.Sprintf(`{"query": {"multi_match": {"query": "%s", "fields": ["full_name"]}}}`, query)
// 	log.Println(userHandler.ESClient)
// 	res, err := userHandler.ESClient.Search(

// 		userHandler.ESClient.Search.WithIndex("users"),
// 		userHandler.ESClient.Search.WithBody(strings.NewReader(body)),
// 		userHandler.ESClient.Search.WithPretty(),
// 		)
// 	var r map[string]interface{}
// 	if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 			return
// 		}
// 		defer res.Body.Close()
// 		if res.IsError() {

// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "yw"}) // ERROR_REQUEST_TO_ELASTRIC_SEARCH
// 			return
// 		}
		
// 		if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}) // ERROR_WHEN_PARSE_RESPONSE_BODY
// 			return
// 		}
// 		c.JSON(http.StatusOK, gin.H{"data": r["hits"]})
	
// }
