package changelog

import (
	"context"
	"fmt"
	"github.com/google/go-github/v50/github"
	"github.com/pixiake/auto-release/pkg/config"
	"golang.org/x/oauth2"
)

func CreateChangeLog() error {

	c, err := config.GetConfig()
	if err != nil {
		return err
	}
	fmt.Println(c)

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: ""},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	commits, _, err := client.Search.Issues(ctx, "is:pr is:open repo:kubesphere/ks-installer label:lgtm", &github.SearchOptions{
		Sort:        "",
		Order:       "",
		TextMatch:   false,
		ListOptions: github.ListOptions{},
	})
	if err != nil {
		return err
	}
	fmt.Println(*commits.Issues[0])
	return nil
}
