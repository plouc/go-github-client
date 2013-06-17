package main

import (
	"flag"
	"fmt"
	"github.com/plouc/go-github-client"
)

func main() {
	help := flag.Bool("help", false, "Show usage")

	var function string
	flag.StringVar(&function, "f", "", "Specify method to retrieve events, available methods:\n" +
									   "  > all\n" +
									   "  > repo\n" +
									   "  > repo_issues\n" +
									   "  > repo_network\n" +
									   "  > user_received\n" +
									   "  > user_received_pub\n" +
									   "  > user_performed\n" +
									   "  > user_performed_pub\n" +
									   "  > org\n" +
									   "  > org_public")

	var user string
	flag.StringVar(&user, "u", "", "Specify user")

	flag.Usage = func() {
		fmt.Printf("Usage:\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	if *help == true || function == "" {
		flag.Usage()
		return
	}

	github := gogithub.NewGithub()

	switch function {
	case "all":
		events, err := github.Events()
		if err != nil {
			fmt.Println(err.Error())
		}
	
		for _, event := range events {
			fmt.Printf("%s > %s\n", event.CreatedAt.Format("Mon 02 Jan 15:04"), event.Message(""))
		}
	case "user_performed":
		if user == "" {
			fmt.Println("Please provide a user name within the -u argument")
			return
		}
		events, err := github.UserPerformedEvents(user)
		if err != nil {
			fmt.Println(err.Error())
		}
	
		for _, event := range events {
			fmt.Printf("%s > %s\n", event.CreatedAt.Format("Mon 02 Jan 15:04"), event.Message(""))
		}
	}
}