package gogithub

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
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
		message = user + " pushed to " + e.Repo.Name
	case "PublicEvent":
		message = user + " open sourced " + e.Repo.Name
	case "CreateEvent":
		message = user + " created " + e.Repo.Name
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
// These are events that you’ve received by watching repos and following users.
// If you are authenticated as the given user, you will see private events.
// Otherwise, you’ll only see public events.
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
// If you are authenticated as the given user, you will see your private events.
// Otherwise, you’ll only see public events.
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
// This is the user’s organization dashboard.
// You must be authenticated as the user to view this.
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