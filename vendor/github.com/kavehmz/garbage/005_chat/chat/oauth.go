package chat

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/google/go-github/github"

	"golang.org/x/oauth2"
	ghoauth "golang.org/x/oauth2/github"
)

var (
	oauthConf = &oauth2.Config{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		Scopes:       []string{"user:email"},
		Endpoint:     ghoauth.Endpoint,
	}
)

type oauthURL struct {
	URL string
}

type githubUser struct {
	GithubUser string `json:"gihthub_user"`
}

//GithubCallback will handle the oauth callback and also broadcast the new user to all serssions
func GithubCallback(w http.ResponseWriter, r *http.Request) {
	state := r.FormValue("state")
	if _, ok := on.connections[state]; !ok {
		return
	}
	code := r.FormValue("code")
	token, err := oauthConf.Exchange(oauth2.NoContext, code)
	if err != nil {
		return
	}

	oauthClient := oauthConf.Client(oauth2.NoContext, token)
	client := github.NewClient(oauthClient)
	user, _, err := client.Users.Get("")
	if err != nil {
		return
	}
	userName, _ := json.Marshal(githubUser{GithubUser: *user.Login})

	c := on.connections[state]

	c.Write(userName)
	c.userID = *user.Login
	on.Lock()
	on.connections[state] = c
	on.Unlock()
	io.WriteString(w, `<script>window.close();</script>`)
}
