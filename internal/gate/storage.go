package gate

import (
	"github.com/serz999/notesServer/pkg/dto"
	"github.com/serz999/notesServer/pkg/storage"
)

type StorageGate struct {
    s storage.Storage
}

func NewStorageGate(s storage.Storage) (*StorageGate, error) {
    return &StorageGate{s}, nil
}

func (g *StorageGate) Add(note dto.Note) (int64, error) {
    id, err := g.s.Add(note)
    return id, err
} 

func (g *StorageGate) GetById(id int64) (dto.Note, error) {
    v, ok := g.s.GetByIndex(id)
    if !ok {
        return dto.Note{}, &dto.NotFoundErr{}
    }

    return v.(dto.Note), nil
}

func (g *StorageGate) Del(id int64) error {
    _, ok := g.s.GetByIndex(id)
    if !ok {
        return &dto.NotFoundErr{}
    }

    g.s.RemoveByIndex(id) 
    return nil
} 
