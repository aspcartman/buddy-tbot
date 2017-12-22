package bot

import (
	"bytes"
	"io/ioutil"
	"text/template"

	"github.com/aspcartman/buddy-tbot/api/buddy"
	"github.com/aspcartman/buddy-tbot/e"
)

var buddyTmpl *template.Template

func (b *Bot) HandleBuddyEvent(event *buddy.WebHookEvent) {
	b.SendMessage(buildUpATGMessage(event))
}

func buildUpATGMessage(event *buddy.WebHookEvent) string {
	buf := &bytes.Buffer{}
	e.Must(prepareTemplate().Execute(buf, event), "failed executing the message template")
	return buf.String()
}

func prepareTemplate() *template.Template {
	if buddyTmpl != nil {
		return buddyTmpl
	}

	tmplText, err := ioutil.ReadFile("resources/msg.template")
	if err != nil {
		e.Throw("failed reading template file", err)
	}

	tmpl, err := template.New("ololo?").Parse(string(tmplText))
	if err != nil {
		e.Throw("failed parsing template", err, e.Map{
			"text": string(tmplText),
		})
	}
	buddyTmpl = tmpl
	return tmpl
}
