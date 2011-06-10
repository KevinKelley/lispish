//
// Stutter!
//
package stutter

import (
	"fmt"
	"os"
	"io"
)

func parse(in io.Reader) SObj { return nil }

func read_lisp() SObj {
	fmt.Print("> ")
	expr := parse(os.Stdin)
	if expr != nil {
		print_lisp(expr)
	}
	return expr
}

func eval_atom(a SObj) SObj {
	if a == error_cell {
		return (error_cell)
	}
	if a.Kind() != AtomTy {
		err("Error: set: first argument is not an atom.")
		return error_cell
	}

	name := a.(AtomCell).name
	cell := context.get(name)

	if cell == nil || cell.(ListCell).cdr == nil {
		err(fmt.Sprintf("Error: unbound atom %s.", name))
		return error_cell
	}
	return cell.(ListCell).cdr
}

func eval_lambda(expr, fn SObj) SObj {
	cell := nil_cell
	argn := fn.(ListCell).car   // arg name
	argv := expr.(ListCell).cdr // arg val

	// Evaluate all of the parameters and hold them in a temporary list.
	// (blist.car is item, blist.cdr is next node)
	blist := nil_cell
	for count := 1; argn != nil && argn != nil_cell; count++ {
		if argv == nil {
			cell = nil_cell
		} else {
			cell = eval_lisp(car(argv))
			if cell == error_cell {
				err("eval_lambda: error evaluating argument #$count")
				break
			}
		}
		bind := makeList(argn.(ListCell).car, cell)
		blist = cons(bind, blist)
		argn = argn.(ListCell).cdr
		if argv != nil {
			argv = argv.(ListCell).cdr
		}
	}

	// Now that all of the formals are evaluated
	// push the bindings into a local evaluation context
	context = new_context(context)

	for blist != nil && blist != nil_cell {
		bind := blist.(ListCell).car
		context.set(bind.(ListCell).car.(AtomCell).name, bind)
		blist = blist.(ListCell).cdr
	}

	if cell != error_cell {
		cell = eval_lisp(fn.(ListCell).cdr)
	}
	// Remove the bindings from this function call.
	context = context.parent

	return cell
}
func eval_lisp(expr SObj) SObj {
	if expr == error_cell {
		return error_cell
	}
	if expr.Kind() == AtomTy {
		return eval_atom(expr)
	}
	if expr.Kind() == ListTy {
		fn := eval_lisp(expr.(ListCell).car)

		if fn == error_cell {
			return error_cell
		}

		if fn.Kind() == VfuncTy {
			a := eval_lisp(car(expr.(ListCell).cdr))
			b := eval_lisp(car(cdr(expr.(ListCell).cdr)))
			vfunc := fn.(VFuncCell).fn
			return vfunc(a, b)
		}
		if fn.Kind() == SfuncTy {
			return fn.(SFuncCell).fn(expr.(ListCell).cdr)
		}
		if fn.Kind() == LambdaTy {
			return eval_lambda(expr, fn)
		}
	}
	return expr
}

func print_lisp(expr SObj) { fmt.Println(expr) }
