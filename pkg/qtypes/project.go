package qtypes

import (
	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
	"net/http"
	"time"
)

type GH4 struct {
	Src        oauth2.TokenSource
	HttpClient *http.Client
	Client     *githubv4.Client
}

type Q struct {
	Repository struct {
		Projects struct {
			TotalCount int
			Edges      []struct {
				Node struct {
					Name    string
					Id      string
					Columns struct {
						Edges []struct {
							Node struct {
								Id         string
								Name       string
								DatabaseId int32
								CreatedAt  time.Time
								UpdatedAt  time.Time

								Cards struct {
									Edges []struct {
										Node struct {
											Id         string
											DatabaseID int32
											Note       string
											CreatedAt  time.Time
											UpdatedAt  time.Time
											Url        string
											IsArchived bool
											Creator    struct {
												Login string
											}
										}
									}
								} `graphql:"cards(last: 100)"`
							}
						}

						TotalCount int
					} `graphql:"columns(last: 10)"`
				}
			}
		} `graphql:"projects(last: 1)"`
	} `graphql:"repository(owner: \"mchirico\", name: \"agil\")"`
}
