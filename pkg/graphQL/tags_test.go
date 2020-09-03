package graphQL

import (
	"fmt"
	testing_graphql "github.com/mchirico/agil/pkg/fixtures/testing-graphql"
	"github.com/mchirico/agil/pkg/qtypes"
	"strings"
	"testing"
)

func Test_FindTags(t *testing.T) {

	q := testing_graphql.MockQueryGraphQL(t)
	r := FindTags(q, `:[a-z|A-Z].*$`)
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

func Test_MarkCmds(t *testing.T) {
	r := testing_graphql.MockResponse()
	result, _ := MarkCmds(r)
	if result.NoteID != "MDExOlByb2plY3RDYXJkNDQ3NTUxMzE=" {
		t.Fatalf("Failed to find slash in %v\n", result.Note)
	}
	if string(result.Tag) != "/status :agil:testing" {
		t.Fatalf("Can't find tag. Got: %v\n", result.Tag)
	}

	r.ProjectCard.Note = `

![img](https://agil.mchirico.io/circle?text=Active&text2=%22a.i.%20bot%22&id=2342&tag=+=vbot)
/status :agil:testing
`
	result, _ = MarkCmds(r)
	if result.NoteID != "" {
		t.Fatalf("should have caught image")
	}

}

func Test_OnUpdateDoCMD(t *testing.T) {
	r := testing_graphql.MockResponse()
	_exFn := func(s0, s1 string, options ...func(*qtypes.GH4) error) {

		img := `![img](https://agil.mchirico.io/circle?text=Active&text2=%22a.i.%20bot%22&id=2342&tag=+=vbot)`
		if !strings.Contains(s0, img) {
			t.Fatalf("img not found in %v\n", s0)
		}
		fmt.Println(s0, s1)
	}

	OnUpdateDoCMD(r, _exFn)
	// OnUpdateDoCMD(r,MutateCard)

}
