package app

import "github.com/andrewarrow/feedback/router"

func AfterCreateUser(c *router.Context, guid string) {
	u := c.LookupUser(guid)

	orgName := u["username"].(string) + "_org"
	params := map[string]any{"name": orgName}
	newGuid := c.Insert("org", params)

	org := c.SelectOne("org", "where guid=$1", []any{newGuid})

	c.UpdateOne("user", "org_id=$1", "id=$2", []any{org["id"], u["id"]})

	params = map[string]any{"org_id": org["id"], "name": orgName + "_project1"}
	c.Insert("project", params)
}
