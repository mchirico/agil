package testing_graphql

import (
	"encoding/json"
	"fmt"
	"github.com/mchirico/agil/pkg/utils"
	"io/ioutil"
	"log"
)

func MockResponse(action string) utils.ProjectCardUpdate {

	actionMap :=map[string]string{}
	actionMap["edited"] = "action-project_card-edited.json"
	actionMap["created"] = "action-project_card-created.json"
	actionMap["moved"] = "action-project_card-moved.json"

	if _,ok := actionMap[action];!ok {
		log.Fatalf("invalid action")
	}

	body := read("../fixtures/testing-graphql/" + actionMap[action])

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
