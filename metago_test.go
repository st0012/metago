package metago

import "testing"

const(
	attrErrorFormat = "Expect subject's attribute '%[1]s' to be: '%[2]s'. got: %[3]v(%[3]T)"
	expectResultErrorFormat = "Expect result to be: '%[1]s'. got: %[2]v(%[2]T)"
)

type Bar struct {
	ptrName string
	valueName string
}

func (b *Bar) PtrName() string {
	return b.ptrName
}

func (b *Bar) SetPtrName(n string) {
	b.ptrName = n
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
		t.Errorf(attrErrorFormat, "ptrName" ,"John", b.ptrName)
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