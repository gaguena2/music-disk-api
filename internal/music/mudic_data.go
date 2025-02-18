package music

import (
	"database/sql"
	"fmt"
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
	// Loop through rows, using Scan to assign column data to struct fields.
	var music MusicData
	for rows.Next() {
		if err := rows.Scan(&music.Code, &music.Name, &music.Albun); err != nil {
			return nil, fmt.Errorf("albumsByArtist %q: %v", code, err)
		}
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", code, err)
	}
	return &music, nil
}
