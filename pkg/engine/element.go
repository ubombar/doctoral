package engine

import (
	"fmt"
)

// This is essentially the same as the string but it represents a valid regex.
type regexstring string

// This is the match expression interface. It contains the methods to linearlize
// the regex implied by this.
type matchableElement interface {
}

// This is the match expression
type MEList struct {
	children []matchableElement
}

// Constant expression
type MEConst struct {
	value string
}

// Match any as in name etc.
type MEAny struct {
	inregex regexstring
}

type MEOr struct {
	children []matchableElement
}

func List(m ...matchableElement) matchableElement {
	return &MEList{children: m}
}

func Or(m ...matchableElement) matchableElement {
	return &MEOr{children: m}
}

func Const(s string) matchableElement {
	return &MEConst{value: s}
}

func Tag(m ...matchableElement) matchableElement {
	return List(Const("#"), List(m...))
}

func It1(m ...matchableElement) matchableElement {
	return List(Const("_"), List(m...), Const("_"))
}

func It2(m ...matchableElement) matchableElement {
	return List(Const("*"), List(m...), Const("*"))
}

func Bo1(m ...matchableElement) matchableElement {
	return List(Const("__"), List(m...), Const("__"))
}

func Bo2(m ...matchableElement) matchableElement {
	return List(Const("**"), List(m...), Const("**"))
}

func H1(m ...matchableElement) matchableElement {
	return List(Const("# "), List(m...))
}

func H2(m ...matchableElement) matchableElement {
	return List(Const("## "), List(m...))
}

func H3(m ...matchableElement) matchableElement {
	return List(Const("### "), List(m...))
}

func H4(m ...matchableElement) matchableElement {
	return List(Const("#### "), List(m...))
}

func H5(m ...matchableElement) matchableElement {
	return List(Const("##### "), List(m...))
}

func Linearlize(e interface{}, nl bool) ([]regexstring, error) {
	yield := []regexstring{}

	switch v := e.(type) {
	case *MEList:
		for _, child := range v.children {
			childYield, err := Linearlize(child, false)

			if err != nil {
				return []regexstring{}, err
			}

			yield = append(yield, childYield...)
		}

		if nl {
			nlYield, err := Linearlize(Const("\n"), false)

			if err != nil {
				return []regexstring{}, err
			}

			yield = append(yield, nlYield...)
		}

		return yield, nil
	case *MEConst:
		yield = append(yield, regexstring(v.value))

		if nl {
			nlYield, err := Linearlize(Const("\n"), false)

			if err != nil {
				return []regexstring{}, err
			}

			yield = append(yield, nlYield...)
		}
	case *MEAny:
		yield = append(yield, regexstring(v.inregex)) // This is a VAR so do more stuff in the future

		if nl {
			nlYield, err := Linearlize(Const("\n"), false)

			if err != nil {
				return []regexstring{}, err
			}

			yield = append(yield, nlYield...)
		}
	case *MEOr:
		leftPar, err := Linearlize(Const("("), false)

		if err != nil {
			return []regexstring{}, err
		}

		rightPar, err := Linearlize(Const(")"), false)

		if err != nil {
			return []regexstring{}, err
		}

		or, err := Linearlize(Const("|"), false)

		if err != nil {
			return []regexstring{}, err
		}

		yield = append(yield, leftPar...)

		for i, child := range v.children {
			childYield, err := Linearlize(child, false)

			if err != nil {
				return []regexstring{}, err
			}

			yield = append(yield, childYield...)

			// Add the or except to the last step
			if i < len(v.children)-1 {
				yield = append(yield, or...)
			}
		}

		yield = append(yield, rightPar...)

		if nl {
			nlYield, err := Linearlize(Const("\n"), false)

			if err != nil {
				return []regexstring{}, err
			}

			yield = append(yield, nlYield...)
		}

		return yield, nil
	default:
		return []regexstring{}, fmt.Errorf("cannot linearize unknown type: %v", v)
	}

	return yield, nil
}
