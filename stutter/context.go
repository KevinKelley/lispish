package stutter

type Context struct {
	parent  *Context        //{ get; private set }
	entries map[string]SObj //Str:SObj entries := Str:SObj[:]
}

// create variable context.
func new_context(enclosing *Context) *Context {
	return &Context{parent: enclosing, entries: map[string]SObj{}}
}

// get variable binding from a ctx (local only)
// returns null if variable doesn't exist
func (self *Context) retr(name string) (val SObj) {
	if val, ok := self.entries[name]; ok {
		return val
	}
	return nil
}

// get variable binding from ctx (with parent fallback)
// returns null if variable doesn't exist
func (self *Context) get(name string) SObj {
	if val, present := self.entries[name]; present {
		return val
	}
	return self.parent.get(name)
}

// set/store a variable binding in ctx
// (updates if the var already exists)
// NOTE val should represent a binding of name to value
func (self *Context) set(name string, val SObj) {
	/*assert(val.car.name == name)*/
	self.entries[name] = val
}


// 'unique' the atom: check context; if it is not found then a new atom with the
// supplied name is created and inserted into this context.
//
// at name, store a binding of an atom.
// Create an atom for the given name;
// wrap it in a list object;
// store the list (bound atom) in context;
// return the atom.
func (self *Context) new_atom(name string) SObj {
	bind := self.retr(name) // get if locally declared

	if bind != nil {
		cell := bind.(*ListCell)
		return cell.Car() // return the atom (the car of the variable binding)
	}
	// Not found so insert it in this context.

	// wrap the atom in a list and store it
	atom := makeAtom(name)
	bind = makeList(atom, nil) // { it.car = atom; it.cdr = null }

	self.set(name, bind)

	return atom
}
