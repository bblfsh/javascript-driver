package normalizer

import (
	"gopkg.in/bblfsh/sdk.v2/uast"
	. "gopkg.in/bblfsh/sdk.v2/uast/transformer"
)

var Preprocess = Transformers([][]Transformer{
	{Mappings(Preprocessors...)},
}...)

var Normalize = Transformers([][]Transformer{
	{Mappings(Normalizers...)},
}...)

// Preprocessors is a block of AST preprocessing rules rules.
var Preprocessors = []Mapping{
	// ObjectToNode defines how to normalize common fields of native AST
	// (like node type, token, positional information).
	//
	// https://godoc.org/gopkg.in/bblfsh/sdk.v2/uast#ObjectToNode
	ObjectToNode{
		InternalTypeKey: "type",
		OffsetKey:       "start",
		EndOffsetKey:    "end",
	}.Mapping(),

	ASTMap("remove loc",
		Part("_", Obj{"loc": AnyNode(nil)}),
		Part("_", Obj{}),
	),
	ASTMap("remove extra",
		Part("_", Obj{"extra": AnyNode(nil)}),
		Part("_", Obj{}),
	),
}

// Normalizers is the main block of normalization rules to convert native AST to semantic UAST.
var Normalizers = []Mapping{
	MapSemantic("", "Identifier", uast.Identifier{}, nil,
		Obj{
			"name": Var("name"),
		},
		Obj{
			"Name": Var("name"),
		},
	),
	MapSemantic("", "StringLiteral", uast.String{}, nil,
		Obj{
			"value": Var("val"),
		},
		Obj{
			"Value": Var("val"),
		},
	),
	MapSemantic("", "CommentLine", uast.Comment{}, nil,
		Obj{
			"value": CommentText([2]string{"", ""}, "comm"),
		},
		CommentNode(false, "comm", nil),
	),
	MapSemantic("", "CommentBlock", uast.Comment{}, nil,
		Obj{
			"value": CommentText([2]string{"", ""}, "comm"),
		},
		CommentNode(true, "comm", nil),
	),
	MapSemantic("", "BlockStatement", uast.Block{}, nil,
		Obj{
			"body":       Var("stmts"),
			"directives": Arr(), // TODO: find an example
		},
		Obj{
			"Statements": Var("stmts"),
		},
	),
	MapSemantic("", "ImportDeclaration", uast.Import{}, nil,
		Obj{
			"source":     Var("path"),
			"specifiers": Arr(),
		},
		Obj{
			"Path": Var("path"),
		},
	),
	MapSemantic("", "ImportDeclaration", uast.Import{}, nil,
		Obj{
			"importKind": String("value"),
			"source":     Var("path"),
			"specifiers": Arr(Obj{ // FIXME: it may appear anywhere in the slice
				uast.KeyType: String("ImportNamespaceSpecifier"),
				uast.KeyPos:  Var("local_pos"),
				"local":      Var("local"),
			}),
		},
		Obj{
			"Path": UASTType(uast.Alias{}, Obj{
				uast.KeyPos: Var("local_pos"),
				"Name":      Var("local"),
				"Node":      Var("path"),
			}),
			"Names": Arr(),
			"All":   Bool(true),
		},
	),
	MapSemantic("", "ImportDeclaration", uast.Import{}, nil,
		Obj{
			"importKind": String("value"),
			"source":     Var("path"),
			"specifiers": Check(Not(Arr()), Var("names")),
		},
		Obj{
			"Path":  Var("path"),
			"Names": Var("names"),
		},
	),
	MapSemantic("", "ImportSpecifier", uast.Alias{}, nil,
		Obj{
			"importKind": Is(nil),
			"local":      Var("local"),
			"imported":   Var("imp"),
		},
		Obj{
			"Name": Var("local"),
			"Node": Var("imp"),
		},
	),
	MapSemantic("", "ImportDefaultSpecifier", uast.Alias{}, nil,
		Obj{
			"local": Var("local"),
		},
		Obj{
			"Name": Var("local"),
			"Node": UASTType(uast.Identifier{}, Obj{
				uast.KeyPos: AnyNode(nil),
				"Name":      String("."), // TODO: scope
			}),
		},
	),
	MapSemantic("", "ImportDeclaration", uast.Import{}, nil,
		Obj{
			"importKind": String("value"),
			"source":     Var("path"),
			"specifiers": Check(Not(Arr()), Var("names")),
		},
		Obj{
			"Path":  Var("path"),
			"Names": Var("names"),
		},
	),
	MapSemantic("", "FunctionDeclaration", uast.FunctionGroup{}, nil,
		Obj{
			"id":        Var("name"),
			"generator": Var("gen"),   // FIXME: define channels in SDK? or return a function?
			"async":     Var("async"), // TODO: async
			"body":      Var("body"),
			"params":    Var("params"), // FIXME: Ident | AssignmentPattern
		},
		Obj{
			"Nodes": Arr(
				Obj{
					"async":     Var("async"),
					"generator": Var("gen"),
				},
				UASTType(uast.Alias{}, Obj{
					"Name": Var("name"),
					"Node": UASTType(uast.Function{}, Obj{
						"Type": UASTType(uast.FunctionType{}, Obj{
							"Arguments": Var("params"),
							"Returns": Arr(
								UASTType(uast.Argument{}, Obj{
									"Init": Is(uast.Identifier{
										Name: "undefined",
									}),
								}),
							),
						}),
						"Body": Var("body"),
					}),
				}),
			),
		},
	),
}
