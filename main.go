package main

import (
	"math/rand"
	"os"
	"talk2track/app"
	"time"

	"github.com/andrewarrow/feedback/router"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	if len(os.Args) == 1 {
		return
	}
	arg := os.Args[1]

	if arg == "init" {
		router.InitNewApp()
	} else if arg == "reset" {
		r := router.NewRouter()
		r.ResetDatabase()
	} else if arg == "run" {
		r := router.NewRouter()
		r.Paths["/"] = app.HandleWelcome
		r.Paths["dashboard"] = app.HandleDashboard
		r.Paths["issues"] = app.HandleIssues
		r.AfterCreate["user"] = app.AfterCreateUser
		r.ListenAndServe(":3000")
	} else if arg == "help" {
	}
}
