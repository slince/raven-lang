package vm

import (
	"github.com/slince/php-plus/vm/object"
)

type BuiltinFunction func(values ...object.Object) (object.Object, error)
