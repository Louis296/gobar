package gobar

import (
	"fmt"
	"testing"
)

func TestLC1024(t *testing.T) {
	//fmt.Println(LC1024([]int{1337,300,13,1},[]byte{'-','-','*'}))
	fmt.Println(LC1024([]int{25, 17, 1, 17, 6, 35, 2, 1337, 14, 30, 1}, []byte{'*', '/', '>'}))
}
