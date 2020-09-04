package firebase

import (
	"errors"
	"io/ioutil"
)

func LocateFile(locations []string) (string, error) {

	for _,file := range locations {
		_,err :=ioutil.ReadFile(file)
		if err == nil {
			return file, nil
		}
	}
	return "", errors.New("Fire not found")
}

func FindCredentials() (string,error) {
	directories :=[]string{"/credentials",
		"../../credentials",
		"../credentials",
		"../../../credentials",
		"/etc/credentials",
	}
	locations := []string{}
	file := "septapig-firebase-adminsdk.json"
	for _,d := range directories {
		locations = append(locations,d+"/"+file)
	}


	loc, err := LocateFile(locations)
	return loc, err
}
