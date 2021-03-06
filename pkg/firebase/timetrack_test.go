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

func Test_Note_in_FB(t *testing.T) {
	r := testing_graphql.MockResponse("moved")
	result, _ := IdentifyCard(r)
	note := `Note:

With returns
`
	result.Note = note
	result.NoteID = "A_test_CARD"
	result.Updates = buildmap()
	InsertUpdateCardIntoFB(result)
	resultFind, err := GetCardInfo("A_test_CARD")
	if err != nil {
		t.Fatalf("Test_Note_in_FB: %v\n", err)
	}
	x := resultFind["Note"].(string)
	if x != note {
		t.Fatalf("Got: ->%v<-\nExpected: ->%v<-\n", x, note)
	}

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

type ff struct{}

func (f *ff) Data() map[string]interface{} {
	m := map[string]interface{}{}
	m["UpdatedAt"] = time.Now()
	m["Action"] = "spud"
	m["Updates"] = map[string]interface{}{}
	return m
}
func Test_updateMainCard(t *testing.T) {

	f := &ff{}
	m := map[string]interface{}{}
	now := time.Now()
	m["UpdatedAt"] = now
	m["Action"] = "no spud"
	m["Note"] = "mNote"
	m["Changes"] = map[string]interface{}{}
	m["Archived"] = "snow"
	m["Updates"] = map[string]interface{}{}

	um := _updateMainCard(f, m)
	if um["Action"] != "no spud" {
		t.Fatalf("_updateMainCard didn't update")
	}

}
