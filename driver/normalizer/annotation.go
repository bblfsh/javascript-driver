package normalizer

import (
	"github.com/bblfsh/javascript-driver/driver/normalizer/babylon"
	"gopkg.in/bblfsh/sdk.v1/uast"
	. "gopkg.in/bblfsh/sdk.v1/uast/ann"
)

// AnnotationRules annotate a UAST with roles.
var AnnotationRules = On(babylon.File).Roles(uast.File).Descendants(
	On(babylon.Program).Roles(uast.Module).Descendants(
		// Statements
		On(babylon.ExpressionStatement).Roles(uast.Statement),

		// Expressions
		On(babylon.Super).Roles(uast.Expression, uast.Identifier),
		On(babylon.Import).Roles(uast.Expression, uast.Import),
		On(babylon.ThisExpression).Roles(uast.Expression, uast.Identifier),
		On(babylon.ArrowFunctionExpression).Roles(uast.Expression, uast.Function),
		On(babylon.YieldExpression).Roles(uast.Expression, uast.Return, uast.Incomplete),
		On(babylon.AwaitExpression).Roles(uast.Expression, uast.Incomplete),
		On(babylon.ArrayExpression).Roles(uast.Expression, uast.Initialization, uast.List, uast.Literal),
		On(babylon.ObjectExpression).Roles(uast.Expression, uast.Initialization, uast.Literal),
		On(babylon.FunctionExpression).Roles(uast.Expression, uast.Function),
		On(babylon.CallExpression).Roles(uast.Expression, uast.Call),
		On(babylon.MemberExpression).Roles(uast.Qualified, uast.Expression, uast.Identifier),
		On(Or(babylon.UnaryExpression, babylon.UpdateExpression)).Roles(uast.Expression, uast.Unary)
		On(babylon.BinaryExpression).Roles(uast.Expression, uast.Binary),

		// Object properties
		On(babylon.ObjectMethod).Roles(uast.Function, uast.Assignment),
		On(babylon.ObjectProperty).Roles(uast.Identifier, uast.Assignment),
	),
)
