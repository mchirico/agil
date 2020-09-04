package firebase

import (
	"github.com/mchirico/agil/pkg/utils"
	"time"
	"errors"
)

type FBTimeStamp struct {
	Note string
	NoteID string
	Action string
	CreatedAt time.Time
	UpdatedAt time.Time

}


func IdentifyCard(r utils.ProjectCardUpdate) (FBTimeStamp, error) {
	if r.Action == "created" || r.Action == "moved" {


		fbt := FBTimeStamp{
			r.ProjectCard.Note,
			r.ProjectCard.NodeID,
			r.Action,
			r.ProjectCard.CreatedAt,
			r.ProjectCard.UpdatedAt,
		}
		return fbt, nil

	}
	return FBTimeStamp{}, errors.New("Can't return data")
}

