package main

import (
	"log"
	"net/http"
	"os"
	"github.com/joho/godotenv"
	"github.com/serz999/notesServer/internal/contrellers"
	"github.com/serz999/notesServer/internal/gate"
	"github.com/serz999/notesServer/internal/usecases"
)

func main() {
    enverr := godotenv.Load()
    if enverr != nil {
        log.Fatal(enverr)
    }

    pgurl := os.Getenv("DB_URL")

    pg, gateerr := gate.NewPgNotesGate(pgurl) 
    if gateerr != nil {
        log.Fatal(gateerr)
    }
    defer pg.Close()

    notesController := contrellers.NewNoteController(
        usecases.NewAddNoteInteractor(pg),
        usecases.NewDelNoteInteractor(pg),
        usecases.NewGetNoteByIdInteractor(pg),
    )
    
    mux := http.NewServeMux()
    mux.Handle("/notes/", notesController)
    serverr := http.ListenAndServe(":" + os.Getenv("NOTES_SERVER_PORT"), mux)
    if serverr != nil {
        log.Fatal(serverr)
    }
}
