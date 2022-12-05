package evaluator

import (
	"fmt"
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

			case *object.Array:
				return &object.Integer{Value: int64(len(a.Elements))}

			default:
				return newError("argument to `len` not supported, got %s", args[0].Type())
			}
		},
	},
	"max": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) < 1 {
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

				case *object.Array:
					for _, e := range b.Elements {
						if ee, ok := e.(*object.Integer); !ok {
							continue
						} else if ee.Value > max {
							max = ee.Value
						}
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

			min := int64(math.MaxInt64)

			for i, a := range args {
				switch b := a.(type) {
				case *object.Integer:
					lhs := a.(*object.Integer)
					if lhs.Value < min {
						min = lhs.Value
					}

				case *object.Array:
					for _, e := range b.Elements {
						if ee, ok := e.(*object.Integer); !ok {
							continue
						} else if ee.Value < min {
							min = ee.Value
						}
					}

				default:
					return newError("argument to `min` not supported, got %s at postition %d", b, i)
				}
			}

			return &object.Integer{Value: int64(min)}
		},
	},
	"first": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d want=1", len(args))
			}

			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `first` must be of type ARRAY, got %s", args[0].Type())
			}

			arr := args[0].(*object.Array)
			if len(arr.Elements) > 0 {
				return arr.Elements[0]
			}

			return NULL
		},
	},
	"last": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d want=1", len(args))
			}

			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `last` must be of type ARRAY, got %s", args[0].Type())
			}

			arr := args[0].(*object.Array)
			if len(arr.Elements) > 0 {
				return arr.Elements[len(arr.Elements)-1]
			}

			return NULL
		},
	},
	"rest": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d want=1", len(args))
			}

			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `rest` must be of type ARRAY, got %s", args[0].Type())
			}

			arr := args[0].(*object.Array)
			length := len(arr.Elements)

			if length > 0 {
				newElements := make([]object.Object, length-1, length-1)
				copy(newElements, arr.Elements[1:length])
				return &object.Array{Elements: newElements}
			}

			return NULL
		},
	},
	"push": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d want=2", len(args))
			}

			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `push` must be of type ARRAY, got %s", args[0].Type())
			}

			arr := args[0].(*object.Array)
			length := len(arr.Elements)

			newElements := make([]object.Object, length-1, length-1)
			copy(newElements, arr.Elements)
			newElements[length] = args[1]

			return &object.Array{Elements: newElements}
		},
	},
	"puts": {
		// variadic function
		Fn: func(args ...object.Object) object.Object {
			for _, arg := range args {
				fmt.Println(arg.Inspect())
			}
			return NULL
		},
	},
}
