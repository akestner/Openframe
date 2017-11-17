package openframe

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Frame struct {
	name       string
	connected  bool
	settings   interface{}
	isPublic   bool
	extensions interface{}
	formats    interface{}
	id         string
	ownerId    string
	created    string
	modified   string
}

func (frame *Frame) Save(filePath string) error {
	bytes, err := json.Marshal(frame)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return ioutil.WriteFile(filePath, bytes, 0644)
}

func LoadFrame(filePath string) (Frame, error) {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err.Error())
		return Frame{}, err
	}

	var frame Frame
	json.Unmarshal(file, &frame)

	return frame, nil
}
