package usecases

import (
    "github.com/serz999/notesServer/pkg/dto"
)

type NotesGate interface {
    Add(note dto.Note) (int64, error); 
    GetById(id int64) (dto.Note, error);
    Del(id int64) error; 
}

type AddNote interface {
    Exec(note dto.Note) (int64, error); 
}

type DelNote interface { 
    Exec(id int64) error;
}

type GetNoteById interface {
    Exec(id int64) (dto.Note, error); 
}
