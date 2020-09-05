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

	r.Action = "nonsense"
	_, err := IdentifyCard(r)
	if err == nil {
		t.Fatalf("Should have thrown error")
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

func buildmap() map[string]interface{} {
	m := map[string]interface{}{}
	t := time.Now().String()
	mm := map[string]string{}
	mm["a"] = "ok"
	m[t] = mm
	return m
}

func TestInsertTimeStamp(t *testing.T) {
	r := testing_graphql.MockResponse("created")
	result, _ := IdentifyCard(r)
	result.NoteID = "A_test_CARD"
	result.Updates = buildmap()
	InsertCreateCardIntoFB(result)
	resultFind, err := GetCardInfo("A_test_CARD")
	if err != nil {
		t.Fatalf("TestInsertTimeStamp")
	}
	createdAt := resultFind["CreatedAt"].(time.Time)
	t.Log(createdAt)
	m := resultFind["Updates"].(map[string]interface{})
	t.Log(m)
}

func TestUpdateTimeStamp(t *testing.T) {
	r := testing_graphql.MockResponse("moved")
	result, _ := IdentifyCard(r)
	result.NoteID = "A_test_CARD"
	result.Updates = buildmap()
	InsertUpdateCardIntoFB(result)
	resultFind, err := GetCardInfo("A_test_CARD")
	if err != nil {
		t.Fatalf("TestInsertTimeStamp")
	}
	createdAt := resultFind["CreatedAt"].(time.Time)
	t.Log(createdAt)
	m := resultFind["Updates"].(map[string]interface{})
	t.Log(m)
}

func TestUpdateTimeStamp_edited(t *testing.T) {
	r := testing_graphql.MockResponse("edited")
	result, _ := IdentifyCard(r)
	result.NoteID = "A_test_CARD"
	result.Updates = buildmap()
	InsertUpdateCardIntoFB(result)
	resultFind, err := GetCardInfo("A_test_CARD")
	if err != nil {
		t.Fatalf("TestInsertTimeStamp")
	}
	createdAt := resultFind["CreatedAt"].(time.Time)
	t.Log(createdAt)
	m := resultFind["Updates"].(map[string]interface{})
	t.Log(m)
}

func TestErrors(t *testing.T) {
	tmp := FILEBASE_TOKEN
	FILEBASE_TOKEN = "junk.json"
	r := testing_graphql.MockResponse("created")
	result, _ := IdentifyCard(r)
	result.NoteID = "A_test_CARD"
	InsertCreateCardIntoFB(result)
	_, err := GetCardInfo("A_test_CARD")
	if err == nil {
		t.Fatalf("Should have errored")
	}

	FILEBASE_TOKEN = tmp

}
