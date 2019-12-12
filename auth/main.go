package auth

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
	"runtime"
)

const configFileName = "ccn.json"

type Auth struct {
	JwtToken string `json:"token"`
}

func (a *Auth) Get() error {
	filePath := configFilePath()

	// if the config file does not exist, return an error
	_, err := os.Stat(filePath)
	if err != nil {
		return err
	}

	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	byteValue, _ := ioutil.ReadAll(f)
	json.Unmarshal(byteValue, &a)
	return nil
}

func (a *Auth) Persist() error {
	filePath := configFilePath()

	json, err := json.MarshalIndent(a, "", " ")
	if err != nil {
		return err
	}
	_, err = os.Stat(filePath)

	// create file if not exists
	if os.IsNotExist(err) {
		f, err := os.Create(filePath)
		if err != nil {
			return err
		}
		defer f.Close()
	}

	f, err := os.OpenFile(filePath, os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(json)
	if err != nil {
		return err
	}

	err = f.Sync()
	if err != nil {
		return err
	}

	return nil
}

func IsAuthenticated() bool {
	auth := Auth{}
	err := auth.Get()
	if err != nil {
		return false
	}

	return true
}

func userHomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	} else if runtime.GOOS == "linux" {
		home := os.Getenv("XDG_CONFIG_HOME")
		if home != "" {
			return home
		}
	}
	return os.Getenv("HOME")
}

func configFilePath() string {
	return path.Join(userHomeDir(), configFileName)

}
