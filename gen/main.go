// Based on https://goa.design/en/design/vendoring/ suggestions
package main

import (
	_ "github.com/goadesign/gorma-cellar/design"

	"github.com/goadesign/goa/design"
	"github.com/goadesign/goa/goagen/codegen"
	genapp "github.com/goadesign/goa/goagen/gen_app"
	genclient "github.com/goadesign/goa/goagen/gen_client"
	genjs "github.com/goadesign/goa/goagen/gen_js"
	genschema "github.com/goadesign/goa/goagen/gen_schema"
	genswagger "github.com/goadesign/goa/goagen/gen_swagger"
	genmeta "github.com/goadesign/goa/goagen/meta"
)

func main() {
	codegen.ParseDSL()
	codegen.Run(
		&genapp.Generator{
			API:    design.Design,
			OutDir: "app",
			Target: "app",
			NoTest: false,
		},
		&genclient.Generator{
			API: design.Design,
		},
		&genswagger.Generator{
			API: design.Design,
		},
		&genjs.Generator{
			API: design.Design,
		},
		&genschema.Generator{
			API: design.Design,
		},
		&genmeta.Generator{
			Genfunc:       "gorma.Generate",
			Imports:       []*codegen.ImportSpec{codegen.SimpleImport("github.com/goadesign/gorma")},
			DesignPkgPath: "github.com/goadesign/gorma-cellar/design",
			OutDir:        "models",
			Flags:         map[string]string{"debug": "true"},
		},
	)
}
