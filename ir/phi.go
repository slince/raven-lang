package ir

import "github.com/slince/php-plus/ir/value"

type Phi struct {
	Variable value.Value
	First    *value.Temporary
	Second   *value.Temporary
	instruction
}
