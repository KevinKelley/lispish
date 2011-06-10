// Copyright 2011 Kevin Kelley. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package stutter

import (
	"fmt"
)

type Err struct{ Msg string }
func (e Err) String() string { return e.Msg }
func err(msg string)         { fmt.Println(msg) }

// ...
//
type SType byte // type tag

type FUNC1 func(a SObj) SObj
type FUNC2 func(a, b SObj) SObj

const (
	LambdaTy SType = iota
	SfuncTy
	VfuncTy
	ListTy
	AtomTy
)

type SObj interface { Kind() SType }

type   AtomCell struct { name string }
type   ListCell struct { car, cdr SObj }
type  SFuncCell struct { fn FUNC1 }
type  VFuncCell struct { fn FUNC2 }
type LambdaCell struct { args, body SObj }

func (my   AtomCell) String() string { return my.name }
func (my   ListCell) String() string { return fmt.Sprintf("(%v . %v)", my.car, my.cdr) }
func (my  SFuncCell) String() string { return "<internal-special-function>" }
func (my  VFuncCell) String() string { return "<internal-value-function>" }
func (my LambdaCell) String() string { return fmt.Sprintf("(lambda %v %v)", my.args, my.body) }

func (  AtomCell) Kind() SType { return AtomTy   }
func (  ListCell) Kind() SType { return ListTy   }
func ( SFuncCell) Kind() SType { return SfuncTy  }
func ( VFuncCell) Kind() SType { return VfuncTy  }
func (LambdaCell) Kind() SType { return LambdaTy }

func (my *ListCell) Car() SObj     { return my.car }
func (my *ListCell) Cdr() SObj     { return my.cdr }
func (my *ListCell) SetCar(a SObj) { my.car = a }
func (my *ListCell) SetCdr(a SObj) { my.cdr = a }


func makeAtom (name string)  *AtomCell { return &AtomCell{ name: name } }
func makeList  (a, b SObj)   *ListCell { return &ListCell{ car: a, cdr: b } }
func makeLambda(a, b SObj) *LambdaCell { return &LambdaCell{ args: a, body: b } }
func makeVFunc2(fn FUNC2)   *VFuncCell { return &VFuncCell{ fn: fn } }
func makeSFunc1(fn FUNC1)   *SFuncCell { return &SFuncCell{ fn: fn } }

// unique an atom in context
func new_atom(name string) SObj { return context.new_atom(name) }

// "globals": these should be accessed thru current-context,
// which is the execution environment; and the env/context
// should be passed as param.
var context = new_context(nil)

var nil_cell   = new_atom("nil")
var true_cell  = new_atom("t")
var error_cell = new_atom("<error>")
var quote_cell = new_atom("quote")

func Init() {
    // assign values to the built-in atoms.  Nil, T, and <error>
    // evaluate to themselves.
    set(nil_cell, nil_cell)
    set(true_cell, true_cell)
    set(error_cell, error_cell)

    // create stutter's built-in functions, and assign them
    // to appropriate names.  eval_lisp will look up the name,
    // get the function, and evaluate it.
    set(new_atom("car"   ), makeVFunc2(func (a,b SObj) SObj { return car(a)     }))
    set(new_atom("cdr"   ), makeVFunc2(func (a,b SObj) SObj { return cdr(a)     }))
    set(new_atom("cons"  ), makeVFunc2(func (a,b SObj) SObj { return cons(a,b)  }))
    set(new_atom("set"   ), makeVFunc2(func (a,b SObj) SObj { return set(a,b)   }))
    set(new_atom("equal" ), makeVFunc2(func (a,b SObj) SObj { return equal(a,b) }))
    set(quote_cell,         makeSFunc1(func (a   SObj) SObj { return quote(a)   }))
    set(new_atom("lambda"), makeSFunc1(func (a   SObj) SObj { return lambda(a)  }))
    set(new_atom("if"    ), makeSFunc1(func (a   SObj) SObj { return lisp_if(a) }))
}
