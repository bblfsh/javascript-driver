package fixtures

import (
	"path/filepath"
	"testing"

	"github.com/bblfsh/javascript-driver/driver/normalizer"
	"gopkg.in/bblfsh/sdk.v2/driver"
	"gopkg.in/bblfsh/sdk.v2/driver/fixtures"
	"gopkg.in/bblfsh/sdk.v2/driver/native"
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
	Docker: fixtures.DockerConfig{
		Image: "node:8",
	},
}

func TestJavascriptDriver(t *testing.T) {
	Suite.RunTests(t)
}

func BenchmarkJavascriptDriver(b *testing.B) {
	Suite.RunBenchmarks(b)
}
