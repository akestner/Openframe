package openframe

import (
	"fmt"
)

type Controller struct {
}

type ControllerConfig struct {
	Username string
	Password string
	Frame    string
}

func (controller *Controller) Init(controllerConfig ControllerConfig) error {
	fmt.Println("Controller.init()")
	fmt.Println("controllerConfig:", controllerConfig)
	return nil
}
