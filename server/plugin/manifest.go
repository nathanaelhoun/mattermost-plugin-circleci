// This file is automatically generated. Do not modify it manually.

package plugin

import (
	"strings"

	"github.com/mattermost/mattermost-server/v5/model"
)

var manifest *model.Manifest

const manifestStr = `
{
  "id": "com.github.nathanaelhoun.plugin-circleci",
  "name": "CircleCI",
  "description": "Interact with CircleCI workflows, jobs and builds with slash commands",
  "icon_path": "public/circleci.png",
  "version": "0.1.0",
  "min_server_version": "5.12.0",
  "server": {
    "executables": {
      "linux-amd64": "server/dist/plugin-linux-amd64",
      "darwin-amd64": "server/dist/plugin-darwin-amd64",
      "windows-amd64": "server/dist/plugin-windows-amd64.exe"
    },
    "executable": ""
  },
  "settings_schema": {
    "header": "This plugin is under development. To contribute or submit a request, visit the [repository on Github](https://github.com/nathanaelhoun/mattermost-plugin-circleci)",
    "footer": "",
    "settings": [
      {
        "key": "WebhooksSecret",
        "display_name": "Webhooks Secret",
        "type": "generated",
        "help_text": "The secret used to authenticate the webhook to Mattermost",
        "regenerate_help_text": "Regenerates the secret for the webhook URL endpoint. Regenerating the secret invalidates your existing CircleCI integrations",
        "placeholder": "",
        "default": null
      },
      {
        "key": "EncryptionKey",
        "display_name": "At Rest Encryption Key",
        "type": "generated",
        "help_text": "The AES encryption key used to encrypt stored access tokens.",
        "placeholder": "",
        "default": null
      }
    ]
  }
}
`

func init() {
	manifest = model.ManifestFromJson(strings.NewReader(manifestStr))
}
