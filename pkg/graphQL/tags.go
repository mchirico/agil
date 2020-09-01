package graphQL

import (
	"fmt"
	"github.com/mchirico/agil/pkg/qtypes"
	"regexp"
)

func FindTags(q qtypes.Q) {
	re := regexp.MustCompile(`:[a-z|A-Z].*$`)
	fmt.Printf("\n  Columns:\n")
	for i, column := range q.Repository.Projects.Edges[0].Node.Columns.Edges {
		if 1==2 {
			fmt.Printf("i:%v v: %v updated: %v      id: %v\n", i,
				column.Node.Name, column.Node.UpdatedAt, column.Node.Id)
		}
		for _, card := range column.Node.Cards.Edges {
			//fmt.Printf("%v\n",card.Node.Note)
			s := card.Node.Note
			ok := re.Find([]byte(s))
			if ok != nil {
				fmt.Printf("Found: %s\n",ok)
			}
		}
	}

}
