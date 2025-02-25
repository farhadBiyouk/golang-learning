package generic

import (
	"github.com/google/go-cmp/cmp"
)

type List[T any] struct {
	Items []T
}

func (list *List[T]) Add(item T) {
	list.Items = append(list.Items, item)
}

func (list *List[T]) InsertAt(index int, item T) {
	list.Items = append(list.Items, item)
	copy(list.Items[index+1:], list.Items[index:])
	list.Items[index] = item
}
func (list *List[T]) RemoveAt(index int) {
	list.Items = append(list.Items[:index], list.Items[index+1:]...)
}
func (list *List[T]) Remove(item T) {
	index := list.Find(item)
	if index != -1 {
		list.RemoveAt(index)
	}

}
func (list *List[T]) Get(index int) T {
	return list.Items[index]
}

func (list *List[T]) Find(item T) int {
	for index, value := range list.Items {
		if cmp.Equal(value, item) {
			return index
		}
	}
	return -1
}
