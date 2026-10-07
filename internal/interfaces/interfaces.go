package interfaces

type MyErr struct{}

func (*MyErr) Error() string { return "boom" }

func do() error {
	var e *MyErr = nil
	return e // iface: {type: *MyErr, value: nil}
}

func doNil() error {
	return nil
}

func Run() bool {
	return do() == nil
}

func RunNil() bool {
	return doNil() == nil
}

// embedding

type Base struct{}

func (Base) Name() string    { return "base" }
func (b Base) Hello() string { return "hi " + b.Name() }

type Child struct{ Base }

func (Child) Name() string { return "child" }

func Embedding() {
	Child{}.Hello() // "hi base" - not "hi child"!
	Child{}.Base.Name()
}

// alignment

type Bad struct { // 24 bytes
	A bool  // 1 + 7 padding
	B int64 // 8
	C bool  // 1 + 7 padding
}

type Good struct { // 16 bytes
	B    int64
	A, C bool // 2 + 6 padding
}
