package contrellers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"github.com/serz999/notesServer/pkg/dto"
	"github.com/serz999/notesServer/internal/usecases"
)

type NoteContreller struct {
    addn usecases.AddNote
    deln usecases.DelNote
    getnbyid usecases.GetNoteById
}

func NewNoteController(
    addn usecases.AddNote,
    deln usecases.DelNote,
    getnbyid usecases.GetNoteById,
) *NoteContreller {
    return &NoteContreller{addn, deln, getnbyid}  
}

func WriteNotFound(w http.ResponseWriter) {
    w.WriteHeader(http.StatusNotFound)
    w.Write([]byte(`{"code":404,"msg":"Not Found"}`))
}

func WriteInternalServerError(w http.ResponseWriter, err error) {
    w.WriteHeader(http.StatusInternalServerError)
    w.Write([]byte(`{"code":500,"msg":"` + err.Error() + `"}`))
}

func (c *NoteContreller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    log.Printf("%s %s", r.Method, r.URL)

    w.Header().Set("content-type", "application/json") 

    url := strings.Split(r.URL.Path, "/")[1:]
    if r.Method == http.MethodGet && len(url) == 2 && url[1] != "" {
        c.GetById(w, r)
    } else if r.Method == http.MethodPost {
        c.Add(w, r)
    } else if r.Method == http.MethodDelete && len(url) == 2 && url[1] != "" {
        c.Del(w, r)
    } else {
        WriteNotFound(w)     
    } 
}

func (c *NoteContreller) Add(w http.ResponseWriter, r *http.Request) {
    var note dto.Note
    err := json.NewDecoder(r.Body).Decode(&note)
    if err != nil {
        log.Println(err)
        WriteInternalServerError(w, err) 
        return
    }
    
    id, err := c.addn.Exec(note)
    if err != nil {
        log.Println(err)
        WriteInternalServerError(w, err) 
        return
    }

    note, geterr := c.getnbyid.Exec(id)
    if geterr != nil {
        log.Println(err)
        WriteInternalServerError(w, err) 
        return
    }

    jsonBytes, err := json.Marshal(note)
    if err != nil {
        log.Println(err)
        WriteInternalServerError(w, err)
        return
    }

    w.Write(jsonBytes)
}

func (c *NoteContreller) Del(w http.ResponseWriter, r *http.Request) {
    id := dto.Id(strings.Split(r.URL.Path, "/")[1:][1]) 
    err := c.deln.Exec(id)
    if err != nil {
        log.Println(err)
        WriteInternalServerError(w, err)
        return
    }

    w.WriteHeader(http.StatusOK)
    w.Write([]byte("{}"))
}

func (c *NoteContreller) GetById(w http.ResponseWriter, r *http.Request) {
    id := dto.Id(strings.Split(r.URL.Path, "/")[1:][1]) 
    note, err := c.getnbyid.Exec(id)
    if err != nil {
        log.Println(err)
        WriteInternalServerError(w, err)
        return
    }

    jsonBytes, err := json.Marshal(note)
    if err != nil {
        log.Println(err)
        WriteInternalServerError(w, err)
        return
    } 

    w.WriteHeader(http.StatusOK)
    w.Write(jsonBytes)
}
