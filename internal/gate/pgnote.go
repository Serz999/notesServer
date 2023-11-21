package gate

import (
	"context"
    "github.com/google/uuid"
	pgxuuid "github.com/jackc/pgx-gofrs-uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/serz999/notesServer/pkg/dto"
)

type PgNotesGate struct {
    dbpool *pgxpool.Pool
}

func NewPgNotesGate(url string) (*PgNotesGate, error) {
    dbconfig, configerr := pgxpool.ParseConfig(url)
    if configerr != nil {
        return nil, configerr 
    }
    
    dbconfig.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
        pgxuuid.Register(conn.TypeMap())
        return nil
    }

    dbpool, connerr := pgxpool.NewWithConfig(context.Background(), dbconfig)
    if connerr != nil {
        return nil, connerr 
    }
 
    return &PgNotesGate{dbpool}, nil
}

func (g *PgNotesGate) Close() {
    g.dbpool.Close()
} 

func (g *PgNotesGate) Add(note dto.Note) (dto.Id, error) {
    query := `INSERT INTO note (id, author_first_name, aurhor_last_name, note) VALUES ($1, $2, $3, $4)`
    
    uuid := uuid.NewString()
    _, err := g.dbpool.Exec(context.Background(), query, 
        uuid,
        note.AuthorFirstName,
        note.AuthorLastName,
        note.Note,
    ) 

    return dto.Id(uuid), err
} 

func (g *PgNotesGate) GetById(id dto.Id) (dto.Note, error) {
    query := `SELECT * FROM note WHERE id = $1`
    
    note := dto.Note{}
    err := g.dbpool.QueryRow(context.Background(), query, id).Scan(
        &note.Id, 
        &note.AuthorFirstName,
        &note.AuthorLastName,
        &note.Note,
    )

    return note, err
}

func (g *PgNotesGate) Del(id dto.Id) error {
    query := `DELETE FROM note WHERE id = @id` 
    args := pgx.NamedArgs{
        "id": id,
    }
    _, err := g.dbpool.Exec(context.Background(), query, args)

    return err
} 
