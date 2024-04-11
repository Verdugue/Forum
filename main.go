package main

import (
	"FORUM/routeur"
	"FORUM/templates"
)

func main() {
	templates.InitTemplate()
	routeur.Initserv()
}
