package workorder

import (
	"database/sql"
	_ "github.com/mattn/go-oci8"
	"log"
	"time"
)

type Workorder struct {
	Hdnum       string
	StsDate     time.Time
	TsrtPartNum string
	HgaPartNum  string
	Radius      string
	Tmwi        string
	LoadDate    time.Time
	ParmId      string
	RefRadius   string
	WorkOrder   string
}

const dsn = ""

func NewWorkorderManager() *Workorder {
	return &Workorder{}
}

func (w *Workorder) All() ([]*Workorder, error) {
	db, err := sql.Open("oci8", dsn)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("select 3.14, 'foo' from dual")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var f1 float64
		var f2 string
		rows.Scan(&f1, &f2)
		println(f1, f2) // 3.14 foo
	}
	return nil, nil
}

func (w *Workorder) Find(Hdnum string) ([]*Workorder, bool, error) {
	db, err := sql.Open("oci8", dsn)
	if err != nil {
		log.Println(err)
		return nil, false, nil
	}
	defer db.Close()

	return nil, false, nil
}
