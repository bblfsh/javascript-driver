package normalizer

import (
	"gopkg.in/bblfsh/sdk.v1/protocol/driver"
	"gopkg.in/bblfsh/sdk.v1/protocol/native"
)

// ToNoder specifies the driver options. Driver programmers should fill it
var ToNoder = &native.ObjectToNoder{
	InternalTypeKey:    "type",
	OffsetKey:          "start",
	EndOffsetKey:       "end",
	TopLevelIsRootNode: true,
	TokenKeys:          map[string]bool{},
	PositionFill:       native.LineColFromOffset,
}

// ParserBuilder creates a parser that transform source code files into *uast.Node.
func ParserBuilder(opts driver.ParserOptions) (parser driver.Parser, err error) {
	parser, err = native.ExecParser(ToNoder, opts.NativeBin)
	if err != nil {
		return
	}

	switch ToNoder.PositionFill {
	case native.OffsetFromLineCol:
		parser = &driver.TransformationParser{
			Parser:         parser,
			Transformation: driver.FillOffsetFromLineCol,
		}
	case native.LineColFromOffset:
		parser = &driver.TransformationParser{
			Parser:         parser,
			Transformation: driver.FillLineColFromOffset,
		}
	}

	return
}
