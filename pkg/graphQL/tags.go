package graphQL

import (
	"fmt"
	"github.com/mchirico/agil/pkg/qtypes"
	"github.com/mchirico/agil/pkg/utils"
	"regexp"
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

func MarkCmds(r utils.ProjectCardUpdate) {
	if r.Action == "edited" || r.Action == "created" {

		fmt.Println(r.ProjectCard.Note)

	}

}
