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
	Changes         int
	CreatedAt       time.Time
	UpdatedAt       time.Time
	Updates         map[string]interface{}
	FirebaseUpdated bool
	Archived        bool
}

func IdentifyCard(r utils.ProjectCardUpdate) (*FBTimeStamp, error) {
	if r.Action == "created" || r.Action == "moved" || r.Action == "edited" {

		fbt := &FBTimeStamp{
			r.ProjectCard.Note,
			r.ProjectCard.NodeID,
			r.Action,
			r.Changes.ColumnID.From,
			r.ProjectCard.CreatedAt,
			r.ProjectCard.UpdatedAt,
			map[string]interface{}{},
			false,
			r.ProjectCard.Archived,
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
	m["Changes"] = fbt.Changes
	m["Updates"] = fbt.Updates
	m["Archived"] = fbt.Archived
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
	updated := m["UpdatedAt"].(time.Time)
	action := m["Action"].(string)

	if noteID == graphQL.PROJECTCARDID {
		return
	}

	toupdate, err := fb.ReadMap(ctx, "Agil", noteID)
	um := toupdate.Data()

	uUpdates := um["Updates"].(map[string]interface{})
	timeStamp := updated.String()
	uUpdates[timeStamp] = action

	um["Note"] = m["Note"]
	um["UpdatedAt"] = m["UpdatedAt"]
	um["Action"] = m["Action"]
	um["Changes"] = m["Changes"]
	um["Archived"] = m["Archived"]

	fb.WriteMap(ctx, um, "Agil", noteID)
	fb.WriteMapCol2Doc2(ctx, m, "Agil", noteID, updated.String(), action)

}

func GetCardInfo(nodeID string) (map[string]interface{}, error) {

	fb, ctx, cancel, err := CreateFB()
	defer cancel()
	if err != nil {
		return map[string]interface{}{}, err
	}
	return fb.Find(ctx, "Agil", "NoteID", "==", nodeID)
}
