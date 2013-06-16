package gogithub

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
	"strconv"
)

const (
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
)

// Event struct
type Event struct {
	Id                string             `json:"id,omitempty"`
	Type              string             `json:"type,omitempty"`
	Public            bool               `json:"public,omitempty"`
	Repo              *SimpleRepo        `json:"repo,omitempty"`
	Actor             *SimpleUser        `json:"actor,omitempty"`
	Org               *SimpleUser        `json:"org,omitempty"`
	CreatedAt         time.Time 
	CreatedAtRaw      string             `json:"created_at,omitempty"`
	Payload           json.RawMessage    `json:"payload,omitempty"`
	Created           *Creation        
	Pushed            *Push
	PullRequestAction *PullRequestAction
}

type Push struct {
	Head         string    `json:"head,omitempty"`          // The SHA of the HEAD commit on the repository.
	Ref          string    `json:"ref,omitempty"`           // The full Git ref that was pushed. Example: “refs/heads/master”
	Size         int       `json:"size,omitempty"`          // The number of commits in the push
	DistinctSize int       `json:"distinct_size,omitempty"`
	Commits      []*Commit `json:"commits,omitempty"`       // The list of pushed commits.
}

type PullRequestAction struct {
	Action string `json:"action,omitempty"` // The action that was performed: “opened”, “closed”, “synchronize”, or “reopened”.
	Number int    `json:"number,omitempty"` // The pull request number.
	// @todo add pull_request object
}

// Represents a created repository, branch, or tag.
type Creation struct {
	RefType      string `json:"ref_type,omitempty"`      // The object that was created: “repository”, “branch”, or “tag”
	Ref          string `json:"ref,omitempty"`           // The git ref (or null if only a repository was created).
	MasterBranch string `json:"master_branch,omitempty"` // The name of the repository’s master branch.
	Description  string `json:"description,omitempty"`   // The repository’s current description.
}

func (e *Event) Message(me string) string {

	var message string

	user := e.Actor.Login
	if me != "" && me == user {
		user = "you"
	}

	switch e.Type {
	default:
		message = user + " - " + e.Type + " - " + e.Repo.Name
	case "PushEvent":
		message = user + " pushed " + strconv.Itoa(e.Pushed.Size) + " commit(s) to repository " + e.Repo.Name + " ("
		for index, commit := range e.Pushed.Commits {
			message = message + commit.Message
			if index < e.Pushed.Size - 1 {
				message = message + ", "
			}
		}
		message = message + ")"
	case "PublicEvent":
		message = user + " open sourced repository " + e.Repo.Name
	case "PullRequestEvent":
		message = user + " " + e.PullRequestAction.Action +
		          " pull request #" + strconv.Itoa(e.PullRequestAction.Number) +
				  " on repository " + e.Repo.Name
	case "CreateEvent":
		message = user + " created " + e.Created.RefType
		if e.Created.RefType == "repository" {
			message = message + " " + e.Repo.Name
		} else {
			if e.Created.RefType == "branch" {
				message = message + " " + e.Created.Ref
			}
			message = message + " on repository " + e.Repo.Name
		}
	}

	return message
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
		t, _ := time.Parse(dateLayout, event.CreatedAtRaw)
    	event.CreatedAt = t

    	switch event.Type {
		default:
		case "CreateEvent":
			var created *Creation
			err := json.Unmarshal(event.Payload, &created)
			if err != nil {
				fmt.Println("%s", err)
			}
			event.Created = created
		case "PublicEvent":
		case "PullRequestEvent":
			var pullRequestAction *PullRequestAction
			err := json.Unmarshal(event.Payload, &pullRequestAction)
			if err != nil {
				fmt.Println("%s", err)
			}
			event.PullRequestAction = pullRequestAction
		case "PushEvent":
			var pushed *Push
			err := json.Unmarshal(event.Payload, &pushed)
			if err != nil {
				fmt.Println("%s", err)
			}
			event.Pushed = pushed
		}    	
	}

	return events
}

// List public events
func (g *Github) Events() []*Event {
	url := g.apiUrl + events_url_public

	return g.GetEvents(url)
}

// List repository events
func (g *Github) RepoEvents(owner string, repo string) []*Event {
	url := strings.Replace(events_url_repo, ":owner", owner, -1)
	url = g.apiUrl + strings.Replace(url, ":repo", repo, -1)

	return g.GetEvents(url)
}

// List issue events for a repository
func (g *Github) RepoIssuesEvents(owner string, repo string) []*Event {
	url := strings.Replace(events_url_repo_issues, ":owner", owner, -1)
	url = g.apiUrl + strings.Replace(url, ":repo", repo, -1)

	return g.GetEvents(url)
}

// List public events for a network of repositories
func (g *Github) RepoNetworkEvents(owner string, repo string) []*Event {
	url := strings.Replace(events_url_network_public, ":owner", owner, -1)
	url = g.apiUrl + strings.Replace(url, ":repo", repo, -1)

	return g.GetEvents(url)
}

// List events that a user has received
// These are events that you’ve received by watching repos and following users.
// If you are authenticated as the given user, you will see private events.
// Otherwise, you’ll only see public events.
func (g *Github) UserReceivedEvents(user string) []*Event {
	url := g.apiUrl + strings.Replace(events_url_user_received, ":user", user, -1)

	return g.GetEvents(url)
}

// List public events that a user has received
func (g *Github) UserReceivedPublicEvents(user string) []*Event {
	url := g.apiUrl + strings.Replace(events_url_user_received_public, ":user", user, -1)

	return g.GetEvents(url)
}

// List events performed by a user
// If you are authenticated as the given user, you will see your private events.
// Otherwise, you’ll only see public events.
func (g *Github) UserPerformedEvents(user string) []*Event {
	url := g.apiUrl + strings.Replace(events_url_user_performed, ":user", user, -1)

	return g.GetEvents(url)
}

// List public events performed by a user
func (g *Github) UserPerformedPublicEvents(user string) []*Event {
	url := g.apiUrl + strings.Replace(events_url_user_performed_public, ":user", user, -1)

	return g.GetEvents(url)
}

// List events for an organization
// This is the user’s organization dashboard.
// You must be authenticated as the user to view this.
func (g *Github) OrgEvents(user string, org string) []*Event {
	url := strings.Replace(events_url_org, ":user", user, -1)
	url = g.apiUrl + strings.Replace(url, ":org", org, -1)

	return g.GetEvents(url)
}

// List public events for an organization
func (g *Github) OrgPublicEvents(org string) []*Event {
	url := g.apiUrl + strings.Replace(events_url_org_public, ":org", org, -1)

	return g.GetEvents(url)
}