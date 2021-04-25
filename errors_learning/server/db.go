package server

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/maaars/Go_Language_Advanced/errors_learning/model"
	"github.com/pkg/errors"
	"github.com/prometheus/common/log"
	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/dialects/mysql"
)

type Employee struct {
	ID     int32   `json:"id"`
	Name   string  `json:"name"`
	Salary float32 `json:"salary"`
}

var db *reform.DB

func Init(dsn string) {
	log.Info("start init database...")
	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("failed init database", err)
	}

	db = reform.NewDB(conn, mysql.Dialect, reform.NewPrintfLogger(log.Infof))
	log.Info("init db end")

}

func NewEmployee(salary float32) (*Employee, error) {
	record, err := db.SelectOneFrom(model.TbEmp1View, fmt.Sprintf(" where salary > %v limit 1", salary))
	if err != nil || err == reform.ErrNoRows {
		return nil, errors.Wrap(err, fmt.Sprintf("SelectAllFrom database error ,salary > %v not find ", salary))
	}
	item := record.(*model.TbEmp1)
	emp := &Employee{
		ID:     *item.ID,
		Name:   *item.Name,
		Salary: *item.Salary,
	}
	return emp, nil
}
