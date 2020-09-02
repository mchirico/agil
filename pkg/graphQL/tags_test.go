package graphQL

import (
	testing_graphql "github.com/mchirico/agil/pkg/fixtures/testing-graphql"
	"testing"
)

func Test_FindTags(t *testing.T) {

	q := testing_graphql.MockQueryGraphQL(t)
	//fmt.Println(q.Repository.Projects.Edges[0].Node.Columns.Edges[0].Node.Name)
	r := FindTags(q)
	found := false
	for _,v := range r {
		if string(v.Tag) == ":agil:testing" {
			found = true
		}
	}
	if !found {
		t.Fatalf("Tag \"agil:testing\" not found\n")
	}
}
