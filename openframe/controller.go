package openframe

import (
	"fmt"
	"github.com/akestner/openframe/openframe/responses"
	"github.com/davecgh/go-spew/spew"
	"gopkg.in/resty.v1"
)

type Controller struct {
	Options   ControllerOptions
	Config    Config
	User      User
	Frame     Frame
	AuthToken string
}

type ControllerOptions struct {
	Username  string
	Password  string
	Frame     string
	ConfigDir string
}

func (controller *Controller) Init() error {
	fmt.Println("Controller.Init()")

	filepaths, err := DefaultFilepaths(controller.Options.ConfigDir)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	config, err := LoadConfig(filepaths.Config)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	controller.Config = config

	frame, err := LoadFrame(filepaths.Frame)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	controller.Frame = frame

	user, err := LoadUser(filepaths.User)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	controller.User = user

	authToken, userId, err := controller.authenticateUser(controller.User)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	controller.AuthToken = authToken
	controller.User.Id = userId

	pubSubUrl, err := controller.getUserConfig()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	controller.Config.PubSubUrl = pubSubUrl
	controller.Config.Save(filepaths.Config)

	spew.Dump(controller)

	return nil
}

func (controller *Controller) authenticateUser(user User) (string, string, error) {
	fmt.Println("Controller.authenticateUser()")

	response, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetBody(`
			{
				"username": "` + user.Username + `",
				"password": "` + user.Password + `"
			}
		`).
		SetResult(&responses.UserLogin{}).
		Post(controller.Config.Network.ApiBaseUrl + "/v0/users/authenticateUser")

	if err != nil {
		fmt.Println(err.Error())
		return "", "", err
	}

	userLoginResponse := response.Result().(*responses.UserLogin)

	return userLoginResponse.Id, userLoginResponse.UserId, nil
}

func (controller *Controller) getUserConfig() (string, error) {
	fmt.Println("Controller.getUserConfig()")

	response, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetAuthToken(controller.AuthToken).
		SetResult(&responses.UserConfig{}).
		Get(controller.Config.Network.ApiBaseUrl + "/v0/users/config")

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return response.Result().(*responses.UserConfig).Config.PubSubUrl, nil
}
