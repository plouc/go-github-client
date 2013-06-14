// Package github implements a simple client to consume github API (V3).
package gogithub

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const (

	apiUrl = "https://api.github.com"

	// Users
	public_user_url  = "/users/:user"
	current_user_url = "/user"
	users_url        = "/users"

	// Repositories
	repos      = "/repositories"      // List all public repositories
	repos_user = "/users/:user/repos" // List public repositories for the specified user.
	repos_org  = "/orgs/:org/repos"   // List repositories for the specified org.

	// Events
	events_url_public                = "/events"                             // List public events
	events_url_repo                  = "/repos/:owner/:repo/events"          // List repository events
	events_url_repo_issues           = "/repos/:owner/:repo/issues/events"   // List issue events for a repository
	events_url_network_public        = "/networks/:owner/:repo/events"       // List public events for a network of repositories
	events_url_user_received         = "/users/:user/received_events"        // List events that a user has received
	events_url_user_received_public  = "/users/:user/received_events/public" // List public events that a user has received
	events_url_user_performed        = "/users/:user/events"                 // List events performed by a user
	events_url_user_performed_public = "/users/:user/events/public"          // List public events performed by a user
	events_url_org                   = "/users/:user/events/orgs/:org"       // List events for an organization
	events_url_org_public            = "/orgs/:org/events"                   // List public events for an organization

	dateLayout = "2006-01-02T15:04:05Z"
)

type Github struct {
	Client *http.Client
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
	Sha       string
	Url       string
	Author    *Author
	Committer *Author
	Message   string
	Tree      *Tree
	Parents   []*Tree
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

type Repo struct {
	SimpleRepo
	Full_name      string
	Description    string
	Owner          *SimpleUser
	Private        bool
	Fork           bool
	Html_url       string
	Clone_url      string
	Git_url        string
	Ssh_url        string
	Svn_url        string
	Mirror_url     string
	Homepage       string
	Language       string
	Forks          int
	Forks_count    int
	Watchers       int
	Watchers_count int
	Size           int
	Master_branch  string
	Open_issues    int
	Pushed_at      string
	Created_at     string
	Updated_at     string
}

type Event struct {
	Id         string
	Type       string
	Public     bool
	Repo       *SimpleRepo
	Actor      *SimpleUser
	Org        *SimpleUser
	CreatedAt  time.Time
	Created_at string
	// @todo add payload field
}

func NewGithub() *Github {

	client := &http.Client{}

	return &Github{
		Client: client,
	}
}

func (g *Github) buildAndExecRequest(method string, url string) []byte {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		panic("Error while building github request")
	}

	resp, err := g.Client.Do(req)
	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%s", err)
	}
	//fmt.Println(string(contents))

	return contents
}

// Get repositories from the given url
func (g *Github) GetRepos(url string) []*Repo {

	contents := g.buildAndExecRequest("GET", url)

	var repos []*Repo
	err := json.Unmarshal(contents, &repos)
	if err != nil {
		fmt.Println("%s", err)
	}

	return repos
}

// List all public repositories
func (g *Github) Repos(user string) []*Repo {
	url := apiUrl + repos

	return g.GetRepos(url)
}

// List public repositories for the specified user.
func (g *Github) UserRepos(user string) []*Repo {
	url := apiUrl + strings.Replace(repos_user, ":user", user, -1)

	return g.GetRepos(url)
}

// List repositories for the specified org.
func (g *Github) OrgRepos(org string) []*Repo {
	url := apiUrl + strings.Replace(repos_org, ":org", org, -1)

	return g.GetRepos(url)
}

// Get events from the given url
func (g *Github) GetEvents(url string) []*Event {
	contents := g.buildAndExecRequest("GET", url)

	var events []*Event
	err := json.Unmarshal(contents, &events)
	if err != nil {
		fmt.Println("%s", err)
	}

	for _, event := range events {
		t, _ := time.Parse(dateLayout, event.Created_at)
    	event.CreatedAt = t
	}

	return events
}

// List public events
func (g *Github) Events() []*Event {
	url := apiUrl + events_url_public

	return g.GetEvents(url)
}

// List repository events
func (g *Github) RepoEvents(owner string, repo string) []*Event {
	url := strings.Replace(events_url_repo, ":owner", owner, -1)
	url = apiUrl + strings.Replace(url, ":repo", repo, -1)

	return g.GetEvents(url)
}

// List issue events for a repository
func (g *Github) RepoIssuesEvents(owner string, repo string) []*Event {
	url := strings.Replace(events_url_repo_issues, ":owner", owner, -1)
	url = apiUrl + strings.Replace(url, ":repo", repo, -1)

	return g.GetEvents(url)
}

// List public events for a network of repositories
func (g *Github) RepoNetworkEvents(owner string, repo string) []*Event {
	url := strings.Replace(events_url_network_public, ":owner", owner, -1)
	url = apiUrl + strings.Replace(url, ":repo", repo, -1)

	return g.GetEvents(url)
}

// List events that a user has received
func (g *Github) UserReceivedEvents(user string) []*Event {
	url := apiUrl + strings.Replace(events_url_user_received, ":user", user, -1)

	return g.GetEvents(url)
}

// List public events that a user has received
func (g *Github) UserReceivedPublicEvents(user string) []*Event {
	url := apiUrl + strings.Replace(events_url_user_received_public, ":user", user, -1)

	return g.GetEvents(url)
}

// List events performed by a user
func (g *Github) UserPerformedEvents(user string) []*Event {
	url := apiUrl + strings.Replace(events_url_user_performed, ":user", user, -1)

	return g.GetEvents(url)
}

// List public events performed by a user
func (g *Github) UserPerformedPublicEvents(user string) []*Event {
	url := apiUrl + strings.Replace(events_url_user_performed_public, ":user", user, -1)

	return g.GetEvents(url)
}

// List events for an organization
func (g *Github) OrgEvents(user string, org string) []*Event {
	url := strings.Replace(events_url_org, ":user", user, -1)
	url = apiUrl + strings.Replace(url, ":org", org, -1)

	return g.GetEvents(url)
}

// List public events for an organization
func (g *Github) OrgPublicEvents(org string) []*Event {
	url := apiUrl + strings.Replace(events_url_org_public, ":org", org, -1)

	return g.GetEvents(url)
}
