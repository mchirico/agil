package firebase

import (
	"context"
	"errors"
	"github.com/mchirico/agil/pkg/graphQL"
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
	Updates         map[string]interface{}
	FirebaseUpdated bool
}

func IdentifyCard(r utils.ProjectCardUpdate) (*FBTimeStamp, error) {
	if r.Action == "created" || r.Action == "moved" || r.Action == "updated" {

		fbt := &FBTimeStamp{
			r.ProjectCard.Note,
			r.ProjectCard.NodeID,
			r.Action,
			r.ProjectCard.CreatedAt,
			r.ProjectCard.UpdatedAt,
			map[string]interface{}{},
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
	m["Updates"] = fbt.Updates
	if len(fbt.NoteID) == 0 {
		return m, errors.New("No NoteID")
	}
	return m, nil
}

func InsertCreateCardIntoFB(fbt *FBTimeStamp) {

	fb, ctx, cancel, err := CreateFB()
	if err != nil {
		return
	}
	defer cancel()

	m, err := BuildMap(fbt)
	if err != nil {
		log.Printf("InsertCreateCardIntoFB: %v\n", err)
		return
	}
	noteID := m["NoteID"].(string)
	fb.WriteMap(ctx, m, "Agil", noteID)

}

func InsertUpdateCardIntoFB(fbt *FBTimeStamp) {

	fb, ctx, cancel, err := CreateFB()
	if err != nil {
		return
	}
	defer cancel()

	m, err := BuildMap(fbt)
	if err != nil {
		log.Printf("InsertUpdateCardIntoFB: %v\n", err)
		return
	}
	noteID := m["NoteID"].(string)
	updated := m["UpdatedAt"].(string)
	action := m["Action"].(string)

	if noteID == graphQL.PROJECTCARDID {
		return
	}

	fb.WriteMap(ctx, m, "Agil", noteID)
	fb.WriteMapCol2Doc2(ctx,m,"Agil",noteID,updated,action)

}



func GetCardInfo(nodeID string) (map[string]interface{}, error) {

	fb, ctx, cancel, err := CreateFB()
	defer cancel()
	if err != nil {
		return map[string]interface{}{}, err
	}
	return fb.Find(ctx, "Agil", "NoteID", "==", nodeID)
}
