package main

import (
	"fmt"
	"github.com/maaars/Go_Language_Advanced/errors_learning/server"
)

func main() {
	server.Init("root:@tcp(127.0.0.1:3306)/demo")
	// 查找工资大于两万的员工，会找不到记录，报错
	employee, err := server.NewEmployee(20000)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(employee.Name)
	}
}
