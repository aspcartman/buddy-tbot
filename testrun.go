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
	"current_date": "2017-12-22T16:02:50.244Z",
	"execution": {
		"url": "https://buddy.aspc.me/api/workspaces/aspc/projects/buddy-tbot/pipelines/3/executions/5a3d2c604754f9001040d774",
		"html_url": "https://buddy.aspc.me/aspc/buddy-tbot/pipelines/pipeline/3/execution/5a3d2c604754f9001040d774",
		"id": 8,
		"start_date": "2017-12-22T16:01:36.097Z",
		"finish_date": "2017-12-22T16:02:49.780Z",
		"mode": "ON_EVERY_PUSH",
		"refresh": false,
		"status": "SUCCESSFUL",
		"comment": "I'm having fun'",
		"branch": {
			"url": "https://buddy.aspc.me/api/workspaces/aspc/projects/buddy-tbot/repository/branches/master",
			"html_url": "https://buddy.aspc.me/aspc/buddy-tbot/repository/branch/master",
			"name": "master",
			"default": false
		},
		"from_revision": {
			"url": "https://buddy.aspc.me/api/workspaces/aspc/projects/buddy-tbot/repository/commits/78d640d86923db7723744d054522cef2735422ab",
			"html_url": "https://buddy.aspc.me/aspc/buddy-tbot/repository/commit/78d640d86923db7723744d054522cef2735422ab",
			"revision": "78d640d86923db7723744d054522cef2735422ab",
			"short_revision": "78d640d",
			"message": "Improved responsiveness",
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
			"url": "https://buddy.aspc.me/api/workspaces/aspc/projects/buddy-tbot/repository/commits/aa90668db5f5ec069777843681168d9520ce46c1",
			"html_url": "https://buddy.aspc.me/aspc/buddy-tbot/repository/commit/aa90668db5f5ec069777843681168d9520ce46c1",
			"revision": "aa90668db5f5ec069777843681168d9520ce46c1",
			"short_revision": "aa90668",
			"message": "Hostname",
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
			"last_execution_status": "SUCCESSFUL",
			"last_execution_revision": "aa90668db5f5ec069777843681168d9520ce46c1",
			"always_from_scratch": false,
			"no_skip_to_most_recent": false,
			"auto_clear_cache": false
		},
		"action_executions": [{
				"status": "SUCCESSFUL",
				"progress": 100,
				"action": {
					"url": "https://buddy.aspc.me/api/workspaces/aspc/projects/buddy-tbot/pipelines/3/actions/4",
					"html_url": "https://buddy.aspc.me/aspc/buddy-tbot/pipelines/pipeline/3/action/4/edit",
					"id": 4,
					"name": "Build binary",
					"type": "BUILD",
					"trigger_time": "ON_EVERY_EXECUTION",
					"last_execution_status": "SUCCESSFUL",
					"disabled": false,
					"trigger_condition": "ALWAYS"
				},
				"log": [
					"Build started.",
					"Fetching changes started.",
					"Fetching changes finished.",
					"Pulling image library/golang:1.9.2-alpine3.7",
					"Creating image started.",
					"Sending build context to Docker daemon  2.048kB",
					"",
					"Step 1/3 : FROM library/golang:1.9.2-alpine3.7",
					" ---> aa1093f7445e",
					"Step 2/3 : RUN apk add --no-cache git make",
					" ---> Using cache",
					" ---> 0097ddee19d6",
					"Step 3/3 : RUN go get -u github.com/golang/dep/cmd/dep",
					" ---> Using cache",
					" ---> ce8f1daffee8",
					"Successfully built ce8f1daffee8",
					"Successfully tagged buddy/e4ce788e990bd36c240991fe10cde9da:1.0",
					"Creating image finished.",
					"",
					"BUDDY_COMMAND: export GOPATH=/",
					"",
					"BUDDY_COMMAND: make build",
					"dep ensure",
					"mkdir -pv build",
					"GOOS=linux GOARCH=amd64 go build -i -o build/buddy_bot main.go",
					"Build finished successfully!."
				],
				"log_url": "https://buddy.aspc.me/aspc/buddy-tbot/pipelines/pipeline/3/execution/5a3d2c604754f9001040d774/action/4/logs"
			},
			{
				"status": "SUCCESSFUL",
				"progress": 100,
				"action": {
					"url": "https://buddy.aspc.me/api/workspaces/aspc/projects/buddy-tbot/pipelines/3/actions/5",
					"html_url": "https://buddy.aspc.me/aspc/buddy-tbot/pipelines/pipeline/3/action/5/edit",
					"id": 5,
					"name": "Build Docker image",
					"type": "DOCKERFILE",
					"trigger_time": "ON_EVERY_EXECUTION",
					"last_execution_status": "SUCCESSFUL",
					"disabled": false,
					"trigger_condition": "ALWAYS"
				},
				"log": [
					"Build started.",
					"",
					"BUDDY_COMMAND: echo ###### | docker login --username=- --password-stdin registry.aspc.me",
					"Login Succeeded",
					"",
					"BUDDY_COMMAND: docker build -m 4g -t registry.aspc.me/aspcartman/buddy-bot:latest --pull -f \"Dockerfile\" .",
					"Sending build context to Docker daemon  19.37MB",
					"",
					"Step 1/5 : FROM alpine:latest",
					"latest: Pulling from library/alpine",
					"Digest: sha256:ccba511b1d6b5f1d83825a94f9d5b05528db456d9cf14a1ea1db892c939cda64",
					"Status: Image is up to date for alpine:latest",
					" ---> e21c333399e0",
					"Step 2/5 : RUN apk add --no-cache ca-certificates",
					" ---> Using cache",
					" ---> d3e3ab99e79b",
					"Step 3/5 : COPY build/buddy_bot /buddy_bot",
					" ---> 2947c3404806",
					"Step 4/5 : RUN chmod +x /buddy_bot",
					" ---> Running in 378c0ddeb442",
					" ---> ab6c5de9ae49",
					"Removing intermediate container 378c0ddeb442",
					"Step 5/5 : ENTRYPOINT /buddy_bot",
					" ---> Running in d333f5c8cbd2",
					" ---> 3dd9ddaab655",
					"Removing intermediate container d333f5c8cbd2",
					"Successfully built 3dd9ddaab655",
					"Successfully tagged registry.aspc.me/aspcartman/buddy-bot:latest",
					"",
					"BUDDY_COMMAND: docker push registry.aspc.me/aspcartman/buddy-bot:latest",
					"The push refers to a repository [registry.aspc.me/aspcartman/buddy-bot]",
					"a590013f3ecf: Preparing",
					"a590013f3ecf: Preparing",
					"52a1cf711ad5: Preparing",
					"04a094fe844e: Preparing",
					"52a1cf711ad5: Layer already exists",
					"04a094fe844e: Layer already exists",
					"a590013f3ecf: Pushed",
					"latest: digest: sha256:63db287975f95ea078c41f648260fedc4024bd22cafc9a477c5fac2dffcea59d size: 1160",
					"Build finished successfully!."
				],
				"log_url": "https://buddy.aspc.me/aspc/buddy-tbot/pipelines/pipeline/3/execution/5a3d2c604754f9001040d774/action/5/logs"
			},
			{
				"status": "SUCCESSFUL",
				"progress": 100,
				"action": {
					"url": "https://buddy.aspc.me/api/workspaces/aspc/projects/buddy-tbot/pipelines/3/actions/6",
					"html_url": "https://buddy.aspc.me/aspc/buddy-tbot/pipelines/pipeline/3/action/6/edit",
					"id": 6,
					"name": "Upload docker-compose.yml",
					"type": "SFTP",
					"trigger_time": "ON_EVERY_EXECUTION",
					"last_execution_status": "SUCCESSFUL",
					"disabled": false
				},
				"log": [
					"Listing files that will be deployed...",
					"List of files is ready. Starting uploading...",
					"Uploading file docker-compose.yml...",
					"Deployment accomplished, great!"
				],
				"log_url": "https://buddy.aspc.me/aspc/buddy-tbot/pipelines/pipeline/3/execution/5a3d2c604754f9001040d774/action/6/logs"
			},
			{
				"status": "SUCCESSFUL",
				"progress": 100,
				"action": {
					"url": "https://buddy.aspc.me/api/workspaces/aspc/projects/buddy-tbot/pipelines/3/actions/7",
					"html_url": "https://buddy.aspc.me/aspc/buddy-tbot/pipelines/pipeline/3/action/7/edit",
					"id": 7,
					"name": "Deploy to aspc.me",
					"type": "SSH_COMMAND",
					"trigger_time": "ON_EVERY_EXECUTION",
					"last_execution_status": "SUCCESSFUL",
					"disabled": false,
					"trigger_condition": "ALWAYS"
				},
				"log": [
					"Resolving host aspc.me...",
					"Executing command: docker-compose pull",
					"latest: Pulling from aspcartman/buddy-bot",
					"Digest: sha256:63db287975f95ea078c41f648260fedc4024bd22cafc9a477c5fac2dffcea59d",
					"Status: Downloaded newer image for registry.aspc.me/aspcartman/buddy-bot:latest",
					"Executing command: docker-compose down",
					"Executing command: docker-compose up -d",
					"Commands accomplished, great!"
				],
				"log_url": "https://buddy.aspc.me/aspc/buddy-tbot/pipelines/pipeline/3/execution/5a3d2c604754f9001040d774/action/7/logs"
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
