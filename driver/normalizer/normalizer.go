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

	Map(
		Part("_", Obj{"loc": AnyNode(nil)}),
		Part("_", Obj{}),
	),
	Map(
		Part("_", Obj{"extra": AnyNode(nil)}),
		Part("_", Obj{}),
	),
}

// Normalizers is the main block of normalization rules to convert native AST to semantic UAST.
var Normalizers = []Mapping{
	MapSemantic("Identifier", uast.Identifier{}, MapObj(
		Obj{
			"name": Var("name"),
		},
		Obj{
			"Name": Var("name"),
		},
	)),
	MapSemantic("StringLiteral", uast.String{}, MapObj(
		Obj{
			"value": Var("val"),
		},
		Obj{
			"Value": Var("val"),
		},
	)),
	MapSemantic("CommentLine", uast.Comment{}, MapObj(
		Obj{
			"value": CommentText([2]string{"", ""}, "comm"),
		},
		CommentNode(false, "comm", nil),
	)),
	MapSemantic("CommentBlock", uast.Comment{}, MapObj(
		Obj{
			"value": CommentText([2]string{"", ""}, "comm"),
		},
		CommentNode(true, "comm", nil),
	)),
	MapSemantic("BlockStatement", uast.Block{}, MapObj(
		Obj{
			"body":       Var("stmts"),
			"directives": Arr(), // TODO: find an example
		},
		Obj{
			"Statements": Var("stmts"),
		},
	)),
	MapSemantic("ImportDeclaration", uast.Import{}, MapObj(
		Obj{
			"source":     Var("path"),
			"specifiers": Arr(),
		},
		Obj{
			"Path": Var("path"),
		},
	)),
	MapSemantic("ImportDeclaration", uast.Import{}, MapObj(
		CasesObj("case",
			// common
			Obj{
				"importKind": String("value"),
				"source":     Var("path"),
			},
			Objs{
				// namespace
				{
					"specifiers": ArrWith(Var("names"), Obj{
						uast.KeyType: String("ImportNamespaceSpecifier"),
						uast.KeyPos:  Var("local_pos"),
						"local":      Var("local"),
					}),
				},
				// normal import
				{
					"specifiers": Check(Not(Arr()), Var("names")),
				},
			},
		),
		CasesObj("case", nil,
			Objs{
				// namespace
				{
					"Path": UASTType(uast.Alias{}, Obj{
						uast.KeyPos: Var("local_pos"),
						"Name":      Var("local"),
						"Node":      Var("path"),
					}),
					"Names": Var("names"),
					"All":   Bool(true),
				},
				// normal import
				{
					"Path":  Var("path"),
					"Names": Var("names"),
					"All":   Bool(false),
				},
			},
		),
	)),
	MapSemantic("ImportSpecifier", uast.Alias{}, MapObj(
		Obj{
			"importKind": Is(nil),
			"local":      Var("local"),
			"imported":   Var("imp"),
		},
		Obj{
			"Name": Var("local"),
			"Node": Var("imp"),
		},
	)),
	MapSemantic("ImportDefaultSpecifier", uast.Alias{}, MapObj(
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
	)),
	MapSemantic("FunctionDeclaration", uast.FunctionGroup{}, MapObj(
		Obj{
			"id":        Var("name"),
			"generator": Var("gen"),   // FIXME: define channels in SDK? or return a function?
			"async":     Var("async"), // TODO: async
			"body":      Var("body"),
			"params": Each("params", Cases("param_case",
				// Identifier
				Check(
					HasType(uast.Identifier{}),
					Var("arg_name"),
				),
				// AssignmentPattern
				Obj{
					uast.KeyType: String("AssignmentPattern"),
					uast.KeyPos:  Var("arg_pos"),
					"left":       Var("arg_name"),
					"right":      Var("arg_init"),
				},
				// RestElement
				Obj{
					uast.KeyType: String("RestElement"),
					uast.KeyPos:  Var("arg_pos"),
					"argument":   Var("arg_name"),
				},
			)),
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
							"Arguments": Each("params", Cases("param_case",
								// Identifier
								UASTType(uast.Argument{}, Obj{
									"Name": Var("arg_name"),
								}),
								// AssignmentPattern
								UASTType(uast.Argument{}, Obj{
									uast.KeyPos: Var("arg_pos"),
									"Name":      Var("arg_name"),
									"Init":      Var("arg_init"),
								}),
								// RestElement
								UASTType(uast.Argument{}, Obj{
									uast.KeyPos: Var("arg_pos"),
									"Name":      Var("arg_name"),
									"Variadic":  Bool(true),
								}),
							)),
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
	)),
}
