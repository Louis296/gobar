package gobar

import (
	"fmt"
	"testing"
)

func TestLC1024(t *testing.T) {
	//fmt.Println(LC1024([]int{1337, 300, 13, 1}, []byte{'-', '-', '*'}))
	//fmt.Println(LC1024([]int{16, 2, 16, 0}, []byte{'<', '*', '+'}))
	//fmt.Println(LC1024([]int{15, 2, 0, 10}, []byte{'%', '>', '<'}))
	//fmt.Println(LC1024([]int{2, 7, 2, 8}, []byte{'&', 'p', '<'}))
	//fmt.Println(LC1024([]int{2, 7, 2, 1337, 2, 5, 8, 35}, []byte{'&', 'p', '<'}))
	fmt.Println(LC1024([]int{25, 17, 1, 17, 6, 35, 2, 1337, 14, 30, 1}, []byte{'*', '/', '>', '*'}))
}
