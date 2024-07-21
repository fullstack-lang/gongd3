package gongd3

import "embed"

// NgDistNg is the export of angular distribution. This allows
// embedding of the pages in the web server
//
//go:embed ng-github.com-fullstack-lang-gongd3/dist/ng-github.com-fullstack-lang-gongd3
var NgDistNg embed.FS
