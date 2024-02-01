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
