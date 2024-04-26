package list

import (
	"fmt"
	"testing"
)

func TestList(t *testing.T) {
	l := New[int32]()
	l.PushBack(1)
	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}
