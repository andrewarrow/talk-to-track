package app

import (
	"net/http"

	"github.com/andrewarrow/feedback/router"
)

func HandleWelcome(c *router.Context, second, third string) {
	if c.User != nil {
		http.Redirect(c.Writer, c.Request, "/dashboard/", 302)
		return
	}
	c.SendContentInLayout("welcome_index.html", nil, 200)
}
