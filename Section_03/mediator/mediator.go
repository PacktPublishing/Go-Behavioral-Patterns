package main

import "fmt"

type One struct{}

func (o *One) OnePlus(n interface{}) interface{} {
	switch n.(type) {
	case One:
		return &Two{}
	case Two:
		return &Three{}
	case Three:
		return &Four{}
	case Four:
		return &Five{}
	case Five:
		return &Six{}
	case Six:
		return &Seven{}
	case Seven:
		return &Eight{}
	case Eight:
		return &Nine{}
	case Nine:
		return [2]interface{}{&One{}, &Zero{}}
	default:
		return fmt.Errorf("Number not found")

	}
}

func Sum(a, b interface{}) interface{}{
	switch a := a.(type) {
	case One:
		switch b := b.(type) {
		case One:
			return &Two{}
		case Two:
			return &Three{}
		case int:
			return b+1
		default:
			return fmt.Errorf("Number not found")

		}
	case Two:
		switch b.(type) {
		case One:
			return &Three{}
		case Two:
			return &Four{}
		default:
			return fmt.Errorf("Number not found")

		}
	case int:
		switch b := b.(type) {
		case One:
			return &Three{}
		case Two:
			return &Four{}
		case int:
			return a + b
		default:
			return fmt.Errorf("Number not found")

		}
	default:
		return fmt.Errorf("Number not found")

	}
}


type Two struct{}
type Three struct{}
type Four struct{}
type Five struct{}
type Six struct{}
type Seven struct{}
type Eight struct{}
type Nine struct{}
type Zero struct{}

func main(){
	fmt.Printf("%#v\n", Sum(One{}, Two{}))
	fmt.Printf("%d\n", Sum(1,2))
	fmt.Printf("%d\n", Sum(One{},2))
}