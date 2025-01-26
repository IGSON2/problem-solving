package reallyspecialsubtree

type interfaceA interface {
	add()
	mul()
}

type interfaceB interface {
	interfaceA
	sub()
}

type testcase struct {
	a, b int
}

func (t testcase) add()  {}
func (t *testcase) sub() {}
func (t testcase) mul()  {}

func testInit(t interfaceB) {}

func test() {
	// 인터페이스 A를 인자로 받는 함수에서 객체 t가 포인터를 리시버로 구현하는 메서드가 하나라도 있다면 인터페이스 A의 인자로 &t를 제공하여야 함.
	// 만약 포인터를 리시버로 구현하는 메서드가 하나도 없더라도 &t를 매개변수로 제공 가능.
	t := testcase{1, 1}
	testInit(&t)
	// testInit(t) -> error
}

type interfaceC interface {
	add2()
	mul2()
}

type testcase2 struct {
	a, b int
}

func (t testcase2) add2() {}
func (t testcase2) mul2() {}

func test2Init(t interfaceC) {}

func test2() {
	t := testcase2{1, 1}
	test2Init(t)
	test2Init(&t)
}
