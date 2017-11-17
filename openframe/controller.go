package openframe

import (
	"fmt"
)

type Controller struct {
}

type ControllerOptions struct {
	Username  string
	Password  string
	Frame     string
	ConfigDir string
}

func (controller *Controller) Init(options ControllerOptions) error {
	fmt.Println("Controller.init()")
	fmt.Println("options:", options)

	filepaths, err := DefaultFilepaths(options.ConfigDir)
	if err != nil {
		fmt.Println(err.Error())
	}

	user, err := LoadUser(filepaths.User)
	if err != nil {
		fmt.Println(err.Error())
	}
	controller.Login(user)

	//config := LoadConfig(filepaths.Config)
	//frame := LoadFrame(filepaths.Frame)

	return nil
}

func (controller *Controller) Login(user User) error {
	fmt.Println("Controller.Login()")
	fmt.Println("user:", user)

	return nil
}
