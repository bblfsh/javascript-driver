package normalizer

import (
	"gopkg.in/bblfsh/sdk.v1/uast"
)

// ToNode is an instance of `uast.ObjectToNode`, defining how to transform an
// into a UAST (`uast.Node`).
//
// https://godoc.org/gopkg.in/bblfsh/sdk.v1/uast#ObjectToNode
var ToNode = &uast.ObjectToNode{
	TopLevelIsRootNode: true,
	InternalTypeKey:    "type",
	OffsetKey:          "start",
	EndOffsetKey:       "end",
	LineKey:            "loc.start.line",
	EndLineKey:         "loc.end.line",
	ColumnKey:          "loc.start.column",
	EndColumnKey:       "loc.end.column",
	TokenKeys: map[string]bool{
		"value": true,
		"name":  true,
	},
	// Several maps with properties are returned by the AST, we should use only
	// those with the key "type"
	IsNode: func(v map[string]interface{}) bool {
		_, ok := v["type"].(string)
		return ok
	},
	// The column value from babylon are based on 0, so we always sum one.
	Modifier: func(n map[string]interface{}) error {
		loc, ok := n["loc"].(map[string]interface{})
		if !ok {
			return nil
		}

		for _, key := range []string{"start", "end"} {
			p, ok := loc[key].(map[string]interface{})
			if !ok {
				continue
			}

			if c, ok := p["column"].(float64); ok {
				p["column"] = c + 1
			}
		}

		return nil
	},
}
