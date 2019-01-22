package fixtures

import (
	"path/filepath"
	"testing"

	"github.com/bblfsh/javascript-driver/driver/normalizer"
	"gopkg.in/bblfsh/sdk.v2/driver"
	"gopkg.in/bblfsh/sdk.v2/driver/fixtures"
	"gopkg.in/bblfsh/sdk.v2/driver/native"
	"gopkg.in/bblfsh/sdk.v2/uast/transformer/positioner"
)

const projectRoot = "../../"

var Suite = &fixtures.Suite{
	Lang: "javascript",
	Ext:  ".js",
	Path: filepath.Join(projectRoot, fixtures.Dir),
	NewDriver: func() driver.Native {
		return native.NewDriverAt(filepath.Join(projectRoot, "build/bin/native"), native.UTF8)
	},
	Transforms: normalizer.Transforms,
	BenchName:  "u2_class_method", // TODO: specify a largest file
	Semantic: fixtures.SemanticConfig{
		BlacklistTypes: []string{
			"Identifier",
			"StringLiteral",
			"CommentLine",
			"CommentBlock",
			"BlockStatement",
			"ImportDeclaration",
			"ImportSpecifier",
			"ImportDefaultSpecifier",
			"ImportNamespaceSpecifier",
			"FunctionDeclaration",
		},
	},
	VerifyTokens: []positioner.VerifyToken{
		{Types: []string{
			// TODO(dennwc): issues with positions in native AST
			//               in some cases a positional info of an
			//               identifier covers the whole parameter
			//               declaration

			//"Identifier",

			// TODO(dennwc): positions doesn't cover the "//" and "/*" tokens

			// "CommentLine",
			// "CommentBlock",

			"StringLiteral",
			"RegExpLiteral",
			"StringLiteral",
			"BooleanLiteral",
			"NumericLiteral",
		}},
	},
}

func TestJavascriptDriver(t *testing.T) {
	Suite.RunTests(t)
}

func BenchmarkJavascriptDriver(b *testing.B) {
	Suite.RunBenchmarks(b)
}
