package evaluator

import (
	"math"

	"github.com/darragh-downey/twerp/pkg/object"
)

var builtins = map[string]*object.Builtin{
	"len": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d want=1", len(args))
			}

			switch a := args[0].(type) {
			case *object.String:
				return &object.Integer{Value: int64(len(a.Value))}

			default:
				return newError("argument to `len` not supported, got %s", args[0].Type())
			}
		},
	},
	"max": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) <= 1 {
				return newError("wrong number of arguments. got=%d want 2 or more", len(args))
			}

			max := int64(math.MinInt64)

			for i, a := range args {
				switch b := a.(type) {
				case *object.Integer:
					lhs := a.(*object.Integer)
					if lhs.Value > max {
						max = lhs.Value
					}
				default:
					return newError("argument to `max` not supported, got %s at postition %d", b, i)
				}
			}

			return &object.Integer{Value: int64(max)}
		},
	},
	"min": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) <= 1 {
				return newError("wrong number of arguments. got=%d want 2 or more", len(args))
			}

			max := int64(math.MaxInt64)

			for i, a := range args {
				switch b := a.(type) {
				case *object.Integer:
					lhs := a.(*object.Integer)
					if lhs.Value < max {
						max = lhs.Value
					}
				default:
					return newError("argument to `min` not supported, got %s at postition %d", b, i)
				}
			}

			return &object.Integer{Value: int64(max)}
		},
	},
}
