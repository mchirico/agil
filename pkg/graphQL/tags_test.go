package graphQL

import (
	testing_graphql "github.com/mchirico/agil/pkg/fixtures/testing-graphql"
	"testing"
)

func Test_FindTags(t *testing.T) {

	q := testing_graphql.MockQueryGraphQL(t)
	//fmt.Println(q.Repository.Projects.Edges[0].Node.Columns.Edges[0].Node.Name)
	FindTags(q)

}
