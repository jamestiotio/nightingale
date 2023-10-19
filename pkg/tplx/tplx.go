package tplx

import (
	"bytes"
	"html/template"
	"net/url"
	"regexp"
	"strings"

	"github.com/toolkits/pkg/logger"
)

var TemplateFuncMap = template.FuncMap{
	"escape":                    url.PathEscape,
	"unescaped":                 Unescaped,
	"urlconvert":                Urlconvert,
	"timeformat":                Timeformat,
	"timestamp":                 Timestamp,
	"args":                      Args,
	"reReplaceAll":              ReReplaceAll,
	"match":                     regexp.MatchString,
	"toUpper":                   strings.ToUpper,
	"toLower":                   strings.ToLower,
	"contains":                  strings.Contains,
	"humanize":                  Humanize,
	"humanize1024":              Humanize1024,
	"humanizeDuration":          HumanizeDuration,
	"humanizeDurationInterface": HumanizeDurationInterface,
	"humanizePercentage":        HumanizePercentage,
	"humanizePercentageH":       HumanizePercentageH,
	"add":                       Add,
	"sub":                       Subtract,
	"mul":                       Multiply,
	"div":                       Divide,
	"now":                       Now,
	"toString":                  ToString,
	"formatDecimal":             FormatDecimal,
}

// ReplaceMacroVariables replaces variables in a template string with values.
//
// It accepts the following parameters:
//
// - name: The name to use when parsing the template
//
// - templateText: The template string containing variables to replace
//
// - macroValue: A struct containing fields to replace the variables
//
// It parses the templateText into a template using template.New and template.Parse.
//
// It executes the parsed template with macroValue as the data, writing the result
// to a bytes.Buffer.
//
// Any {{.Field}} variables in templateText are replaced with values from macroValue.
//
// If there are any errors parsing or executing the template, they are logged and
// the original templateText is returned.
//
// The rendered template string is returned on success.
//
// Example usage:
//
//	type Data struct {
//	  Name string
//	}
//
//	data := Data{"John"}
//
//	output := ReplaceMacroVariables("mytpl", "Hello {{.Name}}!", data)
func ReplaceMacroVariables(name string, templateText string, macroValue any) string {
	tpl, err := template.New(name).Parse(templateText)
	if err != nil {
		logger.Warningf("parse config error: %v", err)
		return templateText
	}
	var body bytes.Buffer
	if err := tpl.Execute(&body, macroValue); err != nil {
		logger.Warningf("execute config error: %v", err)
		return templateText
	}
	return body.String()
}
