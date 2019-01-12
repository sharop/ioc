package ioc

import "github.com/gopub/log"

var logger log.Logger

func init() {
	logger = log.Default().Derive("IoC")
}
