package gogithub

import (
	"encoding/json"
	"fmt"
	"strings"
)

const (
	repos      = "/repositories"      // List all public repositories
	repos_user = "/users/:user/repos" // List public repositories for the specified user.
	repos_org  = "/orgs/:org/repos"   // List repositories for the specified org.
)

// Repository struct
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
//
//     repos := Github.GetRepos()
func (g *Github) Repos() []*Repo {
	url := apiUrl + repos

	return g.GetRepos(url)
}

// List public repositories for the specified user.
//
//     repos := Github.UserRepos("plouc")
func (g *Github) UserRepos(user string) []*Repo {
	url := apiUrl + strings.Replace(repos_user, ":user", user, -1)

	return g.GetRepos(url)
}

// List repositories for the specified org.
//
//     repos := Github.OrgRepos("ekino")
func (g *Github) OrgRepos(org string) []*Repo {
	url := apiUrl + strings.Replace(repos_org, ":org", org, -1)

	return g.GetRepos(url)
}