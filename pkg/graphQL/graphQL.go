package graphQL

import (
	"context"
	"fmt"
	"github.com/mchirico/agil/pkg/qtypes"
	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
	"log"
	"os"
	"time"
	_ "time/tzdata"
)

type StatData struct {
	Count   int
	Updated time.Time
	Created time.Time
}

// Ref: https://github.com/mchirico/agil/issues/2

func MutateCard(status string, projectCardID string, options ...func(*GH4) error) {

	gh4 := DefaultGH4()
	for _, op := range options {
		err := op(gh4)
		if err != nil {
			log.Fatalf("Invalid Option Setup")
		}
	}

	var m struct {
		UpdateProjectCard struct {
			ProjectCard struct {
				ID   string
				Note string
			}
		} `graphql:"updateProjectCard(input: $input)"`
	}
	s := githubv4.String(status)

	input := githubv4.UpdateProjectCardInput{
		ProjectCardID: projectCardID,
		Note:          &s,
	}

	err := gh4.client.Mutate(context.Background(), &m, input, nil)
	if err != nil {
		fmt.Printf("MutateCard err: %v\n", err)
	}
}

func DefaultGH4() *GH4 {
	gh4 := &GH4{}

	gh4.src = oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	gh4.httpClient = oauth2.NewClient(context.Background(), gh4.src)
	gh4.client = githubv4.NewClient(gh4.httpClient)
	return gh4
}

func QueryGraphQL(options ...func(*GH4) error) qtypes.Q {

	gh4 := DefaultGH4()
	for _, op := range options {
		err := op(gh4)
		if err != nil {
			log.Fatalf("Invalid Option Setup")
		}
	}

	q := qtypes.Q{}
	err := gh4.client.Query(context.Background(), &q, nil)
	if err != nil {
		fmt.Printf("\nerror:%v\n", err)
	}

	return q
}

func Columns(q qtypes.Q) {
	fmt.Printf("\n  Columns:\n")
	for i, v := range q.Repository.Projects.Edges[0].Node.Columns.Edges {
		fmt.Printf("i:%v v: %v updated: %v      id: %v\n", i,
			v.Node.Name, v.Node.UpdatedAt, v.Node.Id)
	}

}

func Cards(q qtypes.Q) string {
	s := ""
	s = fmt.Sprintf("\n  Cards:\n")
	for _, columns := range q.Repository.Projects.Edges[0].Node.Columns.Edges {
		s += fmt.Sprintf("\n_________Column: %s __________\n", columns.Node.Name)
		for i, v := range columns.Node.Cards.Edges {
			s += fmt.Sprintf("\n\n%v: %v "+
				"\nupdated: %v"+
				"\ncreated: %v"+
				"\narchived: %v"+
				"\nid: %v"+
				"\n_________________", i,
				v.Node.Note, v.Node.UpdatedAt,
				v.Node.CreatedAt, v.Node.IsArchived, v.Node.Id)
		}
	}
	return s
}

func TotalCardsToday(q qtypes.Q) map[string]StatData {

	m := map[string]StatData{}
	for _, columns := range q.Repository.Projects.Edges[0].Node.Columns.Edges {
		m[columns.Node.Name] = StatData{}

		for _, v := range columns.Node.Cards.Edges {

			updated := v.Node.UpdatedAt
			created := v.Node.CreatedAt
			blank := time.Time{}
			recent := updated
			if recent == blank {
				recent = created
			}
			now := time.Now()

			if now.Sub(recent).Hours() < 12 {
				data := m[columns.Node.Name]
				data.Count += 1
				m[columns.Node.Name] = data
			}

		}
	}
	return m
}

func UpdateEveryNtime() {
	for {
		TimeStampIT()
		time.Sleep(300 * time.Second)
	}
}

var lastUpdate = time.Now()

func TimeStampIT() {
	lastUpdate = time.Now()

	q := QueryGraphQL()

	r := TotalCardsToday(q)
	projectCardID := "MDExOlByb2plY3RDYXJkNDQ2MDU0MzU="

	loc, _ := time.LoadLocation("America/New_York")
	msg := fmt.Sprintf("\nUpdated:\n %s\n\n", time.Now().In(loc).Format("01-02 15:04:05 pm"))
	for k, v := range r {
		if k != "Metrics" {
			msg += fmt.Sprintf("%v: %v\n", k, v.Count)
		}

	}
	msg += "\n Results include archived cards." +
		"\nShowing changes in the last 12 hours."

	MutateCard(msg, projectCardID)

}
