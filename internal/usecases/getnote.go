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

func (getn *GetNoteByIdInteractor) Exec(id int64) (dto.Note, error) {
    return getn.gate.GetById(id) 
}
