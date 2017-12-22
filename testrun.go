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
  "current_date": "2017-12-22T18:27:27.138Z",
  "execution": {
    "url": "https://buddy.aspc.me/api/workspaces/aspc/projects/buddy-tbot/pipelines/3/executions/5a3d4e8e4754f9001040d789",
    "html_url": "https://buddy.aspc.me/aspc/buddy-tbot/pipelines/pipeline/3/execution/5a3d4e8e4754f9001040d789",
    "id": 11,
    "start_date": "2017-12-22T18:27:26.735Z",
    "finish_date": null,
    "mode": "ON_EVERY_PUSH",
    "refresh": false,
    "status": "INPROGRESS",
    "comment": "",
    "branch": {
      "url": "https://buddy.aspc.me/api/workspaces/aspc/projects/buddy-tbot/repository/branches/master",
      "html_url": "https://buddy.aspc.me/aspc/buddy-tbot/repository/branch/master",
      "name": "master",
      "default": false
    },
    "from_revision": {
      "url": "https://buddy.aspc.me/api/workspaces/aspc/projects/buddy-tbot/repository/commits/f76fcad2379e841306b5b94abe32d3913d9e7d81",
      "html_url": "https://buddy.aspc.me/aspc/buddy-tbot/repository/commit/f76fcad2379e841306b5b94abe32d3913d9e7d81",
      "revision": "f76fcad2379e841306b5b94abe32d3913d9e7d81",
      "short_revision": "f76fcad",
      "message": "Showing something actually usefull",
      "committer": {
        "name": "Ruslan Fedorov",
        "title": null,
        "email": "ruslan@tapingo.com"
      },
      "author": {
        "title": null,
        "email": "ruslan@tapingo.com"
      },
      "tags": []
    },
    "to_revision": {
      "url": "https://buddy.aspc.me/api/workspaces/aspc/projects/buddy-tbot/repository/commits/df1ec85660b3fc6ea32d15efabd7abce77202e69",
      "html_url": "https://buddy.aspc.me/aspc/buddy-tbot/repository/commit/df1ec85660b3fc6ea32d15efabd7abce77202e69",
      "revision": "df1ec85660b3fc6ea32d15efabd7abce77202e69",
      "short_revision": "df1ec85",
      "message": "Fix json unmarshalling",
      "committer": {
        "name": "Ruslan Fedorov",
        "title": null,
        "email": "ruslan@tapingo.com"
      },
      "author": {
        "title": null,
        "email": "ruslan@tapingo.com"
      },
      "tags": []
    },
    "creator": {
      "url": "https://buddy.aspc.me/api/workspaces/aspc/members/1",
      "html_url": "https://buddy.aspc.me/aspc/profile/1",
      "id": 1,
      "name": "Aspcartman",
      "avatar_url": "https://buddy.aspc.me/image-server/user/0/0/0/0/0/0/1/d643744fbe5ebf2906a4d075a5b97110/w/32/32/AVATAR.png",
      "title": null,
      "email": "aspcartman@gmail.com"
    },
    "pipeline": {
      "url": "https://buddy.aspc.me/api/workspaces/aspc/projects/buddy-tbot/pipelines/3",
      "html_url": "https://buddy.aspc.me/aspc/buddy-tbot/pipelines/pipeline/3",
      "id": 3,
      "name": "Build",
      "trigger_mode": "ON_EVERY_PUSH",
      "ref_name": "master",
      "last_execution_status": "INPROGRESS",
      "last_execution_revision": "df1ec85660b3fc6ea32d15efabd7abce77202e69",
      "always_from_scratch": false,
      "no_skip_to_most_recent": false,
      "auto_clear_cache": false
    },
    "action_executions": [
      {
        "status": "ENQUEUED",
        "progress": 0,
        "action": {
          "url": "https://buddy.aspc.me/api/workspaces/aspc/projects/buddy-tbot/pipelines/3/actions/4",
          "html_url": "https://buddy.aspc.me/aspc/buddy-tbot/pipelines/pipeline/3/action/4/edit",
          "id": 4,
          "name": "Build binary",
          "type": "BUILD",
          "trigger_time": "ON_EVERY_EXECUTION",
          "last_execution_status": "ENQUEUED",
          "disabled": false,
          "trigger_condition": "ALWAYS"
        },
        "log": [],
        "log_url": "https://buddy.aspc.me/aspc/buddy-tbot/pipelines/pipeline/3/execution/5a3d4e8e4754f9001040d789/action/4/logs"
      },
      {
        "status": "ENQUEUED",
        "progress": 0,
        "action": {
          "url": "https://buddy.aspc.me/api/workspaces/aspc/projects/buddy-tbot/pipelines/3/actions/5",
          "html_url": "https://buddy.aspc.me/aspc/buddy-tbot/pipelines/pipeline/3/action/5/edit",
          "id": 5,
          "name": "Build Docker image",
          "type": "DOCKERFILE",
          "trigger_time": "ON_EVERY_EXECUTION",
          "last_execution_status": "ENQUEUED",
          "disabled": false,
          "trigger_condition": "ALWAYS"
        },
        "log": [],
        "log_url": "https://buddy.aspc.me/aspc/buddy-tbot/pipelines/pipeline/3/execution/5a3d4e8e4754f9001040d789/action/5/logs"
      },
      {
        "status": "ENQUEUED",
        "progress": 0,
        "action": {
          "url": "https://buddy.aspc.me/api/workspaces/aspc/projects/buddy-tbot/pipelines/3/actions/6",
          "html_url": "https://buddy.aspc.me/aspc/buddy-tbot/pipelines/pipeline/3/action/6/edit",
          "id": 6,
          "name": "Upload docker-compose.yml",
          "type": "SFTP",
          "trigger_time": "ON_EVERY_EXECUTION",
          "last_execution_status": "ENQUEUED",
          "disabled": false
        },
        "log_url": "https://buddy.aspc.me/aspc/buddy-tbot/pipelines/pipeline/3/execution/5a3d4e8e4754f9001040d789/action/6/logs"
      },
      {
        "status": "ENQUEUED",
        "progress": 0,
        "action": {
          "url": "https://buddy.aspc.me/api/workspaces/aspc/projects/buddy-tbot/pipelines/3/actions/7",
          "html_url": "https://buddy.aspc.me/aspc/buddy-tbot/pipelines/pipeline/3/action/7/edit",
          "id": 7,
          "name": "Deploy to aspc.me",
          "type": "SSH_COMMAND",
          "trigger_time": "ON_EVERY_EXECUTION",
          "last_execution_status": "ENQUEUED",
          "disabled": false,
          "trigger_condition": "ALWAYS"
        },
        "log_url": "https://buddy.aspc.me/aspc/buddy-tbot/pipelines/pipeline/3/execution/5a3d4e8e4754f9001040d789/action/7/logs"
      }
    ]
  }
}
`

func main() {
	b := bot.Run()

	event := buddy.WebHookEvent{}
	json.Unmarshal([]byte(js), &event)
	b.HandleBuddyEvent(&event)
}
