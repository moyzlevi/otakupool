package models

import 	"database/sql"

type Image struct {
	ID int
	Data []byte
	Mime_type string
	File_name string
}

type ImageModel struct {
	DB *sql.DB
}

func (m *ImageModel) Insert(data []byte, mime_type string, file_name string) (int, error) {
	stmt := `INSERT INTO images  (data,mime_type,file_name)
	VALUES(
		$1, 
		$2,
		$3) RETURNING id`

	var imageId int
	err := m.DB.QueryRow(stmt, data, mime_type, file_name).Scan(&imageId)

	if err != nil {
		return 0, err
	}

	return imageId, nil
}

func (m *ImageModel) Latest() ([]*Image, error) {
	stmt := `SELECT id, data, mime_type, file_name from images
	ORDER BY id 
	LIMIT 10`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	images := []*Image{}

	for rows.Next() {
		s := &Image{}
		err = rows.Scan(&s.ID,
			&s.Data, &s.Mime_type, &s.File_name)
		if err != nil {
			return nil, err
		}
		images = append(images, s)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return images, nil
}