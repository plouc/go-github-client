go-github-client
================

go-github-client is a simple client to consume github API (V3).

Events
------
* **func Events** list public events 
* **func RepoEvents** list repository events
* **func RepoIssuesEvents** list issue events for a repository
* **func RepoNetworkEvents** list public events for a network of repositories
* **func UserReceivedEvents** list events that a user has received
* **func UserReceivedPublicEvents** list public events that a user has received
* **func UserPerformedEvents** list events performed by a user
* **func UserPerformedPublicEvents** list public events performed by a user
* **func OrgEvents** list events for an organization
* **func OrgPublicEvents** list public events for an organization

------

Repositories
------------
* **func Repos** list all public repositories
* **func UserRepos** list public repositories for the specified user.
* **func OrgRepos** list repositories for the specified org

------


Installation
============

To install go-github-client, use `go get`:

    go get github.com/plouc/go-github-client

Import the `go-github-client` package into your code:

    package whatever

    import (
      "github.com/plouc/go-github-client"
    )

------
    
Update
------

To update `go-github-client`, use `go get -u`:

    go get -u github.com/plouc/go-github-client

------



