package graphQL

import (
	"fmt"
	"testing"
	"time"
)

func Test_Cards(t *testing.T) {

	q := QueryGraphQL()
	//fmt.Println(q.Repository.Projects.Edges[0].Node.Columns.Edges[0].Node.Name)
	Columns(q)

	fmt.Printf("%s\n", Cards(q))
}

func Test_Columns(t *testing.T) {

	q := QueryGraphQL()
	//fmt.Println(q.Repository.Projects.Edges[0].Node.Columns.Edges[0].Node.Name)
	Columns(q)

}

func Test_TimeStampIT(t *testing.T) {

	TimeStampIT()
}

func Test_TotalCardsToday(t *testing.T) {
	r := time.Now().Sub(lastUpdate).Seconds()

	fmt.Printf("%v\n", r)
	time.Sleep(3)
	fmt.Printf("%v\n", time.Now().Sub(lastUpdate))

}
