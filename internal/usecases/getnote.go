package usecases

import (
    "github.com/serz999/notesServer/pkg/dto"
)

type GetNoteByIdInteractor struct {
    gate NotesGate 
}

func NewGetNoteByIdInteractor(gate NotesGate) *GetNoteByIdInteractor { 
    return &GetNoteByIdInteractor{gate}
}

func (getn *GetNoteByIdInteractor) Exec(id dto.Id) (dto.Note, error) {
    return getn.gate.GetById(id) 
}

func (getn *GetNoteByIdInteractor) GetNotFoundMsg() string {
    return getn.gate.GetNotFoundMsg()
}
