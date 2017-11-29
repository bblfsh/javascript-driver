package babylon

import (
	"github.com/bblfsh/sdk/uast/ann"
)

var (
	File    = ann.HasInternalType("Type")
	Program = ann.HasInternalType("Program")

	// Statements
	ExpressionStatement = ann.HasInternalType("ExpressionStatement")

	// Expression
	Super                   = ann.HasInternalType("Super")
	Import                  = ann.HasInternalType("Import")
	ThisExpression          = ann.HasInternalType("ThisExpression")
	ArrowFunctionExpression = ann.HasInternalType("ArrowFunctionExpression")
	YieldExpression         = ann.HasInternalType("YieldExpression")
	AwaitExpression         = ann.HasInternalType("AwaitExpression")
	ArrayExpression         = ann.HasInternalType("ArrayExpression")
	ObjectExpression        = ann.HasInternalType("ObjectExpression")
	FunctionExpression      = ann.HasInternalType("FunctionExpression")
	CallExpression          = ann.HasInternalType("CallExpression")
	MemberExpression        = ann.HasInternalType("MemberExpression")

	// Object properties
	ObjectMethod   = ann.HasInternalType("ObjectMethod")
	ObjectProperty = ann.HasInternalType("ObjectProperty")
)
