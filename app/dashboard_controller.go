package app

import (
	"github.com/andrewarrow/feedback/router"
)

func HandleDashboard(c *router.Context, second, third string) {
	if c.User == nil {
		c.UserRequired = true
		return
	}
	params := []any{0}
	rows := c.SelectAll("project", "where org_id=$1", params)
	c.SendContentInLayout("dashboard.html", rows, 200)
}
