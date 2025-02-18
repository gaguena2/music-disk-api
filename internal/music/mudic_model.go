package music

import (
	"context"
	"database/sql"
	"fmt"
	"log"
)

type MusicData struct {
	Code  string
	Name  string
	Albun string
}

type Musics struct {
	db *sql.DB
}

func (m *Musics) FindByCode(code string) (*MusicData, error) {
	rows, err := m.db.Query("SELECT * FROM TB_MUSICS WHERE code = ?", code)
	if err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", code, err)
	}
	defer rows.Close()
	var music MusicData
	for rows.Next() {
		if err := rows.Scan(&music.Code, &music.Name, &music.Albun); err != nil {
			return nil, fmt.Errorf("musics %q: %v", code, err)
		}
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("musics %q: %v", code, err)
	}
	return &music, nil
}

func (m *Musics) Create(data MusicData) {
	query := "INSERT INTO TB_MUSICS(M_CODE, M_NAME, M_ALBUM) VALUES (?, ?, ?)"
	insertResult, err := m.db.ExecContext(context.Background(), query, data.Code, data.Name, data.Albun)
	if err != nil {
		log.Fatalf("impossible insert musics: %s", err)
	}
	id, err := insertResult.LastInsertId()
	if err != nil {
		log.Fatalf("impossible to retrieve last inserted id: %s", err)
	}
	log.Printf("inserted id: %d", id)
}
