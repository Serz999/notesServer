package usecases

type DelNoteInteractor struct {
    gate NotesGate 
}

func NewDelNoteInteractor(gate NotesGate) *DelNoteInteractor { 
    return &DelNoteInteractor{gate}
}

func (deln *DelNoteInteractor) Exec(id int64) (err error) {
    return deln.gate.Del(id) 
}
