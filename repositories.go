package gogithub

import (
	"encoding/json"
	"fmt"
	"strings"
)

const (
	repos      = "/repositories"       // List all public repositories
	repos_user = "/users/:user/repos"  // List public repositories for the specified user.
	repos_org  = "/orgs/:org/repos"    // List repositories for the specified org.
	repo_url   = "/repos/:owner/:name" // Get a repository by its owner/name
)

// Repository struct
type Repo struct {
	SimpleRepo
	Full_name      string      `json:"full_name,omitempty"`
	Description    string      `json:"description,omitempty"`
	Owner          *SimpleUser `json:"owner,omitempty"`
	Private        bool        `json:"private,omitempty"`
	Fork           bool        `json:"fork,omitempty"`
	Html_url       string      `json:"html_url,omitempty"`
	Clone_url      string      `json:"clone_url,omitempty"`
	Git_url        string      `json:"git_url,omitempty"`
	Ssh_url        string      `json:"ssh_url,omitempty"`
	Svn_url        string      `json:"svn_url,omitempty"`
	Mirror_url     string      `json:"mirror_url,omitempty"`
	Homepage       string      `json:"homepage,omitempty"`
	Language       string      `json:"language,omitempty"`
	Forks          int         `json:"forks,omitempty"`
	Forks_count    int         `json:"forks_count,omitempty"`
	Watchers       int         `json:"watchers,omitempty"`
	Watchers_count int         `json:"watchers_count,omitempty"`
	Size           int         `json:"size,omitempty"`
	Master_branch  string      `json:"master_branch,omitempty"`
	Open_issues    int         `json:"open_issues,omitempty"`
	Pushed_at      string      `json:"pushed_at,omitempty"`
	Created_at     string      `json:"created_at,omitempty"`
	Updated_at     string      `json:"updated_at,omitempty"`
}

// Get repositories from the given url
func (g *Github) getRepos(url string) ([]*Repo, error) {

	contents, err := g.buildAndExecRequest("GET", url)

	var repos []*Repo
	err = json.Unmarshal(contents, &repos)
	if err != nil {
		fmt.Println("%s", err)
	}

	return repos, err
}

// Get a repository
func (g *Github) Repo(owner string, name string) (*Repo, error) {
	url := g.apiUrl + strings.Replace(repo_url, ":owner", owner, -1)
	url = strings.Replace(url, ":name", name, -1)

	contents, err := g.buildAndExecRequest("GET", url)

	var repo *Repo
	err = json.Unmarshal(contents, &repo)
	if err != nil {
		fmt.Println("%s", err)
	}

	return repo, err
}

// List all public repositories
//
//     repos := Github.GetRepos()
func (g *Github) Repos() ([]*Repo, error) {
	url := g.apiUrl + repos

	return g.getRepos(url)
}

// List public repositories for the specified user.
//
//     repos := Github.UserRepos("plouc")
func (g *Github) UserRepos(user string) ([]*Repo, error) {
	url := g.apiUrl + strings.Replace(repos_user, ":user", user, -1)

	return g.getRepos(url)
}

// List repositories for the specified org.
//
//     repos := Github.OrgRepos("ekino")
func (g *Github) OrgRepos(org string) ([]*Repo, error) {
	url := g.apiUrl + strings.Replace(repos_org, ":org", org, -1)

	return g.getRepos(url)
}