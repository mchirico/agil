package graphQL

import (
	"fmt"
	testing_graphql "github.com/mchirico/agil/pkg/fixtures/testing-graphql"
	"os"
	"strings"
	"testing"
)

func Test_Cards(t *testing.T) {

	q := testing_graphql.MockQueryGraphQL(t)
	//fmt.Println(q.Repository.Projects.Edges[0].Node.Columns.Edges[0].Node.Name)
	Columns(q)

	fmt.Printf("%s\n", Cards(q))
}

func Test_Columns(t *testing.T) {

	q := testing_graphql.MockQueryGraphQL(t)
	Columns(q)

}

func Test_TimeStampIT(t *testing.T) {
	// FIXME: To run with Mock
	if len(os.Getenv("GITHUB_TOKEN")) != 0 {
		TimeStampIT()
	}
}

// Very simple test
func Test_TotalCardsToday(t *testing.T) {
	q := testing_graphql.MockQueryGraphQL(t)
	r := TotalCardsToday(q)
	if len(r) != 5 {
		t.Fatalf("TotalCardsToday not working")
	}
}

func Test_QueryGraphQL(t *testing.T) {
	q := QueryGraphQL(testing_graphql.MockOpt)
	s := fmt.Sprintf("%v", q)
	if !strings.Contains(s, "istio and service") {
		t.Fatalf("Not picking of Mock for QueryGraphQL")
	}

}
