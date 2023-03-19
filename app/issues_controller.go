package app

import (
	"github.com/andrewarrow/feedback/router"
)

func HandleIssues(c *router.Context, second, third string) {
	if c.User == nil {
		c.UserRequired = true
		return
	}
	if second != "" {
		issue := c.SelectOne("issue", "where guid=$1", []any{second})
		c.SendContentInLayout("issues_show.html", issue, 200)
		return
	}
	c.NotFound = true
}
