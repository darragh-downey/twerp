package object

// Environment keeps track of of variables and functions within its scope
type Environment struct {
	store map[string]Object // a mapping for variables and functions
	outer *Environment      // a pointer to an enclosing environment (think function scope)
}

func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s, outer: nil}
}

func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}

// Get will attempt to locate the value for the given identifier in the current and surrounding/enclosing scope
func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}
