package model

//go:generate reform

// TbEmp1 represents a row in tb_emp1 table.
//reform:tb_emp1
type TbEmp1 struct {
	ID     *int32   `reform:"id"`
	Name   *string  `reform:"name"`
	EptId  *int32   `reform:"eptId"`
	Salary *float32 `reform:"salary"`
}
