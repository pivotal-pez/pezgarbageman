package main

const (
	purge_full_name  string = "purge"
	purge_short_name        = "d"
	purge_usage             = "purge -t uaa -u randomuser@nowhere.com"
	purge_descr             = "allows us to purge the output of a command to a desintaion defined in a config file"
	show_full_name   string = "show"
	show_short_name         = "s"
	show_usage              = "show -t uaa -u randomuser@nowhere.com"
	show_descr              = "allows us to just view the purge candidates"
	usertypeDescr    string = "user type filter (okta or uaa) a managed user is okta, a unmanaged user is uaa"
	usernameDescr    string = "this is the name of the specific user you would like to search for"
)

var (
	usertypeFlag = []string{"usertype", "t"}
	usernameFlag = []string{"username", "u"}
)
