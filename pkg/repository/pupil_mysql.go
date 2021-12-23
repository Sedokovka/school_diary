package repository
import (
  "github.com/jmoiron/sqlx"
  "github.com/Sedokovka/simple-to-do-app"
  "fmt"
  "strconv"
)

type PupilMysql  struct{
  db *sqlx.DB
}

func NewPupilMysql(db *sqlx.DB) *PupilMysql {
  return &PupilMysql{db: db}
}

func (r *PupilMysql) CreatePupil(userId int, pupil crud.Pupil) (int, error) {
  tx, err := r.db.Begin()
  	if err != nil {
  		return 0, err
  	}

  	var id int
  	createPupilQuery := fmt.Sprintf("INSERT INTO %s (name, rfid_card_id) VALUES ('%s', '%s')", pupilTable, pupil.Name, pupil.Rfid_card_id )
    _, err = r.db.Query(createPupilQuery)
    if err != nil {
      tx.Rollback()
      return 0, err
    }
    query := fmt.Sprintf("SELECT id FROM %s WHERE name = '%s' AND rfid_card_id = '%s'", pupilTable, pupil.Name, pupil.Rfid_card_id);
    row := r.db.QueryRow(query)
    if err := row.Scan(&id); err != nil {
      tx.Rollback()
      return 0, err
    }
  	return id, tx.Commit()
}


func (r *PupilMysql) GetPupilById(userId, pupilId int) (crud.Pupil, error) {
  	var pupil crud.Pupil

  	selectPupilQuery := fmt.Sprintf("SELECT * FROM %s WHERE id = %s", pupilTable, strconv.Itoa(pupilId))
    err := r.db.Get(&pupil, selectPupilQuery)
    if err != nil {
      return pupil, err
    }

  	return pupil, nil
}


func (r *PupilMysql) GetAllPupils(userId int) ([]crud.Pupil, error) {
  	var pupils []crud.Pupil
  	selectPupilQuery := fmt.Sprintf("SELECT * FROM %s ", pupilTable)
    err := r.db.Select(&pupils, selectPupilQuery)
    if err != nil {
      return pupils, err
    }

  	return pupils, nil
}
