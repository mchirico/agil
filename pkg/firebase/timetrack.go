package firebase

import (
	"context"
	"errors"
	"github.com/mchirico/agil/pkg/utils"
	"github.com/mchirico/go-firebase/pkg/gofirebase"
	"log"
	"time"
)

type FBTimeStamp struct {
	Note            string
	NoteID          string
	Action          string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	FirebaseUpdated bool
}

func IdentifyCard(r utils.ProjectCardUpdate) (*FBTimeStamp, error) {
	if r.Action == "created" || r.Action == "moved" {

		fbt := &FBTimeStamp{
			r.ProjectCard.Note,
			r.ProjectCard.NodeID,
			r.Action,
			r.ProjectCard.CreatedAt,
			r.ProjectCard.UpdatedAt,
			false,
		}
		return fbt, nil

	}
	return &FBTimeStamp{}, errors.New("Can't return data")
}

func CreateFB() (*gofirebase.FB, context.Context, func(), error) {
	credentials, err := FindCredentials()
	if err != nil {
		log.Printf("Not able to use a credentials file for testing.")

	}
	ctx, cancel := context.WithCancel(context.Background())
	// defer cancel()

	StorageBucket := "septapig.appspot.com"

	fb := &gofirebase.FB{Credentials: credentials, StorageBucket: StorageBucket}
	fb.CreateApp(ctx)
	return fb, ctx, cancel, err
}

func BuildMap(fbt *FBTimeStamp) (map[string]interface{}, error) {
	m := map[string]interface{}{}
	m["Note"] = fbt.Note
	m["NoteID"] = fbt.NoteID
	m["UpdatedAt"] = fbt.UpdatedAt
	m["CreatedAt"] = fbt.CreatedAt
	m["Action"] = fbt.Action
	if len(fbt.NoteID) == 0 {
		return m, errors.New("No NoteID")
	}
	return m, nil
}

func InsertCardIntoFB(fbt *FBTimeStamp) {

	fb, ctx, cancel, err := CreateFB()
	if err != nil {
		return
	}
	defer cancel()

	m, err := BuildMap(fbt)
	if err != nil {
		log.Printf("InsertCardIntoFB: %v\n", err)
		return
	}
	noteID := m["NoteID"].(string)
	fb.WriteMap(ctx, m, "Agil", noteID)

}

func GetCardInfo(nodeID string) (map[string]interface{}, error) {

	fb, ctx, cancel, err := CreateFB()
	defer cancel()
	if err != nil {
		return map[string]interface{}{}, err
	}
	return fb.Find(ctx, "Agil", "NoteID", "==", nodeID)
}
