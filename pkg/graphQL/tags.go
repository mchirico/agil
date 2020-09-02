package graphQL

import (
	"github.com/mchirico/agil/pkg/qtypes"
	"regexp"
	"time"
)

type Tags struct {
	Note     string
	ColID    string
	NodeID    string
	Tag       []byte
	UpdatedAt time.Time
	CreatedAt time.Time
	URL       string
	IsArchived bool
}

func FindTags(q qtypes.Q) []Tags {
	tags := []Tags{}
	re := regexp.MustCompile(`:[a-z|A-Z].*$`)

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


				tags = append(tags,tag)

			}
		}
	}
	return tags
}
