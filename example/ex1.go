package main

import (
	"fmt"
	"github.com/kelleysoft/Lispish/stutter"
)

func main() {
  fmt.Println("sqwauk!")
  stutter.Init()
}
/*
  Void main()
  {
    expr:= SObj<|'(here is a list (with a list (with a list) and another (list)))|>
    echo(">> $expr")

    expr = SObj<|'(a b c)|>
    echo(">> $expr")

    expr = SObj
    <| ;sample stutter script
      'my-name-is-gary
      '(this    is    a better way of writing a string of text)
      '(here is a list(with a list(with a list)and another(list)))
      ()
      nil
      ;this-is-an-unquoted-undefined-atom
      (choke on this list)
      ; comments are text from semi to eol
      (set 'name 'gary)
      name
      (set 'ten '10)
      ten
      (set ten '5-plus-5)
      ten
      10
      ;
      (car '((a b c) x y))
      (cdr '((a b c) x y))
      (car (car '((a b c) x y)))
      (cdr (cdr '((a b c) x y)))
      (car (cdr (cdr (car '((a b c) x y)))))
      (cons 'a nil)
      (cons 'a '(b))
      (cons '(a b c) '(x y))
      (cons '(a b c) (ten))
     |>
    echo(">> $expr")
  }
*/