package tracking

import (
	"github.com/mchirico/agil/pkg/firebase"
	"github.com/mchirico/agil/pkg/utils"
)

func TrackCreateMoved(pc utils.ProjectCardUpdate) {

	result, err := firebase.IdentifyCard(pc)
	if err != nil {
		return
	}
	firebase.InsertCardIntoFB(result)
}