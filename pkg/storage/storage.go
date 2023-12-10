package storage

import "errors"

// Storage - интерфейс, представляющий обобщенное хранилище данных.
type Storage interface {
	Len() int64
	Add(value interface{}) (int64, error)
	RemoveByIndex(id int64)
	RemoveByValue(value interface{})
	RemoveAllByValue(value interface{})
	GetByIndex(id int64) (interface{}, bool)
	GetByValue(value interface{}) (int64, bool)
	GetAllByValue(value interface{}) ([]int64, bool)
	GetAll() ([]interface{}, bool)
	Clear()
	Print()
}

// ErrMismatchType ошибка, возвращаемая методом Add, если тип добавляемого элемента
// не соответствует типу уже присутствующих в хранилище элементов.
var ErrMismatchType = errors.New("mismatched type: the type of the provided value does not match the type of items already in the storage")
