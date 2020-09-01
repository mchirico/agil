package testing_graphql

import (
	"fmt"
	"io/ioutil"
)

func Qtest() string {

	b, err := ioutil.ReadFile("../fixtures/testing-graphql/graphQL-server-response.json")
	if err != nil {
		fmt.Errorf("Error Qtest(): %v\n", err)
	}
	return string(b)
}
