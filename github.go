// Package github implements a simple client to consume github API (V3).
package gogithub

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
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
	apiUrl    string
	client    *http.Client
	RateLimit *RateLimit
}

type Author struct {
	Date  string `json:"date,omitempty"`
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

type Tree struct {
	Url string `json:"url,omitempty"`
	Sha string `json:"sha,omitempty"`
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
	Id   int    `json:"id"`
	Name string `json:"name"`
	Url  string `json:"url"`
}

type RateLimit struct {
	Limit     int64 `json:"limit"`
	Remaining int64 `json:"remaining"`
	Reset     int64 `json:"reset"`
}

func NewGithub() *Github {
	client := &http.Client{}

	return &Github{
		apiUrl:    apiUrl,
		client:    client,
		RateLimit: &RateLimit{},
	}
}

// Build a request and execute it within the curent htto client end returns response content
func (g *Github) buildAndExecRequest(method string, url string) ([]byte, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		panic("Error while building github request")
	}

	resp, err := g.client.Do(req)
	defer resp.Body.Close()

	if resp.StatusCode == 403 {
		// @todo handle rate limit
	}

	limit, err := strconv.ParseInt(resp.Header.Get("X-RateLimit-Limit"), 10, 64)
	if err == nil {
		g.RateLimit.Limit = limit
	}
	remaining, err := strconv.ParseInt(resp.Header.Get("X-RateLimit-Remaining"), 10, 64) 
	if err == nil {
		g.RateLimit.Remaining = remaining
	}
	reset, err := strconv.ParseInt(resp.Header.Get("X-RateLimit-Reset"), 10, 64)
	if err == nil {
		g.RateLimit.Reset = reset	
	}

	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%s", err)
	}

	return contents, err
}