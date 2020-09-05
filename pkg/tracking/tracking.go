package tracking

import (
	"github.com/mchirico/agil/pkg/firebase"
	"github.com/mchirico/agil/pkg/utils"
)

func TrackCreate(pc utils.ProjectCardUpdate) {

	result, err := firebase.IdentifyCard(pc)
	if err != nil {
		return
	}
	firebase.InsertCreateCardIntoFB(result)
}

func TrackUpdate(pc utils.ProjectCardUpdate) {

	result, err := firebase.IdentifyCard(pc)
	if err != nil {
		return
	}
	firebase.InsertUpdateCardIntoFB(result)
}
