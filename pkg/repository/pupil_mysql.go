package repository
import (
  "github.com/jmoiron/sqlx"
  "github.com/Sedokovka/simple-to-do-app"
)

type PupilMysql  struct{
  db *sqlx.DB
}

func NewPupilMysql(db *sqlx.DB) *PupilMysql {
  return &PupilMysql{db: db}
}

func (r *PupilMysql) CretePupil(userId int, pupil crud.Pupil) (int, err) {
  tx, err := r.db.Begin()
  	if err != nil {
  		return 0, err
  	}

  	var id int
  	createPupilQuery := fmt.Sprintf("INSERT INTO %s (name, rfid_card_id) VALUES ('%s', '%s')", pupilTable, pupil.name, pupil.rfid_card_id )
    _, err := r.db.Query(createPupilQuery)
    if err != nil {
      tx.Rollback()
      return 0, err
    }
    query := fmt.Sprintf("SELECT id FROM %s WHERE name = '%s' AND rfid_card_id = '%s'", pupilTable, pupil.name, pupil.rfid_card_id);
    row := r.db.QueryRow(query)
    if err := row.Scan(&id); err != nil {
      tx.Rollback()
      return 0, err
    }

  	return id, tx.Commit()
}
