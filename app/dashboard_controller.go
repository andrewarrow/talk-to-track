package app

import (
	"github.com/andrewarrow/feedback/router"
)

func HandleDashboard(c *router.Context, second, third string) {
	if c.User == nil {
		c.UserRequired = true
		return
	}
	params := []any{c.User["org_id"]}
	org := c.SelectOne("org", "where id=$1", params)

	projects := c.SelectAll("project", "where org_id=$1", params)

	dash := map[string]any{"org": org,
		"projects": projects}
	c.SendContentInLayout("dashboard.html", dash, 200)
}
