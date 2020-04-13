package templates

import "text/template"

var Set = template.Must(template.New("set").Parse(tpl_set))
var SetTest = template.Must(template.New("set_test").Parse(tpl_set_test))
