package ch11

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	re := add(1, 3)
	if re != 4 {
		t.Errorf("expect %d, actual %d", 4, re)
	}
}

func TestAdd2(t *testing.T) {
	fmt.Println("yes1")
	if testing.Short() {
		t.Skip("short模式下跳过单元测试")
	}
	fmt.Println("yes")

	re := add(1, 5)
	if re != 6 {
		t.Errorf("expect %d, actual %d", 3, re)
	}
}

func TestAdd3(t *testing.T) {
	var dataset = []struct {
		a   int
		b   int
		out int
	}{
		{1, 1, 2},
		{8, 8, 16},
		{-9, 8, -1},
		{0, 0, 1},
	}

	for _, value := range dataset {
		re := add(value.a, value.b)
		if re != value.out {
			t.Errorf("expect:%d, actual:%d", value.out, re)
		}
	}
}
