package firebase

import (
	"fmt"
	testing_graphql "github.com/mchirico/agil/pkg/fixtures/testing-graphql"
	"testing"
)

func TestIdentifyCard(t *testing.T) {
	r := testing_graphql.MockResponse("created")
	result, _ := IdentifyCard(r)
	fmt.Println(result.NoteID)
	if result.NoteID != "MDExOlByb2plY3RDYXJkNDQ5MTgxMDk=" {
		t.Fatalf("Failed to find slash in %v\n", result.Note)
	}

}
