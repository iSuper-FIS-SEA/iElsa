package workorder

import (
	"database/sql"
	_ "github.com/mattn/go-oci8"
	"log"
	"time"
)

type Workorder struct {
	Hdnum       string
	TstDate     time.Time
	TsrtPartNum sql.NullString
	HgaPartNum  sql.NullString
	Radius      string
	Tmwi        sql.NullString
	LoadDate    time.Time
	ParmId      sql.NullString
	RefRadius   sql.NullFloat64
	WorkOrder   sql.NullString
	Zipname     sql.NullString
}

type OrderManager struct {
	workorders []*Workorder
}

func connection() *sql.DB {
	db, err := sql.Open("oci8", dsn)
	if err != nil {
		log.Println(err)
		return nil
	}
	return db
}

func NewWorkorderManager() *OrderManager {
	return &OrderManager{}
}

func (w *OrderManager) FindHdnum(hdnum string) []*Workorder {
	db := connection()
	defer db.Close()

	rows, err := db.Query("select hd_num,tst_date,tstr_part_num,hga_part_num,radius,tmwi,load_date,parm_id,ref_radius,work_order,s_zipname from data_he_d2d where hd_num = :1", hdnum)
	if err != nil {
		log.Println(err)
		return nil
	}
	defer rows.Close()

	for rows.Next() {
		o := new(Workorder)
		rows.Scan(&o.Hdnum, &o.TstDate, &o.TsrtPartNum, &o.HgaPartNum, &o.Radius, &o.Tmwi, &o.LoadDate, &o.ParmId, &o.RefRadius, &o.WorkOrder, &o.Zipname)
		w.workorders = append(w.workorders, clone(o))
	}

	return w.workorders
}

func (w *OrderManager) FindOrder(order string) []*Workorder {
	db := connection()
	defer db.Close()

	rows, err := db.Query("select hd_num,tst_date,tstr_part_num,hga_part_num,radius,tmwi,load_date,parm_id,ref_radius,work_order,s_zipname from data_he_d2d where work_order = :1", order)
	if err != nil {
		log.Println(err)
		return nil
	}
	defer rows.Close()

	for rows.Next() {
		o := new(Workorder)
		rows.Scan(&o.Hdnum, &o.TstDate, &o.TsrtPartNum, &o.HgaPartNum, &o.Radius, &o.Tmwi, &o.LoadDate, &o.ParmId, &o.RefRadius, &o.WorkOrder, &o.Zipname)
		w.workorders = append(w.workorders, clone(o))
	}

	return w.workorders
}

func (w *OrderManager) FindFile(zfile string) []*Workorder {
	db := connection()
	defer db.Close()

	rows, err := db.Query("select hd_num,tst_date,tstr_part_num,hga_part_num,radius,tmwi,load_date,parm_id,ref_radius,work_order,s_zipname from data_he_d2d where s_zipname like :1", zfile+"%")
	if err != nil {
		log.Println(err)
		return nil
	}
	defer rows.Close()

	for rows.Next() {
		o := new(Workorder)
		rows.Scan(&o.Hdnum, &o.TstDate, &o.TsrtPartNum, &o.HgaPartNum, &o.Radius, &o.Tmwi, &o.LoadDate, &o.ParmId, &o.RefRadius, &o.WorkOrder, &o.Zipname)
		w.workorders = append(w.workorders, clone(o))
	}

	return w.workorders
}

func clone(w *Workorder) *Workorder {
	c := *w
	return &c
}
