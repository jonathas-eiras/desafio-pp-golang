package users

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/jonathas-eiras/desafio-picpay-golang/pkg/config"
	"github.com/jonathas-eiras/desafio-picpay-golang/pkg/utils/path"
	"github.com/jonathas-eiras/desafio-picpay-golang/services/users"
	"github.com/stretchr/testify/suite"
	"math/rand"
	"net/http"
	"strconv"
	"testing"
)

type UsersSuite struct {
	suite.Suite
	application config.Application
	baseUrl     string
	Client      *resty.Client
	token       string
	id          int
}

func (u *UsersSuite) SetupSuite() {

	u.application = config.GetApplication()
	u.baseUrl = u.application.Service.Users.BasePath
	u.token = u.application.Service.Users.Token
	u.Client = resty.New()

	rand.Seed(1000)
	u.id = rand.Intn(500)

}

func TestUsersSuite(t *testing.T) {
	suite.Run(t, new(UsersSuite))
}

func (u *UsersSuite) TestUsers() {

	userRequest := users.NewUser()

	u.T().Run("1- Create a new user", func(t *testing.T) {
		//create a new user
		response, err := u.Client.R().
			SetHeader("Authorization", u.token).
			SetBody(userRequest).
			Post(u.baseUrl + path.PathUsers())
		if err != nil {
			fmt.Println("Error on Request:", err)
		}

		output := users.UserResponse{}

		err = json.Unmarshal(response.Body(), &output)
		if err != nil {
			fmt.Println("Error unmarshal:", err)
		}

		fmt.Println(output.Data)
		u.Equal(http.StatusCreated, response.StatusCode())
		u.Equal(userRequest.Name, output.Data.Name)
		u.Equal(userRequest.Email, output.Data.Email)
		u.Equal(userRequest.Status, output.Data.Status)
		u.Equal(userRequest.Gender, output.Data.Gender)
		u.Equal(http.StatusCreated, output.Code)

		u.id = output.Data.ID

	})
	u.T().Run("2- Check the users list", func(t *testing.T) {

		//check the users list
		pages := u.returnPages()
		for page := 1; page <= pages; page++ {
			response, err := u.Client.R().
				SetHeader("Authorization", u.token).
				SetQueryParam("page", strconv.Itoa(page)).
				Get(u.baseUrl + path.PathUsers())
			if err != nil {
				fmt.Println("Error on Request:", err)
			}

			output := users.UsersListResponse{}
			err = json.Unmarshal(response.Body(), &output)
			if err != nil {
				fmt.Println("Error unmarshal:", err)
			}

			for _, user := range output.Data {
				if user.ID == u.id {
					u.Equal(user.Name, userRequest.Name)
					u.Equal(user.Email, userRequest.Email)
					u.Equal(user.Gender, userRequest.Gender)
					u.Equal(user.Status, userRequest.Status)
					fmt.Println(user.ID)
					break
				}
			}

		}

	})
	u.T().Run("3- Update a user", func(t *testing.T) {
		//update a new user
		response, err := u.Client.R().
			SetHeader("Authorization", u.token).
			SetBody(userRequest).
			Put(u.baseUrl + path.PathUserById(strconv.Itoa(u.id)))
		if err != nil {
			fmt.Println("Error on Request:", err)
		}

		output := users.UserResponse{}

		err = json.Unmarshal(response.Body(), &output)
		if err != nil {
			fmt.Println("Error unmarshal:", err)
		}

		fmt.Println(output.Data)
		u.Equal(http.StatusOK, response.StatusCode())
		u.Equal(userRequest.Name, output.Data.Name)
	})
	u.T().Run("4- Get user updated by id", func(t *testing.T) {
		//get the user by id
		response, err := u.Client.R().
			SetHeader("Authorization", u.token).
			Get(u.baseUrl + path.PathUserById(strconv.Itoa(u.id)))
		if err != nil {
			fmt.Println("Error on Request:", err)
		}

		output := users.UserResponse{}
		err = json.Unmarshal(response.Body(), &output)
		if err != nil {
			fmt.Println("Error unmarshal:", err)
		}

		u.Equal(http.StatusOK, response.StatusCode())
		u.Equal(output.Data.Name, userRequest.Name)
		u.Equal(output.Data.ID, u.id)
	})
	u.T().Run("5- Delete a user", func(t *testing.T) {
		//delete a new user
		response, err := u.Client.R().
			SetHeader("Authorization", u.token).
			Delete(u.baseUrl + path.PathUserById(strconv.Itoa(u.id)))
		if err != nil {
			fmt.Println("Error on Request:", err)
		}

		fmt.Println(u.id)
		u.Equal(http.StatusOK, response.StatusCode())
	})
	u.T().Run("6- Get user deleted by id", func(t *testing.T) {
		//get the user by id
		response, err := u.Client.R().
			SetHeader("Authorization", u.token).
			Get(u.baseUrl + path.PathUserById(strconv.Itoa(u.id)))
		if err != nil {
			fmt.Println("Error on Request:", err)
		}

		u.Equal(http.StatusNotFound, response.StatusCode())

	})

}

func (u *UsersSuite) returnPages() int {
	response, err := u.Client.R().
		SetHeader("Authorization", u.token).
		Get(u.baseUrl + path.PathUsers())
	if err != nil {
		fmt.Println("Error on Request:", err)
	}
	output := users.UsersListResponse{}
	err = json.Unmarshal(response.Body(), &output)
	if err != nil {
		fmt.Println("Error unmarshal:", err)
	}
	return output.Meta.Pagination.Pages
}
