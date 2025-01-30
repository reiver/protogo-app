package baseui

import (
	_ "embed"
)

const (
	IDContentChat    string = "content-chat"
	IDContentGigs    string = "content-gigs"
	IDContentPeople  string = "content-people"
	IDContentProfile string = "content-profile"
	IDContentSearch  string = "content-search"
)

//go:embed ui.html
var HTML string
