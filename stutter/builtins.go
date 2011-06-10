
package stutter

// built-in functions are either "value" functions: args passed by value;
// or "special" functions: args passed unevaluated, and the function may do
// what it likes.
//

/////////////////////////////////////////////////////////////////////////////
// "value" functions: args pass by value
//
func car(a SObj) SObj {
	if a == error_cell {
		return error_cell
	}
	if a == nil_cell {
		return nil_cell
	}
	if a.Kind() != ListTy {
		err("Error: car: argument is not a list.")
		return error_cell
	}
	return a.(*ListCell).Car()
}

func cdr(a SObj) SObj {
	if a == error_cell {
		return error_cell
	}
	if a == nil_cell {
		return nil_cell
	}
	if a.Kind() != ListTy {
		err("Error: cdr: argument is not a list.")
		return error_cell
	}
	//if a.(ListCell).cdr == nil {
	//	return nil_cell
	//}
	return a.(ListCell).cdr
}
func cons(a, b SObj) SObj {
	if a == error_cell || b == error_cell {
		return error_cell
	}
	if !(b.Kind() == ListTy || b == nil_cell) {
		err("Error: cons: second argument is not a list.")
		return error_cell
	}
	if b == nil_cell {
		return makeList(a, nil)
	}
	return makeList(a, b)
}

func set(a, b SObj) SObj {
	if a == error_cell || b == error_cell {
		return error_cell
	}
	if a.Kind() != AtomTy {
		err("Error: set: first argument is not an atom.")
		return error_cell
	}
	bind := context.get(a.(*AtomCell).name)
	if bind.Kind() == ListTy {
		if bound, ok := bind.(ListCell); ok {
			bound.SetCar(a) // context maps name to a list whose
			bound.SetCdr(b) // car is an atom, and whose cdr is its value
		} else {
			err("Error: set: binder isn't a list")
		}
		return b
	}

	// Not found so insert it in context.
	bind = makeList(a, b)
	context.set(a.(AtomCell).name, bind)

	return b
}

func equal(a, b SObj) SObj {
	if a != nil && a.Kind() == AtomTy && a == b {
		return true_cell
	}
	return nil_cell
}

/////////////////////////////////////////////////////////////////////////////
// "special" functions: args passed unevaluated
//
func quote(expr SObj) SObj { return car(expr) }

func lambda(expr SObj) SObj {
	args := car(expr)
	body := car(cdr(expr))
	if args == error_cell || body == error_cell {
		return error_cell
	}
	ufunc := makeLambda(args, body)

	// Check to make sure that the formal argument list is a simple list.
	if args != nil_cell {
		for args != nil {
			if args.Kind() != ListTy {
				err("Error: bad argument list supplied.")
				return error_cell
			}
			if args.(ListCell).car.Kind() != AtomTy {
				err("Error: bad argument list supplied.")
				return error_cell
			}
			args = args.(ListCell).cdr
		}
	}
	return ufunc
}

func lisp_if(expr SObj) SObj {
	bool_expr := car(expr)
	then_expr := car(cdr(expr))
	else_expr := car(cdr(cdr(expr)))

	if bool_expr == error_cell || then_expr == error_cell || else_expr == error_cell {
		return error_cell
	}

	bool_rslt := eval_lisp(bool_expr)
	if bool_rslt != nil_cell {
		return eval_lisp(then_expr)
	}
	return eval_lisp(else_expr)
}
