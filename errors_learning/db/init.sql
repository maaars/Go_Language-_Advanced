create database demo;
use demo;
CREATE TABLE tb_emp1
(
    id INT(11),
    name VARCHAR(25),
    eptId INT(11),
    salary FLOAT
);

insert into tb_emp1 values(1,'zhangsan',1,9000);
insert into tb_emp1 values(1,'lisi',1,10000);
insert into tb_emp1 values(1,'wangwu',1,11000);