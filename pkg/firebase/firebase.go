package firebase

import (
	"errors"
	"io/ioutil"
)

var FILEBASE_TOKEN = "septapig-firebase-adminsdk.json"

func LocateFile(locations []string) (string, error) {

	for _, file := range locations {
		_, err := ioutil.ReadFile(file)
		if err == nil {
			return file, nil
		}
	}
	return "", errors.New("Fire not found")
}

func FindCredentials(dirs ...string) (string, error) {
	directories := []string{"/credentials",
		"../../credentials",
		"../credentials",
		"../../../credentials",
		"/etc/credentials",
	}

	for i, dir := range dirs {
		if len(directories) > i {
			directories[i] = dir
		} else {
			directories = append(directories, dir)
		}
	}

	locations := []string{}
	file := FILEBASE_TOKEN
	for _, d := range directories {
		locations = append(locations, d+"/"+file)
	}

	loc, err := LocateFile(locations)
	return loc, err
}
