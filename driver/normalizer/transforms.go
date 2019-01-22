package normalizer

import "gopkg.in/bblfsh/sdk.v2/driver"

var Transforms = driver.Transforms{
	Namespace:      "javascript",
	Preprocess:     Preprocess,
	PreprocessCode: PreprocessCode,
	Normalize:      Normalize,
	Annotations:    Native,
	Code:           Code,
}
