package firebase

import (
	"fmt"
	testing_graphql "github.com/mchirico/agil/pkg/fixtures/testing-graphql"
	"testing"
	"time"
)

func TestIdentifyCard(t *testing.T) {
	r := testing_graphql.MockResponse("created")
	result, _ := IdentifyCard(r)
	fmt.Println(result.NoteID)
	if result.NoteID != "MDExOlByb2plY3RDYXJkNDQ5MTgxMDk=" {
		t.Fatalf("Failed to find slash in %v\n", result.Note)
	}

}

func TestBuildMap(t *testing.T) {
	r := testing_graphql.MockResponse("created")
	result, _ := IdentifyCard(r)
	m, err := BuildMap(result)
	if err != nil {
		t.Fatalf("f")
	}
	action := fmt.Sprintf("%v", m["Action"])
	if action != "created" {
		t.Fatalf("Wrong action")
	}

}

func TestInsertTimeStamp(t *testing.T) {
	r := testing_graphql.MockResponse("created")
	result, _ := IdentifyCard(r)
	result.NoteID = "A_test_CARD"
	InsertCardIntoFB(result)
	resultFind, err := GetCardInfo("A_test_CARD")
	if err != nil {
		t.Fatalf("TestInsertTimeStamp")
	}
	createdAt := resultFind["CreatedAt"].(time.Time)
	t.Log(createdAt)

}
