package main

import (
	"encoding/json"

	"github.com/aspcartman/buddy-tbot/api/buddy"
	"github.com/aspcartman/buddy-tbot/bot"
	_ "github.com/aspcartman/buddy-tbot/env"
)

var js = `
{
  "workspace": {
    "url": "https://buddy.aspc.me/api/workspaces/aspc",
    "html_url": "https://buddy.aspc.me/aspc",
    "id": 1,
    "name": "Aspc",
    "domain": "aspc"
  },
  "invoker": {
    "url": "https://buddy.aspc.me/api/workspaces/aspc/members/1",
    "html_url": "https://buddy.aspc.me/aspc/profile/1",
    "id": 1,
    "name": "Aspcartman",
    "avatar_url": "https://buddy.aspc.me/image-server/user/0/0/0/0/0/0/1/d643744fbe5ebf2906a4d075a5b97110/w/32/32/AVATAR.png",
    "title": null,
    "email": "aspcartman@gmail.com"
  },
  "project": {
    "url": "https://buddy.aspc.me/api/workspaces/aspc/projects/buddy-tbot",
    "html_url": "https://buddy.aspc.me/aspc/buddy-tbot",
    "name": "buddy-tbot",
    "display_name": "buddy-tbot",
    "status": "ACTIVE"
  },
  "current_date": "2017-12-22T18:35:57.046Z",
  "repository_url": "ssh://buddy@buddy.aspc.me:7879/aspc/buddy-tbot",
  "ref": "refs/heads/master",
  "push_id": 10,
  "operation": "UPDATE",
  "old_head": "df1ec85660b3fc6ea32d15efabd7abce77202e69",
  "new_head": "5d98cc75dbe3b4bd7c178a2cdb46ae85d2507e27",
  "commits_count": 1,
  "commits": [
    {
      "url": "https://buddy.aspc.me/api/projects/buddy-tbot/repository/commits/5d98cc75dbe3b4bd7c178a2cdb46ae85d2507e27",
      "html_url": "https://buddy.aspc.me/aspc/buddy-tbot/repository/commit/5d98cc75dbe3b4bd7c178a2cdb46ae85d2507e27",
      "revision": "5d98cc75dbe3b4bd7c178a2cdb46ae85d2507e27",
      "author_date": "2017-12-22T18:35:49.000Z",
      "commit_date": "2017-12-22T18:35:49.000Z",
      "message": "Add repos\n",
      "committer": {
        "url": "https://buddy.aspc.me/api/workspaces/aspc/members/1",
        "html_url": "https://buddy.aspc.me/aspc/profile/1",
        "id": 1,
        "name": "Aspcartman",
        "avatar_url": "https://buddy.aspc.me/image-server/user/0/0/0/0/0/0/1/d643744fbe5ebf2906a4d075a5b97110/w/32/32/AVATAR.png",
        "title": null,
        "email": "aspcartman@gmail.com"
      },
      "author": {
        "title": null,
        "email": "ruslan@tapingo.com"
      }
    }
  ],
  "push_url": "https://buddy.aspc.me/aspc/buddy-tbot/repository/comparison?path=%2F&fromRevision=df1ec85660b3fc6ea32d15efabd7abce77202e69&toRevision=5d98cc75dbe3b4bd7c178a2cdb46ae85d2507e27",
  "branch_url": "https://buddy.aspc.me/aspc/buddy-tbot/repository/branch/master"
}
`

func main() {
	b := bot.Run()

	event := buddy.WebHookEvent{}
	json.Unmarshal([]byte(js), &event)
	b.HandleBuddyEvent(&event)
}
