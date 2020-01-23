package a

func F() { // OK
}

func G() {} // want "a.G is not tested"
