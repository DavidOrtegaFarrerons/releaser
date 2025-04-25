package jira

import (
	"bytes"
	"text/template"

	"github.com/spf13/viper"
)

func getJQL(releaseName string) (string, error) {
	tmplString := viper.GetString("JIRA_JQL")
	tmpl, err := template.New("jql").Parse(tmplString)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, map[string]string{
		"releaseName": releaseName,
	})
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
