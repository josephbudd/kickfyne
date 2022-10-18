package screens

import "github.com/josephbudd/kickfyne/source/utils"

type docTemplateData struct {
	PackageName string
	PackageDoc  string
	Funcs       utils.Funcs
}

const (
	docTemplate = `{{ call .Funcs.Comment .PackageDoc }}
package {{ .PackageName }}
`
)
