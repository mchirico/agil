package graphQL

import (
	testing_graphql "github.com/mchirico/agil/pkg/fixtures/testing-graphql"
	"testing"
)

func Test_FindTags(t *testing.T) {

	q := testing_graphql.MockQueryGraphQL(t)
	r := FindTags(q)
	found := false
	for _, v := range r {
		if string(v.Tag) == ":agil:testing" {
			found = true
		}
	}
	if !found {
		t.Fatalf("Tag \"agil:testing\" not found\n")
	}
}

func Test_MockEdited(t *testing.T) {
	r := testing_graphql.MockResponse()
	if r.Action != "edited" {
		t.Fatalf("Not picking up correct json mock?")
	}

}
