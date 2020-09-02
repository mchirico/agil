package testing_graphql

import (
	"encoding/json"
	"fmt"
	"github.com/mchirico/agil/pkg/utils"
	"io/ioutil"
	"log"
)

func MockResponse() utils.ProjectCardUpdate {

	body := read("../fixtures/testing-graphql/" + "action-project_card-edited.json")

	projectCardUpdate := utils.ProjectCardUpdate{}
	if err := json.Unmarshal(body, &projectCardUpdate); err != nil {
		log.Fatalf("Not card action: %v\n", err)
	}
	return projectCardUpdate
}

func read(file string) []byte {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Errorf("Error read(): %v\n", err)
	}
	return b

}
