// Package github implements a simple client to consume github API (V3).
package gogithub

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	apiUrl = "https://api.github.com"

	// Users
	public_user_url  = "/users/:user"
	current_user_url = "/user"
	users_url        = "/users"

	dateLayout = "2006-01-02T15:04:05Z"
)

type Github struct {
	apiUrl string
	client *http.Client
}

type Author struct {
	Date  string
	Name  string
	Email string
}

type Tree struct {
	Url string
	Sha string
}

type Commit struct {
	Sha       string  // The SHA of the commit.
	Url       string  // Points to the commit API resource.
	Author    *Author // The git author of the commit.
	Committer *Author
	Message   string  // The commit message.
	Tree      *Tree
	Parents   []*Tree
	Distinct  bool    // (Available in PushEvent) Whether this commit is distinct from any that have been pushed before.
}

type Object struct {
	Type string
	Sha  string
	Url  string
}

type Tag struct {
	Tag     string
	Sha     string
	Url     string
	Message string
	Tagger  *Author
	Object  *Object
}

type SimpleUser struct {
	Login       string
	Id          int
	Avatar_url  string
	Gravatar_id string
	Url         string
}

type PublicUser struct {
	SimpleUser
	Name         string
	Company      string
	Blog         string
	Location     string
	Email        string
	Hireable     bool
	Bio          string
	Public_repos int
	Public_gists int
	Followers    int
	Following    int
	Html_url     string
	Created_at   string
	Type         string
}

type Plan struct {
	Name          string
	Space         int
	Collaborators int
	Private_repos int
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

type SimpleRepo struct {
	Id   int
	Name string
	Url  string
}

func NewGithub() *Github {
	client := &http.Client{}

	return &Github{
		apiUrl: apiUrl,
		client: client,
	}
}

// Build a request and execute it within the curent htto client end returns response content
func (g *Github) buildAndExecRequest(method string, url string) []byte {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		panic("Error while building github request")
	}

	resp, err := g.client.Do(req)
	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%s", err)
	}

	return contents
}