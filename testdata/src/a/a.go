package a

func F1() {} // OK

func F2() {} // want `a\.F2 is not tested`

func f3() {} // OK

type T struct{}

func (T) M1() {} // OK

func (T) M2() {} // want `\(a\.T\)\.M2 is not tested`

func (T) m3() {} // OK
