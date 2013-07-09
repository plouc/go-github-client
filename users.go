package gogithub

import (
	"encoding/json"
	"fmt"
	"strings"
)

const (
	user_url = "/users/:user" // Get a single user
)

type SimpleUser struct {
	Id          int    `json:"id"`
	Login       string `json:"login"`
	Avatar_url  string `json:"avatar_url"`
	Gravatar_id string `json:"gravatar_id"`
	Url         string `json:"url"`
}

type PublicUser struct {
	SimpleUser
	Name         string `json:"name"`
	Company      string `json:"company"`
	Blog         string `json:"blog"`
	Location     string `json:"location"`
	Email        string `json:"email"`
	Hireable     bool   `json:"hireable"`
	Bio          string `json:"bio"`
	Public_repos int    `json:"public_repos"`
	Public_gists int    `json:"public_gists"`
	Followers    int    `json:"followers"`
	Following    int    `json:"public_repos"`
	Html_url     string `json:"html_url"`
	Created_at   string `json:"created_at"`
	Type         string `json:"type"`
}

type PrivateUser struct {
	PublicUser
	Total_private_repos int
	Owned_private_repos int
	Private_gists       int
	Disk_usage          int
	Collaborators       int
	Plan                *Plan
}

func (g *Github) GetUser(username string) (*PublicUser, error) {
	url := g.apiUrl + strings.Replace(user_url, ":user", username, -1)
	fmt.Println(url)
	contents, err := g.buildAndExecRequest("GET", url)

	user := new(PublicUser)
	err = json.Unmarshal(contents, &user)
	if err != nil {
		fmt.Println("%s", err)
	}

	return user, err
}
