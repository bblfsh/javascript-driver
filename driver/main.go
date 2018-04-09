package main

import (
	_ "github.com/bblfsh/javascript-driver/driver/impl"
	"github.com/bblfsh/javascript-driver/driver/normalizer"

	"gopkg.in/bblfsh/sdk.v1/sdk/driver"
)

func main() {
	driver.Run(driver.Transforms{
		Native: normalizer.Native,
		Code:   normalizer.Code,
	})
}
