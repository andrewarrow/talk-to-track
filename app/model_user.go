package app

import "github.com/andrewarrow/feedback/router"

func AfterCreateUser(c *router.Context, guid string) {
	u := c.LookupUser(guid)

	orgName := u["username"].(string) + "_org"
	c.Params = map[string]any{"name": orgName}
	c.Validate("org")
	c.Insert("org")
	newGuid := c.Params["guid"]

	org := c.SelectOne("org", "where guid=$1", []any{newGuid})

	c.UpdateOne("user", "org_id=$1", "id=$2", []any{org["id"], u["id"]})

	c.Params = map[string]any{"org_id": org["id"], "name": orgName + "_project1"}
	c.Validate("project")
	c.Insert("project")
	newGuid = c.Params["guid"]
	project := c.SelectOne("project", "where guid=$1", []any{newGuid})

	c.Params = map[string]any{"project_id": project["id"], "org_id": org["id"], "title": "sample issue one"}
	c.Validate("issue")
	c.Insert("issue")
}
