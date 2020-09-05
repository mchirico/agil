package actions

import (
	"encoding/json"
	"github.com/mchirico/agil/pkg/graphQL"
	"github.com/mchirico/agil/pkg/http/handles/github"
	"github.com/mchirico/agil/pkg/tracking"
	"github.com/mchirico/agil/pkg/utils"
	"log"
	"os"
)

var secret = os.Getenv("GITHUB_WEBHOOK_SECRET")

func GithubActions() github.GithubData {

	g := github.NewGithubData(secret, func(body []byte) {
		//count += 1

		projectCardUpdate := utils.ProjectCardUpdate{}
		if err := json.Unmarshal(body, &projectCardUpdate); err != nil {
			log.Printf("Not a card action")
			return
		}

		if projectCardUpdate.Action == "created" {
			log.Printf("\nCard Created:\nNode: %v\n",
				projectCardUpdate.ProjectCard.Note)
			graphQL.OnUpdateDoCMD(projectCardUpdate, graphQL.MutateCard)
			tracking.TrackCreate(projectCardUpdate)
		}

		if projectCardUpdate.Action == "edited" {
			log.Printf("\nCard Edited\n")
			graphQL.OnUpdateDoCMD(projectCardUpdate, graphQL.MutateCard)
			tracking.TrackUpdate(projectCardUpdate)
		}

		if projectCardUpdate.Action == "moved" {
			log.Printf("\nCard Moved:\nNode: %v\n, From: %v\n",
				projectCardUpdate.ProjectCard.Note, projectCardUpdate.Changes.ColumnID.From)
			tracking.TrackUpdate(projectCardUpdate)
		}

	})
	return g
}
