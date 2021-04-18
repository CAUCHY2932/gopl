package bits

import (
	// "reflect"
	"testing"
)

// // % 08b
// func TestNot(t *testing.T) {

// 	testCase := []struct {
// 		input  uint32
// 		expect uint32
// 	}{
// 		{0x1234, 0xedca},
// 	}
// 	for i, tt := range testCase {
// 		got := Not(tt.input)

// 		if !reflect.DeepEqual(got, tt.expect) {
// 			t.Errorf("[t] idx: %d, input %0b expect %0b, but got %0b, self %0b\n", i, tt.input, tt.expect, got, -tt.input)
// 		}
// 	}

// }
func TestGetLastOne(t *testing.T) {
	var x uint32 = 0x1234
	var expect uint32 = 0x0004
	got := GetLastOne(x)
	// t.Errorf("==========[t] x %0b not x %0b\n", x, Not(x))

	if got != expect {
		t.Errorf("==========[t] x %0b not x %0b\n", x, Not(x))
		t.Errorf("[t] expect %0b, but got %0b", 0x0004, got)
	}
}
