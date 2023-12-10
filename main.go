package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/serz999/golist"
	"github.com/serz999/gomap"
	"github.com/serz999/notesServer/internal/contrellers"
	"github.com/serz999/notesServer/internal/gate"
	"github.com/serz999/notesServer/internal/usecases"
)

func main() {
    enverr := godotenv.Load()
    if enverr != nil {
        log.Fatal(enverr)
    }

    var storageType string 
    if len(os.Args) > 1 {
        storageType = os.Args[1]
    }

    var g usecases.NotesGate
    if storageType == "list" {
        var err error
        g, err = gate.NewStorageGate(golist.NewList())  
        if err != nil {
            log.Fatal(err)
        }
    } else if storageType == "map" {
        var err error
        g, err = gate.NewStorageGate(gomap.NewMap())  
        if err != nil {
            log.Fatal(err)
        }
    } else {
        var err error
        g, err = gate.NewPgNotesGate(os.Getenv("DB_URL")) 
        if err != nil {
            log.Fatal(err)
        }
        defer g.(*gate.PgNotesGate).Close()
    }

    notesController := contrellers.NewNoteController(
        usecases.NewAddNoteInteractor(g),
        usecases.NewDelNoteInteractor(g),
        usecases.NewGetNoteByIdInteractor(g),
    )
    
    mux := http.NewServeMux()
    mux.Handle("/notes/", notesController)
    serverr := http.ListenAndServe(":" + os.Getenv("NOTES_SERVER_PORT"), mux)
    if serverr != nil {
        log.Fatal(serverr)
    }
}
