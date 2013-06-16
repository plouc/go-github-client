package gogithub

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"fmt"
	"github.com/stretchr/testify/assert"
)

func TestEvents(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `[{
"id":"0123456789",
"type":"PushEvent",
"actor":{
  "id":501642,
  "login":"plouc",
  "gravatar_id":"7070eab7c6206530d3b7820362227fec",
  "url":"https://api.github.com/users/plouc",
  "avatar_url":"https://secure.gravatar.com/avatar/7070eab7c6206530d3b7820362227fec?d=https://a248.e.akamai.net/assets.github.com%2Fimages%2Fgravatars%2Fgravatar-user-420.png"},
"repo":{
  "id":10675738,
  "name":"plouc/go-github-client",
  "url":"https://api.github.com/repos/plouc/go-github-client"},
"payload":{
  "push_id":188809031,
  "size":1,
  "distinct_size":1,
  "ref":"refs/heads/master",
  "head":"cb1964c5e88c976ffff882bd1d331dd6f35fe5bd",
  "before":"6474b21ef162c7c43a4a8bb7196ce937e30cd571",
  "commits":[{
    "sha":"cb1964c5e88c976ffff882bd1d331dd6f35fe5bd",
    "author":{
      "email":"test@test.com",
      "name":"Raphaël Benitte"
    },
    "message":"Add commit messages for PushEvent",
    "distinct":true,
    "url":"https://api.github.com/repos/plouc/go-github-client/commits/cb1964c5e88c976ffff882bd1d331dd6f35fe5bd"}
  ]},
  "public":true,
  "created_at":"2013-06-16T01:06:59Z"
}]`)
	}))
	defer ts.Close()

	github := NewGithub()
	github.apiUrl = ts.URL
	events := github.Events()

	assert.Equal(t, len(events), 1)
	assert.IsType(t, new(Event), events[0])
	assert.Equal(t, events[0].Id, "0123456789")
	assert.Equal(t, events[0].Type, "PushEvent")
	assert.IsType(t, new(SimpleUser), events[0].Actor)
	assert.IsType(t, new(SimpleRepo), events[0].Repo)
	assert.IsType(t, new(Push), events[0].Pushed)
	assert.Equal(t, events[0].Pushed.Size, 1)
	assert.Equal(t, events[0].Pushed.DistinctSize, 1)
	assert.Equal(t, events[0].Pushed.Ref, "refs/heads/master")
	assert.Equal(t, len(events[0].Pushed.Commits), 1)
}

func TestRepoEvents(t *testing.T) {
	t.Log("todo")
}

func TestRepoIssuesEvents(t *testing.T) {
	t.Log("todo")
}

func TestRepoNetworkEvents(t *testing.T) {
	t.Log("todo")
}

func TestUserReceivedEvents(t *testing.T) {
	t.Log("todo")
}

func TestUserReceivedPublicEvents(t *testing.T) {
	t.Log("todo")
}

func TestUserPerformedEvents(t *testing.T) {
	t.Log("todo")
}

func TestUserPerformedPublicEvents(t *testing.T) {
	t.Log("todo")
}

func TestOrgEvents(t *testing.T) {
	t.Log("todo")
}

func TestOrgPublicEvents(t *testing.T) {
	t.Log("todo")
}
