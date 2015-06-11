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

	cfdomainDescr string = "this is the target domain of your cf foundation"
	cfuserDescr   string = "this is the user you will use to login to the targetted foundation"
	cfpassDescr   string = "this is the cfuser's password"

	cfdomainEnv string = "CF_DOMAIN"
	cfuserEnv   string = "CF_USER"
	cfpassEnv   string = "CF_PASS"
)

var (
	usertypeFlag = []string{"usertype", "t"}
	usernameFlag = []string{"username", "u"}
	cfdomainFlag = []string{"cfdomain", "cd"}
	cfuserFlag   = []string{"cfuser", "cu"}
	cfpassFlag   = []string{"cfpass", "cp"}
)
