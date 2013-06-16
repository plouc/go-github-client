#go-github-client
[![Build Status](https://travis-ci.org/plouc/go-github-client.png?branch=master)](https://travis-ci.org/plouc/go-github-client)

go-github-client is a simple client to consume github API (V3).

##features

*	
	###Events [github api doc](http://developer.github.com/v3/activity/events/)
	* list public events 
	* list repository events
	* list issue events for a repository
	* list public events for a network of repositories
	* list events that a user has received
	* list public events that a user has received
	* list events performed by a user
	* list public events performed by a user
	* list events for an organization
	* list public events for an organization

*	
	###Repositories [github api doc](http://developer.github.com/v3/repos/)
	* list all public repositories
	* list public repositories for the specified user.
	* list repositories for the specified org


##Installation

To install go-github-client, use `go get`:

    go get github.com/plouc/go-github-client

Import the `go-github-client` package into your code:

```go
package whatever

import (
    "github.com/plouc/go-github-client"
)
```

    
##Update

To update `go-github-client`, use `go get -u`:

    go get -u github.com/plouc/go-github-client



##Documentation

Visit the docs at http://godoc.org/github.com/plouc/go-github-client



