package graphQL

import (
	"fmt"
	"github.com/mchirico/agil/pkg/qtypes"
	"github.com/mchirico/agil/pkg/utils"
	"log"
	"regexp"
	"strings"
	"time"
)

type Tags struct {
	Note       string
	ColID      string
	NodeID     string
	Tag        []byte
	UpdatedAt  time.Time
	CreatedAt  time.Time
	URL        string
	IsArchived bool
}

// regex := `:[a-z|A-Z].*$`
func FindTags(q qtypes.Q, regex string) []Tags {

	tags := []Tags{}
	re := regexp.MustCompile(regex)

	for _, column := range q.Repository.Projects.Edges[0].Node.Columns.Edges {

		for _, card := range column.Node.Cards.Edges {

			s := card.Node.Note
			ok := re.Find([]byte(s))
			if ok != nil {
				tag := Tags{}
				tag.Note = card.Node.Note
				tag.NodeID = card.Node.Id
				tag.URL = card.Node.Url
				tag.UpdatedAt = card.Node.UpdatedAt
				tag.CreatedAt = card.Node.CreatedAt
				tag.IsArchived = card.Node.IsArchived
				tag.ColID = column.Node.Id
				tag.Tag = ok

				tags = append(tags, tag)

			}
		}
	}
	return tags
}

type NoteToUpdate struct {
	Note   string
	NoteID string
	Tag    []byte
}

func MarkCmds(r utils.ProjectCardUpdate) (NoteToUpdate, error) {
	if r.Action == "edited" || r.Action == "created" {

		log.Printf("MarkCmds:"+
			"\nNote: %v\n", r.ProjectCard.Note)
		_done := func(s string) []byte {
			regex := `tag=\+=vbot`
			re := regexp.MustCompile(regex)
			ok := re.Find([]byte(s))
			fmt.Println(string(ok))
			return ok
		}
		regex := `/[a-z|A-Z].* :[a-z|A-Z].*$`
		re := regexp.MustCompile(regex)
		ok := re.Find([]byte(r.ProjectCard.Note))

		log.Printf("MarkCmds:"+
			"\nNote: %v"+
			"\nok: %v\n", r.ProjectCard.Note, ok)

		if ok != nil {

			if _done(r.ProjectCard.Note) == nil {
				ntu := NoteToUpdate{
					Note:   r.ProjectCard.Note,
					NoteID: r.ProjectCard.NodeID,
					Tag:    ok,
				}
				return ntu, nil
			}
		}
	}
	return NoteToUpdate{}, nil
}

// TODO: Pull this out, when done
func OnUpdateDoCMD(r utils.ProjectCardUpdate, fn func(string, string)) {

	result, err := MarkCmds(r)
	if err != nil {
		return
	}

	log.Printf("OnUpdateDoCMD:"+
		"\nNote: %v"+
		"\n", r.ProjectCard.Note)

	if len(result.NoteID) > 10 {

		img := `![img](https://agil.mchirico.io/circle?text=Active&text2=%22a.i.%20bot%22&id=2342&tag=+=vbot)
`
		newstring := fmt.Sprintf("%s\n%s", img, result.Tag)

		s := strings.ReplaceAll(result.Note, string(result.Tag), newstring)

		fn(s, result.NoteID)

		// MutateCard(s, result.NoteID)

	}
}
