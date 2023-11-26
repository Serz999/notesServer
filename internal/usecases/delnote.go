package usecases

import (
    "github.com/serz999/notesServer/pkg/dto"
)

type DelNoteInteractor struct {
    gate NotesGate 
}

func NewDelNoteInteractor(gate NotesGate) *DelNoteInteractor { 
    return &DelNoteInteractor{gate}
}

func (deln *DelNoteInteractor) Exec(id dto.Id) (err error) {
    return deln.gate.Del(id) 
}

func (deln *DelNoteInteractor) GetNotFoundMsg() string {
    return deln.gate.GetNotFoundMsg()
}
