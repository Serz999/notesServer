package usecases

import (
    "github.com/serz999/notesServer/pkg/dto"
)

type NotesGate interface {
    Add(note dto.Note) (dto.Id, error); 
    GetById(id dto.Id) (dto.Note, error);
    Del(id dto.Id) error; 
    GetNotFoundMsg() string;
}

type AddNote interface {
    Exec(note dto.Note) (dto.Id, error); 
}

type DelNote interface { 
    Exec(id dto.Id) error;
    GetNotFoundMsg() string;
}

type GetNoteById interface {
    Exec(id dto.Id) (dto.Note, error); 
    GetNotFoundMsg() string;
}
