package usecases

import (
    "github.com/serz999/notesServer/internal/dto"
)

type AddNoteInteractor struct {
    gate NotesGate 
}

func NewAddNoteInteractor(gate NotesGate) *AddNoteInteractor { 
    return &AddNoteInteractor{gate}
}

func (addn *AddNoteInteractor) Exec(note dto.Note) (id dto.Id, err error) {
    return addn.gate.Add(note) 
}
