package metago

import "testing"

const (
	attrErrorFormat         = "Expect subject's attribute '%[1]s' to be: '%[2]v'. got: %[3]v(%[3]T)"
	expectResultErrorFormat = "Expect result to be: '%[1]v'. got: %[2]v(%[2]T)"
)

type Bar struct {
	ptrName   string
	valueName string
}

func (b *Bar) PtrName() string {
	return b.ptrName
}

func (b *Bar) SetPtrName(n string) {
	b.ptrName = n
}

func (b *Bar) Foo() (string, int) {
	return "foo", 100
}

func (b Bar) ValueName() string {
	return b.valueName
}

func (b Bar) SetValueName(n string) Bar {
	b.valueName = n
	return b
}

func (b *Bar) Send(methodName string, args ...interface{}) interface{} {
	return CallFunc(b, methodName, args...)
}

func TestCallFuncWithPtrReceiver(t *testing.T) {
	b := &Bar{}
	b.Send("SetPtrName", "John")

	if b.ptrName != "John" {
		t.Errorf(attrErrorFormat, "ptrName", "John", b.ptrName)
	}

	n := b.Send("PtrName")

	if n != "John" {
		t.Errorf(expectResultErrorFormat, "John", n)
	}
}

func TestCallFuncWithValueReceiver(t *testing.T) {
	b := Bar{}
	result := b.Send("SetValueName", "John")

	b, ok := result.(Bar)

	if !ok {
		t.Fatalf("Expect 'SetValueName' results a %T object. got: %T", b, result)
	}

	if b.valueName != "John" {
		t.Fatalf(attrErrorFormat, "valueName", "John", b.valueName)
	}

	n := b.valueName

	if n != "John" {
		t.Fatalf(expectResultErrorFormat, "John", n)
	}
}

func TestCallFuncWithMultipleReturnValue(t *testing.T) {
	b := &Bar{}
	result := b.Send("Foo")
	results := result.([]interface{})
	s := results[0]

	if s != "foo" {
		t.Fatalf(expectResultErrorFormat, "foo", s)
	}

	i := results[1]

	if i != 100 {
		t.Fatalf(expectResultErrorFormat, 100, i)
	}
}
