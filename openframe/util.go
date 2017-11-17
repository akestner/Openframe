package openframe

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"path/filepath"
)

type Filepaths struct {
	Config string
	User   string
	Frame  string
}

func ConfigDir() string {
	homeDir, err := homedir.Dir()
	if err != nil {
		fmt.Println(err.Error())
	}

	return filepath.Join(homeDir, ".openframe")
}

func DefaultFilepaths(configDir string) (Filepaths, error) {
	defaultFilepaths := Filepaths{
		Config: filepath.Join(configDir, ".ofrc"),
		User:   filepath.Join(configDir, "user.json"),
		Frame:  filepath.Join(configDir, "frame.json"),
	}

	return defaultFilepaths, nil
}
