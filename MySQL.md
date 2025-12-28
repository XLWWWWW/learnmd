---

---

# MySQL基础

## 一、基础概念

#### 1、通用语法及其分类

- #### DDL: 数据定义语言，用来定义数据库对象（数据库、表、字段）

  - #### 数据库操作

    - 查询所有数据库：

      `SHOW DATABASES;`

    - 查看当前选择的数据库 

      `SELECT DATABASE();`

    - 创建数据库 

      `CREATE DATABASE [ IF NOT EXISTS ] 数据库名 [ DEFAULT CHARSET 字符集] [COLLATE 排序规则 ];`

    - 删除数据库：

      `DROP DATABASE [ IF EXISTS ] 数据库名;`

    - 使用数据库：

      `USE 数据库名;`

  - #### 表操作

    - 查询当前数据库所有表：
    
      `SHOW TABLES;`
    
    - 查询表结构：
    
      `DESC 表名;`
    
    - 查询指定表的建表语句：
    
      `SHOW CREATE TABLE 表名;`
    
    - 创建表
    
    ```mysql
    CREATE TABLE 表名(
    	字段1 字段1类型 [COMMENT 字段1注释],
    	字段2 字段2类型 [COMMENT 字段2注释],
    	字段3 字段3类型 [COMMENT 字段3注释],
    	...
    	字段n 字段n类型 [COMMENT 字段n注释]
    )[ COMMENT 表注释 ];
    ```
    
    - 添加表字段 
    
       `ALTER TABLE 表名 ADD [column] 字段名 类型(长度) [COMMENT 注释] [约束];`
    
    - 修改数据类型：
    
      `ALTER TABLE 表名 MODIFY 字段名 新数据类型(长度);`
    
    - 修改字段名和字段类型：
    
      `ALTER TABLE 表名 CHANGE 旧字段名 新字段名 类型(长度) [COMMENT 注释] [约束];`
    
    - 删除字段： 
    
      `ALTER TABLE 表名 DROP [column] 字段名;`
    
    - 修改表名字：
    
      `ALTER TABLE 表名 RENAME TO 新表名`
    
    - 删除表：
    
      `DROP TABLE [IF EXISTS] 表名;`

- #### DML: 数据操作语言，用来对数据库表中的数据进行增删改

  - #### 添加数据 

    -  `INSERT INTO 表名 (字段名1, 字段名2, ...) VALUES (值1, 值2, ...);`   更加安全的写法 防止表格结构变化时插入失误

    - `INSERT INTO 表名 VALUES (值1, 值2, ...);`

    - ```mysql
      -- 批量添加数据
      INSERT INTO 表名 (字段名1, 字段名2, ...) VALUES (值1, 值2, ...), (值1, 值2, ...), (值1, 值2, ...);`
      INSERT INTO 表名 VALUES (值1, 值2, ...), (值1, 值2, ...), (值1, 值2, ...);
      ```

  - #### 更新和删除数据

    - 修改数据：

      `UPDATE 表名 SET 字段名1 = 值1, 字段名2 = 值2, ... [ WHERE 条件 ];`
    
    - 删除数据：
    
      `DELETE FROM 表名 [ WHERE 条件 ];`

- #### DQL: 数据查询语言，用来查询数据库中表的记录

  - ```mysql
    SELECT
    	字段列表
    FROM
    	表名字段
    WHERE
    	条件列表
    GROUP BY
    	分组字段列表
    HAVING
    	分组后的条件列表
    ORDER BY
    	排序字段列表
    LIMIT
    	分页参数
    ```

  - #### 基础查询

    - ```mysql
      -- 查询多个字段：
      SELECT 字段1, 字段2, 字段3, ... FROM 表名;
      SELECT * FROM 表名;
      
      -- 设置别名：
      SELECT 字段1 [ AS 别名1 ], 字段2 [ AS 别名2 ], 字段3 [ AS 别名3 ], ... FROM 表名;
      SELECT 字段1 [ 别名1 ], 字段2 [ 别名2 ], 字段3 [ 别名3 ], ... FROM 表名;
      
      -- 转义： 使用/转义
      SELECT * FROM 表名 WHERE name LIKE '/_张三' ESCAPE '/'
      ```
  
  - ####  条件查询
  
    ```mysql
    语法：
    SELECT 字段列表 FROM 表名 WHERE 条件列表;
    ```
  
    | 比较运算符      | 功能                                       |
    | --------------- | ------------------------------------------ |
    | >               | 大于                                       |
    | >=              | 大于等于                                   |
    | <               | 小于                                       |
    | <=              | 小于等于                                   |
    | =               | 等于                                       |
    | <> 或 !=        | 不等于                                     |
    | BETWEEN … AND … | 在某个范围内（含最小、最大值）             |
    | IN(…)           | 在in之后的列表中的值，多选一               |
    | LIKE 占位符     | 模糊匹配（_匹配单个字符，%匹配任意个字符） |
    | IS NULL         | 是NULL                                     |
  
    | 逻辑运算符 | 功能                         |
    | ---------- | ---------------------------- |
    | AND 或 &&  | 并且（多个条件同时成立）     |
    | OR 或 \|\| | 或者（多个条件任意一个成立） |
    | NOT 或 !   | 非，不是                     |
    
  - ####  聚合查询
  
    使用函数 将多行或者多列的信息进行整合
  
    语法：
  
     `SELECT 聚合函数(字段列表) FROM 表名;`
  
  - #### 分组查询
  
    使用 group by 将结果进行分组 并且能够进一步使用`HAVING`对分组得到的结果进行进一步过滤 
  
    语法：
    `SELECT 字段列表 FROM 表名 [ WHERE 条件 ] GROUP BY 分组字段名 [ HAVING 分组后的过滤条件 ];`
    
  - #### 排序查询
  
    语法：
    `SELECT 字段列表 FROM 表名 ORDER BY 字段1 排序方式1, 字段2 排序方式2;`
  
    ASC 	升序 、默认
  
    DESC	降序、需要明确指出
  
  - #### 分页查询
    
    语法：
    `SELECT 字段列表 FROM 表名 LIMIT 起始索引, 查询记录数;`
    
    `SELECT 字段列表 FROM 表名 LIMIT 查询记录数 OFFSET 起始索引;`
    
    `SELECT 字段列表 FROM 表名 LIMIT 查询记录数;` 默认起始索引为0
    
  - DQL执行顺序
    
    FROM -> WHERE -> GROUP BY -> SELECT -> ORDER BY -> LIMIT
    
    这里通过对每个涉及到的东西取别名 来摸清楚整个语句运行的顺序
    
    ![image-20240323000731968](.\asset\images\image-20240323000731968.png)
  
- #### DCL: 数据控制语言，用来创建数据库用户、控制数据库的控制权限

  - #### 管理用户

    - 查询用户：

      ```mysql
      USE mysql;  -- 选择数据库
      SELECT * FROM user; 
      
      select * from mysql.user;
      ```
  
    - 创建用户:
      `CREATE USER '用户名'@'主机名' IDENTIFIED BY '密码';`
  
    - 删除用户：
      `DROP USER '用户名'@'主机名';`
  
    ```mysql
    -- 创建用户test，只能在当前主机localhost访问
    create user 'test'@'localhost' identified by '123456';
    -- 创建用户test，能在任意主机访问
    create user 'test'@'%' identified by '123456';
    create user 'test' identified by '123456';
    -- 修改密码
    alter user 'test'@'localhost' identified with mysql_native_password by '1234';
    -- 删除用户
    drop user 'test'@'localhost';
    ```
  
  - ####  权限控制
  
    - 常用权限：
  
    | 权限                | 说明               |
    | ------------------- | ------------------ |
    | ALL, ALL PRIVILEGES | 所有权限           |
    | SELECT              | 查询数据           |
    | INSERT              | 插入数据           |
    | UPDATE              | 修改数据           |
    | DELETE              | 删除数据           |
    | ALTER               | 修改表             |
    | DROP                | 删除数据库/表/视图 |
    | CREATE              | 创建数据库/表      |
  
    - 查询权限：
      `SHOW GRANTS FOR '用户名'@'主机名';`
    - 授予权限：
      `GRANT 权限列表 ON 数据库名.表名 TO '用户名'@'主机名';`
    - 撤销权限：
      `REVOKE 权限列表 ON 数据库名.表名 FROM '用户名'@'主机名';`
  

#### 2、约束

| 约束                    | 描述                                                         | 关键字      |
| ----------------------- | ------------------------------------------------------------ | ----------- |
| 非空约束                | 限制该字段的数据不能为null                                   | NOT NULL    |
| 唯一约束                | 保证该字段的所有数据都是唯一、不重复的，唯一索引是只允许索引中的每个值对应记录的一行 这里尝试过唯一约束可以为NULL 并且能够插入多个该列为NULL的数据行 | UNIQUE      |
| 主键约束                | 主键是一行数据的唯一标识，要求非空且唯一                     | PRIMARY KEY |
| 默认约束                | 保存数据时，如果未指定该字段的值，则采用默认值               | DEFAULT     |
| 检查约束（8.0.1版本后） | 保证字段值满足某一个条件                                     | CHECK       |
| 外键约束                | 用来让两张图的数据之间建立连接，保证数据的一致性和完整性     | FOREIGN KEY |

##### 常用约束

| 约束条件 | 关键字         |
| -------- | -------------- |
| 主键     | PRIMARY KEY    |
| 自动增长 | AUTO_INCREMENT |
| 不为空   | NOT NULL       |
| 唯一     | UNIQUE         |
| 逻辑条件 | CHECK          |
| 默认值   | DEFAULT        |

```mysql
create table user(
	id int primary key auto_increment,
	name varchar(10) not null unique,
	age int check(age > 0 and age < 120),
	status char(1) default '1',
	gender char(1)
);
```

##### 外键约束

```mysql
-- 添加外键：
CREATE TABLE 表名(
	字段名 字段类型,
	...
	[CONSTRAINT] [外键名称] FOREIGN KEY(外键字段名) REFERENCES 主表(主表列名)
);
ALTER TABLE 表名 ADD CONSTRAINT 外键名称 FOREIGN KEY (外键字段名) REFERENCES 主表(主表列名);

-- 例子
alter table emp add constraint fk_emp_dept_id foreign key(dept_id) references dept(id);

-- 删除外键
ALTER TABLE 表名 DROP FOREIGN KEY 外键名;
```

```mysql
drop table if exists emp;

create table dept(
    id int auto_increment primary key comment 'ID',
    name varchar(50) not null comment '部门名称'
) comment '部门表';

insert into dept(id,name)
values(1, '研发部'), (2, '市场部'),(3, '财务部'), (4,
'销售部'), (5, '总经办');

create table emp(
    id  int auto_increment comment 'ID' primary key,
    name varchar(50) not null comment '姓名',
    age  int comment '年龄',
    job varchar(20) comment '职位',
    salary int comment '薪资',
    entrydate date comment '入职时间',
    managerid int comment '直属领导ID',
    dept_id int comment '部门ID'
 )comment '员工表';

INSERT INTO emp (id, name, age, job,salary, entrydate, managerid, dept_id)
values (1, '金庸', 66, '总裁',20000, '2000-01-01', null,5),(2, '张无忌', 20,'项目经理',12500, '2005-12-05', 1,1),
 (3, '杨逍', 33, '开发', 8400,'2000-11-03', 2,1),(4, '韦一笑', 48, '开发',11000, '2002-02-05', 2,1),
 (5, '常遇春', 43, '开发',10500, '2004-09-07', 3,1),(6, '小昭', 19, '程序员鼓励师',6600, '2004-10-12', 2,1);

-- 添加外键保证两表内数据的完整性。
-- 这里就这么记住吧 referce谁 谁是父表(parent) 父表中与子表关联的数据不能删除
-- 子表只能插入父表中有的项 但是如果子表中没有not null的规定，NULL是能够插进去的！！！
ALTER TABLE emp ADD CONSTRAINT fk_emp_dept_id FOREIGN KEY (dept_id) references dept(id);
-- 父表中与子表关联的数据不能删除
DELETE FROM dept WHERE id = 5;
-- 但是不与子表中关联的数据是能够删除的。
DELETE FROM dept WHERE id = 2;
-- 能够更新吗 同样，需要修改的数据不和子表的数据关联的能够正常修改
UPDATE dept SET id = 10 WHERE id = 1;
UPDATE dept SET id = 10 WHERE id = 2;

-- 解除约束  两种写法都是正确的
ALTER TABLE emp DROP CONSTRAINT fk_emp_dept_id;
ALTER TABLE enp FOREIGN KEY fk_emp_dept_id;
```

##### 约束行为

是指添加外键后，再删除、修改父表数据产生的约束行为，具体的形式分为以下几种，

| 行为        | 说明                                                         |
| ----------- | ------------------------------------------------------------ |
| NO ACTION   | 当在父表中删除/更新对应记录时，首先检查该记录是否有对应外键，如果有则不允许删除/更新（与RESTRICT一致） |
| RESTRICT    | 当在父表中删除/更新对应记录时，首先检查该记录是否有对应外键，如果有则不允许删除/更新（与NO ACTION一致） |
| CASCADE     | 当在父表中删除/更新对应记录时，首先检查该记录是否有对应外键，如果有则也删除/更新外键在子表中的记录 |
| SET NULL    | 当在父表中删除/更新对应记录时，首先检查该记录是否有对应外键，如果有则设置子表中该外键值为null（要求该外键允许为null） |
| SET DEFAULT | 父表有变更时，子表将外键设为一个默认值**（Innodb不支持）**   |

```mysql
-- 更改删除/更新行为：
ALTER TABLE 表名 ADD CONSTRAINT 外键名称 FOREIGN KEY (外键字段) REFERENCES 主表名(主表字段名) ON UPDATE 行为 ON DELETE 行为;

-- 涉及更新和删除的外键处理关系 这里是创建的时候同时说明了更新和删除关系
ALTER TABLE emp ADD CONSTRAINT fk_emp_dept_id FOREIGN KEY (dept_id) REFERENCES dept(id)  ON UPDATE CASCADE ON DELETE CASCADE ;
-- 因为是两个操作的行为 所以 这里创建的时候是支持分别设置的
ALTER TABLE emp ADD CONSTRAINT fk_emp_dept_id1 FOREIGN KEY (dept_id) REFERENCES dept(id)  ON UPDATE CASCADE;
ALTER TABLE emp ADD CONSTRAINT fk_emp_dept_id2 FOREIGN KEY (dept_id) REFERENCES dept(id)  ON UPDATE CASCADE ON DELETE SET NULL;
ALTER TABLE emp DROP FOREIGN KEY fk_emp_dept_id1;
ALTER TABLE emp DROP FOREIGN KEY fk_emp_dept_id2;

-- 级联效果 子表继承父表的操作
UPDATE dept SET id = 6 WHERE id = 1;
DELETE FROM dept WHERE id = 6;

-- SET NULL 关联的子表将父表删除数据置为NULL
ALTER TABLE emp ADD CONSTRAINT fk_emp_dept_id2 FOREIGN KEY (dept_id) REFERENCES dept(id)  ON UPDATE CASCADE ON DELETE SET NULL;
UPDATE dept SET id = 6 WHERE id = 1;
DELETE FROM dept WHERE id = 6;
```



#### 3、多表查询

##### 内连接查询

内连接查询的是两张表交集的部分  显式性能比隐式高

- 隐式内连接：
  `SELECT 字段列表 FROM 表1, 表2 WHERE 条件 ...;`
- 显式内连接：
  `SELECT 字段列表 FROM 表1 [ INNER ] JOIN 表2 ON 连接条件 ...;`

##### 外连接查询

- 左外连接：查询左表所有数据，以及两张表交集部分数据

  `SELECT 字段列表 FROM 表1 LEFT [ OUTER ] JOIN 表2 ON 条件 ...;`

- 右外连接：查询右表所有数据，以及两张表交集部分数据
  `SELECT 字段列表 FROM 表1 RIGHT [ OUTER ] JOIN 表2 ON 条件 ...;`

##### 自连接查询

当前表与自身的连接查询，自连接必须使用表别名

语法：
`SELECT 字段列表 FROM 表A 别名A JOIN 表A 别名B ON 条件 ...;`

**自连接查询，可以是内连接查询，也可以是外连接查询**

```mysql
-- 查询员工及其所属领导的名字
select a.name, b.name from employee a, employee b where a.manager = b.id;
-- 没有领导的也查询出来
select a.name, b.name from employee a left join employee b on a.manager = b.id;
```

##### 联合查询 union, union all

把多次查询的结果合并，形成一个新的查询集

```mysql
SELECT 字段列表 FROM 表A ...
UNION [ALL]
SELECT 字段列表 FROM 表B ...
```

- UNION ALL 会有重复结果，UNION 不会
- 联合查询比使用or效率高，不会使索引失效

##### 子查询

SQL语句中嵌套SELECT语句，称谓嵌套查询，又称子查询。

- #### 标量子查询

子查询返回的结果是单个值（数字、字符串、日期等）。常用操作符：- < > > >= < <=

- #### 列子查询

| 操作符 | 描述                                   |
| ------ | -------------------------------------- |
| IN     | 在指定的集合范围内，多选一             |
| NOT IN | 不在指定的集合范围内                   |
| ANY    | 子查询返回列表中，有任意一个满足即可   |
| SOME   | 与ANY等同，使用SOME的地方都可以使用ANY |
| ALL    | 子查询返回列表的所有值都必须满足       |



```mysql
-- 查询销售部和市场部的所有员工信息
select * from employee where dept in (select id from dept where name = '销售部' or name = '市场部');
-- 查询比财务部所有人工资都高的员工信息 all表示返回列表的值都需要满足
select * from employee where salary > all(select salary from employee where dept = (select id from dept where name = '财务部'));
-- 查询比研发部任意一人工资高的员工信息
select * from employee where salary > any (select salary from employee where dept = (select id from dept where name = '研发部'));
```

- #### 行子查询

返回的结果是一行（可以是多列）。常用操作符：=, <, >, IN, NOT IN

```mysql
-- 查询与xxx的薪资及直属领导相同的员工信息
select * from employee where (salary, manager) = (12500, 1);
select * from employee where (salary, manager) = (select salary, manager from employee where name = 'xxx');
```

- #### 表子查询 

  返回的结果是多行多列  常用操作符：IN

  ```mysql
  -- 查询与xxx1，xxx2的职位和薪资相同的员工
  select * from employee where (job, salary) in (select job, salary from employee where name = 'xxx1' or name = 'xxx2');
  -- 查询入职日期是2006-01-01之后的员工，及其部门信息
  select e.*, d.* from (select * from employee where entrydate > '2006-01-01') as e left join dept as d on e.dept = d.id;
  ```

#### 4、事务

事务是一组操作的集合，事务会把所有操作作为一个整体一起向系统提交或撤销操作请求，即这些操作要么同时成功，要么同时失败。

```mysql
-- 查看事务提交方式
SELECT @@AUTOCOMMIT;
-- 设置事务提交方式，1为自动提交，0为手动提交，该设置只对当前会话有效
SET @@AUTOCOMMIT = 0;
-- 提交事务
COMMIT;
-- 回滚事务
ROLLBACK;

-- 开启事务：
START TRANSACTION 或 BEGIN TRANSACTION;
-- 提交事务：
COMMIT;
-- 回滚事务：
ROLLBACK;

start transaction;
select * from account where name = '张三';
update account set money = money - 1000 where name = '张三';
update account set money = money + 1000 where name = '李四';
commit;
```

##### 四大特性ACID

- 原子性(Atomicity)：事务是不可分割的最小操作但愿，要么全部成功，要么全部失败
- 一致性(Consistency)：事务完成时，必须使所有数据都保持一致状态
- 隔离性(Isolation)：数据库系统提供的隔离机制，保证事务在不受外部并发操作影响的独立环境下运行
- 持久性(Durability)：事务一旦提交或回滚，它对数据库中的数据的改变就是永久的

##### 并发事务

- ##### 问题

| 问题       | 描述                                                         |
| ---------- | ------------------------------------------------------------ |
| 脏读       | 一个事务读到另一个事务还没提交的数据                         |
| 不可重复读 | 一个事务先后读取同一条记录，但两次读取的数据不同             |
| 幻读       | 一个事务按照条件查询数据时，没有对应的数据行，但是再插入数据时，又发现这行数据已经存在 |

- **并发事务隔离级别：**

  - √表示在当前隔离级别下该问题会出现

  - Serializable 性能最低；Read uncommitted 性能最高，数据安全性最差

  - | 隔离级别              | 脏读 | 不可重复读 | 幻读 |
    | --------------------- | ---- | ---------- | ---- |
    | Read uncommitted      | √    | √          | √    |
    | Read committed        | ×    | √          | √    |
    | Repeatable Read(默认) | ×    | ×          | √    |
    | Serializable          | ×    | ×          | ×    |

- 查看事务隔离级别：
  `SELECT @@TRANSACTION_ISOLATION;`

- 设置事务隔离级别：
  `SET [ SESSION | GLOBAL ] TRANSACTION ISOLATION LEVEL {READ UNCOMMITTED | READ COMMITTED | REPEATABLE READ | SERIALIZABLE };`

  **SESSION 是会话级别，表示只针对当前会话有效，GLOBAL 表示对所有会话有效**



## 二、语法语句

#### 1、选择列及显示（select ... from ...）

```mysql
# lesson 1

select prod_name from products;

select  prod_id,vend_id from products;

#减配所有的列
select * from  products;

select vend_id from products;
select distinct vend_id from  products;

SELECT DISTINCT vend_id, prod_price FROM Products;

SELECT DISTINCT vend_id, prod_price FROM Products;

SELECT vend_id, prod_price FROM Products;

#显示顺序
select  prod_name from products limit  5;

-- offset从哪里开始 limit 限制多少个
select  prod_name from products limit  5 offset 6;

select  prod_name from products limit 6,5; -- 简化形式 limit(offset,limit);

select  cust_id from customers;
select  distinct  prod_id from orderitems  limit 7;

select * from products;
```

#### 2、显示顺序(order by  列名 desc/asc)

```mysql
#lesson 2
select prod_name from products order by prod_name;
select prod_name from products order by vend_id;

select prod_name,prod_id from products order by  prod_name,prod_id;-- 首先按照什么再按照什么排序
-- ASC 默认 DESC 需要单独指明
select prod_name from products order by prod_name desc ;
select prod_name from products order by prod_name asc ;
```

#### 3、过滤数据(where)

```mysql
#lesson 3
select  prod_name,prod_price from products where prod_price = 3.49;

select  prod_name,prod_price from products where prod_price = 3.49 order by  prod_name;

select prod_name,prod_price from products where prod_price <= 10;

select vend_id,prod_name from products where vend_id != 'DLL01';

select vend_id,prod_name from products where vend_id <> 'DLL01';

select prod_name,prod_price from products where prod_price between 5 and 10;

select prod_id,prod_name from products where prod_price = 9.49;

select prod_id,prod_name,prod_price from products where prod_price >=9;

select distinct order_num from orderitems  limit 100;

select  prod_name,prod_price from products where prod_price between 3 and 6 order by prod_price desc ;

select  prod_id,prod_price,prod_name from products where vend_id ='DLL01' and prod_price <= 4;

select prod_id,prod_price,prod_name,vend_id from products where vend_id = 'DLL01' or vend_id = 'BSR01';

-- SQL语言存在优先级 先处理AND 因此需要使用小括号以保证处理流程正确性
SELECT prod_name, prod_price
FROM Products
WHERE vend_id = 'DLL01' OR vend_id = 'BRS01'
AND prod_price >= 10;

SELECT prod_name, prod_price
FROM Products
WHERE (vend_id = 'DLL01' OR vend_id = 'BRS01')
AND prod_price >= 10;
```

#### 4、高级过滤  (in not的使用)

```mysql
# lesson 4
select prod_name,prod_price from products where vend_id in ('DLL01','BRS01') order by  prod_name;

SELECT prod_name, prod_price
FROM Products
WHERE vend_id = 'DLL01' OR vend_id = 'BRS01'
ORDER BY prod_name;

select prod_name from products where not vend_id = 'DLL01' order by prod_name;

select prod_name from products where  vend_id != 'DLL01' order by prod_name;

select vend_name from vendors where vend_country = 'USA' and vend_state = 'CA';

select order_num,prod_id from orderitems where prod_id in ('BR01','BR02','BR03') and quantity >=100;

select prod_name,prod_price from products where  prod_price >= 3 and prod_price <=6 order by prod_price desc ;
```

#### 5、通配符(like 、% 、_的使用)

%匹配0个或多个字符，_只匹配一个字符

```mysql
# lesson 5
select prod_id,prod_name from products where prod_name like 'fish%'

select prod_id,prod_name from products where prod_name like '%bean bag%';
select prod_id,prod_name from products where prod_name like 'f%y';

select prod_id,prod_name from products where prod_name like '__ inch teddy bear';
select prod_id,prod_name from products where prod_name like '_ inch teddy bear';

select  cust_contact from customers where cust_contact like '[jm]%' order by cust_contact;

select prod_name,prod_desc from products where prod_desc like '%toy%';

select prod_name,prod_desc from products where not prod_desc like '%toy%' order by prod_name desc;

select prod_name,prod_desc from products where  prod_desc like '%toy%' and prod_desc like '%carrots%';

select prod_name,prod_desc from products where  prod_desc like '%toy%carrots%';
```

#### 6、创建计算字段(concat,表达式，as起别名)

```mysql
# lesson 6
select distinct concat(vend_name,'(',vend_country,')') from vendors order by vend_name;

select distinct concat(rtrim(vend_name),'(',rtrim(vend_country),')') from vendors order by vend_name;

select distinct concat(rtrim(vend_name),'(',rtrim(vend_country),')') as vend_title from vendors order by vend_name;

select distinct concat(rtrim(vend_name),'(',rtrim(vend_country),')') vend_title from vendors order by vend_name;

select prod_id,quantity,item_price from orderitems where order_num = 20008;

select prod_id,quantity,item_price ,quantity*item_price as expended_prices from orderitems
where order_num = 20008;

select vend_id,vend_name as vname,vend_address as vaddress,vend_city as vcity from vendors;

select prod_id,prod_price,prod_price*0.9 as sale_prices from products;

```

#### 7、函数

每一个DBMS都有特定的函数

- rtrim()去除i字符串右边的空格

- upper将字符串转换为大写

- substr()求子串  substr(str,startindem,num) 从startindex开始取num位  

- LEFT(add,index) 从index开始向左截取

- RIGHT(add,index)从index开始向右截取

  - 这里字符串的起始位置都是 1 

- soundex 匹配发音类似的

- ### 字符串函数

| 函数                             | 功能                                                      |
| -------------------------------- | --------------------------------------------------------- |
| CONCAT(s1, s2, …, sn)            | 字符串拼接，将s1, s2, …, sn拼接成一个字符串               |
| LOWER(str)                       | 将字符串全部转为小写                                      |
| UPPER(str)                       | 将字符串全部转为大写                                      |
| LPAD(str, n, pad)                | 左填充，用字符串pad对str的左边进行填充，达到n个字符串长度 |
| RPAD(str, n, pad)                | 右填充，用字符串pad对str的右边进行填充，达到n个字符串长度 |
| TRIM(str)                        | 去掉字符串头部和尾部的空格                                |
| SUBSTRING(str, start, len)       | 返回从字符串str从start位置起的len个长度的字符串           |
| REPLACE(column, source, replace) | 替换字符串                                                |

- ### 数值函数

| 函数        | 功能                             |
| ----------- | -------------------------------- |
| CEIL(x)     | 向上取整                         |
| FLOOR(x)    | 向下取整                         |
| MOD(x, y)   | 返回x/y的模                      |
| RAND()      | 返回0~1内的随机数                |
| ROUND(x, y) | 求参数x的四舍五入值，保留y位小数 |

- ### 日期函数

| 函数                               | 功能                                              |
| ---------------------------------- | ------------------------------------------------- |
| CURDATE()                          | 返回当前日期                                      |
| CURTIME()                          | 返回当前时间                                      |
| NOW()                              | 返回当前日期和时间                                |
| YEAR(date)                         | 获取指定date的年份                                |
| MONTH(date)                        | 获取指定date的月份                                |
| DAY(date)                          | 获取指定date的日期                                |
| DATE_ADD(date, INTERVAL expr type) | 返回一个日期/时间值加上一个时间间隔expr后的时间值 |
| DATEDIFF(date1, date2)             | 返回起始时间date1和结束时间date2之间的天数        |

- ### 流程函数

| 函数                                                         | 功能                                                    |
| ------------------------------------------------------------ | ------------------------------------------------------- |
| IF(value, t, f)                                              | 如果value为true，则返回t，否则返回f                     |
| IFNULL(value1, value2)                                       | 如果value1不为空，返回value1，否则返回value2            |
| CASE WHEN [ val1 ] THEN [ res1 ] … ELSE [ default ] END      | 如果val1为true，返回res1，… 否则返回default默认值       |
| CASE [ expr ] WHEN [ val1 ] THEN [ res1 ] … ELSE [ default ] END | 如果expr的值等于val1，返回res1，… 否则返回default默认值 |

```mysql
# lesson 7
select vend_name,upper(vend_name) as vend_nama_upper from vendors;

select cust_name,cust_contact from customers where cust_contact = 'Michael green';

select cust_name,cust_contact from customers where soundex(cust_contact) = soundex('Michael green');

select order_num from orders where year(order_date) = 2020;

select upper(cust_id),upper(cust_name),
       upper(concat(substring(cust_contact,1,2),substring(cust_city,1,3))) as user_login from customers;

select order_num,order_date from orders where year(order_date) = 2020 and month(order_date) =1 order by order_date;
```

#### 8、汇总数据

| 函数  | 功能     |
| ----- | -------- |
| count | 统计数量 |
| max   | 最大值   |
| min   | 最小值   |
| avg   | 平均值   |
| sum   | 求和     |

```mysql
#lesson 8
select avg(prod_price) as avg_prices from products;

select avg(prod_price) as avg_prices from products where vend_id = 'DLL01';

select count(*) from products;
select count(cust_email) as num_cust from customers;

select max(prod_price) from products;

#当用于文本数据时，返回该列排序后的最后一行
select max(prod_name) from products;

select min(prod_price) from products;
select min(prod_name) from products;

select sum(quantity) as items_ordered from orderitems where order_num = 20005;

select sum(item_price*orderitems.quantity) as total_prices from orderitems where order_num = 20005;


select avg(distinct prod_price) from products where vend_id = 'DLL01';

select count(*) as num_items,min(prod_price) as prices_min,max(prod_price) as prices_max from products;

select sum(quantity) from orderitems;
select sum(quantity) from orderitems where prod_id = 'BR01';
select max(prod_price) as max_prices from products where prod_price <=10;
```

#### 9、数组分组(group by 和 having)

where 过滤行 having 过滤分组

**WHERE 在数据分组前进行过滤，HAVING 在数据分组后进行过滤。 经过where筛选的数据不可能出现在分组中**

![image-20220923233250161](C:\Users\xlwww\AppData\Roaming\Typora\typora-user-images\image-20220923233250161.png)

```mysql
# lesson 9
select vend_id ,count(*) as num_prods from products group by vend_id;

select cust_id,count(*) as orders from orders group by cust_id having count(*) >=2;
select vend_id,count(*) as nums_prods from products where prod_price >=4 group by vend_id having count(*)>=2;

select order_num,count(*) as items from orderitems group by order_num having count(*)>=3;


--  SQL规范没有要求分组顺序 因此需要使用ORDER BY规定顺序
select order_num,count(*) as items from orderitems group by order_num having count(*)>=3 order by items,order_num;

select order_num,count(*) as order_lines from orderitems group by order_num order by order_lines;

select vend_id,min(prod_price) as cheapest_item  from products group by vend_id order by cheapest_item;

select order_num from orderitems where quantity >=100;

select order_num from orderitems where item_price*quantity >=1000;

select order_num,count(*) as items from orderitems group by order_num having count(*)>=3 order by items,order_num;
```

#### 10、子查询

就是查询中套查询，把一条select中的结果用于另一条select语句的where语句

##### 使用子查询进行过滤

有点小饶，其实分为两种，一是一层一层的往上查询的，**称为使用子查询进行过滤**

```mysql
# lesson 10
select cust_id from orders where order_num in
                                 (select order_num from orderitems where prod_id = 'RGAN01');
-- 上一句晚会结果为1000000004,1000000005
#接着检索其顾客信息
select cust_name,cust_contact from customers where  cust_id in(1000000004,1000000005);
#合并为一句
select cust_name,cust_contact from customers where cust_id in(
                                                 select cust_id from orders where order_num in
                                                                            (select order_num from orderitems where prod_id = 'RGAN01'));

```

##### 作为计算字段使用子查询

二是利用这种特性作为返回值，**一般会使用完全限定列名(Table.column)**

```mysql
-- (1) 从 Customers 表中检索顾客列表；
-- (2) 对于检索出的每个顾客，统计其在 Orders 表中的订单数目。
SELECT cust_name, cust_state,
       (SELECT COUNT(*)  FROM Orders WHERE Orders.cust_id = Customers.cust_id) AS orders
FROM Customers
ORDER BY cust_name;
```

课后题答案

```mysql
SELECT cust_id
FROM orders
WHERE order_num IN(SELECT order_num
                   FROM orderitems
                   WHERE item_price>=10);

SELECT cust_id,order_date
FROM orders
WHERE order_num IN(SELECT order_num
                   FROM orderitems
                   WHERE prod_id = 'BR01')
ORDER BY order_date;

SELECT cust_email
FROM customers
WHERE cust_id IN (SELECT cust_id
                  FROM orders
                  WHERE order_num IN (SELECT order_num
                                      FROM orderitems
                                      WHERE prod_id = 'BR01'));

SELECT cust_id,
       -- 这里是sum 因为在orders中顾客的order_num不止一个
       -- 不加sum 返回就不是单列了
       (SELECT SUM(quantity*item_price)
        FROM orderitems
        WHERE order_num IN(SELECT order_num
                           FROM orders
                           WHERE orders.cust_id = customers.cust_id)) AS total_ordered
FROM customers
ORDER BY total_ordered DESC;


select prod_name,
       (select sum(quantity) from orderitems where orderitems.prod_id = products.prod_id) as quant_sold
from products;
```

#### 11、联结表

##### 1、等值联结

直接使用where子句创建联结，在子句中使用完全限定列名进行匹配

```mysql
#lesson 11
select vend_name,prod_name,prod_price from vendors,products where vendors.vend_id = products.vend_id;

# 不匹配的话返回的是笛卡尔积 得到的总个数位 左表数量*右表数量个结果
select vend_name,prod_name,prod_price from vendors,products;

#简单等值语句实现筛选得到内连结的效果，
select vend_name,prod_name,prod_price from vendors,products  where vendors.vend_id = products.vend_id;
```

##### 2、内连接

使用 inner join ... on... 的方法

```mysql
#简单等值语句实现筛选得到内连结的效果， 编译器会自动提供联结关系建议
select vend_name,prod_name,prod_price from vendors,products  where vendors.vend_id = products.vend_id;
#使用内连接关键字 inner join ... on ...
select vend_name,prod_name,prod_price from vendors inner join products on vendors.vend_id = products.vend_id;
select vend_name,prod_name,prod_price from products inner join vendors on products.vend_id = vendors.vend_id; 
-- 这两个实际的结果是一致的 以为获取的是两表间相同关联的行
```

##### 3、联结多个表

```mysql
select prod_name,vend_name,prod_price,quantity from orderitems,products,vendors where products.vend_id = vendors.vend_id
and orderitems.prod_id = products.prod_id and order_num = 20007;

select  cust_name,cust_contact from customers,orders,orderitems where
                                                                    orders.cust_id = customers.cust_id and
                                                                    orders.order_num = orderitems.order_num and
                                                                    orderitems.prod_id = 'RGAN01';

select cust_name,order_num from customers,orders
                           where customers.cust_id = orders.cust_id order by  cust_name,order_num;
```

```mysql
select cust_name,order_num,(select orderitems.quantity*orderitems.item_price as ordertotal from orderitems,orders where orderitems.order_num = orders.order_num)
from customers,orders where  customers.cust_id = orders.cust_id order by  cust_name,order_num;

-- 这是错误的形式 因为sum(orderitems.quantity*orderitems.item_price)每个cust_name,orders.order_num可能对应好几个这样计算出来的结果
-- 因此 这样会报错 这个表综合不到一起
select cust_name,orders.order_num,sum(orderitems.quantity*orderitems.item_price) ordertotal from customers,orders,orderitems where customers.cust_id = orders.cust_id and orderitems.order_num = orders.order_num 
-- group by cust_name, orders.order_num order by  cust_name,order_num;


-- 这里使用联结 联结好的部分需要指定让返回形式匹配才不会报错 因此需要对SUM里面的东西规定返回形式 使用GROUP BY cust_name,orders.order_num
-- 是的最后返回的结果 能够以表的形式呈现
select cust_name,orders.order_num,sum(orderitems.quantity*orderitems.item_price) ordertotal from customers,orders,orderitems where customers.cust_id = orders.cust_id and orderitems.order_num = orders.order_num group by cust_name, orders.order_num order by  cust_name,order_num;

select  cust_name,order_num,(select sum(quantity*item_price) from orderitems where orderitems.order_num = orders.order_num) as ordertotal from orders,customers where customers.cust_id = orders.cust_id order by  cust_name,order_num;

select cust_id,order_date from orders,orderitems where orders.order_num = orderitems.order_num and orderitems.prod_id = 'BR01';

-- inner join... on 的多表联结形式
select cust_email from customers
    inner join orders on orders.cust_id = orders.cust_id
    inner join orderitems on orderitems.order_num = orders.order_num where orderitems.prod_id = 'BR01';

select cust_name from customers
    inner join orders on orders.cust_id = customers.cust_id
    inner join orderitems on orders.order_num = orderitems.order_num
    group by cust_name having sum(item_price*quantity)>=1000  -- 这一句很关键，使用组的形式进行结果筛选
order by cust_name;

select cust_name,sum(item_price*quantity) as total_prices from customers,orders,orderitems
             where customers.cust_id = orders.cust_id and orders.order_num = orderitems.order_num
             group by cust_name having sum(item_price*quantity)>=1000
             order by cust_name;
```

#### 12、创建高级联结

```mysql
# lesson 12

# 自查询，通过对自己的表起不同的别名，进而在自己的表中得到结果并且继续再自己的表中提取其他信息
select cust_id,cust_name,cust_contact from customers where cust_name =(
select cust_name from customers where soundex(cust_contact) = soundex('JIm jones'));

select c1.cust_id,c1.cust_name,c1.cust_contact from customers as c1,customers as c2
                                               where c1.cust_name = c2.cust_name and c2.cust_contact = 'Jim Jones';

# 自然联结，就是删除掉回收结果的相同的列,使每一列只出现一次 使用select *  说实话 这里真没整明白
SELECT C.*, O.order_num, O.order_date,
 OI.prod_id, OI.quantity, OI.item_price
FROM Customers AS C, Orders AS O,
 OrderItems AS OI
WHERE C.cust_id = O.cust_id
AND OI.order_num = O.order_num
AND prod_id = 'RGAN01';

SELECT C.*, O.order_num, O.order_date,oi.order_num,
 OI.prod_id, OI.quantity, OI.item_price
FROM Customers AS C, Orders AS O,
 OrderItems AS OI
WHERE C.cust_id = O.cust_id
AND OI.order_num = O.order_num
AND prod_id = 'RGAN01';

# 外连接，既返回关联的行也返回不关联的行 left/right outer join == left/right join
-- 这是一个普通的内连接
-- 但是只返回了存在联结的表项，会将没有联结的选项忽略
select customers.cust_id,orders.order_num from customers inner join orders on customers.cust_id = orders.cust_id;

--  左外联结 遍历左边的所有cust_id 并且列出与之关联的右边的数据
select customers.cust_id,orders.order_num from customers left outer join orders on customers.cust_id = orders.cust_id;
--  右外联结 便利右边的所有cust_id 并且列出与之关联的左边的数据
select customers.cust_id,orders.order_num from customers right outer join orders on customers.cust_id = orders.cust_id;

# 全外联结，检索两个表中的所有行并关联那些可以关联的行
# full outer join mysql支持
-- select customers.cust_id,orders.order_num from customers full outer join orders on customers.cust_id = orders.cust_id;

select customers.cust_id,count(orders.order_num) as num_ord from customers inner join orders on customers.cust_id = orders.cust_id  group by  orders.cust_id;
select customers.cust_id,count(orders.order_num) as num_ord from customers left outer join orders on customers.cust_id = orders.cust_id  group by  orders.cust_id;
```

#### 

#### 13、组合查询(union /union all)

```mysql
# lesson 13

select cust_name,cust_contact,cust_email from customers where cust_state in ('il','in','mi');

select cust_name,cust_contact,cust_email from customers where cust_name = 'Fun4All';
#使用union进行组合查询,从查询结果中自动去掉重复的行
select cust_name,cust_contact,cust_email from customers where cust_state in ('il','in','mi') 
union
select cust_name,cust_contact,cust_email from customers where cust_name = 'Fun4All';
-- 其实也可以使用条where首先相同的结果
-- 任何具有多个 WHERE 子句的 SELECT 语句都可以作为一个组合查询，
select cust_name,cust_contact,cust_email from customers where cust_state in ('il','in','mi') or cust_name = 'Fun4All';

-- 使用union时回去掉重复的行，也可使用 union all 来将两个select 语句的结果进行一个简单的聚合,那union all的功能就不能直接使用where来进行替换了
select cust_name,cust_contact,cust_email from customers where cust_state in ('il','in','mi') union all
                                                                                             select cust_name,cust_contact,cust_email from customers where cust_name = 'Fun4All';
                                                                                             
-- 起别名 union的几个部分都起了别名的话，默认使用第一个语句的别名
SELECT cust_name AS c_n1,cust_contact,cust_email FROM customers
WHERE cust_state IN('IL','IN','MI')
UNION
SELECT cust_name AS c_n2,cust_contact,cust_email FROM customers
WHERE cust_name = 'Fun4All';                                                                                             
                                                                                             
                                                                                             

-- 在使用union时只能使用一条order by 语句，必须位于最后一条select语句之后
select cust_name,cust_contact,cust_email from customers where cust_state in ('IL','IN','MI')
union
select cust_name,cust_contact,cust_email from customers where  cust_name = 'Fun4All'
order by  cust_name,cust_contact;

--
select prod_id,quantity,prod_id from orderitems where quantity>=100 or prod_id like ('BNBG%');

select prod_id,quantity from orderitems where  quantity >=100
union
select prod_id,quantity from orderitems where prod_id like ('BNBG%');

select prod_name from products
union
select cust_name from customers
order by prod_name ;
```

#### 14、插入数据

insert into table (....)

values(...)

```mysql
# Lesson 14

insert into customers
values (1000000006,
 'Toy Land',
 '123 Any Street',
 'New York',
 'NY',
 '11111',
 'USA',
 NULL,
 NULL)

-- 更安全的插入语句 表明了插入数据的顺序关系
-- 保证了即使表的结构改变，这条INSERT 语句仍然能正确工作。
insert into customers(cust_id,
 cust_name,
cust_address,
 cust_city,
cust_state,
 cust_zip,
 cust_country,
 cust_contact,
 cust_email)
VALUES(1000000006,
 'Toy Land',
 '123 Any Street',
 'New York',
 'NY',
 '11111',
 'USA',
 NULL,
 NULL);
 
-- 插入部分行
-- 省略的行 必须是允许null 或者是存在默认值 若是不满足以上条件，如果省略表中的值则会报错

INSERT INTO Customers(cust_id,
 cust_name,
cust_address,
 cust_city,
cust_state,
 cust_zip,
 cust_country)
VALUES(1000000006,
 'Toy Land',
 '123 Any Street',
 'New York',
 'NY',
 '11111',
 'USA');

INSERT INTO Customers(cust_id,
 cust_contact,
 cust_email,
 cust_name,
 cust_address,
 cust_city,
cust_state,
 cust_zip,
 cust_country)
SELECT cust_id,
 cust_contact,
 cust_email,
 cust_name,
 cust_address,
 cust_city,
 cust_state,
 cust_zip,
 cust_country
FROM CustNew;

#使用create table ... select .... 将另一个表中的数据复制到全新的表中
create  table custcopy as select * from customers;
```

#### 15、更新数据

```mysql
# lesson 15
# 更新数据 update .... set.... 
update customers set cust_email = 'kim@thetoystore.com' where cust_id = 1000000005;

update customers set cust_contact = 'sam roberts',cust_email = 'sam@toyland.com'
where cust_id = 1000000006;
-- -- DELETE 删除数据 删除表的内容而不是表
delete from customers where cust_id = 1000000006;
# 从表中删除所有行 使用 truncate table/delete 函数 因为不会记录操作过程 因此速度很快
```

#### 16、新建表

```mysql
#lesson 16
-- 每个表列要么是 NULL 列，要么是 NOT NULL 列，这种状态在创建时由表的定义规定。
-- 默认是NULL
create table wuhu(
    prod_id char(10) not null,
    vend_id char(10) not null ,
    prod_name char(254) not null,
    prod_prices decimal(8,2) not null,
    prod_desc varchar(1000) null
);

insert into wuhu(prod_id,vend_id,prod_name,prod_prices,prod_desc)
values ('dada','asdad','asdasdasd',3.14,'sadasd');

-- 只有不允许NULL值的列可作为主键，允许 NULL 值的列不能作为唯一标识。
-- 使用DEFAULT可以指定默认值
-- MySQL CURRENT_DATE()函数获取当前时间
create table orderitems1
(
    order_num int not null ,
    order_item int not null ,
    prod_id char(10) not null ,
    quantity int not null  default 1,
    item_price decimal(8,2) not null
);
```

#### 17、修改表

```mysql
-- 删除表
DROP TABLE custcpy;
-- 加列
alter table orderitems1 add vend_phone char(20);
-- 删除列
alter table  orderitems1 drop column vend_phone;
-- 
drop table orderitems1;
-- 重命名表
rename table wuhu to fifei;


alter table vendors add vend_wab char(200);

alter table vendors drop column vend_wab;
```

#### 18、视图

实际上就是创建了一个临时的表供使用

**视图本身不包含数据，因此返回的数据是从其他表中检索出来的。在添加或更改这些表中的数据时，视图将返回改变过的数据。**

```mysql
# lesson 18
-- CREATE VIWE.... AS SELECT.......
-- DROP VIEW....

select cust_name,cust_contact from customers,orders,orderitems
where customers.cust_id = orders.cust_id
  and orderitems.order_num = orders.order_num
   and prod_id = 'RGAN01';

-- 视图 就是虚拟的表 只包含使用时动态检索数据的查询。
-- 视图本身不包含数据，因此返回的数据是从其他表中检索出来的。在添加或更改这些表中的数据时，视图将返回改变过的数据。
-- 视图可以嵌套，即可以利用从其他视图中检索数据的查询来构造视图 但是视图的嵌套可能会严重降低查询的性能。
create view productcustomers as select cust_name,cust_contact,prod_id
from customers,orders,orderitems
where customers.cust_id = orders.cust_id and orderitems.order_num = orders.order_num;

select cust_name,cust_contact from  productcustomers where prod_id = 'RGAN01';

drop view productcustomers;

select * from vendors;

select concat(rtrim(vend_name),'(',rtrim(vend_country),')') as vend_title from vendors order by  vend_title;

create view vendorLocation as
    select concat(rtrim(vend_name),'(',rtrim(vend_country),')') as vend_title from vendors;

select * from vendorlocation;
drop view vendorlocation;

create view customerEmaillist as
select cust_id,cust_name,cust_email from customers where cust_email is not null;

select * from customerEmaillist;
drop view customerEmaillist;

select prod_id,quantity,item_price,quantity*item_price as expanded_price from orderitems where order_num = 20008;

create view OrderItemsExpandend as
    select order_num,prod_id,quantity,item_price,quantity*item_price as expanded_prices
from OrderItems;
select * from orderitemsexpandend;
drop view orderitemsexpandend;

-- 这一句就返回两个相同的 orders.cust_id和c.cust_id
select * from customers inner join orders on orders.cust_id = customers.cust_id;
-- 解决方法1 使用自然联结
create view CustomerWithOthers as
    select c.* from customers c inner join orders on orders.cust_id = c.cust_id;

drop view CustomerWithOthers;

-- 表明需要返回的列
CREATE VIEW CustomersWithOrders1 AS
SELECT Customers.cust_id,
       Customers.cust_name,
       Customers.cust_address,
       Customers.cust_city,
       Customers.cust_state,
       Customers.cust_zip,
       Customers.cust_country,
       Customers.cust_contact,
       Customers.cust_email
FROM Customers
JOIN Orders ON Customers.cust_id = Orders.cust_id;

SELECT * FROM CustomersWithOrders;
drop view CustomersWithOrders1;

CREATE VIEW OrderItemsExpanded AS
SELECT order_num,
       prod_id,
       quantity,
       item_price,
       quantity*item_price AS expanded_price
FROM OrderItems
ORDER BY order_num;
drop view orderitemsexpanded;
```

#### 19、存储过程

```mysql
# 存储过程就是为了以后使用而保存的一条或者多条SQL，说白了就是一条或者多条SQl语句封装在一起
# 并且创建和使用就像是函数一样
# 暂时没有例子

```

#### 20、事务

使用事务处理（transaction processing），通过确保成批的 SQL 操作要么 完全执行，要么完全不执行，来维护数据库的完整性。

#### 21、游标

SQL 检索操作返回一组称为结果集的行，这组返回的行都是与 SQL 语句 相匹配的行（零行到多行）。

但是的得到的结果都是完整的查询，可能包含0或者多行，

简单地使用 SELECT 语句，没有办法得到第 一行、下一行或前 10 行。有时，需要在检索出来的行中前进或后退一行或多行，这就是游标的用途所在。

游标（cursor）是一个存储在 DBMS 服务器上的数据库查询， 它不是一条 SELECT 语句，而是被该语句检索出来的结果集。

#### 22、高级SQL特性

```mysql
# lesson 20
# 1、约束，就是对出入的数据进行要求，其保证数据的合法性

-- 1、主键，任意两行的主键值都不相同，每行都有一个主键值 主键值不能重用 包含主键值的列从不修改或更新。
create table vendors1
(
    vend_id char(10) not null primary key,
    vend_name char(12) not null,
    vend_address char(50) null,
    vend_city char(50) null,
    vend_state char(50) null
);
-- 或者使用
alter table vendors1
add constraint primary key (vend_id);

-- 2、外键 一个表的某一列的值必须在另一个表中的主键中
create table orders1
(
    order_name int not null primary key ,
    order_data datetime not null ,
    cust_id char(10) not null references customers(cust_id)
);
-- 或者使用
alter table orders1
add constraint 键名 foreign key (cust_id) references customers(cust_id);
-- 对于删除
-- 外键的另一个作用是能防止意外删除，例如上一个例子，orders1中的cust_id来自于customer中的cust_id,
-- 那么若customer中需要删除某cust_id时，因为该cust_id与orders1中的相关联，所以不能删除，需要等到删除订单中的相关
-- cust_id对应的订单信后才能继续删除，即orders中的能直接删除，而customers中的不能直接删除

-- 对于插入，在orders1中插入的所有cust_id要求只有当在customers中有的话，才能插入，而customers中可以随便插入


-- 3、唯一约束，即保障某一列中的数据是唯一的且可以包含null项
create table employee
(
    emp_id char(10) not null primary key ,
    emp_nub char(50) null unique ,
    emp_sal char(10) not null
);

alter table employee
add constraint unique (emp_nub);

alter table employee
add column  num_data char(100) null ;
alter table employee
drop column num_data;


-- 4、约束检查，保证所有的值在某一个范围或者是特定的值
create table orderitems1
(
    order_num int not null ,
    order_item int not null ,
    prod_id char(10) not null ,
    quantity int not null check (quantity>0 )
    -- state char(1) not null  check ( in'[ST]' )
)
alter table orderitems1
add column state char(1) not null  check ( state like '[MF]')


create table orderitems2
(
    order_num int not null ,
    order_item int not null ,
    prod_id char(10) not null ,
    quantity int not null check ( quantity>0 ),
    state char(1) not null  check ( state like'[ST]' )
)

# 2、索引，强制对某一列进行排序以提高检索速率，主键默认进行排序
-- 虽然提高了检索的速率，但是因为在插入时需要维护这个索引关系，所以，会增加插入和删除的效率
-- CREATE INDEX .... ON TABEL(COLUMN)

create table orderitems3
(
    order_num int not null ,
    order_item int not null ,
    prod_id char(10) not null ,
    quantity int not null check ( quantity>0 ),
    state char(1) not null  check ( state like'[ST]' )
);

create index prod_id_index on orderitems3(prod_id);

# 3、触发器，
-- 在特定的数据库发生作用是自动触发
-- 触发器与单个的表相关联。
-- 这是一个样例，设置触发器，当其在插入新的数据时使cust_state 变为大写，
-- 这是 SQLserver的，在mysql中有些问题
create trigger customers_state on customers for insert,update as
    update customers set cust_state = upper(cust_state) where customers.cust_state = inserted.cust_id;

# 4、数据库安全，
-- 使用grant和revoke来管理，一般是管理员来进行设置
```





# MySQL进阶

## 存储引擎

MySQL体系结构：

- 连接层
- 服务层
- 引擎层
- 存储层

<img src=".\asset\images\image-20240325160345095.png" alt="image-20240325160345095" style="zoom:67%;" />

![存储引擎层](C:\Users\xlwwwww\Downloads\mysql_note-main\mysql_note-main\images\artic.png)

存储引擎就是存储数据、建立索引、更新/查询数据等技术的实现方式。**存储引擎是基于表而不是基于库的**，(一个数据库中的不同表能够使用不同类别的存储引擎），所以存储引擎也可以被称为表引擎。  
**默认存储引擎是InnoDB**。  

相关操作：  

```mysql
-- 查询建表语句
show create table account;  
-- 建表时指定存储引擎
CREATE TABLE 表名(
	...
) ENGINE=INNODB;
-- 查看当前数据库支持的存储引擎
show engines;
```

### InnoDB

InnoDB 是一种**兼顾高可靠性和高性能**的通用存储引擎，在 MySQL 5.5 之后，InnoDB 是默认的 MySQL 引擎。  

特点：

- DML 操作遵循 ACID 模型，支持**事务** 
- **行级锁**，提高并发访问性能
- 支持**外键**约束，保证数据的完整性和正确性

文件：

- xxx.ibd: xxx代表表名，**InnoDB 引擎的每张表都会对应这样一个表空间文件**，存储该表的表结构（frm、sdi）、数据和索引。

参数：innodb_file_per_table，决定多张表共享一个表空间还是每张表对应一个表空间

知识点：

查看 Mysql 变量：
`show variables like 'innodb_file_per_table';`

从idb文件提取表结构数据：
（在cmd运行）  
`ibd2sdi xxx.ibd`

InnoDB 逻辑存储结构：  

<img src=".\asset\images\image-20240325162519611.png" alt="image-20240325162519611" style="zoom: 80%;" />

- 表空间 : InnoDB存储引擎逻辑结构的最高层，`ibd`文件其实就是表空间文件，在表空间中可以 包含多个Segment段
- 段 : 表空间是由各个段组成的， 常见的段有数据段、索引段、回滚段等。InnoDB中对于段的管 理，都是引擎自身完成，不需要人为对其控制，一个段中包含多个区。
- 区 : 区是表空间的单元结构，每个区的大小为1M。 默认情况下， InnoDB存储引擎页大小为 16K， 即一个区中一共有64个连续的页。
- 页 : 页是组成区的最小单元，页也是InnoDB 存储引擎磁盘管理的最小单元，每个页的大小默 认为 16KB。**为了保证页的连续性，InnoDB 存储引擎每次从磁盘申请 4-5 个区。**
- 行 : InnoDB 存储引擎是面向行的，也就是说数据是按行进行存放的，在每一行中除了定义表时 所指定的字段以外，还包含两个隐藏字段(后面会详细介绍)。行中存储的数据就是表中存储的每一行的数据。

### MyISAM

MyISAM 是 MySQL 早期的默认存储引擎。

特点：

- 不支持事务，不支持外键
- 支持表锁，不支持行锁
- 访问速度快

文件：

- xxx.sdi: 存储表结构信息 存储形式为JSON文本
- xxx.MYD: 存储数据
- xxx.MYI: 存储索引

### Memory

Memory 引擎的表数据是存储在内存中的，受硬件问题、断电问题的影响，只能将这些表作为临时表或缓存使用。

特点：

- 存放在内存中，速度快
- hash索引（默认）

文件：

- xxx.sdi: 存储表结构信息 数据都存储在内存中

### 存储引擎特点

| 特点         | InnoDB              | MyISAM | Memory |
| ------------ | ------------------- | ------ | ------ |
| 存储限制     | 64TB                | 有     | 有     |
| 事务安全     | **支持**            | -      | -      |
| 锁机制       | **行锁**            | 表锁   | 表锁   |
| B+tree索引   | 支持                | 支持   | 支持   |
| Hash索引     | -                   | -      | 支持   |
| 全文索引     | 支持（5.6版本之后） | 支持   | -      |
| 空间使用     | 高                  | 低     | N/A    |
| 内存使用     | 高                  | 低     | 中等   |
| 批量插入速度 | 低                  | 高     | 高     |
| 支持外键     | **支持**            | -      | -      |

### 存储引擎的选择

在选择存储引擎时，应该根据应用系统的特点选择合适的存储引擎。对于复杂的应用系统，还可以根据实际情况选择多种存储引擎进行组合。

- InnoDB: 如果应用对事物的完整性有比较高的要求，在并发条件下要求数据的一致性，数据操作除了插入和查询之外，还包含很多的更新、删除操作，则 InnoDB 是比较合适的选择
- MyISAM: 如果应用是以读操作和插入操作为主，只有很少的更新和删除操作，并且对事务的完整性、并发性要求不高，那这个存储引擎是非常合适的。
- Memory: 将所有数据保存在内存中，访问速度快，通常用于临时表及缓存。Memory 的缺陷是对表的大小有限制，太大的表无法缓存在内存中，而且无法保障数据的安全性

电商中的足迹和评论适合使用 MyISAM 引擎，缓存适合使用 Memory 引擎。

## 性能分析

### 查看执行频次

查看当前数据库的 INSERT, UPDATE, DELETE, SELECT 访问频次：  
`SHOW GLOBAL STATUS LIKE 'Com_______';` 或者 `SHOW SESSION STATUS LIKE 'Com_______';`  
例：`show global status like 'Com_______'`  

### 慢查询日志

慢查询日志记录了所有执行时间超过指定参数（long_query_time，单位：秒，默认10秒）的所有SQL语句的日志。  
MySQL的慢查询日志默认没有开启，需要在MySQL的配置文件（/etc/my.cnf）中配置如下信息：  
	# 开启慢查询日志开关  
	slow_query_log=1  
	# 设置慢查询日志的时间为2秒，SQL语句执行时间超过2秒，就会视为慢查询，记录慢查询日志  
	long_query_time=2  
更改后记得重启MySQL服务，日志文件位置：/var/lib/mysql/localhost-slow.log  

查看慢查询日志开关状态：  
`show variables like 'slow_query_log';`  

### profile

show profile 能在做SQL优化时帮我们了解时间都耗费在哪里。通过 have_profiling 参数，能看到当前 MySQL 是否支持 profile 操作：  
`SELECT @@have_profiling;`  
profiling 默认关闭，可以通过set语句在session/global级别开启 profiling：  
`SET profiling = 1;`  
查看所有语句的耗时：  
`show profiles;`  
查看指定query_id的SQL语句各个阶段的耗时：  
`show profile for query query_id;`  
查看指定query_id的SQL语句CPU的使用情况  
`show profile cpu for query query_id;`  

### explain

EXPLAIN 或者 DESC 命令获取 MySQL 如何执行 SELECT 语句的信息，包括在 SELECT 语句执行过程中表如何连接和连接的顺序。  
语法：  
	# 直接在select语句之前加上关键字 explain / desc  
	EXPLAIN SELECT 字段列表 FROM 表名 HWERE 条件;  

EXPLAIN 各字段含义：  

- id：select 查询的序列号，表示查询中执行 select 子句或者操作表的顺序（id相同，执行顺序从上到下；id不同，值越大越先执行）  
- select_type：表示 SELECT 的类型，常见取值有 SIMPLE（简单表，即不适用表连接或者子查询）、PRIMARY（主查询，即外层的查询）、UNION（UNION中的第二个或者后面的查询语句）、   SUBQUERY（SELECT/WHERE之后包含了子查询）等  
- type：表示连接类型，性能由好到差的连接类型为 NULL、system、const、eq_ref、ref、range、index、all
- possible_key：可能应用在这张表上的索引，一个或多个
- Key：实际使用的索引，如果为 NULL，则没有使用索引
- Key_len：表示索引中使用的字节数，该值为索引字段最大可能长度，并非实际使用长度，在不损失精确性的前提下，长度越短越好
- rows：MySQL认为必须要执行的行数，在InnoDB引擎的表中，是一个估计值，可能并不总是准确的
- filtered：表示返回结果的行数占需读取行数的百分比，filtered的值越大越好

## 索引

索引是帮助 MySQL **高效获取数据**的**数据结构（有序）**。在数据之外，数据库系统还维护着满足特定查找算法的数据结构，这些数据结构以某种方式引用（指向）数据，这样就可以在这些数据结构上实现高级查询算法，这种数据结构就是索引。

**索引是再存储引擎层实现的。**

优缺点：

优点：

- 提高数据检索效率，降低数据库的IO成本  
- 通过索引列对数据进行排序，降低数据排序的成本，降低CPU的消耗

缺点：

- 索引列也是要占用空间的
- 索引大大提高了查询效率，但降低了更新的速度，比如 INSERT、UPDATE、DELETE

### 索引结构

| 索引结构            | 描述                                                         |
| ------------------- | ------------------------------------------------------------ |
| B+Tree              | 最常见的索引类型，大部分引擎都支持B+树索引                   |
| Hash                | 底层数据结构是用哈希表实现，只有精确匹配索引列的查询才有效，不支持范围查询 |
| R-Tree(空间索引)    | 空间索引是 MyISAM 引擎的一个特殊索引类型，主要用于地理空间数据类型，通常使用较少 |
| Full-Text(全文索引) | 是一种通过建立倒排索引，快速匹配文档的方式，类似于 Lucene, Solr, ES |

| 索引       | InnoDB        | MyISAM | Memory |
| ---------- | ------------- | ------ | ------ |
| B+Tree索引 | 支持          | 支持   | 支持   |
| Hash索引   | 不支持        | 不支持 | 支持   |
| R-Tree索引 | 不支持        | 支持   | 不支持 |
| Full-text  | 5.6版本后支持 | 支持   | 不支持 |

#### 二叉树

<figure class="half">
    <img src=".\asset\images\image-20240325170645732.png" style="zoom: 65%;" \>
    <img src=".\asset\images\image-20240325170700235.png" style="zoom: 63%;" \>
</figure>



这里实际上构建的是二叉搜索树，当最坏的情况下，所有二叉树的结点都在同一侧，就变成了链表。查询就没有优势可言

所以，如果选择二叉树作为索引结构，会存在以下缺点：

- 顺序插入时，会形成一个链表，查询性能大大降低。 
- 大数据量情况下，层级较深，检索速度慢

二叉树形成链表的缺点可以用红黑树来解决：红黑树是一颗自平衡二叉树，那这样即使是顺序插入数据，最终形成的数据结构也是一颗平衡的二叉树

<img src=".\asset\images\image-20240325171706442.png" alt="红黑树" title="红黑树" style="zoom:50%;" />

#### B-Tree

红黑树也存在大数据量情况下，层级较深，检索速度慢的问题。

为了解决上述问题，可以使用 B-Tree 结构。
B-Tree (多路平衡查找树) 以一棵最大度数（max-degree，指一个节点的子节点个数）为5（5阶）的 b-tree 为例（每个节点最多存储4个key，5个指针）

![image-20240325171905888](.\asset\images\image-20240325171905888.png)

> B-Tree 的数据插入过程动画参照：https://www.bilibili.com/video/BV1Kr4y1i7ru?p=68
> 演示地址：https://www.cs.usfca.edu/~galles/visualization/BTree.html
>
> 当节点处插入的数据多余树的度的要求时，会使中间节点向上分裂，将分裂出的节点加入上层，一层一层递归向上调整，直至所有节点关系都满足树的度的约束。

#### B+Tree

结构图：

![image-20240325172615971](.\asset\images\image-20240325172615971.png)

> 演示地址：https://www.cs.usfca.edu/~galles/visualization/BPlusTree.html

与 B-Tree 的区别：

- 所有的数据都会出现在叶子节点
- 叶子节点形成一个单向链表

MySQL 索引数据结构对经典的 B+Tree 进行了优化。在原 B+Tree 的基础上，增加一个指向相邻叶子节点的链表指针，就形成了带有顺序指针的 B+Tree，提高区间访问的性能。

![image-20240325172954139](.\asset\images\image-20240325172954139.png)

#### Hash

哈希索引就是采用一定的hash算法，将键值换算成新的hash值，映射到对应的槽位上，然后存储在hash表中。
如果两个（或多个）键值，映射到一个相同的槽位上，他们就产生了hash冲突（也称为hash碰撞），可以通过链表来解决。

![image-20240325173019646](.\asset\images\image-20240325173019646.png)

特点：

- Hash索引只能用于对等比较（=、in），不支持范围查询（between、>、<、...）
- 无法利用索引完成排序操作
- 查询效率高，通常只需要一次检索就可以了，效率通常要高于 B+Tree 索引

存储引擎支持：

- Memory
- InnoDB: 具有自适应hash功能，hash索引是存储引擎根据 B+Tree 索引在指定条件下**自动构建**的

#### 面试题

1. 为什么 InnoDB 存储引擎选择使用 B+Tree 索引结构？

- 相对于二叉树，层级更少，搜索效率高
- 对于 B-Tree，无论是叶子节点还是非叶子节点，都会保存数据，这样导致一页中存储的键值减少，指针也跟着减少，要同样保存大量数据，只能增加树的高度，导致性能降低
- 相对于 Hash 索引，B+Tree 支持范围匹配及排序操作

### 索引分类

| 分类     | 含义                                                 | 特点                                                    | 关键字   |
| -------- | ---------------------------------------------------- | ------------------------------------------------------- | -------- |
| 主键索引 | 针对于表中主键创建的索引                             | 默认**自动创建**，只能有一个                            | PRIMARY  |
| 唯一索引 | 避免同一个表中某数据列中的值重复，**null不算重复**   | 可以有多个,限定为唯一约束的字段，会**自动创建唯一索引** | UNIQUE   |
| 常规索引 | 快速定位特定数据                                     | 可以有多个                                              |          |
| 全文索引 | 全文索引查找的是文本中的关键词，而不是比较索引中的值 | 可以有多个                                              | FULLTEXT |

在 InnoDB 存储引擎中，根据索引的存储形式，又可以分为以下两种：

| 分类                      | 含义                                                         | 特点                 |
| ------------------------- | ------------------------------------------------------------ | -------------------- |
| 聚集索引(Clustered Index) | 将数据存储与索引放一块，索引结构的叶子节点**保存了行数据和主键** | 必须有，而且只有一个 |
| 二级索引(Secondary Index) | 将数据与索引分开存储，索引结构的叶子节点关联的是对应的**主键** | 可以存在多个         |

演示图：

![image-20240325174428640](.\asset\images\image-20240325174428640.png)



- 聚集索引的叶子节点下挂的是这一行的数据 。
- **二级索引的叶子节点下挂的是该字段值对应的主键值**
- 因为二级索引叶子节点并未挂载数据，所以通过二级索引的查询需要再通过聚集索引进行数据提取，这个现象称为回表查询。





![演示图](https://learning-logs-1253130399.cos.ap-guangzhou.myqcloud.com/editor/%E6%BC%94%E7%A4%BA%E5%9B%BE_20220319215403721066.png)

聚集索引选取规则：

- 如果存在主键，主键索引就是聚集索引
- 如果不存在主键，将使用第一个唯一(UNIQUE)索引作为聚集索引
- 如果表没有主键或没有合适的唯一索引，则 InnoDB 会自动生成一个 rowid 作为隐藏的聚集索引

#### 思考题

1\. 以下 SQL 语句，哪个执行效率高？为什么？

```mysql
select * from user where id = 10;
select * from user where name = 'Arm';
-- 备注：id为主键，name字段创建的有索引
```

答：第一条语句，因为第二条需要回表查询，相当于两个步骤。

2\. InnoDB 主键索引的 B+Tree 高度为多少？

答：假设一行数据大小为1k，一页中可以存储16行这样的数据。InnoDB 的指针占用6个字节的空间，主键假设为bigint，占用字节数为8.  
可得公式：`n * 8 + (n + 1) * 6 = 16 * 1024`，其中 8 表示 bigint 占用的字节数，n 表示当前节点存储的key的数量，(n + 1) 表示指针数量（比key多一个）。算出n约为1170。  

如果树的高度为2，那么他能存储的数据量大概为：`1171 * 16 = 18736`；  
如果树的高度为3，那么他能存储的数据量大概为：`1171 * 1171 * 16 = 21939856`。  

另外，如果有成千上万的数据，那么就要考虑分表，涉及运维篇知识。  

### 语法

创建索引：  
`CREATE [ UNIQUE | FULLTEXT ] INDEX index_name ON table_name (index_col_name, ...);`  
如果不加 CREATE 后面不加索引类型参数，则创建的是常规索引  

查看索引：  
`SHOW INDEX FROM table_name;`  

删除索引：  
`DROP INDEX index_name ON table_name;`

案例：

```mysql
-- name字段为姓名字段，该字段的值可能会重复，为该字段创建索引
create index idx_user_name on tb_user(name);

-- phone手机号字段的值非空，且唯一，为该字段创建唯一索引
create unique index idx_user_phone on tb_user (phone);

-- 为profession, age, status创建联合索引
create index idx_user_pro_age_stat on tb_user(profession, age, status);

-- 为email建立合适的索引来提升查询效率
create index idx_user_email on tb_user(email);

-- 删除索引  
drop index idx_user_email on tb_user;  
```

### 性能分析工具

#### 服务器状态信息

**`show [session|global] status`**

```mysql
-- session 是查看当前会话 ;
-- global 是查询全局数据 ;

SHOW  GLOBAL STATUS LIKE  'Com_______'; 

-- Com_delete: 删除次数
-- Com_insert: 插入次数
-- Com_select: 查询次数
-- Com_update: 更新次数
```

![image-20240325185942001](.\asset\images\image-20240325185942001.png)

#### 慢查询日志

慢查询日志记录了所有执行时间超过指定参数（long_query_time，单位：秒，默认10秒）的所有SQL语句的日志。

```mysql
-- 查询慢查询日志默认没有开启
show variables like 'slow_query_log';

-- 开启MySQL慢日志查询开关 慢查询日志记录在 /var/lib/mysql/机器名-slow.log文件中
set global slow_query_log=1

# 设置慢日志的时间，这里设置为SQL语句执行时间超过2秒，就会视为慢查询，记录慢查询日志
long_query_time=2
```

通过慢查询日志定位查询效率低的语句进行优化。

#### profile详情

```mysql
-- 当前MySQL是否支持profile操作
SELECT  @@have_profiling ;

-- 开启profiling：可以选择 session/global 级别 但这里是临时的，mysql重启后需重新设置
SET  profiling = 1;

-- 查看每一条SQL的耗时基本情况
show profiles;

-- 查看指定query_id的SQL语句各个阶段的耗时情况
show profile  for  query query_id;

-- 查看指定query_id的SQL语句CPU的使用情况
show profile  cpu for  query query_id;
```

#### explain/DESC

EXPLAIN 或者 DESC命令**获取 MySQL 如何执行SELECT 语句的信息**，包括在 SELECT 语句执行 过程中表如何连接和连接的顺序。

```mysql
-- 直接在select语句之前加上关键字 explain / desc
 EXPLAIN SELECT 字段列表 FROM   表名   WHERE  条件 ;
```

![image-20240325222425020](.\asset\images\image-20240325222425020.png)

各字段的意义：

- | 字段         | 含义                                                         |
  | ------------ | ------------------------------------------------------------ |
  | id           | select查询的序列号，表示查询中执行select子句或者是操作表的顺序 (id相同，执行顺序从上到下；id不同，值越大，越先执行)。 |
  | select_type  | 表示 SELECT 的类型，常见的取值有 SIMPLE（简单表，即不使用表连接或者子查询）、PRIMARY（主查询，即外层的查询）、 UNION（UNION 中的第二个或者后面的查询语句）、 SUBQUERY（SELECT/WHERE之后包含了子查询）等 |
  | type         | 表示连接类型，性能由好到差的连接类型为NULL、system、const、 eq_ref、ref、range、 index、all 。                 其中，不用查询表时性能最好，为null；![image-20240325222650348](.\asset\images\image-20240325222650348.png)查询索引为unique或primary时，为`const`![image-20240325222846572](.\asset\images\image-20240325222846572.png)普通索引，为`ref`![image-20240325223038185](.\asset\images\image-20240325223038185.png)不使用索引时，最差，为`ALL`![image-20240325223201200](.\asset\images\image-20240325223201200.png) |
  | possible_key | 显示可能应用在这张表上的索引，一个或多个。                   |
  | key          | 实际使用的索引，如果为NULL，则没有使用索引。                 |
  | key_len      | 表示索引中使用的字节数， 该值为索引字段最大可能长度，并非实际使用长 度，在不损失精确性的前提下， 长度越短越好。这个一定程度上与创建索引的列信息有关，该列下数据越复杂，创建的索引的使用的字节数越大。 |
  | rows         | MySQL认为必须要执行查询的行数，在innodb引擎的表中，是一个估计值， 可能并不总是准确的。 |
  | filtered     | 表示返回结果的行数占需读取行数的百分比， filtered 的值越大越好。 |
  | extra        | 字段含义：<br/>`using index condition`：查找使用了索引，但是需要回表查询数据<br/>`using where; using index;`：查找使用了索引，但是需要的数据都在索引列中能找到，所以不需要回表查询这里的显示不准 |



### 使用规则

#### 最左前缀法则

如果索引关联了多列（联合索引），要遵守**最左前缀法则**，

- 最左前缀法则指的是查询**从索引的最左列开始，并且不跳过索引中的列**。  
- 如果跳跃某一列，索引将部分失效（后面的字段索引失效）。跳过的话，后面的排序就无从说起了。
- **最左前缀法则在用select的时候，和放的位置是没有关系的，只要存在就行。**  



#### 索引失效情况

1. 联合索引中，出现范围查询（<, >），范围查询右侧的列索引失效。可以用>=或者<=来规避索引失效问题

   ![image-20240325225931670](.\asset\images\image-20240325225931670.png)

2. **在索引列上进行运算操作，索引将失效**。如：`explain select * from tb_user where substring(phone, 10, 2) = '15';` 换成 `explain select * from tb_user where phone = '17799990015';`这是可以的。  

   ![image-20240325230410646](.\asset\images\image-20240325230410646.png)

3. 字符串类型字段使用时，不加引号，查询没有问题，但是索引将失效。因为规定了列，mysql会实施内置类型转换。如：`explain select * from tb_user where phone = 17799990015;`，此处phone的值没有加引号  

   ![image-20240325230506035](.\asset\images\image-20240325230506035.png)

4. 模糊查询中，

   - 如果仅仅是尾部模糊匹配，索引不会是失效；

   - 如果是头部模糊匹配，索引失效。如：`explain select * from tb_user where profession like '%工程';`，
   - 前后都有 % 也会失效。`explain select * from tb_user where profession like '软件%';` 这个是不会失效的，只有前面加了%才会失效。  

   ![image-20240325230828221](.\asset\images\image-20240325230828221.png)

5. 用 or 分割开的条件，如果 or 其中一个条件的列没有索引，那么涉及的索引都不会被用到。

   ![image-20240325231146688](.\asset\images\image-20240325231146688.png)

6. 如果 MySQL 评估使用索引比全表更慢，则不使用索引。因为只要有一个没有索引，另外一个用不用索引都没有意义，都要进行全表扫描。所以就无需用索引。MySQL的评估与表的数据结构有关。

   ![image-20240325231411089](.\asset\images\image-20240325231411089.png)

#### SQL 提示

是优化数据库的一个重要手段，简单来说，就是在SQL语句中加入一些人为的提示来达到优化操作的目的。  

例如，使用索引：
`explain select * from tb_user use index(idx_user_pro) where profession="软件工程";`
不使用哪个索引：
`explain select * from tb_user ignore index(idx_user_pro) where profession="软件工程";`
必须使用哪个索引：
`explain select * from tb_user force index(idx_user_pro) where profession="软件工程";`

use 是建议，实际使用哪个索引 MySQL 还会自己权衡运行速度去更改，force就是无论如何都强制使用该索引。

![image-20240325232303052](.\asset\images\image-20240325232303052.png)

#### 覆盖索引&回表查询

尽量使用覆盖索引（查询使用了索引，并且需要返回的列，在该索引中已经全部能找到），减少 select *。

![image-20240326005321605](.\asset\images\image-20240326005321605.png)

`using index;using where`  :表示使用了索引，并且需要的数据都在索引列能找到，不需要回表查询(二级索引，叶子节点存储的是主键值【id】，因此这里需要的信息都全了，不需要回表)

`using index condition `:查找使用了索引，但是需要回表查询数据

覆盖索引：
如果在生成的二级索引（辅助索引）中可以一次性获得select所需要的字段，不需要回表查询。

- 在聚集索引中，叶子节点保存的是数据行，能够直接能找到对应的数据，则直接返回行数据，只需要一次查询，哪怕是select \*；  
- 在辅助索引（二级索引）中，叶子节点存储的是主键值，能够找到对应的id，profession和age,states，只需要一次查询； 
- 如果是通过辅助索引查找其他字段，则需要回表查询，如`select id, name, gender from xxx where name='xxx';`  

**所以尽量不要用`select *`，容易出现回表查询，降低效率，除非有联合索引包含了所有字段**  

面试题：一张表，有四个字段（id, username, password, status），由于数据量大，需要对以下SQL语句进行优化，该如何进行才是最优方案：
`select id, username, password from tb_user where username='itcast';`

解：给username和password字段建立联合索引，则不需要回表查询，直接覆盖索引。
username和password字段建立联合索引的叶子节点挂的就是 id 所以不需要三者同时建索引。

#### 前缀索引

- 当字段类型为字符串（varchar, text等）时，有时候需要索引很长的字符串，这会让索引变得很大，查询时，浪费大量的磁盘IO，影响查询效率，此时可以只降字符串的一部分前缀，建立索引，这样可以大大节约索引空间，从而提高索引效率。


- 语法：`create index idx_xxxx on table_name(columnn(n));`  
- 前缀长度：可以根据索引的选择性来决定，而选择性是指**不重复的索引值（基数）和数据表的记录总数的比值**，索引选择性越高则查询效率越高，唯一索引的选择性是1，这是最好的索引选择性，性能也是最好的。  

- 求选择度公式：  


```mysql
select count(distinct email) / count(*) from tb_user;
select count(distinct substring(email, 1, 5)) / count(*) from tb_user;
```

前缀索引中是有可能碰到相同的索引的情况的（因为选择性可能不为1），所以使用前缀索引进行查询的时候，mysql 会有一个回表查询的过程，确定是否为所需数据。

如图中的查询到`lvbu6`之后还要进行回表，得到完整的email数据进行比对，与查询的要求相符，继续查询索引中下一个位置的信息，即`xiaoy`，看到xiaoy是不需要的数据，则停止查下一个。

![pre](C:\Users\xlwwwww\Downloads\mysql_note-main\mysql_note-main\images\pre.png)

show index 里面的sub_part可以看到接取的长度

#### 单列索引&联合索引

单列索引：即一个索引只包含单个列
联合索引：即一个索引包含了多个列
在业务场景中，如果存在多个查询条件，考虑针对于查询字段建立索引时，建议建立联合索引，而非单列索引。

单列索引情况：

`explain select id, phone, name from tb_user where phone = '17799990010' and name = '韩信';`  
phone 和 name 都建立了索引情况下，这句只会用到phone索引字段。

##### 注意事项

- 多条件联合查询时，MySQL优化器会评估哪个字段的索引效率更高，会选择该索引完成本次查询。

### 设计原则

1. 针对于数据量较大，且**查询比较频繁**的表建立索引 （100W）
2. 针对于常作为查询条件（where）、排序（order by）、分组（group by）操作的字段建立索引
3. 尽量选择区分度高的列作为索引，尽量建立唯一索引，区分度越高，使用索引的效率越高
4. 如果是字符串类型的字段，字段长度较长，可以针对于字段的特点，建立前缀索引
5. 尽量使用联合索引，减少单列索引，查询时，联合索引很多时候可以覆盖索引，节省存储空间，避免回表，提高查询效率
6. 要控制索引的数量，索引并不是多多益善，索引越多，维护索引结构的代价就越大，会影响增删改的效率
7. 如果索引列不能存储NULL值，请在创建表时使用NOT NULL约束它。当优化器知道每列是否包含NULL值时，它可以更好地确定哪个索引最有效地用于查询

## SQL 优化

### 插入数据

普通插入：

1. 采用批量插入（一次插入的数据不建议超过1000条，500 - 1000 为宜）

2. 手动提交事务

   每一次1行一行的输入insert 实际上是创建了隐含的事务，每次一条一条的输入会频繁的开启关闭事务

3. 主键顺序插入（主键顺序插入的效率大于乱序插入

4. 大批量插入：
   如果一次性需要插入大批量数据，使用insert语句插入性能较低，此时可以使用MySQL数据库提供的load指令插入。

```mysql
# 客户端连接服务端时，加上参数 --local-infile（这一行在bash/cmd界面输入）
mysql --local-infile -u root -p
# 设置全局参数local_infile为1，开启从本地加载文件导入数据的开关
set global local_infile = 1;
select @@local_infile;
# 执行load指令将准备好的数据，加载到表结构中，先要把表建立起来。   
-- fields terminated by ',' 每一个字段的数据以','分割
-- 每一行的数据以 '\n' 分割
load data local infile '/root/sql1.log' into table 'tb_user' fields terminated by ',' lines terminated by '\n';
```

### 主键优化

数据组织方式：在InnoDB存储引擎中，**表数据都是根据主键顺序组织存放的**，这种存储方式的表称为索引组织表（Index organized table, IOT）

主键的顺序的插入过程如下：  

如果是顺序插入，每一页中存储的数据行是按顺序填满的，因此这一页存满后，下一页的页申请是连续的，

而如果是乱序插入，在这一页数据存满后，如果需要向这一页中间的内容插入数据，那么会发生页分裂的过程。

页分裂：

- 包括申请页
- 满页以50%分裂 拷贝到新申请的页
- 修改指针指向

页合并：当删除一行记录时，实际上记录并没有被物理删除，只是记录被标记（flaged）为删除并且它的空间变得允许被其他记录声明使用。当**页中删除的记录到达 MERGE_THRESHOLD**（默认为页的50%），InnoDB会开始寻找最靠近的页（前后）看看是否可以将这两个页合并以优化空间使用。

MERGE_THRESHOLD：合并页的阈值，可以自己设置，在创建表或创建索引时指定

> 文字说明不够清晰明了，具体可以看视频里的PPT演示过程：https://www.bilibili.com/video/BV1Kr4y1i7ru?p=90

主键设计原则：

- 满足业务需求的情况下，**尽量降低主键的长度**，二级索引的叶子节点保存的就是主键，所以主键小占用的空间也就会少。
- 插入数据时，尽量**选择顺序插入**，选择使用 **AUTO_INCREMENT 自增主键**
- 尽量不要使用 UUID (生成随机，实际上没顺序）做主键或者是其他的自然主键，如身份证号，占用的空间大。
- 业务操作时，避免对主键的修改

### order by优化

1. Using filesort：通过表的索引或全表扫描，读取满足条件的数据行，然后**在排序缓冲区 sort buffer 中完成排序操作**，所有不是通过索引直接返回排序结果的排序都叫 FileSort 排序
2. Using index：通过**有序索引顺序扫描直接返回有序数据**，这种情况即为 using index，不需要额外排序，操作效率高

如果order by字段**全部使用升序排序或者降序排序**，则都会走索引，但是如果一个字段升序排序，另一个字段降序排序，则不会走索引，explain的extra信息显示的是`Using index, Using filesort`

![image-20240326110709245](.\asset\images\image-20240326110709245.png)

如果要优化掉Using filesort，则需要另外再创建一个索引，

如：`create index idx_user_age_phone_ad on tb_user(age asc, phone desc);`，此时使用`select id, age, phone from tb_user order by age asc, phone desc;`会全部走索引。

![image-20240326110959821](.\asset\images\image-20240326110959821.png)

总结：

- 根据排序字段建立合适的索引，多字段排序时，也遵循最左前缀法则

- 尽量使用覆盖索引

- 多字段排序，一个升序一个降序，此时需要注意联合索引在创建时的规则（ASC/DESC）

- 如果不可避免出现filesort，大数据量排序时，可以适当增大排序缓冲区大小 sort_buffer_size（默认256k）

  `show variables like 'sort_buffer_size';`

### group by优化

- 在分组操作时，可以通过索引来提高效率
- 分组操作时，索引的使用也是满足最左前缀法则的

如索引为`idx_user_pro_age_stat`，则句式可以是`select ... where profession = 'xxx" order by age`，这样也符合最左前缀法,相当于先使用了profession的索引，再使用age的索引。

![image-20240326112341302](.\asset\images\image-20240326112341302.png)

### limit优化

大数据量情况下，offset 起始位置约往后，需要的处理的时间越长

常见的问题如`limit 2000000, 10`，此时需要 MySQL 排序前2000000条记录，但仅仅返回2000000 - 2000010的记录，其他记录丢弃，查询排序的代价非常大。
优化方案：一般分页查询时，通过创建覆盖索引能够比较好地提高性能，可以通过覆盖索引加子查询形式进行优化

例如：

```mysql
-- 此语句耗时很长
select * from tb_sku limit 9000000, 10;
-- 通过覆盖索引加快速度，直接通过主键索引进行排序及查询
select id from tb_sku order by id limit 9000000, 10;

-- 下面的语句是错误的，因为 MySQL 不支持 in 里面使用 limit
-- select * from tb_sku where id in (select id from tb_sku order by id limit 9000000, 10);

-- 通过连表查询即可实现第一句的效果，并且能达到第二句的速度
select * from tb_sku as s, (select id from tb_sku order by id limit 9000000, 10) as a where s.id = a.id;
```

### count优化

**MyISAM 引擎把一个表的总行数存在了磁盘上，因此执行 count(\*) 的时候会直接返回这个数，效率很高**（前提是不使用where，返回的是整个表的总行数）；  
InnoDB 在执行 count(\*) 时，需要把数据一行一行地从引擎里面读出来，然后累计计数。  
优化方案：自己计数，如创建key-value表存储在内存或硬盘，或者是用redis  

count的几种用法：

- 如果count函数的参数（count里面写的那个字段）不是NULL（字段值不为NULL），累计值就加一，最后返回累计值 
- 用法：count(\*)、count(主键)、count(字段)、count(1)
- count(主键)跟count(\*)一样，因为主键不能为空；count(字段)只计算字段值不为NULL的行；count(1)引擎会为每行添加一个1，然后就count这个1，返回结果也跟count(\*)一样；count(null)返回0

各种用法的性能：

- count(主键)：InnoDB引擎会遍历整张表，把每行的主键id值都取出来，返回给服务层，服务层拿到主键后，直接按行进行累加（主键不可能为空）
- count(字段)：
  - 没有not null约束的话，InnoDB引擎会遍历整张表把每一行的字段值都取出来，返回给服务层，服务层判断是否为null，**不为null，计数累加**；
  - 有not null约束的话，InnoDB引擎会遍历整张表把每一行的字段值都取出来，返回给服务层，直接按行进行累加
- count(1)：InnoDB 引擎遍历整张表，但不取值。服务层对于返回的每一层，放一个数字 1 进去，直接按行进行累加
- count(\*)：InnoDB 引擎并不会把全部字段取出来，而是专门做了优化，不取值，服务层直接按行进行累加

按效率排序：count(字段) < count(主键) < count(1) < count(\*)，所以尽量使用 count(\*)

### update优化（避免行锁升级为表锁）

InnoDB 的行锁是**针对索引加的锁，不是针对记录加的锁**，并且该索引不能失效，否则会从行锁升级为表锁。

如以下两条语句：
`update student set no = '123' where id = 1;`，这句由于id有主键索引，所以只会锁这一行；  
`update student set no = '123' where name = 'test';`，这句由于name没有索引，所以会把整张表都锁住进行数据更新，解决方法是给name字段添加索引，就可以由表锁变成行锁。  


## 视图

视图（View）是一种**虚拟存在的表**。视图中的数据并不在数据库中实际存在，**行和列数据来自定义视图的查询中使用的表**（基表），并且是在使用视图时动态生成的。  
通俗的讲，视图只保存了查询的SQL逻辑，不保存查询结果。所以我们在创建视图的时候，主要的工作就落在创建这条SQL查询语句上。  

### 创建视图

`CREATE [ OR REPLACE ] VIEW 视图名称[（列名列表）] AS SELECT 语句 [ WITH [ CASCADED | LOCAL ] CHECK OPTION ]`  

> 例子： ` create or replace view stu_wll as select id,name from student where id<=10; `

### 查询视图

查看创建视图语句： `SHOW CREATE VIEW `视图名称；  

查看视图数据：`SELECT * FROM ` 视图名称；  
`show create view stu_v_1;`

### 修改视图

方式一：`CREATE[OR REPLACE] VIEW 视图名称[（列名列表)）] AS SELECT 语句[ WITH[ CASCADED | LOCAL ] CHECK OPTION ]`

方式二：`ALTER VIEW 视图名称 [（列名列表)] AS SELECT语句 [WITH [CASCADED | LOCAL] CHECK OPTION]`

### 删除视图

能够一次删除多个视图

`DROP VIEW [IF EXISTS] 视图名称 [视图名称]`

### 视图检查选项

- 当使用WITH CHECK QPTION子句创建视图时，MySQL会通过视图检查正在更改的每个行，例如插入，更新，删除，以使其符合视图的定义。

- MySQL允许基于另一个视图创建视图，它还会检查依赖视图中的规则以保持一致性。

  为了确定检查的范围，mysql提供了两个选项：**CASCADED 和 LOCAL ，默认值为 CASCADED。**

  - CASCADED

    对于最后的视图，再该视图插入数据时，SQL会检查建立的所有依赖的视图，需要满足所有视图的条件才能插入。

  - LOCAL

NOTE：如果没有开检查选项就不会进行检查。不同版本是不同含义的，要看版本。

#### CASCADED

`级联，一旦选择了这个选项，除了会检查创建视图时候的条件，还会检查所依赖视图的条件。`

> 比如下面的例子：创建stu_v_1 视图，id是小于等于 20的。  
> `create or replace view stu_V_l as select id,name from student where id <=20;`
> 再创建 stu_v_2 视图，20 >= id >=10。  
> `create or replace view stu_v_2 as select id,name from stu_v_1 where id >=10 with cascaded check option;`
> 再创建 stu_v_3 视图。     
> `create or replace view stu_v_3 as select id,name from stu_v_2 where id<=15;`  
> 这条数据能够成功，stu_v_3 没有开检查选项所以不会 去判断 id 是否小于等于15,  with cascaded check option 会作用于使用该语句的那一层并附加给该视图依赖的所有视图。
>
> <img src=".\asset\images\image-20240326170454912.png" alt="image-20240326170454912" style="zoom:67%;" />
>
> <img src=".\asset\images\image-20240326170536072.png" alt="image-20240326170536072" style="zoom: 80%;" />
>
> v2视图是基于v1视图的，如果在v2视图创建的时候指定了检查选项为 cascaded，但是v1视图创建时未指定检查选项。 则在执行检查时，不仅会检查v2，还会级联检查v2的关联视图v1。在v3创建时基于v2视图，但是v3没有check选项，因此v3的条件不需要满足，但是v2和v1由于v2的cascaded，都需要满足。
>
> **with cascaded check option 会作用于使用该语句的那一层视图并附加给该视图依赖的所有视图。**

#### LOCAL

>  本地的条件也会检查，**也还会向上检查**。在向上找的时候，就要看是否上面开了检查选项，如果没开就不检查。和 CASCADED 的区别就是 CASCADED 不管上面开没开检查选项都会进行检查。
>
>  <img src=".\asset\images\image-20240326172328350.png" alt="image-20240326172328350" style="zoom:67%;" />
>
>  ![image-20240326172528441](.\asset\images\image-20240326172528441.png)
>
>  v2视图是基于v1视图的，如果在v2视图创建的时候指定了检查选项为 local ，但是v1视图创建时未指定检查选项。 则在执行检查时，知会检查v2，不会检查v2的关联视图v1。v3视图是基于v2视图的，执行检查时，v3没有check选项不进行检查，向上检查，2视图创建的时候指定了检查选项为 local，检查v2；v1没有检查选项，不检查v1。
>
>  

### 更新及作用

要使视图可更新，**视图中的行与基础表中的行之间必须存在一对一的关系**。如果视图包含以下任何一项，则该视图不可更新

1. 聚合函数或窗口函数 ( SUM()、MIN()、MAX()、COUNT() 等 )
2. DISTINCT
3. GROUP BY
4. HAVING
5. UNION 或者UNION ALL

> 例子： 使用了聚合函数，插入会失败。
> `create view stu_v_count as select count(*) from student;`  
> `insert into stu_v_count values(10);`  
>
> <img src=".\asset\images\image-20240326173412059.png" alt="image-20240326173412059" style="zoom: 80%;" />

作用

>- 简单
>
>  视图不仅可以简化用户对数据的理解，也可以简化他们的操作。那些被经常使用的查询可以被定义为视图，从而使得用户不必为以后的操作每次指定全部的条件。
>
>- 安全
>
>  数据库可以授权，但**不能授权到数据库特定行和特定的列上**。通过视图用户只能查询和修改他们所能见到的数据  
>
>- 数据独立
>
>  视图可帮助用户屏蔽真实表结构变化带来的影响。  

总而言之 类似于给表加上了一个外壳，通过这个外壳访问表的时候，只能按照所设计的方式进行访问与更新。

```mysql
-- 为了保证数据库表的安全性，开发人员在操作tb_user表时，只能看到的用户的基本字段，屏蔽手机号和邮箱两个字段。
create view tb_user_view as select id,name,profession,age,gender,status,createtime FROM tb_user;

SELECT * from tb_user_view;

-- 查询每个学生所选修的课程（三张表联查），这个功能在很多的业务中都有使用到，为了简化操作，定义一个视图。
select s.name student_name,s.no student_num,c.name course_name
from student s,student_course sc,course c where s.id = sc.studentid and sc.courseid = c.id;

create view stu_course_view as
    select s.name student_name,s.no student_num,c.name course_name
    from student s,student_course sc,course c where s.id = sc.studentid and sc.courseid = c.id;

select * from stu_course_view;
```



## 存储过程

存储过程是事先经过编译并**存储在数据库中**的**一段SQL 语句的集合**，调用存储过程可以简化应用开发人员的很多工作，减少数据在数据库和应用服务器之间的传输，对于提高数据处理的效率是有好处的。
存储过程思想上很简单，就是数据库SQL 语言层面的代码封装与重用。

特点

1. 封装
2. 复用
3. 可以接收参数，也可以返回数据减少网络交互，效率提升

### 基本语法

#### 创建

> ```mysql
> CREATE PROCEDURE 存储过程名称( [参数列表] ) 
> BEGIN
> SQL 语句 
> END;
> ```
>
> NOTE: 在命令行中，执行创建存储过程的SQL时，需要通过关键字delimiter 指定SQL语句的结束符。因为默认是分号作为结束符，当出现分号时，系统判定语句结束进行提交，因此需要设置delimiter $$ ，则 $$ 符作为结束符。
>
> ![image-20240326175922421](.\asset\images\image-20240326175922421.png)



#### 调用

> `CALL 名称 ( [参数])`

#### 查看

> - 查询指定数据库的存储过程及状态信息
>
>
> `SELECT* FROM INFORMATION_SCHEMA.ROUTINES WHERE ROUTINE_SCHEMA = 'xxx' `
>
> - 存储过程名称；--查询某个存储过程的定义
>
>
> `SHOW CREATE PROCEDURE`

#### 删除

> `DROP PROCEDURE [ IFEXISTS ] 存储过程名称`

### 变量

#### 系统变量

> 是MySQL服务器提供，不是用户定义的，属于服务器层面。分为全局变量（GLOBAL）、会话变量（SESSION）,默认是会话变量。
>
> ##### 查看系统变量
>
> ```mysql
> -- 查看所有系统变量
> SHOW  [ SESSION | GLOBAL ]   VARIABLES ; 
> 
> -- 查看全局变量，会话变量
> show session variables;
> show global variables ;
> 
> -- like 模糊匹配
> SHOW  [ SESSION | GLOBAL ]   VARIABLES  LIKE  '......'; 
> 
> -- 使用 select @@ 查看指定的系统变量的值
> SELECT  @@[SESSION.| GLOBAL.]系统变量名; 
> ```
>
> ##### 设置系统变量
>
> ```mysql
> -- 设置系统变量的值，默认是session变量 只在本query、本会话窗口中起作用，
> -- 全局变量针对于所有的会话。
> -- 但是在服务器重启后，无论是session变量或是global变量都不起作用了 希望永久起作用需要在my.con 文件中定义
> 
> SET  [SESSION | GLOBAL ] 系统变量名 = 值 ; 
> SET  @@[SESSION | GLOBAL] 系统变量名 = 值 ; 
> 
> ```

#### 用户自定义变量

 是用户根据需要自己定义的变量，用户变量不用提前声明，在用的时候直接用 "@变量名" 使用就可以。其作用域为当前连接。

> ##### 赋值
>
> - 能够使用 ''='' 和 '':=''进行赋值
>
>   @@ 指的是系统变量 @指的是用户变量 其作用域为当前连接。
>
>   `SET   @var_name = expr  [, @var_name = expr] ... ;`  
>
>   `SET   @var_name := expr  [, @var_name := expr] ... ;` 
>
> - 使用select进行赋值 也能将查询的结果进行赋值
>
>   `SELECT   @var_name := expr  [, @var_name := expr] ... ;`  
>
>   `SELECT  字段名  INTO @var_name  FROM  表名;`
>
>   ```mysql
>   -- set
>   -- 能够使用 ''='' 和 '':=''进行赋值
>   -- @@ 指的是系统变量 @指的是用户变量 其作用域为当前连接。
>   SET @var_name = expr  [, @var_name = expr] ... ;
>   SET @var_name := expr  [, @var_name := expr] ... ;
>                                       
>   set @myname = 'itcast';
>   set @myage := 10;
>   set @mygender := '男',@myhobby := 'fuck';
>                                       
>   -- select 使用select赋值 
>   SELECT   @var_name := expr  [, @var_name := expr] ... ;
>   SELECT  字段名  INTO @var_name  FROM  表名;
>                                       
>   select @mycolor := 'yellow';
>   -- 将查询的值进行赋值 必须使用 := 因为select语句在判断时是否相等时 使用的也是 = ，不然就区分不开了
>   select count(*) into @myaccount from tb_user;
>   ```
>
> ##### 使用
>
> ​	用户定义的变量无需对其进行声明或初始化，只不过获取到的值为NULL。
>
> ​	对于同一个变量，也能进行对此赋值
>
> ​	`SELECT   @var_name ;` 
>
> ```mysql
> -- 查看定义的变量
> SELECT   @var_name ;
> 
> select @myage,@mygender,@myhobby,@mycolor,@myaccount;
> 
> -- 对于没有赋值的用户自定义变量，也能够获取其值，为null
> select @funk;
> ```

#### 局部变量

> 局部变量 是根据需要定义的在局部生效的变量，访问之前，需要DECLARE声明。可用作存储过程内的 局部变量和输入参数，局部变量的范围是在其内声明的BEGIN ... END块
>
> - 声明
>
>    `DECLARE 变量名 变量类型 [DEFAULT ... ] ;`
>
> - 赋值
>
>   `SET  变量名 = 值 ;` 
>
>   `SET  变量名 := 值 ;` 
>
>   `SELECT  字段名  INTO  变量名  FROM  表名 ... ;` 
>
> ```mysql
> create procedure p2()
> begin
>     declare printinfo varchar(10);
>     declare stu_count int default 0;
>     set printinfo := 'fuck';
>     select count(*) into stu_count from student;
>     select concat(printinfo,' ',stu_count);
> end;
> 
> call p2();
> ```

### 语法结构

#### 选择结构

- #### IF 

```mysql
-- 语法结构
IF  条件1  THEN 
	。。。
ELSEIF 条件2  THEN       -- 可选
	。。。
ELSE                     -- 可选
	。。。
END IF;
```

- #### 参数

参数的类型，主要分为以下三种：IN、OUT、INOUT。 具体的含义如下

| 参数  | 含义                                         | 备注    |
| ----- | -------------------------------------------- | ------- |
| IN    | 该类参数作为输入，也就是需要调用时传入值     | default |
| OUT   | 该类参数作为输出，也就是该参数可以作为返回值 |         |
| INOUT | 既可以作为输入参数，也可以作为输出参数       |         |

```mysql
CREATE PROCEDURE 存储过程名称 ([IN/OUT/INOUT  参数名  参数类型])
BEGIN
    -- SQL语句
END ;


-- eg
create procedure p4(in score int,out degree varchar(10))
begin
    if score >= 85 then
        set degree = '优秀';
    elseif score >= 60 then
        set degree = '及格';
    else
        set degree = '不及格';
    end if;
end;

call p4(68,@degree);
select @degree;

create procedure p5(inout score double)
begin
    set score = score * 0.5;
end;

set @score = 79;
call p5(@score);
select @score;

```

- #### CASE

```mysql
-- 含义： 当case_value的值为when_value1时，执行statement_list1，
--       当值为 when_value2时，执行statement_list2， 否则就执行 statement_list
 CASE  case_value
    WHEN  when_value1  THEN  statement_list1
   [ WHEN  when_value2  THEN  statement_list2] ...
   [ ELSE  statement_list ]
 END  CASE;
 
-- 含义： 当条件search_condition1成立时，执行statement_list1，
--       当条件search_condition2成立时，执行statement_list2， 否则就执行 statement_list
CASE
  WHEN  search_condition1  THEN  statement_list1
  [WHEN  search_condition2  THEN  statement_list2] ...
  [ELSE  statement_list]
END CASE;
```

#### 循环结构

- #### while

先判定条件，如果条件为true，则执行逻辑

```mysql
-- 先判定条件，如果条件为true，则执行逻辑，否则，不执行逻辑
WHILE  条件  DO
    SQL逻辑...    
END WHILE;
```

- #### repeat

repeat是有条件的循环控制语句, 先执行一次，当满足until声明的条件的时候，则退出循环 。

```mysql
-- 先执行一次逻辑，然后判定UNTIL条件是否满足，如果满足，则退出。如果不满足，则继续下一次循环
REPEAT
    SQL逻辑...  
    UNTIL  条件
END REPEAT;
```

- #### loop

LOOP 实现简单的循环，如果不在SQL逻辑中增加退出循环的条件，可以用其来实现简单的死循环。

LOOP可以配合以下两个语句使用：

1. LEAVE ：配合循环使用，退出循环。
2. ITERATE：必须用在循环中，作用是跳过当前循环剩下的语句，直接进入下一次循环。

```mysql
[label:]  LOOP
    -- SQL逻辑...  
END  LOOP  [end_label];
LEAVE  label;   -- 退出指定标记的循环体
ITERATE  label; -- 直接进入下一次循环

-- eg
create procedure p9(in num int,out res int)
begin
    set res = 0;
    sum:loop
         if num<=0 then
             leave sum;
         end if;
         set res = res + num;
         set num = num - 1;
    end loop sum;
end;

call p9(1000,@res);
select @res;
```



### 游标

> SQL 检索操作返回一组称为结果集的行，这组返回的行都是与 SQL 语句 相匹配的行（零行到多行）。
>
> 但是的得到的结果都是完整的查询，可能包含0或者多行，
>
> 简单地使用 SELECT 语句，没有办法得到第 一行、下一行或前 10 行。有时，需要在检索出来的行中前进或后退一行或多行，这就是游标的用途所在。
>
> 游标（cursor）是一个存储在 DBMS 服务器上的数据库查询， 它不是一条 SELECT 语句，而是被该语句检索出来的结果集。

游标（CURSOR）是用来存储查询结果集的数据类型，在存储过程和函数中可以使用游标对结果集进行循环的处理。游标的使用包括游标的声明、OPEN、FETCH和CLOSE，其语法分别如下。

#### 基础语法

>声明游标：
>`DECLARE 游标名称 CURSOR FOR 查询语句`

>打开游标：
>`OPEN 游标名称`

>获取游标记录：
>`FETCH 游标名称 INTO 变量[变量]`

> 关闭游标：
>
> `CLOSE 游标名称。`

```mysql
create procedure p11(in age int)
begin
	-- NOTE：要先声明普通变量，再申请游标。
	
    -- 定义局部变量 接住遍历游标中得到的内容 
    declare u_name varchar(100);
    declare u_pro varchar(100);

    -- 创建游标
    declare u_cusor cursor for select name,profession FROM tb_user where tb_user.age <= age;
    -- 创建结果表
    create table if not exists tb_user_pro(
        id int primary key auto_increment,
        name varchar(100),
        profession varchar(100)
    );
    -- 打开游标
    open u_cusor;
    -- 使用循环得到内容
    -- 这里不知道怎么判断游标内有无数据 只能一直循环 在调用时会报错
    while true do
        -- 将游标的内容放到两个变量中
        fetch u_cusor into u_name,u_pro;
        -- 插入数据
        insert into tb_user_pro values(null,u_name,u_pro);
    end while;
    -- 关闭游标
    close u_cusor;

end;
```

#### 条件处理程序

条件处理程序（Handler）可以用来**定义在流程控制结构执行过程中遇到问题时相应的处理步骤**。具体语法为：  

> `DECLARE handler_action HANDLER FOR condition_value [condition value]..statement;`  

- `handler_action` 
  - `CONTINUE`：继续执行当前程序
  - `EXIT`：终止执行当前程序

- `condition_value` :
  - `SQLSTATE sqlstate_value`：状态码，如02000
  - ``SQLWARNING`：所有以01开头的SQLSTATE代码的简写
  - `NOT FOUND`：所有以02开头的SQLSTATE代码的简写
  - `SQLEXCEPTION`：所有没有被SQLWARNING或NOT FOUND捕获的SQLSTATE代码的简写

例子：

NOTE：要先声明普通变量，再申请游标。

要求：
`根据传入的参数uage，来查询用户表tb_user中，所有的用户年龄小于等于uage的用户姓名（name）和专业（profession），并将用户的姓名和专业插入到所创建的一张新表（id，name，profession）中。`

```mysql
create procedure p12(in age int)
begin

    -- 定义局部变量 接住遍历游标中得到的内容 要先声明变量 后声明游标
    declare u_name varchar(100);
    declare u_pro varchar(100);

    -- 创建游标
    declare u_cusor cursor for select name,profession FROM tb_user where tb_user.age <= age;

    -- 创建hander
    -- 当条件处理程序的处理的状态码为02000的时候，就会退出。并且捎带关闭 u_cusor;
    declare exit handler for SQLSTATE '02000' close u_cusor;
    
    -- 这里这样写也行
	-- 当条件处理程序的处理的状态码为02开头的时候，就会退出。并且捎带关闭 u_cusor;
    # declare exit handler for NOT FOUND close u_cusor;
    
    -- 创建结果表
    create table if not exists tb_user_pro(
        id int primary key auto_increment,
        name varchar(100),
        profession varchar(100)
    );
    -- 打开游标
    open u_cusor;
    -- 使用循环得到内容
    -- 这里不知道怎么判断游标内有无数据 只能一直循环 在调用时会报错
    while true do
        -- 将游标的内容放到两个变量中
        fetch u_cusor into u_name,u_pro;
        -- 插入数据
        insert into tb_user_pro values(null,u_name,u_pro);
    end while;
    -- 关闭游标
    close u_cusor;

end;

call p12(46);
```

## 存储函数

存储函数是有返回值的存储过程，存储函数的**参数只能是IN类型**的。具体语法如下

基础语法：

```mysql
CREATE  FUNCTION   存储函数名称 ([ 参数列表 ])
RETURNS  type  [characteristic ...]
BEGIN -- SQL语句
RETURN ...;
END ;

-- 其中 characteristic有以下几种类型
- DETERMINISTIC：相同的输入参数总是产生相同的结果\
- NO SQL ：不包含 SQL 语句。
- READS SQL DATA：包含读取数据的语句，但不包含写入数据的语句

```

characteristic说明：

- DETERMINISTIC：相同的输入参数总是产生相同的结果\
- NO SQL ：不包含 SQL 语句。
- READS SQL DATA：包含读取数据的语句，但不包含写入数据的语句。

## 触发器

触发器是与表有关的数据库对象，**指在insert/update/delete之前或之后，触发并执行触发器中定义的SQL语句集合。**

触发器的这种特性可以协助应用在数据库端确保数据的完整性，日志记录，数据校验等操作。

使用别名OLD和NEW来引用触发器中发生变化的记录内容，这与其他的数据库是相似的。

现在触发器还只支持行级触发（比如说一条语句影响了 5 行 则会被触发 5 次），不支持语句级触发（比如说一条语句影响了 5 行，则会被触发 1 次）。  

| 触发器类型 | NEW 和 OLD                                           |
| ---------- | ---------------------------------------------------- |
| INSERT     | NEW 表示将要或者已经新增的数据                       |
| UPDATE     | OLD表示修改之前的数据，NEW表示将要或已经修改后的数据 |
| DELETE     | OLD表示将要或者已经删除的数据                        |

### 基本用法

```mysql
-- 创建触发器

CREATE TRIGGER trigger_name 
BEFORE/AFTER INSERT/UPDATE/DELETE
ON tbl_name FOR EACH ROW  -- MYSQL只支持行级触发器
BEGIN
	trigger_stmt ; -- 触发器实际内容
END;

-- 查看触发器
SHOW TRIGGERS;

-- 删除触发器
DROP TRIGGER [schema_name.]trigger_name ;  -- 如果没有指定 schema_name，默认为当前数据库 。
```

```mysql
-- eg
-- 触发器
create table user_log(
    id int(11) not null auto_increment,
    operation varchar(20) not null comment '操作类型，insert/update/delete',
    operation_time datetime not null comment '操作时间',
    operation_id int(11) not null comment '操作id',
    operate_params varchar(500) comment '操作参数',
    primary key (id)
)engine = innodb default charset =utf8;

-- insert触发器
create trigger tb_user_insert_triger
    after insert on tb_user for each row
begin
    insert into user_log
    values (null,'insert',now(),new.id,concat('插入的数据内容为,id :',NEW.id,' name :',NEW.name,', phone : ',NEW.phone,' ,profession : ',NEW.profession));
end;
-- 查看
show triggers;
-- 删除
Drop trigger tb_user_insert_triger;


-- 插入数据到tb_user
insert into tb_user(id, name, phone, email, profession, age, gender, status, createtime)
VALUES (26,'三皇子','18809091212','erhuangzi@163.com','软件工程',23,'1','1',now());

insert into tb_user(id, name, phone, email, profession, age, gender, status, createtime)
VALUES (25,'二皇子','18809091212','erhuangzi@163.com','软件工程',23,'1','1',now());

-- insert触发器
create trigger tb_user_update_triger
    after update on tb_user for each row
begin
    insert into user_log
    values (null,'update',now(),new.id,concat('更新前的数据内容为,id :',OLD.id,' name :',OLD.name,', phone : ',OLD.phone,' ,profession : ',OLD.profession,
                                       '\n更新后的数据内容为,id :',NEW.id,' name :',NEW.name,', phone : ',NEW.phone,' ,profession : ',NEW.profession));
end;

show triggers;

update tb_user set profession = '会计' where id = 23;

-- 触发几次行的改变 触发几次 行级
update tb_user set profession = '会计' where id <= 5;

create trigger tb_user_delete_triger
    after delete on tb_user for each row
begin
    insert into user_log
    values (null,'delete',now(),OLD.id,concat('删除的的数据内容为,id :',OLD.id,' name :',OLD.name,', phone : ',OLD.phone,' ,profession : ',OLD.profession));
end;

delete from tb_user where id = 26;

```



## 锁

锁是计算机协调多个进程或线程并发访问某一资源的机制。在数据库中，除传统的计算资源（CPU、RAM、I/O）的争用以外，数据也是一种供许多用户共享的资源。如何保证数据并发访问的一致性、有效性是所有数据库必须解决的一个问题，锁冲突也是影响数据库并发访问性能的一个重要因素。从这个角度来说，锁对数据库而言显得尤其重要，也更加复杂。

NOTE : 针对事物才有加锁的意义。

分类：MySQL中的锁，按照锁的粒度分，分为以下三类：

1. 全局锁：锁定数据库中的所有表。
2. 表级锁：每次操作锁住整张表。
3. 行级锁：每次操作锁住对应的行数据。

### 全局锁：

全局锁就是**对整个数据库实例加锁**，加锁后整个实例就处于只读状态，后续的**DML的写语句，DDL语句，已经更新操作的事务提交语句都将被阻塞。**  
其典型的使用场景是做全库的逻辑备份，对所有的表进行锁定，从而获取一致性视图，保证数据的完整性。

```mysql
-- 数据备份：
-- 对整个数据库使用全局锁，锁住了全部的表，另外的连接只能DQL 该数据库中的表
-- DDL DML 会阻塞 直至解开该锁
flush tables with read lock;

-- mysqldump 是mysql提供的工具 在命令行使用
$$ mysqldump -hXXXXXX -uXXXX -pXXXXXX >存储位置 


-- 释放该
unlock tables;
```

  在innodb中，现已经支持不加锁完成一致性备份，在备份时加上参数 --single-transaction 参数即可

```bash
mysqldump  --single-transaction  -uroot –p123456  itcast > itcast.sql
```



### 表级锁：

表级锁，每次操作锁住整张表。锁定粒度大，发生锁冲突的概率最高，并发度最低。应用在MyISAM、InnoDB、BDB等存储引擎中。  

对于表级锁，主要分为以下三类：

1. 表锁：对于表锁，分为两类：

   - 表共享读锁（read lock）**所有的事物都只能读**（当前加锁的客户端也只能读，不能写），不能写 .

      ***\< 本客户端写报错，其他客户端写阻塞直至该锁释放。>***
   
   - 表独占写锁（write lock），**对当前加锁的客户端，可读可写**，对于其他的客户端，不可读也不可写,***读写阻塞直至该锁释放***

**读锁不会阻塞其他客户端的读，但是会阻塞写。写锁既会阻塞其他客户端的读，又会阻塞其他客户端的写。**

   > 语法
   >
   > 加锁 `lock tables 表名... read/write;`
   >
   > 释放锁 `unlock tables;` 
   >
   > 客户端断开连接时也会自动解锁。

2. 元数据锁（meta data lock，MDL）

   **MDL加锁过程是系统自动控制，无需显式使用**，在访问一张表的时候会自动加上。

   MDL锁主要作用是维护表元数据的数据一致性，***\< 在表上有未提交事务时，控制其他客户端或事务对该表的操作权限。>***

   在MySQL5.5中引入了MDL，

   - 当对一张表进行查的时候，加MDL读锁（SHARED_READ），其他事务或客户端能对该表进行读写操作，但不能修改表(alter table...)
   - 当对一张表进行增删改的时候，加MDL写锁（SHARED_WRITE），其他事务或客户端能对该表进行读写操作，但不能修改表(alter table...)

   | 锁       | 对应SQL                            | 锁类型                                  | 说明                                               |
   | -------- | ---------------------------------- | --------------------------------------- | -------------------------------------------------- |
   | 表锁     | **`lock tables xxx read / write`** | SHARED_READ_ONLY / SHARED_NO_READ_WRITE |                                                    |
   |          |                                    |                                         |                                                    |
   | 元数据锁 | select                             | SHARED_READ                             | 与SHARED_READ、SHARED_WRITE兼容，与 EXCLUSIVE互斥  |
   |          | insert 、update、 delete           | SHARED_WRITE                            | 与SHARED_READ、 SHARED_WRITE兼容，与 EXCLUSIVE互斥 |
   |          | alter table ...                    | EXCLUSIVE                               | 与其他的MDL都互斥                                  |

   ```mysql
   -- 我们在操作过程中，可以通过该SQL语句，来查看元数据锁的加锁情况
   select object_type,object_schema,object_name,lock_type,lock_duration from performance_schema.metadata_locks ;
   ```

   

3. 意向锁: 

   为了避免DML在执行时，加的行锁与表锁的冲突，在InnoDB中引入了意向锁，使得表锁不用检查每行数据是否加锁，使用意向锁来减少表锁的检查。  
   一个客户端对某一行加上了行锁，**那么系统也会对该表加上一个意向锁**，当别的客户端来想要对其加*上表锁*时，便会检查意向锁是否兼容，若是不兼容，便会阻塞直到意向锁释放。  

   初衷是为了，当一个事务对一个表中的一行数据加了行锁，另一个事务或是客户端需要对该表加表锁，但是需要搞清楚这个表的加锁情况，还得一行一行的检索该表数据的加锁情况，效率很低。因此，为了方便别的事务的处理，当一个事务对一个表中的一行数据加了行锁时会附带对整个表加意向锁。

> 意向锁兼容性：
>
> 表锁  `lock tables 表名... read/write;`
>
> 1. 意向共享锁（IS）：与表锁共享锁（read）兼容，与表锁排它锁（write）互斥。
>
>      `select * from student lock in share mode;` 加行锁与意向共享锁   与表锁（write）写互斥。
>
> 2. 意向排他锁（lX）：与表锁共享锁（read）及排它锁（write）都互斥。意向锁之间不会互斥。  
>
>    `update student set name = 'sdawdas' where id = 1;`  加行锁与意向排他锁   与表锁读写（read,write）都互斥
>
> ```mysql
> -- 我们在操作过程中，可以通过该SQL语句，来查看意向锁及行锁的加锁情况。
> select object_schema,object_name,index_name,lock_type,lock_mode,lock_data from 
> performance_schema.data_locks;
> ```



### 行级锁：

行级锁，**每次操作锁住对应的行数据**。锁定粒度最小，发生锁冲突的概率最低，并发度最高。应用在InnoDB存储引擎中。  
InnoDB的数据是基于索引组织的，**行锁是通过对索引上的索引项加锁来实现的，而不是对记录加的锁。**

对于行级锁，主要分为以下三类：  

1. 行锁（Record Lock）：锁定单个行记录的锁，防止其他事务**对此行**进行update和delete。

   在RC（read commit ）、RR（repeat read）隔离级别下都支持。

   ![image-20240327111104601](.\asset\images\image-20240327111104601.png)

2. 间隙锁（GapLock）：锁定索引记录间隙（两个记录之间的间隙，不含该记录），确保索引记录间隙不变，防止其他事务在这个间隙进行insert，产生幻读。

   在RR隔离级别下都支持。比如说 两个临近叶子节点为 15 23，那么间隙就是指 [15 , 23],锁的是这个间隙。

   ![image-20240327111215269](.\asset\images\image-20240327111215269.png)

3. 临键锁（Next-Key Lock）：行锁和间隙锁组合，同时锁住数据，并锁住**数据前面的间隙Gap**。在RR隔离级别下支持。

   ![image-20240327111255050](.\asset\images\image-20240327111255050.png)

InnoDB实现了以下两种类型的行锁：  

1. 共享锁（S）：允许一个事务去读一行，阻止其他事务获得相同数据集的排它锁。
2. 排他锁（X）：允许获取排他锁的事务更新数据，阻止其他事务获得相同数据集的共享锁和排他锁。 增删改

| SQL                       | 行锁类型   | 说明                                     |
| ------------------------- | ---------- | ---------------------------------------- |
| insert                    | 排他锁     | 自动加锁                                 |
| update                    | 排他锁     | 自动加锁                                 |
| delete                    | 排他锁     | 自动加锁                                 |
| select                    | 不加任何锁 |                                          |
| select lock in share mode | 共享锁     | 需要手动在SELECT之后加LOCK IN SHARE MODE |
| select for update         | 排他锁     | 需要手动在SELECT之后加FOR UPDATE         |

行锁 - 演示  

默认情况下，InnoDB在REPEATABLE READ事务隔离级别运行，InnoDB使用**临键锁**进行搜索和索引扫描，以防止幻读。

1. 针对唯一索引进行检索时，对已存在的记录进行等值匹配时，将会自动优化为行锁。

   ![image-20240327123545610](.\asset\images\image-20240327123545610.png)

2. InnoDB的行锁是针对于索引加的锁，不通过索引条件检索数据，那么InnoDB将对表中的所有记录加锁，此时就会升级为表锁。

   ![image-20240327125401342](.\asset\images\image-20240327125401342.png)

   ![image-20240327125941285](.\asset\images\image-20240327125941285.png)



间隙锁/临键锁-演示

默认情况下，InnoDB在REPEATABLE READ事务隔离级别运行，InnoDB使用**临键锁**进行搜索和索引扫描，以防止幻读。

1. 索引上的等值查询（唯一索引），给不存在的记录加锁时，优化为间隙锁。

   ![image-20240327130519040](.\asset\images\image-20240327130519040.png)

2. 索引上的等值查询（**普通索引，非唯一索引**），向右遍历时最后一个值不满足查询需求时，**临键锁退化为间隙锁**。

   因为是非唯一索引，以后还可能插入相同的值，因此，要锁住该索引+前后的间隙。

   ![image-20240327132252040](.\asset\images\image-20240327132252040.png)

3. 索引上的**范围查询**（唯一索引）--会访问到不满足条件的第一个值为止。

   ![image-20240327132923877](.\asset\images\image-20240327132923877.png)



注意：间隙锁唯一目的是防止其他事务插入间隙。间隙锁可以共存，一个事务采用的间隙锁不会阻止另一个事务在同一间隙上采用间隙锁。

## InnoDB 引擎

### 逻辑存储结构

<img src=".\asset\images\image-20240327154505900.png" alt="image-20240327154505900" style="zoom:67%;" />

- 表空间（ibd文件），一个mysql实例可以对应多个表空间，用于存储记录、索引等数据。 /var/lib/mysql


- 段，分为数据段（Leaf node segment）、索引段（Non-leaf node segment）、回滚段（Rollback segment），InnoDB是索引组织表，**数据段就是B+树的叶子节点，索引段即为B+树的非叶子节点**。段用来管理多个Extent（区）。


- 区，表空间的单元结构，每个区的大小为1M。默认情况下，InnoDB存储引擎页大小为16K，即一个区中一共有64个连续的页。


- 页，是InnoDB存储引擎磁盘管理的最小单元，每个页的大小默认为16KB。为了保证页的连续性，InnoDB存储引擎每从磁盘申请4-5个区。一页包含若干行。


- 行，InnoDB存储引擎数据是按进行存放的。

  `TRX id` 最后一次事务操作的id

  Roll pointer 指针，通过它来找到该记录修改前的信息。

### 架构

<img src=".\asset\images\image-20240327155312795.png" alt="image-20240327155312795" style="zoom:67%;" />

#### 内存结构

Buffer Pool：缓冲池是主内存中的一个区 域，里面可以缓存磁盘上经常操作的真实数据，在执行增删改查操作时，先操作缓冲池中的数据（若缓冲池没有数据，则从磁盘加载并缓存），然后再以一定频率刷新到磁盘，从而减少磁盘I0，加快处理速度。

<img src=".\asset\images\image-20240327160124083.png" alt="image-20240327160124083" style="zoom:67%;" />

<img src=".\asset\images\image-20240327155812192.png" alt="image-20240327155812192" style="zoom:67%;" />

!<img src=".\asset\images\image-20240327155900093.png" alt="image-20240327155900093" style="zoom:67%;" />
<img src=".\asset\images\image-20240327160059388.png" alt="image-20240327160059388" style="zoom:67%;" />

#### 磁盘架构：

<img src=".\asset\images\image-20240327160323499.png" alt="image-20240327160323499" style="zoom:67%;" />
<img src=".\asset\images\image-20240327160718943.png" alt="image-20240327160718943" style="zoom:67%;" />
<img src=".\asset\images\image-20240327161008630.png" alt="image-20240327161008630" style="zoom:67%;" />

InnoDB的整个体系结构为：

当业务操作的时候直接操作的是内存缓冲区，如果缓冲区当中没有数据，则会从磁盘中加载到缓冲区，增删改查都是在缓冲区的，后台线程以一定的速率刷新到磁盘。

后台线程：

在InnoDB的后台线程中，分为4类，分别是：Master Thread 、IO Thread、Purge Thread、 Page Cleaner Thread。

-  Master Thread

  核心后台线程，负责调**度其他线程**，还负责**将缓冲池中的数据异步刷新到磁盘中**, 保持数据的一致性，  还包括**脏页的刷新、合并插入缓存、undo页的回收** 。

- IO Thread

  在InnoDB存储引擎中大量使用了AIO（异步IO）来处理IO请求, 这样可以极大地提高数据库的性能，而IO  Thread主要负责这些IO请求的回调。

  > | 线程类型             | 默认个数 | 职责                         |
  > | -------------------- | -------- | ---------------------------- |
  > | Read thread          | 4        | 负责读操作                   |
  > | Write thread         | 4        | 负责写操作                   |
  > | Log thread           | 1        | 负责将日志缓冲区刷新到磁盘   |
  > | Insert buffer thread | 1        | 负责将写缓冲区内容刷新到磁盘 |

- Purge Thread

  主要用于回收事务已经提交了的undo log，在事务提交之后，undo log可能不用了，就用它来回收。

- Page Cleaner Thread

  协助 Master Thread 刷新脏页到磁盘的线程，它可以减轻 Master Thread 的工作压力，减少阻塞。

## 事务原理

事务是一组操作的集合，它是一个不可分割的工作单位，事务会把所有的操作作为一个整体一起向系统提交或撤销操作请求，即这些操作要么同时成功，要么同时败。具有ACID四大特征。

-  原子性（Atomicity）：事务是不可分割的最小操作单元，要么全部成功，要么全部失败。 
- 一致性（Consistency）：事务完成时，必须使所有的数据都保持一致状态。
- 隔离性（Isolation）：数据库系统提供的隔离机制，保证事务在不受外部并发操作影响的独立环境下运行。
- 持久性（Durability）：事务一旦提交或回滚，它对数据库中的数据的改变就是永久的。

原子性，一致性，持久性这三大特性 由 redo log 和 undo log 日志来保证的。
隔离性 是由锁机制和MVCC保证的。

redo log:

重做日志，**记录的是事务提交时数据页的物理修改**，记录的是实际的数据，是用来实现事务的**持久性**。
该日志文件由两部分组成：重做日志缓冲（redo log buffer）以及重做日志文件（redo log file），前者是在内存中，后者在磁盘中。当事务提交之后会把所有修改信息都存到该日志文件中，并刷新到磁盘文件中，用于在**刷新脏页到磁盘，发生错误时，进行数据恢复使用。**

> 个人理解： 事物每次提交的时候都会将数据刷到redo log中而不是直接将buffer pool中的数据直接刷到磁盘中（ibd文件中），是因为redo log 日志文件是追加的，为顺序写，性能处理的够快；而将数据刷到ibd中，在实际应用中一般是随机写，性能慢。所以脏页是在下一次读的时候，或者后台线程采用一定的机制进行刷盘到ibd中。

undo log:  
回滚日志，**用于记录数据被修改前的信息**，作用包含两个：**提供回滚**和**MVCC**（多版本并发控制）。

undo log和redo log记录物理日志不一样，它是逻辑日志，记录操作。

可以认为当delete一条记录，undo log中会记录一条对应的insert记录，反之亦然，当update一条记录时，它记录一条对应相反的update记录。当执行rollback时，就可以从undo log中的逻辑记录读取到相应的内容并进行回滚。  

> Undo log销毁：undo log在事务执行时产生，事务提交时，并不会立即删除undo log，因为这些日志可能还用于MVCC。
> Undo log存储：undo log采用段的方式进行管理和记录，存放在前面介绍的rollback segment回滚段中，内部包含1024个undo log segment。

## MVCC

### 当前读:

**读取的是记录的最新版本**，读取时还要保证其他并发事务不能修改当前记录，会对读取的记录进行加锁。对于我们日常的操作，如：

* select...lock in share mode（共享锁）。
* select..…for update、update、insert、delete（排他锁）都是一种当前读。

### 快照读:

简单的select（不加锁）就是快照读，快照读，读取的是记录数据的可见版本，**有可能是历史数据**，不加锁，是非阻塞读。

* Read Committed：每次select，都生成一个快照读。
* Repeatable Read：开启事务后**第一个select语句才是快照读的地方。**
* Serializable：快照读会退化为当前读。

### MVCC: 

**主要用于快照读时查找对应历史版本。**

全称Multi-Version Concurrency Control，多版本并发控制。指维护一个数据的多个版本，**使得读写操作没有冲突**，快照读为MySQL实现MVCC提供了一个非阻塞读功能。

MVCC的具体实现，还需要依赖于：数据库记录中的三个隐式字段、undo log日志、readView。

#### MVCC 实现原理:

- ##### **三个隐藏的字段:**

<img src=".\asset\images\image-20240327164436202.png" alt="image-20240327164436202" style="zoom:67%;" />

- ##### **undo log 版本链：**

  undo log：回滚日志，在insert、update、delete的时候产生的便于数据回滚的日志。

  - 当insert的时候，产生的undo log日志只在回滚时需要，在事务提交后，可被立即删除。

  - 而update、delete的时候，产生的undo log日志不仅在回滚时需要，在快照读时也需要，不会立即被删除。

  undo log日志会记录原来的版本的数据，因为是通过undo log 日志进行回滚的。

  ***版本链原理***

  <img src=".\asset\images\image-20240327170304385.png" alt="image-20240327170304385" style="zoom:67%;" />

  [`undolog` 版本链更新原理](https://www.bilibili.com/video/BV1Kr4y1i7ru/?p=143&spm_id_from=pageDriver&vd_source=58c9363bd239890befbdc2caffaa8925)

- ##### **ReadView**

  ReadView（读视图）是 快照读 SQL执行时MVCC提取数据的依据，记录并维护系统当前活跃的事务 （未提交的）id。

  快照读如何确定返回哪一个版本 这是由**read view**决定返回 undo log 中的哪一个版本。包含四个核心字段

  | 字段           | 含义                                                 |
  | -------------- | ---------------------------------------------------- |
  | m_ids          | 当前活跃的事务ID集合                                 |
  | min_trx_id     | 最小活跃事务ID                                       |
  | max_trx_id     | 预分配事务ID，当前最大事务ID+1（因为事务ID是自增的） |
  | creator_trx_id | ReadView创建者的事务ID                               |

  <img src=".\asset\images\image-20240327171312381.png" alt="image-20240327171312381" style="zoom:67%;" />

  > - RC(Read Commit，读提交模式）
  >
  >   RC隔离级别下，在事务中每一次执行快照读时生成ReadView。按undo log链的顺序进行条件匹配。
  >
  >   也很好理解，读提交，即能够读取已经提交的数据，所以需要对每个查询挨个进行匹配。
  >
  > <img src=".\asset\images\image-20240327172305203.png" alt="image-20240327172305203" style="zoom:67%;" />
  >
  > [RC级别 readview 选取流程](https://www.bilibili.com/video/BV1Kr4y1i7ru/?p=145&spm_id_from=pageDriver&vd_source=58c9363bd239890befbdc2caffaa8925)
  >
  > - RR（Read Repeat ，可重复读）
  >
  >   RR 隔离级别下，**在事务中第一次执行快照读时生成ReadView，后续会复用。**  保证了一个事务中多次读取数据显示相同。
  >

![image-20240327173911361](.\asset\images\image-20240327173911361.png)

MVCC 靠 隐藏字段 , undo log 版本链 , read view 实现的。

* 原子性-undo log 
* 持久性-redo log
* 一致性-undo log + redo log
* 隔离性-锁 + MVCC

# MySQL管理

## 系统数据库

Mysql数据库安装完成后，自带了一下四个数据库，具体作用如下：

| 数据库             | 含义                                                         |
| ------------------ | ------------------------------------------------------------ |
| mysql              | 存储MySQL服务器正常运行所需要的各种信息 （时区、主从、用 户、权限等） |
| information_schema | 提供了访问数据库元数据的各种表和视图，包含数据库、表、字段类 型及访问权限等 |
| performance_schema | 为MySQL服务器运行时状态提供了一个底层监控功能，主要用于收集 数据库服务器性能参数 |
| sys                | 包含了一系列方便 DBA 和开发人员利用 performance_schema 性能数据库进行性能调优和诊断的视图 |

## 常用工具

命令行直接输入

- ### mysql

```shell
#语法 ：    
	mysql   [options]   [database]
#选项 ： 
    -u, --user=name         	#指定用户名
    -p, --password[=name]       #指定密码
    -h, --host=name         	#指定服务器IP或域名
    -P, --port=port             #指定连接端口
    -e, --execute=name          #执行SQL语句并退出
#eg
	mysql -uroot –p123456 db01 -e "select * from stu";
```



- #### mysqladmin

mysqladmin 是一个执行管理操作的客户端程序。可以用它来检查服务器的配置和当前状态、创建并 删除数据库等。

```shell
#通过帮助文档查看选项：
	mysqladmin --help
```



- #### mysqlbinlog

由于服务器生成的二进制日志文件以二进制格式保存，所以如果想要检查这些文本的文本格式，就会使 用到mysqlbinlog 日志管理工具。

```shell
#语法 ：    
	mysqlbinlog [options]  log-files1 log-files2 ...
#选项 ：
	-d, --database=name		#指定数据库名称，只列出指定的数据库相关操作。
    -o, --offset=#			#忽略掉日志中的前n行命令。
    -r,--result-file=name	#将输出的文本格式日志输出到指定文件。
    -s, --short-form		#显示简单格式， 省略掉一些信息            
	--start-datatime=date1	--stop-datetime=date2		#指定日期间隔内的所有日志。
	--start-position=pos1 --stop-position=pos2			#指定位置间隔内的所有日志。          
```

- #### mysqlshow

mysqlshow 客户端对象查找工具，用来很快地查找存在哪些数据库、数据库中的表、表中的列或者索引。

```shell
#语法 ：    
    mysqlshow [options] [db_name [table_name [col_name]]]
#选项 ： 
    --count     #显示数据库及表的统计信息（数据库，表 均可以不指定）
    -i      	#显示指定数据库或者指定表的状态信息
#示例：
    #查询test库中每个表中的字段书，及行数
    mysqlshow -uroot -p2143 test --count
    #查询test库中book表的详细情况
    mysqlshow -uroot -p2143 test book --count
    
    mysqlshow -uroot -p13141567 -i
    mysqlshow -uroot -p13141567 itcast stu -i
    mysqlshow -uroot -p13141567 itcast stu id -i
```

- #### mysqldump

mysqldump 客户端工具用来备份数据库或在不同数据库之间进行数据迁移。备份内容包含创建表，及插入表的SQL语句。

```shell
#语法 ：    
    mysqldump [options] db_name [tables]
    mysqldump [options] --database/-B db1 [db2 db3...]
    mysqldump [options] --all-databases/-A
#连接选项 ：  
    -u, --user=name                 #指定用户名
    -p, --password[=name]           #指定密码
    -h, --host=name                 #指定服务器ip或域名
    -P, --port=#                    #指定连接端口
#输出选项：
    --add-drop-database         #在每个数据库创建语句前加上 drop database 语句
    --add-drop-table            #在每个表创建语句前加上 drop table 语句 , 默认开启 ; 不开启 (--skip-add-drop-table)
    -n, --no-create-db          #不包含数据库的创建语句
    -t, --no-create-info        #不包含数据表的创建语句
    -d --no-data                #不包含数据
    -T, --tab=name              #自动生成两个文件：一个.sql文件，创建表结构的语句；一个.txt文件，数据文件
```

- ####  mysqlimport/source

  -  mysqlimport

   mysqlimport 是客户端数据导入工具，用来导入mysqldump 加 -T 参数后导出的文本文件。

  ```shell
  #语法 ：    
  mysqlimport [options]  db_name  textfile1  [textfile2...]
  #示例 ： 
  mysqlimport -uroot -p2143 test /tmp/city.txt
  ```

  - 如果需要导入sql文件,可以使用mysql中的source 指令 : 

  ```mysql
  # 语法 ：    
  source /xxxx/xxxxx.sql
  ```

  

# 数据类型

## 整型

| 类型名称      | 取值范围                                  | 大小    |
| ------------- | ----------------------------------------- | ------- |
| TINYINT       | -128〜127                                 | 1个字节 |
| SMALLINT      | -32768〜32767                             | 2个宇节 |
| MEDIUMINT     | -8388608〜8388607                         | 3个字节 |
| INT (INTEGHR) | -2147483648〜2147483647                   | 4个字节 |
| BIGINT        | -9223372036854775808〜9223372036854775807 | 8个字节 |

无符号在数据类型后加 unsigned 关键字。

## 浮点型

| 类型名称            | 说明               | 存储需求   |
| ------------------- | ------------------ | ---------- |
| FLOAT               | 单精度浮点数       | 4 个字节   |
| DOUBLE              | 双精度浮点数       | 8 个字节   |
| DECIMAL (M, D)，DEC | 压缩的“严格”定点数 | M+2 个字节 |

## 日期和时间

| 类型名称  | 日期格式            | 日期范围                                          | 存储需求 |
| --------- | ------------------- | ------------------------------------------------- | -------- |
| YEAR      | YYYY                | 1901 ~ 2155                                       | 1 个字节 |
| TIME      | HH:MM:SS            | -838:59:59 ~ 838:59:59                            | 3 个字节 |
| DATE      | YYYY-MM-DD          | 1000-01-01 ~ 9999-12-3                            | 3 个字节 |
| DATETIME  | YYYY-MM-DD HH:MM:SS | 1000-01-01 00:00:00 ~ 9999-12-31 23:59:59         | 8 个字节 |
| TIMESTAMP | YYYY-MM-DD HH:MM:SS | 1980-01-01 00:00:01 UTC ~ 2040-01-19 03:14:07 UTC | 4 个字节 |

## 字符串

| 类型名称   | 说明                                         | 存储需求                                                   |
| ---------- | -------------------------------------------- | ---------------------------------------------------------- |
| CHAR(M)    | 固定长度非二进制字符串                       | M 字节，1<=M<=255                                          |
| VARCHAR(M) | 变长非二进制字符串                           | L+1字节，在此，L< = M和 1<=M<=255                          |
| TINYTEXT   | 非常小的非二进制字符串                       | L+1字节，在此，L<2^8                                       |
| TEXT       | 小的非二进制字符串                           | L+2字节，在此，L<2^16                                      |
| MEDIUMTEXT | 中等大小的非二进制字符串                     | L+3字节，在此，L<2^24                                      |
| LONGTEXT   | 大的非二进制字符串                           | L+4字节，在此，L<2^32                                      |
| ENUM       | 枚举类型，只能有一个枚举字符串值             | 1或2个字节，取决于枚举值的数目 (最大值为65535)             |
| SET        | 一个设置，字符串对象可以有零个或 多个SET成员 | 1、2、3、4或8个字节，取决于集合 成员的数量（最多64个成员） |

## 二进制类型

| 类型名称       | 说明                 | 存储需求               |
| -------------- | -------------------- | ---------------------- |
| BIT(M)         | 位字段类型           | 大约 (M+7)/8 字节      |
| BINARY(M)      | 固定长度二进制字符串 | M 字节                 |
| VARBINARY (M)  | 可变长度二进制字符串 | M+1 字节               |
| TINYBLOB (M)   | 非常小的BLOB         | L+1 字节，在此，L<2^8  |
| BLOB (M)       | 小 BLOB              | L+2 字节，在此，L<2^16 |
| MEDIUMBLOB (M) | 中等大小的BLOB       | L+3 字节，在此，L<2^24 |
| LONGBLOB (M)   | 非常大的BLOB         | L+4 字节，在此，L<2^32 |

# 权限一览表

> 具体权限的作用详见[官方文档](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html "官方文档")

GRANT 和 REVOKE 允许的静态权限

| Privilege                                                    | Grant Table Column           | Context                               |
| :----------------------------------------------------------- | :--------------------------- | :------------------------------------ |
| [`ALL [PRIVILEGES]`](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html#priv_all) | Synonym for “all privileges” | Server administration                 |
| [`ALTER`](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html#priv_alter) | `Alter_priv`                 | Tables                                |
| [`ALTER ROUTINE`](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html#priv_alter-routine) | `Alter_routine_priv`         | Stored routines                       |
| [`CREATE`](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html#priv_create) | `Create_priv`                | Databases, tables, or indexes         |
| [`CREATE ROLE`](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html#priv_create-role) | `Create_role_priv`           | Server administration                 |
| [`CREATE ROUTINE`](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html#priv_create-routine) | `Create_routine_priv`        | Stored routines                       |
| [`CREATE TABLESPACE`](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html#priv_create-tablespace) | `Create_tablespace_priv`     | Server administration                 |
| [`CREATE TEMPORARY TABLES`](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html#priv_create-temporary-tables) | `Create_tmp_table_priv`      | Tables                                |
| [`CREATE USER`](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html#priv_create-user) | `Create_user_priv`           | Server administration                 |
| [`CREATE VIEW`](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html#priv_create-view) | `Create_view_priv`           | Views                                 |
| [`DELETE`](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html#priv_delete) | `Delete_priv`                | Tables                                |
| [`DROP`](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html#priv_drop) | `Drop_priv`                  | Databases, tables, or views           |
| [`DROP ROLE`](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html#priv_drop-role) | `Drop_role_priv`             | Server administration                 |
| [`EVENT`](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html#priv_event) | `Event_priv`                 | Databases                             |
| [`EXECUTE`](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html#priv_execute) | `Execute_priv`               | Stored routines                       |
| [`FILE`](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html#priv_file) | `File_priv`                  | File access on server host            |
| [`GRANT OPTION`](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html#priv_grant-option) | `Grant_priv`                 | Databases, tables, or stored routines |
| [`INDEX`](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html#priv_index) | `Index_priv`                 | Tables                                |
| [`INSERT`](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html#priv_insert) | `Insert_priv`                | Tables or columns                     |
| [`LOCK TABLES`](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html#priv_lock-tables) | `Lock_tables_priv`           | Databases                             |
| [`PROCESS`](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html#priv_process) | `Process_priv`               | Server administration                 |
| [`PROXY`](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html#priv_proxy) | See `proxies_priv` table     | Server administration                 |
| [`REFERENCES`](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html#priv_references) | `References_priv`            | Databases or tables                   |
| [`RELOAD`](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html#priv_reload) | `Reload_priv`                | Server administration                 |
| [`REPLICATION CLIENT`](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html#priv_replication-client) | `Repl_client_priv`           | Server administration                 |
| [`REPLICATION SLAVE`](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html#priv_replication-slave) | `Repl_slave_priv`            | Server administration                 |
| [`SELECT`](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html#priv_select) | `Select_priv`                | Tables or columns                     |
| [`SHOW DATABASES`](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html#priv_show-databases) | `Show_db_priv`               | Server administration                 |
| [`SHOW VIEW`](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html#priv_show-view) | `Show_view_priv`             | Views                                 |
| [`SHUTDOWN`](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html#priv_shutdown) | `Shutdown_priv`              | Server administration                 |
| [`SUPER`](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html#priv_super) | `Super_priv`                 | Server administration                 |
| [`TRIGGER`](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html#priv_trigger) | `Trigger_priv`               | Tables                                |
| [`UPDATE`](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html#priv_update) | `Update_priv`                | Tables or columns                     |
| [`USAGE`](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html#priv_usage) | Synonym for “no privileges”  | Server administration                 |

GRANT 和 REVOKE 允许的动态权限

| Privilege                                                    | Context                                           |
| :----------------------------------------------------------- | :------------------------------------------------ |
| [`APPLICATION_PASSWORD_ADMIN`](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html#priv_application-password-admin) | Dual password administration                      |
| [`AUDIT_ABORT_EXEMPT`](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html#priv_audit-abort-exempt) | Allow queries blocked by audit log filter         |
| [`AUDIT_ADMIN`](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html#priv_audit-admin) | Audit log administration                          |
| [`AUTHENTICATION_POLICY_ADMIN`](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html#priv_authentication-policy-admin) | Authentication administration                     |
| [`BACKUP_ADMIN`](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html#priv_backup-admin) | Backup administration                             |
| [`BINLOG_ADMIN`](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html#priv_binlog-admin) | Backup and Replication administration             |
| [`BINLOG_ENCRYPTION_ADMIN`](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html#priv_binlog-encryption-admin) | Backup and Replication administration             |
| [`CLONE_ADMIN`](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html#priv_clone-admin) | Clone administration                              |
| [`CONNECTION_ADMIN`](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html#priv_connection-admin) | Server administration                             |
| [`ENCRYPTION_KEY_ADMIN`](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html#priv_encryption-key-admin) | Server administration                             |
| [`FIREWALL_ADMIN`](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html#priv_firewall-admin) | Firewall administration                           |
| [`FIREWALL_EXEMPT`](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html#priv_firewall-exempt) | Firewall administration                           |
| [`FIREWALL_USER`](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html#priv_firewall-user) | Firewall administration                           |
| [`FLUSH_OPTIMIZER_COSTS`](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html#priv_flush-optimizer-costs) | Server administration                             |
| [`FLUSH_STATUS`](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html#priv_flush-status) | Server administration                             |
| [`FLUSH_TABLES`](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html#priv_flush-tables) | Server administration                             |
| [`FLUSH_USER_RESOURCES`](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html#priv_flush-user-resources) | Server administration                             |
| [`GROUP_REPLICATION_ADMIN`](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html#priv_group-replication-admin) | Replication administration                        |
| [`GROUP_REPLICATION_STREAM`](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html#priv_group-replication-stream) | Replication administration                        |
| [`INNODB_REDO_LOG_ARCHIVE`](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html#priv_innodb-redo-log-archive) | Redo log archiving administration                 |
| [`NDB_STORED_USER`](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html#priv_ndb-stored-user) | NDB Cluster                                       |
| [`PASSWORDLESS_USER_ADMIN`](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html#priv_passwordless-user-admin) | Authentication administration                     |
| [`PERSIST_RO_VARIABLES_ADMIN`](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html#priv_persist-ro-variables-admin) | Server administration                             |
| [`REPLICATION_APPLIER`](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html#priv_replication-applier) | `PRIVILEGE_CHECKS_USER` for a replication channel |
| [`REPLICATION_SLAVE_ADMIN`](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html#priv_replication-slave-admin) | Replication administration                        |
| [`RESOURCE_GROUP_ADMIN`](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html#priv_resource-group-admin) | Resource group administration                     |
| [`RESOURCE_GROUP_USER`](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html#priv_resource-group-user) | Resource group administration                     |
| [`ROLE_ADMIN`](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html#priv_role-admin) | Server administration                             |
| [`SESSION_VARIABLES_ADMIN`](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html#priv_session-variables-admin) | Server administration                             |
| [`SET_USER_ID`](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html#priv_set-user-id) | Server administration                             |
| [`SHOW_ROUTINE`](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html#priv_show-routine) | Server administration                             |
| [`SYSTEM_USER`](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html#priv_system-user) | Server administration                             |
| [`SYSTEM_VARIABLES_ADMIN`](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html#priv_system-variables-admin) | Server administration                             |
| [`TABLE_ENCRYPTION_ADMIN`](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html#priv_table-encryption-admin) | Server administration                             |
| [`VERSION_TOKEN_ADMIN`](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html#priv_version-token-admin) | Server administration                             |
| [`XA_RECOVER_ADMIN`](https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html#priv_xa-recover-admin) | Server administration                             |

# 图形化界面工具

- Workbench(免费): http://dev.mysql.com/downloads/workbench/
- navicat(收费，试用版30天): https://www.navicat.com/en/download/navicat-for-mysql
- Sequel Pro(开源免费，仅支持Mac OS): http://www.sequelpro.com/
- HeidiSQL(免费): http://www.heidisql.com/
- phpMyAdmin(免费): https://www.phpmyadmin.net/
- SQLyog: https://sqlyog.en.softonic.com/

# 安装

# 小技巧

1. 在SQL语句之后加上`\G`会将结果的表格形式转换成行文本形式
2. 查看Mysql数据库占用空间：

```mysql
SELECT table_schema "Database Name"
     , SUM(data_length + index_length) / (1024 * 1024) "Database Size in MB"
FROM information_schema.TABLES
GROUP BY table_schema;
```

# 参考文献

> https://dhc.pythonanywhere.com/entry/share/?key=3ad29aad765a2b98b2b2a745d71bef715507ee9db8adbec98257bac0ad84cbe4#h1-u6743u9650u4E00u89C8u8868
> 这篇笔记是在别人的基础上完善而来，感谢B站的黑马程序员up主，也感谢路途博客。





# 面试题目

## 1、执行一条 SQL 查询语句，期间发生了什么？

<img src="https://cdn.xiaolincoding.com/gh/xiaolincoder/mysql/sql%E6%89%A7%E8%A1%8C%E8%BF%87%E7%A8%8B/mysql%E6%9F%A5%E8%AF%A2%E6%B5%81%E7%A8%8B.png" alt="查询语句执行流程" style="zoom:67%;" />

- 连接器：建立连接，管理连接、校验用户身份；

  与客户端进行 TCP 三次握手建立连接；

  校验客户端的用户名和密码，如果用户名或密码不对，则会报错；

  如果用户名和密码都对了，会读取该用户的权限，然后后面的权限逻辑判断都基于此时读取到的权限；

  >- 如果一个用户已经建立了连接，即使管理员中途修改了该用户的权限，也不会影响已经存在连接的权限。修改完成后，只有再新建的连接才会使用新的权限设置。
  >- 如果你想知道当前 MySQL 服务被多少个客户端连接了，你可以执行 `show processlist` 命令进行查看。
  >- MySQL 定义了空闲连接的最大空闲时长，由 `wait_timeout` 参数控制的，默认值是 8 小时（28880秒） 。同样，我们也能够手动断开空闲的连接，使用的是 kill connection + id 的命令。
  >- MySQL 服务支持的最大连接数由 max_connections 参数控制，比如我的 MySQL 服务默认是 151 个,超过这个值，系统就会拒绝接下来的连接请求，并报错提示“Too many connections”。
  >
  >

- 查询缓存：查询语句如果命中查询缓存则直接返回，否则继续往下执行。**MySQL 8.0 已删除该模块；**

  如果 SQL 是查询语句（select 语句），MySQL 就会先去查询缓存（ Query Cache ）里查找缓存数据，看看之前有没有执行过这一条命令，这个查询缓存是以 key-value 形式保存在内存中的，key 为 SQL 查询语句，value 为 SQL 语句查询的结果。

- 解析 SQL，

  通过解析器对 SQL 查询语句进行**词法分析**、**语法分析**，然后**构建语法树**，方便后续模块读取表名、字段、语句类型；

- 执行 SQL：执行 SQL 共有三个阶段：

  - 预处理阶段：1、检查表或字段是否存在；2、将 `select *` 中的 `*` 符号扩展为表上的所有列。

  - 优化阶段：基于查询成本的考虑， 选择查询成本最小的执行计划；

  - 执行阶段：根据执行计划执行 SQL 查询语句，从存储引擎读取记录，返回给客户端

    > Server 层每从存储引擎读到一条记录就会发送给客户端，之所以客户端显示的时候是直接显示所有记录的，是因为客户端是等查询语句查询完成后，才会显示出所有的记录

## 2、MySQL 一行记录的存储结构

###  MySQL 的数据存放在哪个文件？

```sql
SHOW VARIABLES LIKE 'datadir';
```

一般来说，会存在三个文件，

- `db.opt`，用来存储当前数据库的默认字符集和字符校验规则。
- `t_order.frm` ，t_order 的**表结构**会保存在这个文件。在 MySQL 中建立一张表都会生成一个.frm 文件，该文件是用来保存每个表的元数据信息的，主要包含表结构定义。
- `t_order.ibd`，t_order 的**表数据**会保存在这个文件。表数据既可以存在共享表空间文件（文件名：ibdata1）里，也可以存放在独占表空间文件（文件名：表名字.ibd）。**这个行为是由参数 innodb_file_per_table 控制的，若设置了参数 innodb_file_per_table 为 1，则会将存储的数据、索引等信息单独存储在一个独占表空间**

### 表空间文件的结构

<img src="https://cdn.xiaolincoding.com/gh/xiaolincoder/mysql/row_format/%E8%A1%A8%E7%A9%BA%E9%97%B4%E7%BB%93%E6%9E%84.drawio.png" alt="img" style="zoom:67%;" />

- #### 行（row）

  数据库表中的记录都是按行（row）进行存放的，每行记录根据不同的行格式，有不同的存储结构。

- #### 页（page）

  记录是按照行来存储的，但是数据库的读取并不以「行」为单位，否则一次读取（也就是一次 I/O 操作）只能处理一行数据，效率会非常低。因此，**InnoDB 的数据是按「页」为单位来读写的**。

- #### 区（extent）

  B+ 树中每一层都是通过双向链表连接起来的，如果是以页为单位来分配存储空间，那么链表中相邻的两个页之间的物理位置并不是连续的，可能离得非常远，那么磁盘查询时就会有大量的随机I/O，随机 I/O 是非常慢的。

  B+ 树中每一层都是通过双向链表连接起来的，如果是以页为单位来分配存储空间，那么链表中相邻的两个页之间的物理位置并不是连续的，可能离得非常远，那么磁盘查询时就会有大量的随机I/O，随机 I/O 是非常慢的。

- #### 段 segment

  表空间是由各个段（segment）组成的，段是由多个区（extent）组成的。段一般分为数据段、索引段和回滚段等。

  - 索引段：存放 B + 树的非叶子节点的区的集合；
  - 数据段：存放 B + 树的叶子节点的区的集合；
  - 回滚段：存放的是回滚数据的区的集合，之前讲[事务隔离 (opens new window)](https://xiaolincoding.com/mysql/transaction/mvcc.html)的时候就介绍到了 MVCC 利用了回滚段实现了多版本查询数据。

### InnoDB 行格式

行格式（row_format），就是一条记录的存储结构

InnoDB 提供了 4 种行格式，分别是 Redundant、Compact、Dynamic和 Compressed 行格式。

- Redundant 是很古老的行格式了， MySQL 5.0 版本之前用的行格式，现在基本没人用了。
- 由于 Redundant 不是一种紧凑的行格式，所以 MySQL 5.0 之后引入了 Compact 行记录存储方式，Compact 是一种紧凑的行格式，设计的初衷就是为了让一个数据页中可以存放更多的行记录，从 MySQL 5.1 版本之后，行格式默认设置成 Compact。
- Dynamic 和 Compressed 两个都是紧凑的行格式，它们的行格式都和 Compact 差不多，因为都是基于 Compact 改进一点东西。从 MySQL5.7 版本之后，默认使用 Dynamic 行格式。

### COMPACT 行格式

![compact行格式](https://cdn.xiaolincoding.com/gh/xiaolincoder/mysql/row_format/COMPACT.drawio.png)

可以看到，一条完整的记录分为「记录的额外信息」和「记录的真实数据」两个部分。

#### 记录的额外信息

- **变长字段长度列表**

  这些变长字段的真实数据占用的字节数会按照列的顺序**逆序存放**

- **NULL值列表**

  表中的某些列可能会存储 NULL 值，如果把这些 NULL 值都放到记录的真实数据中会比较浪费空间，所以 Compact 行格式把这些值为 NULL 的列存储到 NULL值列表中。

  如果存在允许 NULL 值的列，则每个列对应一个二进制位（bit），二进制位按照列的顺序逆序排列。

  - 二进制位的值为`1`时，代表该列的值为NULL。
  - 二进制位的值为`0`时，代表该列的值不为NULL。

  ![img](https://cdn.xiaolincoding.com/gh/xiaolincoder/mysql/row_format/null%E5%80%BC%E5%88%97%E8%A1%A85.png)

- **记录头信息**

  记录头信息中包含的内容很多，我就不一一列举了，这里说几个比较重要的：

  - delete_mask ：标识此条数据是否被删除。从这里可以知道，我们执行 detele 删除记录的时候，并不会真正的删除记录，只是将这个记录的 delete_mask 标记为 1。
  - next_record：下一条记录的位置。从这里可以知道，记录与记录之间是通过链表组织的。在前面我也提到了，指向的是下一条记录的「记录头信息」和「真实数据」之间的位置，这样的好处是向左读就是记录头信息，向右读就是真实数据，比较方便。
  - record_type：表示当前记录的类型，0表示普通记录，1表示B+树非叶子节点记录，2表示最小记录，3表示最大记录

  

#### 记录的真实数据

  ![img](https://cdn.xiaolincoding.com/gh/xiaolincoder/mysql/row_format/%E8%AE%B0%E5%BD%95%E7%9A%84%E7%9C%9F%E5%AE%9E%E6%95%B0%E6%8D%AE.png)

  

  - row_id

    如果我们建表的时候指定了主键或者唯一约束列，那么就没有 row_id 隐藏字段了。如果既没有指定主键，又没有唯一约束，那么 InnoDB 就会为记录添加 row_id 隐藏字段。占用 6 个字节。

  - trx_id

    事务id，表示这个数据是由哪个事务生成的。 trx_id是必需的，占用 6 个字节。

  - roll_pointer

    这条记录上一个版本的指针。roll_pointer 是必需的，占用 7 个字节。

  这里的trx_id和roll_pointer 用于MVCC的readview控制

###  varchar(n) 中 n 最大取值为多少？

**MySQL 规定除了 TEXT、BLOBs 这种大对象类型之外，其他所有的列（不包括隐藏列和记录头信息）占用的字节长度加起来不能超过 65535 个字节**。

- 单字段情况

  这是因为我们存储字段类型为 varchar(n) 的数据时，其实分成了三个部分来存储：

  - 真实数据
  - 真实数据占用的字节数
  - NULL 标识，如果不允许为NULL，这部分不需要

​	所以在算 varchar(n) 中 n 最大值时，需要减去 「变长字段长度列表」和 「NULL 值列表」所占用的字节数的。

- 多字段情况

  **如果有多个字段的话，要保证所有字段的长度 + 变长字段字节数列表所占用的字节数 + NULL值列表所占用的字节数 <= 65535**。



### 行溢出后，MySQL 是怎么处理的？

MySQL 中磁盘和内存交互的基本单位是页，一个页的大小一般是 `16KB`，也就是 `16384字节`，而一个 varchar(n) 类型的列最多可以存储 `65532字节`，一些大对象如 TEXT、BLOB 可能存储更多的数据，这时一个页可能就存不了一条记录。这个时候就会**发生行溢出，多的数据就会存到另外的「溢出页」中**。

如果一个数据页存不了一条记录，InnoDB 存储引擎会自动将溢出的数据存放到「溢出页」中。在一般情况下，InnoDB 的数据都是存放在 「数据页」中。但是当发生行溢出时，溢出的数据会存放到「溢出页」中。

当发生行溢出时，在记录的真实数据处只会保存该列的一部分数据，而把剩余的数据放在「溢出页」中，然后真实数据处用 20 字节存储指向溢出页的地址，从而可以找到剩余数据所在的页。大致如下图所示。

<img src="https://cdn.xiaolincoding.com/gh/xiaolincoder/mysql/row_format/%E8%A1%8C%E6%BA%A2%E5%87%BA.png" alt="img" style="zoom:67%;" />



## 3、索引相关

### 索引的分类

- 按「数据结构」分类：**B+tree索引、Hash索引、Full-text索引**。

  - 如果有主键，默认会使用主键作为聚簇索引的索引键（key）；
  - 如果没有主键，就选择第一个**不包含 NULL 值的唯一列**作为聚簇索引的索引键（key）；
  - 在上面两个都没有的情况下，InnoDB 将自动生成一个隐式自增 id 列作为聚簇索引的索引键（key）；

  > #### 为什么 MySQL InnoDB 选择 B+tree 作为索引的数据结构？
  >
  > - ***B+Tree vs B Tree***
  >
  >   B+Tree 只在叶子节点存储数据，而 B 树 的非叶子节点也要存储数据，所以相同数量数据下，B+树的树的层数更低，所以，在相同的磁盘 I/O 次数下，就能查询更多的节点。
  >
  >   另外，B+Tree 叶子节点采用的是双链表连接，适合 MySQL 中常见的基于范围的顺序查找，而 B 树无法做到这一点。
  >
  > - ***B+Tree vs 二叉树***
  >
  >   二叉树的每个父节点的儿子节点个数只能是 2 个  而B+Tree的叶子节点个数能为N 个叶子节点
  >
  >   那么 建树层数 B+Tree更低 那么查询数据使用的磁盘IO次数就更低。
  >
  > - ***B+Tree vs Hash***
  >
  >   之前提到了，hash索引在进行等值查询时，能够一次完成匹配，但是不支持范围查询
  >
  >   因此 B+Tree 索引要比 Hash 表索引有着更广泛的适用场景的原因。

- 按「物理存储」分类：**聚簇索引（主键索引）、二级索引（辅助索引）**。

  - 主键索引的 B+Tree 的叶子节点存放的是实际数据，所有完整的用户记录都存放在主键索引的 B+Tree 的叶子节点里；
  - 二级索引的 B+Tree 的叶子节点存放的是主键值，而不是实际数据。因此在某些查询时，需要进行回表查询

- 按「字段特性」分类：**主键索引、唯一索引、普通索引、前缀索引**。

  - 主键索引就是建立在主键字段上的索引，自动建立

  - 唯一索引建立在 UNIQUE 字段上的索引，一张表可以有多个唯一索引，索引列的值必须唯一，但是允许有空值。自动建立

  - 普通索引就是建立在普通字段上的索引，既不要求字段为主键，也不要求字段为 UNIQUE。

    需要自己建立 `create index index_name on table_name(column name)`

  - 前缀索引是指对字符类型字段的前几个字符建立的索引，而不是在整个字段上建立的索引 适用于较长的字段

    ```sql
    CREATE INDEX index_name
    ON table_name(column_name(length)); 
    ```

- 按「字段个数」分类：**单列索引、联合索引**。

  联合索引进行查询数据时，先按照一个字段的值进行比较，若是前一个字段相同的情况下在按照下一个字段比较。

  > 这里更深层的原因是，是因为`(a, b, c)` 联合索引，是先按 a 排序，在 a 相同的情况再按 b 排序，在 b 相同的情况再按 c 排序。所以，**b 和 c 是全局无序，局部相对有序的**，
  >
  > <img src="https://cdn.xiaolincoding.com/gh/xiaolincoder/mysql/%E7%B4%A2%E5%BC%95/%E8%81%94%E5%90%88%E7%B4%A2%E5%BC%95%E6%A1%88%E4%BE%8B.drawio.png" alt="img" style="zoom:67%;" />
  >
  > 即，只有在a相同的情况下，b才是有序的 即局部有序 
  >
  > 而在全局情况下，b是无序的
  >
  > **利用索引的前提是索引里的 key 是有序的**。
  >
  > 因此，这里也就能进一步解释联合索引中，为什么使用范围查询会导致右侧的列索引失效的问题了(**也就是范围查询的字段可以用到联合索引，但是在范围查询字段的后面的字段无法用到联合索引**。)
  >
  > 其中，假设ab字段创建了联合索引，
  >
  > - `select * from t_table where a > 1 and b = 2`
  >
  >   该语句单独使用>或者<，第一字段a是有序的，所以 a 字段可以在联合索引的 B+Tree 中进行索引查询。
  >
  >   **但是在符合 a > 1 条件的二级索引记录的范围里，b 字段的值是无序的**。
  >
  >   因此，**Q1 这条查询语句只有 a 字段用到了联合索引进行索引查询，而 b 字段并没有使用到联合索引**。
  >
  > 
  >
  > - `select * from t_table where a >= 1 and b = 2`
  >
  >   首先，a肯定不用说，因为是联合索引的第一个，因此，所以 a 字段可以在联合索引的 B+Tree 中进行索引查询。
  >
  >   虽然在符合 a>= 1 条件的二级索引记录的范围里，b 字段的值是「无序」的，**但是对于符合 a = 1 的二级索引记录的范围里，b 字段的值是「有序」的**，所以，在确定需要扫描的二级索引的范围时，当二级索引记录的 a 字段值为 1 时，可以通过 b = 2 条件减少需要扫描的二级索引记录范围，也就是说，从符合 a = 1 and b = 2 条件的第一条记录开始扫描，而不需要从第一个 a 字段值为 1 的记录开始扫描。
  >
  >   所以，**Q2 这条查询语句 a 和 b 字段都用到了联合索引进行索引查询**。
  >
  >   这里的区别 能够使用explain 命令通过 key_len字段的长度进行判断。
  >
  > - 若是 `SELECT * FROM t_table WHERE a BETWEEN 2 AND 8 AND b = 2`呢
  >
  >   `a BETWEEN 2 AND 8` 的意思是查询 a 字段的值在 2 和 8 之间的记录。在Mysql中，BETWEEN 包含了 value1 和 value2 边界值，类似于 >= and =<。所以这条语句类似上一条语句中的<=/>= 因此 **Q3 这条查询语句 a 和 b 字段都用到了联合索引进行索引查询**。
  >
  > - `SELECT * FROM t_user WHERE name like 'j%' and age = 22`，联合索引（name, age）哪一个字段用到了联合索引的 B+Tree？
  >
  >   由于联合索引（二级索引）是先按照 name 字段的值排序的，所以前缀为 ‘j’ 的 name 字段的二级索引记录都是相邻的， 于是在进行索引扫描的时候，可以定位到符合前缀为 ‘j’ 的 name 字段的第一条记录，然后沿着记录所在的链表向后扫描，直到某条记录的 name 前缀不为 ‘j’ 为止。
  >
  >   ![img](https://cdn.xiaolincoding.com/gh/xiaolincoder/mysql/%E7%B4%A2%E5%BC%95/q4-1.drawio.png)
  >
  >   虽然在符合前缀为 ‘j’ 的 name 字段的二级索引记录的范围里，age 字段的值是「无序」的，**但是对于符合 name = j 的二级索引记录的范围里，age字段的值是「有序」的**。
  >
  >   于是，在确定需要扫描的二级索引的范围时，当二级索引记录的 name 字段值为 ‘j’ 时，可以通过 age = 22 条件减少需要扫描的二级索引记录范围（age 字段可以利用联合索引进行索引查询的意思）。也就是说，从符合 `name = 'j' and age = 22` 条件的第一条记录时开始扫描，而不需要从第一个 name 为 j 的记录开始扫描 。如下图的右边：
  >
  >   ![img](https://cdn.xiaolincoding.com/gh/xiaolincoder/mysql/%E7%B4%A2%E5%BC%95/q4-2.drawio.png)
  
  
  
### 索引下推

  对于联合索引（a, b），在执行 `select * from table where a > 1 and b = 2` 语句的时候，只有 a 字段能用到索引，那在联合索引的 B+Tree 找到第一个满足条件的主键值（ID 为 2）后，还需要判断其他条件是否满足（看 b 是否等于 2），**那是在联合索引里判断？还是回主键索引去判断呢？**

  - 在 MySQL 5.6 之前，只能从 ID2 （主键值）开始一个个回表，到「主键索引」上找出数据行，再对比 b 字段值。
  - 而 MySQL 5.6 引入的**索引下推优化**（index condition pushdown)， **可以在联合索引遍历过程中，对联合索引中包含的字段先做判断，直接过滤掉不满足条件的记录，减少回表次数**。



### 索引区分度

建立联合索引时的字段顺序，对索引效率也有很大影响。越靠前的字段被用于索引过滤的概率越高，实际开发工作中**建立联合索引时，要把区分度大的字段排在前面，这样区分度大的字段越有可能被更多的 SQL 使用到**。

<img src="https://cdn.xiaolincoding.com/gh/xiaolincoder/mysql/%E7%B4%A2%E5%BC%95/%E5%8C%BA%E5%88%86%E5%BA%A6.png" alt="区分度计算公式" style="zoom:50%;" />





### 索引使用场景

#### 什么时候适用索引？

- 字段有唯一性限制的，比如商品编码；
- 经常用于 `WHERE` 查询条件的字段，这样能够提高整个表的查询速度，如果查询条件不是一个字段，可以建立联合索引。
- 经常用于 `GROUP BY` 和 `ORDER BY` 的字段，这样在查询的时候就不需要再去做一次排序了，因为我们都已经知道了建立索引之后在 B+Tree 中的记录都是排序好的。

#### 什么时候不需要创建索引？

- `WHERE` 条件，`GROUP BY`，`ORDER BY` 里用不到的字段，索引的价值是快速定位，如果起不到定位的字段通常是不需要创建索引的，因为索引是会占用物理空间的。
- 字段中存在大量重复数据，不需要创建索引，比如性别字段，只有男女，如果数据库表中，男女的记录分布均匀，那么无论搜索哪个值都可能得到一半的数据。在这些情况下，还不如不要索引，因为 MySQL 还有一个查询优化器，查询优化器发现某个值出现在表的数据行中的百分比很高的时候，它一般会忽略索引，进行全表扫描。
- 表数据太少的时候，不需要创建索引；
- 经常更新的字段不用创建索引，比如不要对电商项目的用户余额建立索引，因为索引字段频繁修改，由于要维护 B+Tree的有序性，那么就需要频繁的重建索引，这个过程是会影响数据库性能的。



### 索引优化策略

####  前缀索引优化

使用前缀索引是为了减小索引字段大小，可以增加一个索引页中存储的索引值，有效提高索引的查询速度。

不过，前缀索引有一定的局限性，例如：

- order by 就无法使用前缀索引；
- 无法把前缀索引用作覆盖索引；

#### 覆盖索引优化

#### 主键索引最好是自增的

**如果我们使用自增主键**，那么每次插入的新数据就会按顺序添加到当前索引节点的位置，不需要移动已有的数据，当页面写满，就会自动开辟一个新页面。因为每次**插入一条新记录，都是追加操作，不需要重新移动数据**，因此这种插入数据的方法效率非常高。

**如果我们使用非自增主键**，由于每次插入主键的索引值都是随机的，因此每次插入新的数据时，就可能会插入到现有数据页中间的某个位置，这将不得不移动其它数据来满足新数据的插入，甚至需要从一个页面复制数据到另外一个页面，我们通常将这种情况称为**页分裂**。**页分裂还有可能会造成大量的内存碎片，导致索引结构不紧凑，从而影响查询效率**。

需要 1、叶分裂

​	2、拷贝和插入

​	3、修改指针



#### 索引最好设置为 NOT NULL

- 第一原因：索引列存在 NULL 就会导致优化器在做索引选择的时候更加复杂，更加难以优化，因为可为 NULL 的列会使索引、索引统计和值比较都更复杂，比如进行索引统计时，count 会省略值为NULL 的行。
- 第二个原因：NULL 值是一个没意义的值，但是它会占用物理空间，所以会带来的存储空间的问题，因为 InnoDB 存储记录的时候，如果表中存在允许为 NULL 的字段，那么[行格式 (opens new window)](https://xiaolincoding.com/mysql/base/row_format.html#innodb-行格式有哪些)中**至少会用 1 字节空间存储 NULL 值列表**，如下图的紫色部分：

#### 防止索引失效

- 当我们使用左或者左右模糊匹配的时候，也就是 `like %xx` 或者 `like %xx%`这两种方式都会造成索引失效；
- 当我们在查询条件中对索引列做了计算、函数、类型转换操作，这些情况下都会造成索引失效；
- 联合索引要能正确使用需要遵循最左匹配原则，也就是按照最左优先的方式进行索引的匹配，否则就会导致索引失效。
- 在 WHERE 子句中，如果在 OR 前的条件列是索引列，而在 OR 后的条件列不是索引列，那么索引会失效。



> 在使用explain中，type字段表示了 扫描类型的执行效率
>
> type 字段就是描述了找到所需数据时使用的扫描方式是什么，常见扫描类型的**执行效率从低到高的顺序为**：
>
> - All（全表扫描）；all 是最坏的情况，因为采用了全表扫描的方式。i
> - index（全索引扫描）； index 对索引表进行全扫描
> - range（索引范围扫描）；
> - ref（非唯一索引扫描）；
> - eq_ref（唯一索引扫描）；
> - const（结果只有一条的主键或唯一索引扫描）。
>
> extra 字段显示的结果中
>
> - Using filesort ：当查询语句中包含 group by 操作，而且无法利用索引完成排序操作的时候， 这时不得不选择相应的排序算法进行，甚至可能会通过文件排序，效率是很低的，所以要避免这种问题的出现。
> - Using temporary：使了用临时表保存中间结果，MySQL 在对查询结果排序时使用临时表，常见于排序 order by 和分组查询 group by。效率低，要避免这种问题的出现。
> - Using index：所需数据只需在索引即可全部获得，不须要再到表中取数据，也就是使用了覆盖索引，避免了回表操作，效率不错。



## 4、从数据页的角度看B+树

MySQL 里 InnoDB 存储引擎是采用 B+ 树来组织数据的。

### InnoDB 是如何存储数据的？

记录是按照行来存储的，但是数据库的读取并不以「行」为单位，否则一次读取（也就是一次 I/O 操作）只能处理一行数据，效率会非常低。

因此，**InnoDB 的数据是按「数据页」为单位来读写的**，也就是说，当需要读一条记录的时候，并不是将这个记录本身从磁盘读出来，而是以页为单位，将其整体读入内存。

数据库的 I/O 操作的最小单位是页，**InnoDB 数据页的默认大小是 16KB**，意味着数据库每次读写都是以 16KB 为单位的，一次最少从磁盘中读取 16K 的内容到内存中，一次最少把内存中的 16K 内容刷新到磁盘中。

数据页包括七个部分，结构如下图：

<img src="https://cdn.xiaolincoding.com//mysql/other/fabd6dadd61a0aa342d7107213955a72.png" alt="图片" style="zoom:67%;" />



一、不同数据页间，各个数据页在 File Header(文件头)中有两个指针，分别指向上一个数据页和下一个数据页，连接起来的页相当于一个双向的链表，如下图所示：

<img src="https://cdn.xiaolincoding.com//mysql/other/557d17e05ce90f18591c2305871af665.png" alt="图片" style="zoom:67%;" />





采用链表的结构是让数据页之间不需要是物理上的连续的，而是逻辑上的连续。



二、同一**数据页中的记录按照「主键」顺序组成单向链表**，单向链表的特点就是插入、删除非常方便，但是检索效率不高，最差的情况下需要遍历链表上的所有节点才能完成检索。

三、因此，数据页中有一个**页目录**，起到记录的索引作用，就像我们书那样，针对书中内容的每个章节设立了一个目录，想看某个章节的时候，可以查看目录，快速找到对应的章节的页数，而数据页中的页目录就是为了能快速找到记录。

<img src="https://cdn.xiaolincoding.com//mysql/other/261011d237bec993821aa198b97ae8ce.png" alt="图片" style="zoom:67%;" />

页目录创建的过程如下：

1. 将所有的记录划分成几个组，这些记录包括最小记录和最大记录，但不包括标记为“已删除”的记录；
2. 每个记录组的最后一条记录就是组内最大的那条记录，并且最后一条记录的头信息中会存储该组一共有多少条记录，作为 n_owned 字段（上图中粉红色字段）
3. 页目录用来存储每组最后一条记录的地址偏移量，这些地址偏移量会按照先后顺序存储起来，每组的地址偏移量也被称之为槽（slot），**每个槽相当于指针指向了不同组的最后一个记录**。

从图可以看到，**页目录就是由多个槽组成的，槽相当于分组记录的索引**。然后，因为记录是按照「主键值」从小到大排序的，所以**我们通过槽查找记录时，可以使用二分法快速定位要查询的记录在哪个槽（哪个记录分组），定位到槽后，再遍历槽内的所有记录，找到对应的记录**，无需从最小记录开始遍历整个页中的记录链表。

> 以上面那张图举个例子，5 个槽的编号分别为 0，1，2，3，4，我想查找主键为 11 的用户记录：
>
> - 先二分得出槽中间位是 (0+4)/2=2 ，2号槽里最大的记录为 8。因为 11 > 8，所以需要从 2 号槽后继续搜索记录；
> - 再使用二分搜索出 2 号和 4 槽的中间位是 (2+4)/2= 3，3 号槽里最大的记录为 12。因为 11 < 12，所以主键为 11 的记录在 3 号槽里；
> - 这里有个问题，**「槽对应的值都是这个组的主键最大的记录，如何找到组里最小的记录」**？比如槽 3 对应最大主键是 12 的记录，那如何找到最小记录 9。解决办法是：通过槽 3 找到 槽 2 对应的记录，也就是主键为 8 的记录。主键为 8 的记录的下一条记录就是槽 3 当中主键最小的 9 记录，然后开始向下搜索 2 次，定位到主键为 11 的记录，取出该条记录的信息即为我们想要查找的内容。
> - InnoDB 对每个分组中的记录条数都是有规定的，槽内的记录就只有几条：
>   - 第一个分组中的记录只能有 1 条记录；
>   - 最后一个分组中的记录条数范围只能在 1-8 条之间；
>   - 剩下的分组中记录条数范围只能在 4-8 条之间。

总结来说

1、不同数据页间通过使用指针分别指向上一个数据页和下一个数据页，连接起来的页相当于一个双向的链表，如下图所示。

2、一个数据页内，**数据页中的记录按照「主键」顺序组成单向链表**

3、因为数据页内数据组织形式为链表，只能顺序查询，页目录通过将所有数据分组，存储每组最后一条记录的地址偏移量，对分组记录进行索引，称为槽(slot)，在查找记录时能够通过二分法快速定位要查询的记录在哪个槽），定位到槽后，再遍历槽内的所有记录，找到对应的记录，无需从最小记录开始遍历整个页中的记录链表。

4、页目录中记录包括最小记录和最大记录



#### B+ 树是如何进行查询的

一个数据页中数据和主键是有限的，

但是 当我们需要存储大量的记录时，就需要多个数据页，这时我们就需要考虑如何建立合适的索引，才能方便定位记录所在的页。

为了解决这个问题，**InnoDB 采用了 B+ 树作为索引**。磁盘的 I/O 操作次数对索引的使用效率至关重要，因此在构造索引的时候，我们更倾向于采用“矮胖”的 B+ 树数据结构，这样所需要进行的磁盘 I/O 次数更少，而且 B+ 树 更适合进行关键字的范围查询。

InnoDB 里的 B+ 树中的**每个节点都是一个数据页**，结构示意图如下：

<img src="https://cdn.xiaolincoding.com//mysql/other/7c635d682bd3cdc421bb9eea33a5a413.png" alt="图片" style="zoom:67%;" />





通过上图，我们看出 B+ 树的特点：

- 只有叶子节点（最底层的节点）才存放了数据，非叶子节点（其他上层节）仅用来存放目录项作为索引。
- 非叶子节点分为不同层次，通过分层来降低每一层的搜索量；
- 所有节点按照索引键大小排序，构成一个双向链表，便于范围查询；

> 我们再看看 B+ 树如何实现快速查找主键为 6 的记录，以上图为例子：
>
> - 从根节点开始，通过二分法快速定位到符合页内范围包含查询值的页，因为查询的主键值为 6，在[1, 7)范围之间，所以到页 30 中查找更详细的目录项；
> - 在非叶子节点（页30）中，继续通过二分法快速定位到符合页内范围包含查询值的页，主键值大于 5，所以就到叶子节点（页16）查找记录；
> - 接着，在叶子节点（页16）中，通过槽查找记录时，使用二分法快速定位要查询的记录在哪个槽（哪个记录分组），定位到槽后，再遍历槽内的所有记录，找到主键为 6 的记录。
>
> 可以看到，在定位记录所在哪一个页时，也是通过二分法快速定位到包含该记录的页。定位到该页后，又会在该页内进行二分法快速定位记录所在的分组（槽号），最后在分组内进行遍历查找。



## 5、为什么MySQL采用B+树作为索引

MySQL 的数据是持久化的，意味着数据（索引+记录）是保存到磁盘上的，因为这样即使设备断电了，数据也不会丢失。

磁盘读写的最小单位是**扇区**，扇区的大小只有 `512B` 大小，操作系统一次会读写多个扇区，所以**操作系统的最小读写单位是块（Block）。Linux 中的块大小为 `4KB`**，也就是一次磁盘 I/O 操作会直接读写 8 个扇区。

由于数据库的索引是保存到磁盘上的，因此当我们通过索引查找某行数据的时候，**就需要先从磁盘读取索引到内存**，**再通过索引从磁盘中找到某行数据，然后读入到内存**，也就是说查询过程中会发生多次磁盘 I/O，而磁盘 I/O 次数越多，所消耗的时间也就越大。

### 1、二分查找树 = 二叉搜索树

用数组来实现线性排序的数据虽然简单好用，但是插入新元素的时候性能太低。

因为插入一个元素，需要将这个元素之后的所有元素后移一位，如果这个操作发生在磁盘中呢？这必然是灾难性的。因为磁盘的速度比内存慢几十万倍，所以我们不能用一种线性结构将磁盘排序。

所以这里使用树的结构将前后物理结构不连续的磁盘连接在一起，

**二叉查找树的特点是一个节点的左子树的所有节点都小于这个节点，右子树的所有节点都大于这个节点**，因此，其中序遍历是有序的。

因此，我们有了两个好处，

1、能够通过值对比的方式执行类似二分查找

2、节点通过指针连接，解决了连续结构插入新数据开销大的问题

但是，**当每次插入的元素都是二叉查找树中最大的元素，二叉查找树就会退化成了一条链表，查找数据的时间复杂度变成了 O(n)**，如下动图演示：

![请添加图片描述](https://cdn.xiaolincoding.com//mysql/other/320c5c2653e74f9e9ebfcd8ca9681905.gif)二叉查找树由于存在退化成链表的可能性，会使得查询操作的时间复杂度从 O(logn) 升为 O(n)。而且会随着插入的元素越多，树的高度也变高，意味着需要磁盘 IO 操作的次数就越多，这样导致查询性能严重下降，再加上不能范围查询，所以不适合作为数据库的索引结构。



### 2、自平衡二叉树

在二叉搜索树的基础上，增加了新的约束。**每个节点的左子树和右子树的高度差不能超过 1**。解决了成为链表的可能性。

除了平衡二叉查找树，还有很多自平衡的二叉树，比如红黑树，它也是通过一些约束条件来达到自平衡，

**不管平衡二叉查找树还是红黑树，都会随着插入的元素增多，而导致树的高度变高，这就意味着磁盘 I/O 操作次数多，会影响整体数据查询的效率**。



### 3、B树

之前二叉树的设计中，每个节点最多有2个子节点，那么当节点个数越多的时候，树的高度也会相应变高，这样就会增加磁盘的 I/O 次数，从而影响数据查询的效率。

为了解决降低树的高度的问题，后面就出来了 B 树，它不再限制一个节点就只能有 2 个子节点，而是允许 M 个子节点 (M>2)，从而降低树的高度。B 树的每一个节点最多可以包括 M 个子节点，M 称为 B 树的阶，所以 B 树就是一个多叉树。

这是一个3阶B树的例子，可以看到该树最多有三个节点，并且有着左小右大的约束。

这是其建树过程。

![图片](https://cdn.xiaolincoding.com//mysql/other/9a96956de3be0614f7ec2344741b4dcc.gif)

同样，因为其节点数据的大小关系，查询实际是和二叉搜索树类似。

![图片](https://cdn.xiaolincoding.com//mysql/other/341cedca6863fcd0624febd835165ae2.gif)

但是 B 树的每个节点都包含数据（索引+记录），而用户的记录数据的大小很有可能远远超过了索引数据，这就需要花费更多的磁盘 I/O 操作次数来读到「有用的索引数据」。

而且，在我们查询位于底层的某个节点（比如 A 记录）过程中，「非 A 记录节点」里的记录数据会从磁盘加载到内存，但是这些记录数据是没用的，我们只是想读取这些节点的索引数据来做比较查询，而「非 A 记录节点」里的记录数据对我们是没用的，这样不仅增多磁盘 I/O 操作次数，也占用内存资源。

另外，**如果使用 B 树来做范围查询的话，需要使用中序遍历，这会涉及多个节点的磁盘 I/O 问题，从而导致整体速度下降。**



### 4、B+树

B+ 树就是对 B 树做了一个升级，MySQL 中索引的数据结构就是采用了 B+ 树，B+ 树结构如下图：

<img src="https://cdn.xiaolincoding.com//mysql/other/b6678c667053a356f46fc5691d2f5878.png" alt="图片" style="zoom:67%;" />

B+ 树与 B 树差异的点，主要是以下这几点：

- 叶子节点（最底部的节点）才会存放实际数据（索引+记录），非叶子节点只会存放索引；
- 所有索引都会在叶子节点出现，叶子节点之间构成一个有序链表；
- **非叶子节点的索引也会同时存在在子节点中**，并且是在子节点中所有索引的最大（或最小）。
- **非叶子节点中有多少个子节点，就有多少个索引；**



在各种应用场景中，B+树都比B树更加有优势

- 单点查询

  虽然，B 树进行单个索引查询时，最快可以在 O(1) 的时间代价内就查到。直接根节点就是需要的数据，而从平均时间代价来看，会比 B+ 树稍快一些。

  但是 B 树的查询波动会比较大，因为每个节点即存索引又存记录，所以有时候访问到了非叶子节点就可以找到索引，而有时需要访问到叶子节点才能找到索引。

  **B+ 树的非叶子节点不存放实际的记录数据，仅存放索引，因此数据量相同的情况下，相比存储即存索引又存记录的 B 树，B+树的非叶子节点可以存放更多的索引，因此 B+ 树可以比 B 树更「矮胖」，查询底层节点的磁盘 I/O次数会更少**。

- 插入和删除效率

  B+ 树有大量的冗余节点，这样使得删除一个节点的时候，可以直接从叶子节点中删除，甚至可以不动非叶子节点，这样删除非常快，

<figure class="half">
    <img src="https://cdn.xiaolincoding.com//mysql/other/25508b0cd9c44ef6937fdd737020a7f1.gif" style="zoom: 120%;" \>
    <img src="https://cdn.xiaolincoding.com//mysql/other/2be62679487640bbaac663fa96c7f35f.gif" style="zoom: 120%;" \>
</figure>

​	图中，左边为 B+树的删除过程，由于存在冗余的节点，所以不会发生复杂的树的变形，

​	而右图中，B树的删除，需要可能涉及复杂的树的变形

​	B+ 树的插入也是一样，有冗余节点，插入可能存在节点的分裂（如果节点饱和），但是最多只涉及树的一条路径。而且 B+ 树会自	动平衡，不需要像更多复杂的算法，类似红黑树的旋转操作等。

​	因此，**B+ 树的插入和删除效率更高**。3



- 范围查询

  B 树和 B+ 树等值查询原理基本一致，先从根节点查找，然后对比目标数据的范围，最后递归的进入子节点查找。

  但是因为 **B+ 树所有叶子节点间还有一个链表进行连接，这种设计对范围查找非常有帮助**。

  而 B 树没有将所有叶子节点用链表串联起来的结构，因此只能通过树的遍历来完成范围查询，这会涉及多个节点的磁盘 I/O 操作，范围查询效率不如 B+ 树。

  

但是 Innodb 使用的 B+ 树有一些特别的点，比如：

- B+ 树的叶子节点之间是用「双向链表」进行连接，这样的好处是既能向右遍历，也能向左遍历。
- B+ 树点节点内容是数据页，数据页里存放了用户的记录以及各种信息，每个数据页默认大小是 16 KB。



## 6、MySQL 单表不要超过 2000W 行，靠谱吗？

Innodb 的索引内部用的是 B+ 树。

开启 `innodb_file_per_table`后每张表创建后都会生成一个ibd文件，虽然数据表中，他们看起来是一条连着一条，但是实际上在文件中它被分成很多小份的数据页，而且每一份都是 16K。

<img src="https://cdn.xiaolincoding.com//mysql/other/ee29f7137057b226658627ff61a6a514.png" alt="图片" style="zoom:67%;" />



而一个数据页的结构如下:

<img src="https://cdn.xiaolincoding.com//mysql/other/c34b589e12e5bc0855c9bdeab0c63a88.png" alt="图片" style="zoom:67%;" />





一个 InnoDB 数据页的存储空间大致被划分成了 7 个部分，有的部分占用的字节数是确定的，有的部分占用的字节数是不确定的。

在页的 7 个组成部分中，我们自己存储的记录会按照我们指定的行格式存储到 `User Records` 部分。

但是在一开始生成页的时候，其实并没有 User Records 这个部分，每当我们插入一条记录，都会从 Free Space 部分，也就是尚未使用的存储空间中申请一个记录大小的空间划分到 User Records 部分。当 Free Space 部分的空间全部被 User Records 部分替代掉之后，也就意味着这个页使用完了，如果还有新的记录插入的话，就需要去申请新的页了。

<img src="https://cdn.xiaolincoding.com//mysql/other/ea5cc8c67b7656d3f2a11e42293a0244.png" alt="图片" style="zoom:67%;" />



众所周知，innode中的索引使用的是B+树，每一个节点都是一个数据页，叶子阶段存放数据，非叶子节点在它的内容区中有 id 和 页号地址两部分：

- id ：对应页中记录的最小记录 id 值；(这里是根据索引的类型来的，看你创建的索引的排序基础是什么，之后按照顺序)
- 页号：地址是指向对应页的指针；

例子如下:

<img src="https://cdn.xiaolincoding.com//mysql/other/585429e5078566bda9b2fa18f85215af.png" alt="图片" style="zoom:67%;" />



比如说我们需要查找一个 id=6 的行数据：

- 因为在非叶子节点中存放的是页号和该页最小的 id，所以我们从顶层开始对比，首先看页号 10 中的目录，有 [id=1, 页号 = 20],[id=5, 页号 = 30], 说明左侧节点最小 id 为 1，右侧节点最小 id 是 5。6>5, 那按照二分法查找的规则，肯定就往右侧节点继续查找；
- 找到页号 30 的节点后，发现这个节点还有子节点（非叶子节点），那就继续比对，同理，6>5 && 6<7, 所以找到了页号 60；
- 找到页号 60 之后，发现此节点为叶子节点（数据节点），于是将此页数据加载至内存进行一一对比，结果找到了 id=6 的数据行。

> 从上述的过程中发现，我们为了查找 id=6 的数据，总共查询了三个页，如果三个页都在磁盘中（未提前加载至内存），那么最多需要经历三次的磁盘 IO。



所以对于一个同样一个 16K 的页，非叶子节点里的每条数据都指向新的页，而新的页有两种可能

- 如果是叶子节点，那么里面就是一行行的数据
- 如果是非叶子节点的话，那么就会继续指向新的页

对于非叶子节点，

​	首先需要刨除数据页基本格式的磁盘开销， File Header (38 byte)、Page Header (56 Byte)、Infimum + Supermum（26 byte）、	File Trailer（8byte）, 再加上页目录，大概 1k 左右。

​	主键我们假设是 Bigint (8 byte), 而页号也是固定的（4Byte）, 那么索引页中的一条数据也就是 12byte，15*1024/12 ≈1280行

​	那么非叶子节点最多能够记录1280行索引，



对于叶子节点，

​	这边我们暂时按一条行数据 1k 来算，那一页就能存下 15 条，Y = 15*1024/1000 ≈15。



- 假设 B+ 树是两层，那么一个B+树索引记录的数据行为 Total = （1280 ^1 ）*15 = 19200
- 假设 B+ 树是三层，那么一个B+树索引记录的数据行为， Total = （1280 ^2） *15 = 24576000 （约 2.45kw）

那么三层B+树，兼顾数据量和访问速度，一行数据为 1k Byte来算，大约就是2000w行数据



所以，这个数据实际上是需要通过实际情况来计算的，比如我实际当行的数据占用空间不是 1K , 而是 5K, 那么单个数据页最多只能放下 3 条数据。同样，还是按照 三层B+树的值来计算，那 Total = （1280 ^2） *3 = 4915200 （近 500w）



## 7、索引失效的情况

虽然，InnoDB 和 MyISAM 都支持 B+ 树索引，但是它们数据的存储结构实现方式不同。不同之处在于：

- InnoDB 存储引擎：B+ 树索引的叶子节点保存数据本身；
- MyISAM 存储引擎：B+ 树索引的叶子节点保存数据的物理地址；



### 1、模糊匹配左边出现通配符时会出现索引失效的可能性，

```mysql
// name 字段为二级索引
select * from t_user where name like '%林';
```

**因为索引 B+ 树是按照「索引值」有序排列存储的，只能根据前缀进行比较。**

<img src="https://cdn.xiaolincoding.com//mysql/other/b161c1d88f978e42077f8c14e10972a7.png" alt="图片" style="zoom:67%;" />

### 2、对索引字段值使用函数

因为索引保存的是索引字段的原始值，而不是经过函数计算后的值，自然就没办法走索引了。

不过，从 MySQL 8.0 开始，索引特性增加了函数索引，即可以针对函数计算后的值建立一个索引，也就是说该索引的值是函数计算后的值，所以就可以通过扫描索引来查询数据。

```mysql
-- 只能以这种方式加 给name字段的长度做一个索引 那么在length(name)的查询中就能使用索引了
alter table t_user add key idx_name_length ((length(name)));
```

### 3、对索引进行表达式计算

```mysql
-- 不走索引
explain select * from t_user where id + 1 = 10;
-- 走索引
explain select * from t_user where where id = 10 - 1
```

这里的原因和使用函数的原因是类似的，

因为索引保存的是索引字段的原始值，而不是 id + 1 表达式计算后的值，所以无法走索引，只能通过把索引字段的取值都取出来，然后依次进行表达式的计算来进行条件判断，因此采用的就是全表扫描的方式。

### 4、对索引字段进行隐式类型转换

如果索引字段是字符串类型，但是在条件查询中，输入的参数是整型的话，你会在执行计划的结果发现这条语句会走全表扫描。

但是，这个例子

```mysql
 -- id 是整型，但是下面这条语句还是走了索引扫描的。
 explain select * from t_user where id = '1';
```

> **MySQL 在遇到字符串和数字比较的时候，会自动把字符串转为数字，然后再进行比较**。
>
> 例如比较 select "10">9
>
> - 如果规则是 MySQL 会将自动「字符串」转换成「数字」，就相当于 select 10 > 9，这个就是数字比较，所以结果应该是 1；
> - 如果规则是 MySQL 会将自动「数字」转换成「字符串」，就相当于 select "10" > "9"，这个是字符串比较，字符串比较大小是逐位从高位到低位逐个比较（按ascii码） ，那么"10"字符串相当于 “1”和“0”字符的组合，所以先是拿 “1” 字符和 “9” 字符比较，因为 “1” 字符比 “9” 字符小，所以结果应该是 0。
>
> 
>
> ```mysql
> -- 例子一的查询语句
> select * from t_user where phone = 1300000001;
> -- 这是因为 phone 字段为字符串，所以 MySQL 要会自动把字符串转为数字，所以这条语句相当于：
> select * from t_user where CAST(phone AS signed int) = 1300000001;
> -- CAST 函数是作用在了 phone 字段，而 phone 字段是索引，也就是对索引使用了函数！而前面我们也说了，对索引使用函数是会导致索引失效的。
> 
> 
> -- 在例子二中，
> select * from t_user where id = "1";
> -- 这时因为字符串部分是输入参数，也就需要将字符串转为数字，所以这条语句相当于：
> -- 索引字段并没有用任何函数，CAST 函数是用在了输入参数，因此是可以走索引扫描的。
> select * from t_user where id = CAST("1" AS signed int);
> 
> 
> ```



### 5、联合索引非最左匹配

在联合索引的情况下，数据是按照索引第一列排序，第一列数据相同时才会按照第二列排序。即后面的字段中实际上处于局部有序，全局无序的状态，因此，若跳过左边字段名，后面无序状态的索引实际上是用不到的。



> 这里再提及一下，索引下推的概念
>
> 例如 创建联合索引 create index idx_a_b_c on tab(a,b,c)
>
> 例如执行 where a = 1 and c = 3 
>
> - MySQL 5.5 的话，前面 a 会走索引，在联合索引找到主键值后，开始回表，到主键索引读取数据行，Server 层从存储引擎层获取到数据行后，然后在 Server 层再比对 c 字段的值。
> - 从 MySQL 5.6 之后，有一个**索引下推功能**，可以在存储引擎层进行索引遍历过程中，**对索引中包含的字段先做判断**，直接过滤掉不满足条件的记录，再返还给 Server 层，从而减少回表次数。



### 6、where字句中的or 

在 WHERE 子句中，如果在 OR 前的条件列是索引列，而在 OR 后的条件列不是索引列，那么索引会失效。

这是因为 OR 的含义就是两个只要满足一个即可，因此只有一个条件列是索引列是没有意义的，只要有条件列不是索引列，就会进行全表扫描。



## 8、一些特殊情况

之前提到过，在like使用左匹配和非最左匹配法则时，会发生索引失效的情况。

但是，这里也存在着一些特殊情况

题目二

一个表有两个字段，其中name是索引字段，id拥有自增主键索引。

如果执行以下查询语句

- `select * from s where name like 'xxx';`
- `select * from s where name like 'xxx%';`
- `select * from s where name like '%xxx';`
- `select * from s where name like '%xxx%';`

其中，毫无疑问，前两条能够使用索引

![图片](https://cdn.xiaolincoding.com//mysql/other/d250a6ba3068ef41da9039974dad206a.png)

但是，在这个问题中，后两条也能够使用索引。

![图片](https://cdn.xiaolincoding.com//mysql/other/948ac3e63c36a93101860e7da11ddc42.png)



这张表中，首先，这张表的字段没有「非索引」字段，所以 `select *` 相当于 `select id,name`，然后**这个查询的数据都在二级索引的 B+ 树，因为二级索引的 B+ 树的叶子节点包含「索引值+主键值」，所以查二级索引的 B+ 树就能查到全部结果了，这个就是覆盖索引。**

但是执行计划里的 type 是 `index`，这代表着是通过全扫描二级索引的 B+ 树的方式查询到数据的，也就是遍历了整颗索引树。

而第一和第二条查询语句的执行计划中 type 是 `range`，表示对索引列进行范围查询，也就是利用了索引树的有序性的特点，通过查询比较的方式，快速定位到了数据行。

所以，type=range 的查询效率会比 type=index 的高一些。

> 1、为什么选择全扫描二级索引树，而不扫描聚簇索引树呢？
>
> 因为二级索引树的记录东西很少，就只有「索引列+主键值」，而聚簇索引记录的东西会更多，比如聚簇索引中的叶子节点则记录了主键值、事务 id、用于事务和 MVCC 的回滚指针以及所有的剩余列。
>
> 再加上，这个 select * 不用执行回表操作。
>
> 所以， MySQL 优化器认为直接遍历二级索引树要比遍历聚簇索引树的成本要小的多，因此 MySQL 选择了「全扫描二级索引树」的方式查询数据。
>
> 2、为什么这个数据表加了非索引字段，执行同样的查询语句后，怎么变成走的是全表扫描呢？
>
> 加了其他字段后，`select * from t_user where name like "%xx";` 要查询的数据就不能只在二级索引树里找了，得需要回表操作才能完成查询的工作，再加上是左模糊匹配，无法利用索引树的有序性来快速定位数据，所以得在二级索引树逐一遍历，获取主键值后，再到聚簇索引树检索到对应的数据行，这样实在太累了。
>
> 所以，优化器认为上面这样的查询过程的成本实在太高了，所以直接选择全表扫描的方式来查询数据。
>
> 



总结来说：

1、如果数据库表中的字段只有主键+二级索引，那么即使使用了左模糊匹配，也不会走全表扫描（type=all），而是走全扫描二级索引树(type=index)。

2、情况相似，但是如果数据库表中的字段都是索引的话，即使查询过程中，没有遵循最左匹配原则，也是走全扫描二级索引树(type=index)

> <img src="https://cdn.xiaolincoding.com//mysql/other/35d04bff09bb638727245c7f9aa95b5c.png" alt="图片" style="zoom:67%;" />



## 9、几种count的实现和效率

当我们对一张数据表中的记录进行统计的时候，习惯都会使用 count 函数来统计，但是 count 函数传入的参数有很多种，比如 count(1)、count(`*`)、count(字段) 等。

其中，**按照性能排序，count(*)=count(1)>count(主键)>count(字段)**

### count函数

count() 是一个聚合函数，函数的参数不仅可以是字段名，也可以是其他任意表达式，该函数作用是**统计符合查询条件的记录中，函数指定的参数不为 NULL 的记录有多少个**

- 假设 count() 函数的参数是字段名，如下：

  ```mysql
  select count(name) from t_order;
  ```

  这条语句是统计「 t_order 表中，name 字段不为 NULL 的记录」有多少个。也就是说，如果某一条记录中的 name 字段的值为 NULL，则就不会被统计进去。

- 再来假设 count() 函数的参数是数字 1 这个表达式，如下：

  ```mysql
  select count(name) from t_order;
  ```

  这条语句是统计「 t_order 表中，1 这个表达式不为 NULL 的记录」有多少个。

  1 这个表达式就是单纯数字，它永远都不是 NULL，所以上面这条语句，其实是在统计 t_order 表中有多少个记录。



### count(主键字段)

在通过 count 函数统计有多少个记录时，MySQL 的 server 层会维护一个名叫 count 的变量。

server 层会循环向 InnoDB 读取一条记录，如果 count 函数指定的参数不为 NULL，那么就会将变量 count 加 1，直到符合查询的全部记录被读完，就退出循环。最后将 count 变量的值发送给客户端。

InnoDB 是通过 B+ 树来记录索引的，根据索引的类型又分为聚簇索引和二级索引，它们区别在于，聚簇索引的叶子节点存放的是实际数据，而二级索引的叶子节点存放的是主键值，而不是实际数据。

- 如果表里只有主键索引，没有二级索引时，

  那么，InnoDB 循环遍历聚簇索引，将读取到的记录返回给 server 层，然后读取记录中的 id 值，就会 id 值判断是否为 NULL，如果不为 NULL，就将 count 变量加1.

- 如果表里有二级索引时，**InnoDB 循环遍历的对象就不是聚簇索引，而是二级索引。**

  这是因为相同数量的二级索引记录可以比聚簇索引记录占用更少的存储空间，所以二级索引树比聚簇索引树小，这样遍历二级索引的 I/O 成本比遍历聚簇索引的 I/O 成本小，因此「优化器」优先选择的是二级索引。

<img src="https://cdn.xiaolincoding.com//mysql/other/aac550602ef1022e0b45020dbe0f716a.png" alt="图片" style="zoom:67%;" />

### count(1)

- 如果表里只有主键索引，没有二级索引时。

  那么，InnoDB 循环遍历聚簇索引（主键索引），将读取到的记录返回给 server 层，**但是不会读取记录中的任何字段的值**，因为 count 函数的参数是 1，不是字段，所以不需要读取记录中的字段值。参数 1 很明显并不是 NULL，因此 server 层每从 InnoDB 读取到一条记录，就将 count 变量加 1。

- 如果表里有二级索引时，InnoDB 循环遍历的对象就二级索引了。

![图片](https://cdn.xiaolincoding.com//mysql/other/01e83441a7721f0864deb1ac14ad8ea6.png)

**可以看到，count(1)比count(主键)少了一个步骤，不需要读取数据，因此其效率更高。**



### count(*)

**count(`*`) 其实等于 count(`0`)**，也就是说，当你使用 count(`*`) 时，MySQL 会将 `*` 参数转化为参数 0 来处理。

所以，**count(\*) 执行过程跟 count(1) 执行过程基本一样的**，性能没有什么差异。

> 在 MySQL 5.7 的官方手册中有这么一句话：
>
> *InnoDB handles SELECT COUNT(`*`) and SELECT COUNT(`1`) operations in the same way. There is no performance difference.*
>
> *翻译：InnoDB以相同的方式处理SELECT COUNT（`\*`）和SELECT COUNT（`1`）操作，没有性能差异。*



### count(字段)

count(字段)的效率是最差的

一是，向上面几种情况中记录的那样，count(字段)读取回数据到server层时需要进行是否为null的判断，但是主键不为null。

因此 count(字段)的效率应该时最低的

二是，在count(1),count(主键)中都通过遍历索引来进行计数，你这里有没有能用的索引还是未知吧，因此效率会比较低，可能会使用全表查询的方式来遍历。



> 1
>
> count(1)、 count(*)、 count(主键字段)在执行的时候，如果表里存在二级索引，优化器就会选择二级索引进行扫描。
>
> 应为同样数据表下的二级索引比聚集索引占用更少的存储空间，这样遍历二级索引的 I/O 成本比遍历聚簇索引的 I/O 成本小。
>
> 因为叶子节点的组织方式是数据页，连续的数据页处理肯定比链接的数据页处理IO成本小。
>
> 所以，如果要执行 count(1)、 count(*)、 count(主键字段) 时，尽量在数据表上建立二级索引，这样优化器会自动采用 key_len 最小的二级索引进行扫描，相比于扫描主键索引效率会高一些。
>
> 2
>
> 在MyISAM引擎中，执行 count 函数只需要 O(1 )复杂度，这是因为每张 MyISAM 的数据表都有一个 meta 信息有存储了row_count值。
>
> 而 InnoDB 存储引擎是支持事务的，同一个时刻的多个查询，由于多版本并发控制（MVCC）的原因，InnoDB 表“应该返回多少行”也是不确定的，所以无法像 MyISAM一样，只维护一个 row_count 变量。



### count优化

前面提到过，使用count(*)，count(1)效率比较高，

但是对于一张比较大的表，就算使用了二级索引，花费的时间也比较多

所以这里也有一些优化方法

- ##### 近似值

  explain语句在之前一直用于统计语句性能，

  如果你的业务对于统计个数不需要很精确，我们就可以使用 show table status 或者 explain 命令来表进行估算

  执行 explain 命令效率是很高的，因为它并不会真正的去查询，在rows字段能够显示对行数的估算

- ##### 额外表保存计数值

  如果是想精确的获取表的记录总数，我们可以将这个计数值保存到单独的一张计数表中。

  当我们在数据表插入一条记录的同时，将计数表中的计数字段 + 1。也就是说，在新增和删除操作时，我们需要额外维护这个计数表。这里直接定义一个触发器解决问题。

  ```mysql
  create trigger XXX after/before insert/delete for each row
  begin
  
  
  end;
  
  show triggers;
  drop trigger xxx;
  ```





## 10、事务有那些特性

事务看起来感觉简单，但是要实现事务必须要遵守 4 个特性，分别如下：

- **原子性（Atomicity）**：一个事务中的所有操作，要么全部完成，要么全部不完成，不会结束在中间某个环节，而且事务在执行过程中发生错误，会被回滚到事务开始前的状态，就像这个事务从来没有执行过一样，就好比买一件商品，购买成功时，则给商家付了钱，商品到手；购买失败时，则商品在商家手中，消费者的钱也没花出去。
- **一致性（Consistency）**：是指事务操作前和操作后，数据满足完整性约束，数据库保持一致性状态。比如，用户 A 和用户 B 在银行分别有 800 元和 600 元，总共 1400 元，用户 A 给用户 B 转账 200 元，分为两个步骤，从 A 的账户扣除 200 元和对 B 的账户增加 200 元。一致性就是要求上述步骤操作后，最后的结果是用户 A 还有 600 元，用户 B 有 800 元，总共 1400 元，而不会出现用户 A 扣除了 200 元，但用户 B 未增加的情况（该情况，用户 A 和 B 均为 600 元，总共 1200 元）。
- **隔离性（Isolation）**：数据库允许多个并发事务同时对其数据进行读写和修改的能力，隔离性可以防止多个事务并发执行时由于交叉执行而导致数据的不一致，因为多个事务同时使用相同的数据时，不会相互干扰，每个事务都有一个完整的数据空间，对其他并发事务是隔离的。也就是说，消费者购买商品这个事务，是不影响其他消费者购买的。
- **持久性（Durability）**：事务处理结束后，对数据的修改就是永久的，即便系统故障也不会丢失。

事务隔离级别是怎么实现的

- 持久性是通过 redo log （重做日志）来保证的；
- 原子性是通过 undo log（回滚日志） 来保证的；
- 隔离性是通过 MVCC（多版本并发控制） 或锁机制来保证的；
- 一致性则是通过持久性+原子性+隔离性来保证；

## 11、并发事务会引起什么问题

MySQL 服务端是允许多个客户端连接的，这意味着 MySQL 会出现同时处理多个事务的情况。当多个事务处理相同的数据时，就可能出现**就可能出现脏读（dirty read）、不可重复读（non-repeatable read）、幻读（phantom read）**。

- 脏读

  **如果一个事务「读到」了另一个「未提交事务修改过的数据」，就意味着发生了「脏读」现象。**

- 不可重复读

  **在一个事务内多次读取同一个数据，如果出现前后两次读到的数据不一样的情况，就意味着发生了「不可重复读」现象。**

- 幻读

  **在一个事务内多次查询某个符合查询条件的「记录数量」，如果出现前后两次查询到的记录数量不一样的情况，就意味着发生了「幻读」现象。**

严重性排序：脏读>不可重复读>幻读

## 12、事务的隔离级别

当多个事务并发执行时可能会遇到「脏读、不可重复读、幻读」的现象，这些现象会对事务的一致性产生不同程序的影响。

从严重性角度来看，脏读>不可重复读>幻读

SQL 标准提出了四种隔离级别来规避这些现象，隔离级别越高，性能效率就越低，这四个隔离级别如下：

- **读未提交（\*read uncommitted\*）**，指一个事务还没提交时，它做的变更就能被其他事务看到；
- **读提交（\*read committed\*）**，指一个事务提交之后，它做的变更才能被其他事务看到；
- **可重复读（\*repeatable read\*）**，指一个事务执行过程中看到的数据，一直跟这个事务启动时看到的数据是一致的，**MySQL InnoDB 引擎的默认隔离级别**；
- **串行化（\*serializable\* ）**；会对记录加上读写锁，在多个事务对这条记录进行读写操作时，如果发生了读写冲突的时候，后访问的事务必须等前一个事务执行完成，才能继续执行；
- <img src=".\asset\images\image-20240502084535821.png" alt="image-20240502084535821" style="zoom:67%;" />



**MySQL InnoDB 引擎的默认隔离级别虽然是「可重复读」，但是它很大程度上避免幻读现象（并不是完全解决了，详见这篇[文章 (opens new window)](https://xiaolincoding.com/mysql/transaction/phantom.html)）**，在可重复读隔离级别中解决的方案有两种：

- 针对**快照读**（普通 select 语句），是**通过 MVCC 方式解决了幻读**，因为可重复读隔离级别下，事务执行过程中看到的数据，一直跟这个事务启动时看到的数据是一致的，即使中途有其他事务插入了一条数据，是查询不出来这条数据的，所以就很好了避免幻读问题。
- 针对**当前读**（select ... for update 等语句），每次读都是拿到最新版本的数据，是**通过 next-key lock（记录锁+间隙锁）方式解决了幻读**，因为当执行 select ... for update 语句的时候，会加上 next-key lock，如果有其他事务在 next-key lock 锁范围内插入了一条记录，那么这个插入语句就会被阻塞，无法成功插入，所以就很好了避免幻读问题。

这四种隔离级别具体是如何实现的呢？

- 对于「读未提交」隔离级别的事务来说，因为可以读到未提交事务修改的数据，所以直接读取最新的数据就好了；
- 对于「串行化」隔离级别的事务来说，通过加读写锁的方式来避免并行访问；
- 对于「读提交」和「可重复读」隔离级别的事务来说，它们是通过 **Read View *来实现的，它们的区别在于创建 Read View 的时机不同，大家可以把 Read View 理解成一个数据快照，就像相机拍照那样，定格某一时刻的风景。\**「读提交」隔离级别是在「每个语句执行前」都会重新生成一个 Read View，而「可重复读」隔离级别是「启动事务时」生成一个 Read View，然后整个事务期间都在用这个 Read View**。

>注意，执行「开始事务」命令，并不意味着启动了事务。在 MySQL 有两种开启事务的命令，分别是：
>
>- 第一种：begin/start transaction 命令；
>
>  执行了 begin/start transaction 命令后，并不代表事务启动了。只有在执行这个命令后，执行了第一条 select 语句，才是事务真正启动的时机；
>
>- 第二种：start transaction with consistent snapshot 命令；
>
>  执行了 start transaction with consistent snapshot 命令，就会马上启动事务。



## 13、readview

就上面的内容可知，为了维护读提交和可重复读，MySQL使用了readview，相当于数据快照，但是具体读取哪一个版本的数据，即readview数据的选取还依赖于几个标识

- Read View 的四个字段

  - m_ids ：指的是在创建 Read View 时，当前数据库中「活跃事务」的**事务 id 列表**，注意是一个列表，**“活跃事务”指的就是，启动了但还没提交的事务**。
  - min_trx_id ：指的是在创建 Read View 时，当前数据库中「活跃事务」中事务 **id 最小的事务**，也就是 m_ids 的最小值。
  - max_trx_id ：这个并不是 m_ids 的最大值，而是**创建 Read View 时当前数据库中应该给下一个事务的 id 值**，也就是全局事务中最大的事务 id 值 + 1；
  - creator_trx_id ：指的是**创建该 Read View 的事务的事务 id**。

- 聚簇索引记录中两个跟事务有关的隐藏列；

  对于使用 InnoDB 存储引擎的数据库表，它的聚簇索引记录中都包含下面两个隐藏列，

  <img src=".\asset\images\image-20240502085709684.png" alt="image-20240502085709684" style="zoom:67%;" />

  - trx_id，当一个事务对某条聚簇索引记录进行改动时，就会**把该事务的事务 id 记录在 trx_id 隐藏列里**；
  - roll_pointer，每次对某条聚簇索引记录进行改动时，都会把旧版本的记录写入到 undo 日志中，然后**这个隐藏列是个指针，指向每一个旧版本记录**，于是就可以通过它找到修改前的记录

  > 所以这些条件组合下来，就有这几种情况
  >
  > trx_id 修改当前数据的事务id
  >
  > min_trx_id 指的是在创建 Read View 时，当前数据库中「活跃事务」中事务 id 最小的事务
  >
  > max_trx_id：活跃事务最大id+1
  >
  > m_ids ：指的是在创建 Read View 时，当前数据库中「活跃事务」的事务 id 列表
  >
  > <img src=".\asset\images\image-20240502090059400.png" alt="image-20240502090059400" style="zoom:67%;" />
  >
  > - 当min_trx_id > trx_id 时，说明当前事务在创建read_view时已经提交，所以该版本的记录对当前事务**可见**。
  >
  > - 当trx_id 值大于等于 Read View 中的 `max_trx_id` 值，表示这个版本的记录是在创建 Read View **后**才启动的事务生成的，所以该版本的记录对当前事务**不可见**。
  > - 如果记录的 trx_id 值在 Read View 的 `min_trx_id` 和 `max_trx_id` 之间，需要判断 trx_id 是否在 m_ids 列表中：
  >   - 如果记录的 trx_id **在** `m_ids` 列表中，表示生成该版本记录的活跃事务依然活跃着（还没提交事务），所以该版本的记录对当前事务**不可见**。
  >   - 如果记录的 trx_id **不在** `m_ids`列表中，表示生成该版本记录的活跃事务已经被提交，所以该版本的记录对当前事务**可见**。

  **这种通过「版本链」来控制并发事务访问同一个记录时的行为就叫 MVCC（多版本并发控制）。**



## 14、读提交和可重复读下的readview

**由于隔离级别是「读提交」，所以事务 B 在每次读数据的时候，会重新创建 Read View**

**由于隔离级别时「可重复读」，所以事务再次读取记录时，还是基于第一次查询时创建的 Read View 。**

检查步骤一般如下：

1、检查修改当前数据的事务trx_id，是否位于min_trx_id和max_trx_id 之间

- trx_id<min_trx_id 已经提交了，能够读取
- trx_id>=min_trx_id 表示这个版本的记录是在创建 Read View **后**才启动的事务生成的,不可见，检查上一个版本的记录

2、如果 trx_id 在最大最小值之间，检查trx_id 是否在**活跃事务列表**m_ids中，

- 如果在，那么表示该事务还并未体提交，因此，不能读取这条数据，继续找上一个版本的数据
- 如果不在，说明已经提交了，能够读取





## 15、可重复读的快照读和当前读解决了幻读的问题吗？

解决了一部分的幻读问题，但是还有其他的问题，

例如

![image-20240502100244580](.\asset\images\image-20240502100244580.png)

可重复读的快照读，**后续沿用的时第一次生成的readview并不是数据**，后续仍需要根据条件进行判断。

再比如：这个场景也会发生幻读现象。

- T1 时刻：事务 A 先执行「快照读语句」：select * from t_test where id > 100 得到了 3 条记录。
- T2 时刻：事务 B 往插入一个 id= 200 的记录并提交；
- T3 时刻：事务 A 再执行「当前读语句」 select * from t_test where id > 100 for update 就会得到 4 条记录，此时也发生了幻读现象。

**MySQL 可重复读隔离级别并没有彻底解决幻读，只是很大程度上避免了幻读现象的发生。**

要避免这类特殊场景下发生幻读的现象的话，就是尽量在开启事务之后，马上执行 select ... for update 这类当前读的语句，因为它会对记录加 next-key lock，从而避免其他事务插入一条新记录。





## 16、MySQL中的锁

### 全局锁

```mysql
flush tables with read lock;

unlock tables;
```

- 执行后，**整个数据库就处于只读状态了**，业务只能读数据，而不能更新数据，这时其他线程执行以下操作，都会被阻塞：

  - 对数据的增删改操作，比如 insert、delete、update等语句；
  - 对表结构的更改操作，比如 alter table、drop table 等语句。

  

- 应用场景

  全局锁主要应用于做**全库逻辑备份**，这样在备份数据库期间，不会因为数据或表结构的更新，而出现备份文件的数据与预期的不一样。

  备份数据库的工具是 mysqldump，在使用 mysqldump 时加上 `–single-transaction` 参数的时候，就会在备份数据库之前先开启事务。这种方法只适用于支持「可重复读隔离级别的事务」的存储引擎。

### 表级锁

MySQL中主要包含这几种表级锁

表锁，元数据锁(MDL) 意向锁  auto-INC锁

#### 表锁

锁除了会限制别的线程的读写外，也会限制本线程接下来的读写操作。

```mysql
//表级别的共享锁，也就是读锁； 
-- 那么本线程接下来如果要对学生表执行写操作的语句，是会被阻塞的，当然其他线程对学生表进行写操作时也会被阻塞，直到锁被释放。
lock tables t_student read;

//表级别的独占锁，也就是写锁；
-- **对当前加锁的客户端，可读可写，对于其他的客户端，不可读也不可写,读写阻塞直至该锁释放
lock tables t_stuent write;

-- 解锁 
unlock tables
```



#### 元数据锁

对数据库表进行操作时，会自动给这个表加上 MDL：

- 对一张表进行 CRUD 操作时，加的是 **MDL 读锁**；
- 对一张表做结构变更操作的时候，加的是 **MDL 写锁**

| 锁       | 对应SQL                            | 锁类型                                  | 说明                                               |
| -------- | ---------------------------------- | --------------------------------------- | -------------------------------------------------- |
| 表锁     | **`lock tables xxx read / write`** | SHARED_READ_ONLY / SHARED_NO_READ_WRITE |                                                    |
|          |                                    |                                         |                                                    |
| 元数据锁 | select                             | SHARED_READ                             | 与SHARED_READ、SHARED_WRITE兼容，与 EXCLUSIVE互斥  |
|          | insert 、update、 delete           | SHARED_WRITE                            | 与SHARED_READ、 SHARED_WRITE兼容，与 EXCLUSIVE互斥 |
|          | alter table ...                    | EXCLUSIVE                               | 与其他的MDL都互斥                                  |

注意：

- MDL 是在事务提交后才会释放，这意味着**事务执行期间，MDL 是一直持有的**。
- 申请 MDL 锁的操作会形成一个队列，队列中**写锁获取优先级高于读锁**，一旦出现 MDL 写锁等待，会阻塞后续该表的所有 CRUD 操作。



#### 意向锁

如果没有「意向锁」，那么加「独占表锁」时，就需要遍历表里所有记录，查看是否有记录存在独占锁，这样效率会很慢。

那么有了「意向锁」，由于在对记录加独占锁前，先会加上表级别的意向独占锁，那么在加「独占表锁」时，直接查该表是否有意向独占锁，如果有就意味着表里已经有记录被加了独占锁，这样就不用去遍历表里的记录。



- **意向共享锁和意向独占锁是表级锁，不会和行级的共享锁和独占锁发生冲突，**

- **而且意向锁之间也不会发生冲突，只会和共享表锁（\*lock tables ... read\*）和独占表锁（\*lock tables ... write\*）发生冲突。**

- 在使用 InnoDB 引擎的表里对某些记录加上「共享锁」之前，需要先在表级别加上一个「意向共享锁」；

  与表级共享锁read兼容，与write互斥

- 在使用 InnoDB 引擎的表里对某些纪录加上「独占锁」之前，需要先在表级别加上一个「意向独占锁」；

  与表锁 read、write都互斥。

这是行锁的加锁语句

| SQL                       | 行锁类型   | 说明                                     |
| ------------------------- | ---------- | ---------------------------------------- |
| insert                    | 独占锁     | 自动加锁                                 |
| update                    | 独占锁     | 自动加锁                                 |
| delete                    | 独占锁     | 自动加锁                                 |
| select                    | 不加任何锁 |                                          |
| select lock in share mode | 共享锁     | 需要手动在SELECT之后加LOCK IN SHARE MODE |
| select for update         | 独占锁     | 需要手动在SELECT之后加FOR UPDATE         |



#### auto-inc锁

表里的主键通常都会设置成自增的，这是通过对主键字段声明 `AUTO_INCREMENT` 属性实现的。之后可以在插入数据时，可以不指定主键的值，数据库会自动给主键赋值递增的值，这主要是通过 **AUTO-INC 锁**实现的。防止连续插入的数据自增主键不连续。

- AUTO-INC 锁是特殊的表锁机制，锁**不是再一个事务提交后才释放，而是再执行完插入语句后就会立即释放**。

  **在插入数据时，会加一个表级别的 AUTO-INC 锁**，然后为被 `AUTO_INCREMENT` 修饰的字段赋值递增的值，等插入语句执行完成后，才会把 AUTO-INC 锁释放掉。

  那么，一个事务在持有 AUTO-INC 锁的过程中，其他事务的如果要向该表插入语句都会被阻塞，从而保证插入数据时，被 `AUTO_INCREMENT` 修饰的字段的值是连续递增的。

  但是， AUTO-INC 锁再对大量数据进行插入的时候，一个insert后跟着很多value值，那么此时因为另一个事务中的插入会被阻塞。会影响插入性能。



> 因此， 在 MySQL 5.1.22 版本开始，InnoDB 存储引擎提供了一种**轻量级的锁**来实现自增。
>
> 一样也是在插入数据的时候，会为被 `AUTO_INCREMENT` 修饰的字段加上轻量级锁，**然后给该字段赋值一个自增的值，就把这个轻量级锁释放了，而不需要等待整个插入语句执行完后才释放锁**。
>
> InnoDB 存储引擎提供了个 innodb_autoinc_lock_mode 的系统变量，是用来控制选择用 AUTO-INC 锁，还是轻量级的锁。
>
> - 当 innodb_autoinc_lock_mode = 0，就采用 AUTO-INC 锁，语句执行结束后才释放锁；
> - 当 innodb_autoinc_lock_mode = 2，就采用轻量级锁，申请自增主键后就释放锁，并不需要等语句执行后才释放。
> - 当 innodb_autoinc_lock_mode = 1：
>   - 普通 insert 语句，自增锁在申请之后就马上释放；
>   - 类似 insert … select 这样的批量插入数据的语句，自增锁还是要等语句结束后才被释放；





### 行级锁

InnoDB 引擎是支持行级锁的，而 MyISAM 引擎并不支持行级锁。

普通的select语句是不加任何锁的，使用的是快照读。

共享锁之间共享，共享锁与独占锁互斥，独占锁之间互斥；

| SQL                       | 行锁类型    | 说明                                     |
| ------------------------- | ----------- | ---------------------------------------- |
| insert                    | 独占锁（X） | 自动加锁                                 |
| update                    | 独占锁（X） | 自动加锁                                 |
| delete                    | 独占锁（X） | 自动加锁                                 |
| select                    | 不加任何锁  |                                          |
| select lock in share mode | 共享锁（S） | 需要手动在SELECT之后加LOCK IN SHARE MODE |
| select for update         | 独占锁（X） | 需要手动在SELECT之后加FOR UPDATE         |

行级锁的类型主要有三类：

- Record Lock，记录锁，也就是仅仅把一条记录锁上；
- Gap Lock，间隙锁，锁定一个范围，但是不包含记录本身；
- Next-Key Lock：Record Lock + Gap Lock 的组合，锁定一个范围，并且锁定记录本身

#### Record lock

Record Lock 称为记录锁，锁住的是一条记录。而且记录锁是有 S 锁和 X 锁之分的：

- 当一个事务对一条记录加了 S 型记录锁后，其他事务也可以继续对该记录加 S 型记录锁（S 型与 S 锁兼容），但是不可以对该记录加 X 型记录锁（S 型与 X 锁不兼容）;

- 当一个事务对一条记录加了 X 型记录锁后，其他事务既不可以对该记录加 S 型记录锁（S 型与 X 锁不兼容），也不可以对该记录加 X 型记录锁（X 型与 X 锁不兼容）。

- 当事务执行 commit 后，事务过程中生成的锁都会被释放。

#### Gap Lock

Gap Lock 称为间隙锁，**只存在于可重复读隔离级别**，目的是为了解决可重复读隔离级别下幻读的现象。

**间隙锁之间是兼容的，即两个事务可以同时持有包含共同间隙范围的间隙锁，并不存在互斥关系，因为间隙锁的目的是防止插入幻影记录而提出的**。

#### Next-Key Lock

Next-Key Lock 称为临键锁，是 Record Lock + Gap Lock 的组合，锁定记录本身和之前的间隙。

#### 插入意向锁

一个事务在插入一条记录的时候，需要判断插入位置是否已被其他事务加了间隙锁（next-key lock 也包含间隙锁）。

如果有的话，插入操作就会发生**阻塞**，直到拥有间隙锁的那个事务提交为止（释放间隙锁的时刻），在此期间会生成一个**插入意向锁**，表明有事务想在某个区间插入新记录，但是现在处于等待状态。



## 17、MySQL是怎么加行级锁的

首先明确一点，MySQL中**加锁的对象是索引，加锁的基本单位是 next-key lock**，

它是由记录锁和间隙锁组合而成的，**next-key lock 是前开后闭区间**（意思说是锁的是当前的数据和数据之前的间隙），而**间隙锁是前开后开区间**（只锁住间隙）。

但是，next-key lock 在一些场景下会退化成记录锁或间隙锁。

**在能使用记录锁或者间隙锁就能避免幻读现象的场景下， next-key lock 就会退化成记录锁或间隙锁**。

> 幻读的定义就是，当一个事务前后两次查询的结果集，不相同时，就认为发生幻读。所以，要避免幻读就是**避免结果集某一条记录被其他事务删除**，或者**有其他事务插入了一条新记录**，这样前后两次查询的结果集就不会出现不相同的情况。

#### 唯一索引等值查询

- 当查询的记录是「存在」的，在索引树上定位到这一条记录后，将该记录的索引中的 next-key lock 会**退化成「记录锁」**

  以主键为例

  - 由于主键具有唯一性，所以**其他事务插入 id = 1 的时候，会因为主键冲突，导致无法插入 id = 1 的新记录**。这样事务 A 在多次查询 id = 1 的记录的时候，不会出现前后两次查询的结果集不同，也就避免了幻读的问题。
  - 由于对 id = 1 加了记录锁，**其他事务无法删除该记录**，这样事务 A 在多次查询 id = 1 的记录的时候，不会出现前后两次查询的结果集不同，也就避免了幻读的问题

  所以，唯一索引等值查询却存在时，只使用一个记录锁就行

- 当查询的记录是「不存在」的，在索引树找到第一条大于该查询记录的记录后，将该记录的索引中的 next-key lock 会**退化成「间隙锁」**。

> 如果是用二级索引（不管是不是非唯一索引，还是唯一索引）进行锁定读查询的时候，除了会对二级索引项加行级锁（如果是唯一索引的二级索引，加锁规则和主键索引的案例相同），而且还会对查询到的记录的主键索引项上加「记录锁」。
>
> 能够使用 `select * from performance_schema.data_locks` 查看加锁的状态

| 字段名    | 选项                                  |
| --------- | ------------------------------------- |
| lock      | S/X 行级锁，共享锁和独占锁            |
|           | IS/IX  意向锁，意向共享锁和意向独占锁 |
| lock_mode | X   行级锁，next_key 锁               |
|           | X, REC_NOT_GAP  行级锁，记录锁        |
|           | X,GAP  行级锁  间隙锁                 |

#### 唯一索引范围查询

当唯一索引进行范围查询时，**会对每一个扫描到的索引加 next-key 锁，然后如果遇到下面这些情况，会退化成记录锁或者间隙锁**：

- 情况一：针对「大于等于」的范围查询，因为存在等值查询的条件，那么如果等值查询的记录是存在于表中，那么该记录的索引中的 next-key 锁会**退化成记录锁**。

> 说人话，就是>=执行时，如果等于的边界存在，这里对这个边界加的是记录锁(X, REC_NOT_GAP),对于其他范围内的记录，加的是临键锁（X） 

- 情况二：针对「小于或者小于等于」的范围查询，要看条件值的记录是否存在于表中：

  - 当条件值的记录不在表中，那么不管是「小于」还是「小于等于」条件的范围查询，**扫描到终止范围查询的记录时，该记录的索引的 next-key 锁会退化成间隙锁**，其他扫描到的记录，都是在这些记录的索引上加 next-key 锁。

  - 当条件值的记录在表中，如果是「小于」条件的范围查询，**扫描到终止范围查询的记录时，该记录的索引的 next-key 锁会退化成间隙锁**，其他扫描到的记录，都是在这些记录的索引上加 next-key 锁；如果「小于等于」条件的范围查询，扫描到终止范围查询的记录时，该记录的索引 next-key 锁不会退化成间隙锁。其他扫描到的记录，都是在这些记录的索引上加 next-key 锁。



#### 非唯一索引等值查询

当我们用非唯一索引进行等值查询的时候，**因为存在两个索引，一个是主键索引，一个是非唯一索引（二级索引），所以在加锁时，同时会对这两个索引都加锁，但是对主键索引加锁的时候，只有满足查询条件的记录才会对它们的主键索引加锁**。

- 当查询的记录「存在」时，由于不是唯一索引，所以肯定存在索引值相同的记录，于是**非唯一索引等值查询的过程是一个扫描的过程，直到扫描到第一个不符合条件的二级索引记录就停止扫描，然后在扫描的过程中，对扫描到的二级索引记录加的是 next-key 锁，而对于第一个不符合条件的二级索引记录，该二级索引的 next-key 锁会退化成间隙锁。同时，在符合查询条件的记录的主键索引上加记录锁**。
- 当查询的记录「不存在」时，**扫描到第一条不符合条件的二级索引记录，该二级索引的 next-key 锁会退化成间隙锁。因为不存在满足查询条件的记录，所以不会对主键索引加锁**

> 对于二级索引加入的间隙锁，其实这个间隙是能够插入的，因为二级索引的结构是按照索引字段排序，在相同的二级索引值情况下， 再按主键 id 的顺序存放。
>
> 所以例如 我们在原来的数据上设置间隙锁 二级索引为age，间隙锁范围是 (22, 39)（其中 age=22，id=10；age = 39,id=20)
>
> 那么：
>
> - 当其他事务插入一条 age = 22，id = 3 的记录的时候，
>
>   **该位置的下一条是 id = 10、age = 22 的记录，该记录的二级索引上没有间隙锁，所以这条插入语句可以执行成功**。
>
> - 当其他事务插入一条 age = 22，id = 12
>
>   **该位置的下一条是 id = 20、age = 39 的记录，正好该记录的二级索引上有间隙锁，所以这条插入语句会被阻塞，无法插入成功**。
>
> - 当其他事务插入一条 age = 39，id = 3 的记录的时候
>
>   **该位置的下一条是 id = 20、age = 39 的记录，正好该记录的二级索引上有间隙锁，所以这条插入语句会被阻塞，无法插入成功**。
>
> - 当其他事务插入一条 age = 39，id = 21 的记录的时候
>
>   **该位置的下一条记录不存在，也就没有间隙锁了，所以这条插入语句可以插入成功**。

> 在上一个问题的查看中，实际上 `LOCK_DATA：39，20` 是什么意思？
>
> - LOCK_DATA 第一个数值，也就是 39， 它代表的是 age 值。从前面我们也知道了，LOCK_DATA 第一个数值是 next-key 锁和间隙锁**锁住的范围的右边界值**。
> - LOCK_DATA 第二个数值，也就是 20， 它代表的是主键值（id)。
>
> 之所以 LOCK_DATA 要多显示一个数值（ID值），是因为针对「当某个事务持有非唯一索引的 (22, 39) 间隙锁的时候，其他事务是否可以插入 age = 39 新记录」的问题，还需要考虑插入记录的 id 值。而 **LOCK_DATA 的第二个数值，就是说明在插入 age = 39 新记录时，哪些范围的 id 值是不可以插入的**。
>
> 因此， `LOCK_DATA：39，20` + `LOCK_MODE : X, GAP` 的意思是，事务 A 在 age = 39 记录的二级索引上（INDEX_NAME: index_age ），加了 age 值范围为 (22, 39) 的 X 型间隙锁，**同时针对其他事务插入 age 值为 39 的新记录时，不允许插入的新记录的 id 值小于 20 **。如果插入的新记录的 id 值大于 20，则可以插入成功。



#### 非唯一索引范围查询

**唯一索引范围查询，索引的 next-key lock 不会有退化为间隙锁和记录锁的情况**，也就是非唯一索引进行范围查询时，对二级索引记录加锁都是加 next-key 锁。(因为二级索引是非唯一索引，只加记录锁（记录锁无法防止插入，只能防止删除或者修改）就会导致其他事务插入一条 age = 22 的记录，这样前后两次查询的结果集就不相同了)



#### 没有加索引的查询

普通select不加任何锁，使用的方法是快照读。

**行级锁中提到的update、delete、select ... for update 等具有加锁性质的语句，没有使用索引列作为查询条件，或者查询语句没有走索引查询，导致扫描是全表扫描。那么，所有记录都会被加 next-key 锁，这样就相当于锁住的全表，这时如果其他事务对该表进行增、删、改操作的时候，都会被阻塞**。

> **如何避免这种事故的发生？**
>
> 我们可以将 MySQL 里的 `sql_safe_updates` 参数设置为 1，开启安全更新模式，这样可以预防 update 操作时 where 条件没有带上索引列
>
> 如果 where 条件带上了索引列，但是优化器最终扫描选择的是全表，而不是索引的话，我们可以使用 `force index([index_name])` 可以告诉优化器使用哪个索引，以此避免有几率锁全表带来的隐患。



### 18、**MySQL 记录锁+间隙锁可以防止删除操作而导致的幻读吗？**



在对二级索引进行加锁时，除了会对二级索引加记录锁+间隙锁，还会对符合条件的主键索引加记录锁，因此锁住的相关记录是无法被删除的。



### 19、MySQL死锁

假设这样一个例子，

普通的 select 语句是不会对记录加锁的，因为它是通过 MVCC 的机制实现的快照读

<img src="https://cdn.xiaolincoding.com//mysql/other/54fc00f9f87a60ab7b5ba92d824a892d.png" alt="图片" style="zoom:67%;" />

一个事务要插入订单 1007 ，另外一个事务要插入订单 1008，因为需要对订单做幂等性校验，所以两个事务先要查询该订单是否存在，不存在才插入记录，过程如下：

<img src="https://cdn.xiaolincoding.com//mysql/other/90c1e01d0345de639e3426cea0390e80.png" alt="img" style="zoom:67%;" />





- **此时事务 A 在二级索引（INDEX_NAME : index_order）上加的是 X 型的 next-key 锁，锁范围是`(1006, +∞]`**。
- 而间隙锁是不是互斥的，因此 事务B也获得了next-key 锁，锁范围是`(1006, +∞]`。
- 因为 +∞ 并不是一个真实的记录，所以两个事务是可以同时持有`(1006, +∞]`的

案例中的事务 A 和事务 B 在执行完后 `select ... for update` 语句后都持有范围为`(1006,+∞]`的next-key 锁，而接下来的插入操作为了获取到插入意向锁，都在等待对方事务的间隙锁释放，于是就造成了循环等待，导致死锁。

> 插入意向锁与间隙锁的另一个非常重要的差别是：尽管「插入意向锁」也属于间隙锁，但两个事务却不能在同一时间内，一个拥有间隙锁，另一个拥有该间隙区间内的插入意向锁（当然，插入意向锁如果不在间隙锁区间内则是可以的）。
>
> 每插入一条新记录，都需要看一下待插入记录的下一条记录上是否已经被加了间隙锁，如果已加间隙锁，此时会生成一个插入意向锁，然后锁的状态设置为等待状态（*PS：MySQL 加锁时，是先生成锁结构，然后设置锁的状态，如果锁状态是等待状态，并不是意味着事务成功获取到了锁，只有当锁状态为正常状态时，才代表事务成功获取到了锁*），现象就是 Insert 语句会被阻塞。

#### 如何避免死锁



在数据库层面，有两种策略通过「打破循环等待条件」来解除死锁状态：

- **设置事务等待锁的超时时间**。当一个事务的等待时间超过该值后，就对这个事务进行回滚，于是锁就释放了，另一个事务就可以继续执行了。在 InnoDB 中，参数 `innodb_lock_wait_timeout` 是用来设置超时时间的，默认值时 50 秒。

  当发生超时后，就出现下面这个提示：

![图片](https://cdn.xiaolincoding.com//mysql/other/c296c1889f0101d335699311b4ef20a8.png)

- **开启主动死锁检测**。主动死锁检测在发现死锁后，主动回滚死锁链条中的某一个事务，让其他事务得以继续执行。将参数 `innodb_deadlock_detect` 设置为 on，表示开启这个逻辑，默认就开启。

  当检测到死锁后，就会出现下面这个提示：

![图片](https://cdn.xiaolincoding.com//mysql/other/f380ef357d065498d8d54ad07f145e09.png)



### 20、Insert 语句是怎么加行级锁的？

Insert 语句在正常执行时是不会生成锁结构的，它是靠聚簇索引记录自带的 trx_id 隐藏列来作为**隐式锁**来保护记录的。

当事务需要加锁的时，如果这个锁不可能发生冲突，InnoDB会跳过加锁环节，这种机制称为隐式锁。隐式锁是 InnoDB 实现的一种延迟加锁机制，其特点是只有在可能发生冲突时才加锁，从而减少了锁的数量，提高了系统整体性能。

隐式锁就是在 Insert 过程中不加锁，只有在特殊情况下，才会将隐式锁转换为显示锁，这里我们列举两个场景。

- 如果记录之间加有间隙锁，为了避免幻读，此时是不能插入记录的；

  每插入一条新记录，都需要看一下待插入记录的下一条记录上是否已经被加了间隙锁，如果已加间隙锁，**此时会生成一个插入意向锁，然后锁的状态设置为等待状态**（*PS：**MySQL 加锁时，是先生成锁结构，然后设置锁的状态，如果锁状态是等待状态，并不是意味着事务成功获取到了锁，只有当锁状态为正常状态时，才代表事务成功获取到了锁***），现象就是 Insert 语句会被阻塞。

- 如果 Insert 的记录和已有记录存在唯一键冲突，此时也不能插入记录；

  - 如果主键索引重复，插入新记录的事务会给已存在的主键值重复的聚簇索引记录**添加 S 型记录锁**。
  - 如果唯一二级索引重复，插入新记录的事务都会给已存在的二级索引列值重复的二级索引记录**添加 S 型 next-key 锁**



在隔离级别可重复读的情况下，开启两个事务，前后执行相同的 Insert 语句，此时**事务 B 的 Insert 语句会发生阻塞**。

<img src="https://cdn.xiaolincoding.com/gh/xiaolincoder/mysql/%E9%94%81/%E5%94%AF%E4%B8%80%E7%B4%A2%E5%BC%95%E5%8A%A0%E9%94%81.drawio.png" alt="img" style="zoom:67%;" />

两个事务的加锁过程：

- 事务 A 先插入 order_no 为 1006 的记录，可以插入成功，此时对应的唯一二级索引记录被「隐式锁」保护，此时还没有实际的锁结构（执行完这里的时候，你可以看查 performance_schema.data_locks 信息，可以看到这条记录是没有加任何锁的）；
- 接着，事务 B 也插入 order_no 为 1006 的记录，由于事务 A 已经插入 order_no 值为 1006 的记录，所以事务 B 在插入二级索引记录时会遇到重复的唯一二级索引列值，此时事务 B 想获取一个 S 型 next-key 锁，但是事务 A 并未提交，**事务 A 插入的 order_no 值为 1006 的记录上的「隐式锁」会变「显示锁」且锁类型为 X 型的记录锁，所以事务 B 向获取 S 型 next-key 锁时会遇到锁冲突，事务 B 进入阻塞状态**。



### 21、MySQL日志

MySQL中有以下三种日志:

- **undo log（回滚日志）**：是 Innodb 存储引擎层生成的日志，实现了事务中的**原子性**，主要**用于事务回滚和 MVCC**。
- **redo log（重做日志）**：是 Innodb 存储引擎层生成的日志，实现了事务中的**持久性**，主要**用于掉电等故障恢复**；
- **binlog （归档日志）**：是 Server 层生成的日志，主要**用于数据备份和主从复制**；





#### Undo log

实现这一机制就是 **undo log（回滚日志），它保证了事务的 [ACID 特性 (opens new window)](https://xiaolincoding.com/mysql/transaction/mvcc.html#事务有哪些特性)中的原子性（Atomicity）**。

undo log 是一种用于撤销回退的日志。在事务没提交之前，MySQL 会先记录更新前的数据到 undo log 日志文件里面，当事务回滚时，可以利用 undo log 来进行回滚。

每当 InnoDB 引擎对一条记录进行操作（修改、删除、新增）时，要把回滚时需要的信息都记录到 undo log 里，比如：

- 在**插入**一条记录时，要把这条记录的主键值记下来，这样之后回滚时只需要把这个主键值对应的记录**删掉**就好了；
- 在**删除**一条记录时，要把这条记录中的内容都记下来，这样之后回滚时再把由这些内容组成的记录**插入**到表中就好了；
- 在**更新**一条记录时，要把被更新的列的旧值记下来，这样之后回滚时再把这些列**更新为旧值**就好了。

> 在发生回滚时，就读取 undo log 里的数据，然后做原先相反操作。比如当 delete 一条记录时，undo log 中会把记录中的内容都记下来，然后执行回滚操作的时候，就读取 undo log 里的数据，然后进行 insert 操作。

另外，**undo log 还有一个作用，通过 ReadView + undo log 实现 MVCC（多版本并发控制）**。

一条记录的每一次更新操作产生的 undo log 格式都有一个 roll_pointer 指针和一个 trx_id 事务id：

- 通过 trx_id 可以知道该记录是被哪个事务修改的；
- 通过 roll_pointer 指针可以将这些 undo log 串成一个链表，这个链表就被称为版本链；

<img src="https://cdn.xiaolincoding.com/gh/xiaolincoder/mysql/how_update/%E7%89%88%E6%9C%AC%E9%93%BE.png?image_process=watermark,text_5YWs5LyX5Y-377ya5bCP5p6XY29kaW5n,type_ZnpsdHpoaw,x_10,y_10,g_se,size_20,color_0000CD,t_70,fill_0" alt="版本链" style="zoom:67%;" />



所以总结来说，undo log有两大作用，

- **实现事务回滚，保障事务的原子性**。事务处理过程中，如果出现了错误或者用户执 行了 ROLLBACK 语句，MySQL 可以利用 undo log 中的历史数据将数据恢复到事务开始之前的状态。
- **实现 MVCC（多版本并发控制）关键因素之一**。MVCC 是通过 ReadView + undo log 实现的。undo log 为每条记录保存多份历史数据，MySQL 在执行快照读（普通 select 语句）的时候，会根据事务的 Read View 里的信息，顺着 undo log 的版本链找到满足其可见性的记录。



> 同时，这里提一下，undo log 和数据页的刷盘策略是一样的，都需要通过 redo log 保证持久化。
>
> 在buffer pool中存在 undo 页，对 undo 页的修改也都会记录到 redo log。redo log 会每秒刷盘，提交事务时也会刷盘，数据页和 undo 页都是靠这个机制保证持久化的。



#### redo log

Buffer Pool 是提高了读写效率，但是Buffer Pool 是基于内存的，而内存总是不可靠，万一断电重启，还没来得及落盘的脏页数据就会丢失。

为了防止断电导致数据丢失的问题，当有一条记录需要更新的时候，InnoDB 引擎就会先更新内存（同时标记为脏页），然后将本次对这个页的修改以 redo log 的形式记录下来，**这个时候更新就算完成了**。

后续，InnoDB 引擎会在适当的时候，由后台线程将缓存在 Buffer Pool 的脏页刷新到磁盘里，这就是 **WAL （Write-Ahead Logging）技术**。

<img src="https://cdn.xiaolincoding.com/gh/xiaolincoder/mysql/how_update/wal.png?image_process=watermark,text_5YWs5LyX5Y-377ya5bCP5p6XY29kaW5n,type_ZnpsdHpoaw,x_10,y_10,g_se,size_20,color_0000CD,t_70,fill_0" alt="img" style="zoom: 50%;" />

- redo log 是物理日志，记录了某个数据页做了什么修改。

- 在事务提交时，只要先将 redo log 持久化到磁盘即可，可以不需要等到将缓存在 Buffer Pool 里的脏页数据持久化到磁盘。

  当系统崩溃时，虽然脏页数据没有持久化，但是 redo log 已经持久化，接着 MySQL 重启后，可以根据 redo log 的内容，将所有数据恢复到最新的状态。

- Undo 页面被修改后，也要记录对应redo log

  开启事务后，InnoDB 层更新记录前，首先要记录相应的 undo log，如果是更新操作，需要把被更新的列的旧值记下来，也就是要生成一条 undo log，undo log 会写入 Buffer Pool 中的 Undo 页面。

  不过，**在内存修改该 Undo 页面后，需要记录对应的 redo log**。

  > - redo log 记录了此次事务「**完成后**」的数据状态，记录的是更新**之后**的值；
  > - undo log 记录了此次事务「**开始前**」的数据状态，记录的是更新**之前**的值；
  >
  > **事务提交之前发生了崩溃，重启后会通过 undo log 回滚事务，事务提交之后发生了崩溃，重启后会通过 redo log 恢复事务**
  >
  > 所以有了 redo log，再通过 WAL 技术，InnoDB 就可以保证即使数据库发生异常重启，之前已提交的记录都不会丢失，这个能力称为 **crash-safe**（崩溃恢复）。可以看出来， **redo log 保证了事务四大特性中的持久性**。

- 写入 redo log 的方式使用了追加操作， 所以磁盘操作是**顺序写**，而写入数据需要先找到写入位置，然后才写到磁盘，所以磁盘操作是**随机写**。

  磁盘的「顺序写 」比「随机写」 高效的多，因此 redo log 写入磁盘的开销更小。

- 产生的 redo log也不是直接写入磁盘的，

  redo log 也有自己的缓存—— **redo log buffer**，每当产生一条 redo log 时，会先写入到 redo log buffer，后续在持久化到磁盘。但是肯定要比数据的持久化到磁盘更加勤快。

  主要有下面几个时机：

  - MySQL 正常关闭时；

  - 当 redo log buffer 中记录的写入量大于 redo log buffer 内存空间的一半时，会触发落盘；

  - InnoDB 的后台线程每隔 1 秒，将 redo log buffer 持久化到磁盘。

  - 每次事务提交时都将缓存在 redo log buffer 里的 redo log 直接持久化到磁盘（这个策略可由 innodb_flush_log_at_trx_commit 参数控制）。

    > 1、redo log buffer 默认大小 16 MB，可以通过 `innodb_log_Buffer_size` 参数动态的调整大小，增大它的大小可以让 MySQL 处理「大事务」是不必写入磁盘，进而提升写 IO 性能
    >
    > 2、innodb_flush_log_at_trx_commit 参数
    >
    > <img src="https://cdn.xiaolincoding.com/gh/xiaolincoder/mysql/how_update/innodb_flush_log_at_trx_commit.drawio.png?image_process=watermark,text_5YWs5LyX5Y-377ya5bCP5p6XY29kaW5n,type_ZnpsdHpoaw,x_10,y_10,g_se,size_20,color_0000CD,t_70,fill_0" alt="img" style="zoom: 67%;" />
    >
    > 3、加入后台线程后，后台线程以1s的间隔将redo log buffer中的数据刷新到磁盘中
    >
    > <img src="https://cdn.xiaolincoding.com/gh/xiaolincoder/mysql/how_update/innodb_flush_log_at_trx_commit2.drawio.png?image_process=watermark,text_5YWs5LyX5Y-377ya5bCP5p6XY29kaW5n,type_ZnpsdHpoaw,x_10,y_10,g_se,size_20,color_0000CD,t_70,fill_0" alt="img" style="zoom:67%;" />
    >
    > 所以在后台线程的情况下；
    >
    > 参数为0时，若是MySQL系统崩溃，那么会导致 前一秒中所有的事务数据丢失
    >
    > 参数为2时，安全性会比参数为1时更高，只有操作系统崩溃时才会丢失上一秒钟所有的事务数据。

- 默认情况下， InnoDB 存储引擎有 1 个重做日志文件组( redo log Group），「重做日志文件组」由有 2 个 redo log 文件组成，这两个 redo 日志的文件名叫 ：`ib_logfile0` 和 `ib_logfile1` 。

  在重做日志组中，每个 redo log File 的大小是固定且一致的，假设每个 redo log File 设置的上限是 1 GB，那么总共就可以记录 2GB 的操作。

  重做日志文件组是以**循环写**的方式工作的，从头开始写，写到末尾就又回到开头，相当于一个环形。

  

所以综合来说，redo log 

- **实现事务的持久性，让 MySQL 有 crash-safe 的能力**，能够保证 MySQL 在任何时间段突然崩溃，重启后之前已提交的记录都不会丢失；
- **将写操作从「随机写」变成了「顺序写」**，提升 MySQL 写入磁盘的性能。





#### bin log

undo log 和 redo log 这两个日志都是 Innodb 存储引擎生成的。

MySQL 在完成一条更新操作后，Server 层还会生成一条 binlog，等之后事务提交的时候，会将该事物执行过程中产生的所有 binlog 统一写 入 binlog 文件。

binlog 文件是记录了**所有数据库表结构变更和表数据修改的日志**，不会记录查询类的操作，比如 SELECT 和 SHOW 操作。

##### redo log 和 binlog 有什么区别？

这两个日志有四个区别。

*1、适用对象不同：*

- binlog 是 MySQL 的 Server 层实现的日志，所有存储引擎都可以使用；
- redo log 是 Innodb 存储引擎实现的日志；

*2、文件格式不同：*

- binlog 有 3 种格式类型，分别是 STATEMENT（默认格式）、ROW、 MIXED，区别如下：
  - STATEMENT：每一条修改数据的 SQL 都会被记录到 binlog 中（相当于记录了逻辑操作，所以针对这种格式， binlog 可以称为逻辑日志），主从复制中 slave 端再根据 SQL 语句重现。但 STATEMENT 有动态函数的问题，比如你用了 uuid 或者 now 这些函数，你在主库上执行的结果并不是你在从库执行的结果，这种随时在变的函数会导致复制的数据不一致；
  - ROW：记录行数据最终被修改成什么样了（这种格式的日志，就不能称为逻辑日志了），不会出现 STATEMENT 下动态函数的问题。但 ROW 的缺点是每行数据的变化结果都会被记录，比如执行批量 update 语句，更新多少行数据就会产生多少条记录，使 binlog 文件过大，而在 STATEMENT 格式下只会记录一个 update 语句而已；
  - MIXED：包含了 STATEMENT 和 ROW 模式，它会根据不同的情况自动使用 ROW 模式和 STATEMENT 模式；
- redo log 是物理日志，记录的是在某个数据页做了什么修改，比如对 XXX 表空间中的 YYY 数据页 ZZZ 偏移量的地方做了AAA 更新；

*3、写入方式不同：*

- binlog 是追加写，写满一个文件，就创建一个新的文件继续写，不会覆盖以前的日志，保存的是全量的日志。
- redo log 是循环写，日志空间大小是固定，全部写满就从头开始，保存未被刷入磁盘的脏页日志。

*4、用途不同：*

- binlog 用于备份恢复、主从复制；
- redo log 用于掉电等故障恢复。
- redo log 可以在事务没提交之前持久化到磁盘，但是 binlog 必须在事务提交之后，才可以持久化到磁盘。

> 所以如果将数据库的数据被删除了
>
> 因为redo log是采用循环写的方法，对于已经持久化到磁盘的数据，redo log就删除了，
>
> 而binlog 文件保存的是全量的日志，也就是保存了所有数据变更的情况，**理论上**只要记录在 binlog 上的数据，都可以恢复，所以如果不小心整个数据库的数据被删除了，得用 binlog 文件恢复数据。



##### 主从复制

MySQL 的主从复制依赖于 binlog ，复制的过程就是将 binlog 中的数据从主库传输到从库上。

这个过程一般是**异步**的，也就是主库上执行事务操作的线程不会等待复制 binlog 的线程同步完成。

MySQL 集群的主从复制过程梳理成 3 个阶段：

- MySQL 主库在收到客户端提交事务的请求之后，会先写入 binlog，再提交事务，更新存储引擎中的数据，事务提交完成后，返回给客户端“操作成功”的响应。
- 从库会创建一个专门的 I/O 线程，连接主库的 log dump 线程，来接收主库的 binlog 日志，再把 binlog 信息写入 relay log 的中继日志里，再返回给主库“复制成功”的响应。
- 从库会创建一个用于回放 binlog 的线程，去读 relay log 中继日志，然后回放 binlog 更新存储引擎中的数据，最终实现主从的数据一致性。

**在完成主从复制之后，你就可以在写数据时只写主库，在读数据时只读从库，这样即使写请求会锁表或者锁记录，也不会影响读请求的执行。**

> MySQL 主从复制还有哪些模型？
>
> 主要有三种：
>
> - **同步复制**：MySQL 主库提交事务的线程要等待所有从库的复制成功响应，才返回客户端结果。这种方式在实际项目中，基本上没法用，原因有两个：一是性能很差，因为要复制到所有节点才返回响应；二是可用性也很差，主库和所有从库任何一个数据库出问题，都会影响业务。
> - **异步复制**（默认模型）：MySQL 主库提交事务的线程并不会等待 binlog 同步到各从库，就返回客户端结果。这种模式一旦主库宕机，数据就会发生丢失。
> - **半同步复制**：MySQL 5.7 版本之后增加的一种复制方式，介于两者之间，事务线程不用等待所有的从库复制成功响应，只要一部分复制成功响应回来就行，比如一主二从的集群，只要数据成功复制到任意一个从库上，主库的事务线程就可以返回给客户端。这种**半同步复制的方式，兼顾了异步复制和同步复制的优点，即使出现主库宕机，至少还有一个从库有最新的数据，不存在数据丢失的风险**。

##### binlog 什么时候刷盘？

事务执行过程中，先把日志写到 binlog cache（Server 层的 cache），事务提交的时候，再把 binlog cache 写到 binlog 文件中。

一个事务的 binlog 是不能被拆开的，因此无论这个事务有多大（比如有很多条语句），也要保证一次性写入。这是因为有一个线程只能同时有一个事务在执行的设定，所以每当执行一个 begin/start transaction 的时候，就会默认提交上一个事务，这样如果一个事务的 binlog 被拆开的时候，在备库执行就会被当做多个事务分段自行，这样破坏了原子性，是有问题的。

所以在事务提交的时候，执行器把 binlog cache 里的完整事务写入到 binlog 文件中，并清空 binlog cache。



> 具体更新一条记录 `UPDATE t_user SET name = 'xiaolin' WHERE id = 1;` 的流程如下:
>
> 1. 执行器负责具体执行，会调用存储引擎的接口，通过主键索引树搜索获取 id = 1 这一行记录：
>    - 如果 id=1 这一行所在的数据页本来就在 buffer pool 中，就直接返回给执行器更新；
>    - 如果记录不在 buffer pool，将数据页从磁盘读入到 buffer pool，返回记录给执行器。
> 2. 执行器得到聚簇索引记录后，会看一下更新前的记录和更新后的记录是否一样：
>    - 如果一样的话就不进行后续更新流程；
>    - 如果不一样的话就把更新前的记录和更新后的记录都当作参数传给 InnoDB 层，让 InnoDB 真正的执行更新记录的操作；
> 3. 开启事务， InnoDB 层更新记录前，首先要记录相应的 undo log，因为这是更新操作，需要把被更新的列的旧值记下来，也就是要生成一条 undo log，undo log 会写入 Buffer Pool 中的 Undo 页面，不过在内存修改该 Undo 页面后，需要记录对应的 redo log。
> 4. InnoDB 层开始更新记录，会先更新内存（同时标记为脏页），然后将记录写到 redo log 里面，这个时候更新就算完成了。为了减少磁盘I/O，不会立即将脏页写入磁盘，后续由后台线程选择一个合适的时机将脏页写入到磁盘。这就是 **WAL 技术**，MySQL 的写操作并不是立刻写到磁盘上，而是先写 redo 日志，然后在合适的时间再将修改的行数据写到磁盘上。
> 5. 至此，一条记录更新完了。
> 6. 在一条更新语句执行完成后，然后开始记录该语句对应的 binlog，此时记录的 binlog 会被保存到 binlog cache，并没有刷新到硬盘上的 binlog 文件，在事务提交时才会统一将该事务运行过程中的所有 binlog 刷新到硬盘。
> 7. 两阶段提交。
>    - **prepare 阶段**：将 redo log 对应的事务状态设置为 prepare，然后将 redo log 刷新到硬盘；
>    - **commit 阶段**：将 binlog 刷新到磁盘，接着调用引擎的提交事务接口，将 redo log 状态设置为 commit（将事务设置为 commit 状态后，刷入到磁盘 redo log 文件）；



#### 两阶段提交（这部分没有看完——

事务提交后，redo log 和 binlog 都要持久化到磁盘，但是这两个是独立的逻辑，可能出现半成功的状态，这样就造成两份日志之间的逻辑不一致。

> 举个例子，假设 id = 1 这行数据的字段 name 的值原本是 'jay'，然后执行 `UPDATE t_user SET name = 'xiaolin' WHERE id = 1;` 如果在持久化 redo log 和 binlog 两个日志的过程中，出现了半成功状态，那么就有两种情况：
>
> - **如果在将 redo log 刷入到磁盘之后， MySQL 突然宕机了，而 binlog 还没有来得及写入**。那么此时，MySQL重启后能够通过redo log恢复事务提交的数据，但是**binlog**中的数据是不完整的，在主从架构中，binlog 会被复制到从库，由于 binlog 丢失了这条更新语句，从库的这一行 name 字段是旧值 jay，与主库的值不一致性；
>
> - **如果在将 binlog 刷入到磁盘之后， MySQL 突然宕机了，而 redo log 还没有来得及写入**。
>
>   同理，那么MySQL重启后是旧值，但是复制到从库值是新值，那么主库从库的值还是不匹配。

因此，**MySQL 为了避免出现两份日志之间的逻辑不一致的问题，使用了「两阶段提交」来解决**，

**两阶段提交把单个事务的提交拆分成了 2 个阶段，分别是「准备（Prepare）阶段」和「提交（Commit）阶段」**（注意，不要把提交（Commit）阶段和 commit 语句混淆了，commit 语句执行的时候，会包含提交（Commit）阶段。）

过程：当客户端执行 commit 语句或者在自动提交的情况下，MySQL 内部开启一个 XA 事务，**分两阶段来完成 XA 事务的提交**，如下图：

<img src="https://cdn.xiaolincoding.com/gh/xiaolincoder/mysql/how_update/%E4%B8%A4%E9%98%B6%E6%AE%B5%E6%8F%90%E4%BA%A4.drawio.png?image_process=watermark,text_5YWs5LyX5Y-377ya5bCP5p6XY29kaW5n,type_ZnpsdHpoaw,x_10,y_10,g_se,size_20,color_0000CD,t_70,fill_0" alt="两阶段提交" style="zoom: 50%;" />



从图中可看出，事务的提交过程有两个阶段，就是**将 redo log 的写入拆成了两个步骤：prepare 和 commit，中间再穿插写入binlog**，具体如下：

- **prepare 阶段**：将 XID（内部 XA 事务的 ID） 写入到 redo log，同时将 redo log 对应的事务状态设置为 prepare，然后将 redo log 持久化到磁盘（innodb_flush_log_at_trx_commit = 1 的作用）；
- **commit 阶段**：把 XID 写入到 binlog，然后将 binlog 持久化到磁盘（sync_binlog = 1 的作用），接着调用引擎的提交事务接口，将 redo log 状态设置为 commit，此时该状态并不需要持久化到磁盘，只需要 write 到文件系统的 page cache 中就够了，因为只要 binlog 写磁盘成功，就算 redo log 的状态还是 prepare 也没有关系，一样会被认为事务已经执行成功；



### 22、Buffer pool

MySQL 的数据都是存在磁盘中的，那么我们要更新一条记录的时候，得先要从磁盘读取该记录，然后在内存中修改这条记录。那修改完这条记录是选择直接写回到磁盘，还是选择缓存起来呢？

当然是缓存起来好，这样下次有查询语句命中了这条记录，直接读取缓存中的记录，就不需要从磁盘获取数据了。

为此，Innodb 存储引擎设计了一个**缓冲池（Buffer Pool）**，来提高数据库的读写性能。

有了 Buffer Pool 后：

- 当读取数据时，如果数据存在于 Buffer Pool 中，客户端就会直接读取 Buffer Pool 中的数据，否则再去磁盘中读取。
- 当修改数据时，如果数据存在于 Buffer Pool 中，那直接修改 Buffer Pool 中数据所在的页，然后将其页设置为脏页（该页的内存数据和磁盘上的数据已经不一致），为了减少磁盘I/O，不会立即将脏页写入磁盘，后续由后台线程选择一个合适的时机将脏页写入到磁盘

> buffer pool是一片完整的空间，默认情况下，uffer Pool 只有 `128MB` 。
>
> 可以通过调整 `innodb_buffer_pool_size` 参数来设置 Buffer Pool 的大小



在 MySQL 启动的时候，**InnoDB 会为 Buffer Pool 申请一片连续的内存空间，然后按照默认的`16KB`的大小划分出一个个的页， Buffer Pool 中的页就叫做缓存页**。此时这些缓存页都是空闲的，之后随着程序的运行，才会有磁盘上的页被缓存到 Buffer Pool 中。所以，MySQL 刚启动的时候，你会观察到使用的虚拟内存空间很大，而使用到的物理内存空间却很小，这是因为只有这些虚拟内存被访问后，操作系统才会触发缺页中断，申请物理内存，接着将虚拟地址和物理地址建立映射关系。

Buffer Pool 除了缓存「索引页」和「数据页」，还包括了 Undo 页，插入缓存、自适应哈希索引、锁信息等等。

为了更好的管理这些在 Buffer Pool 中的缓存页，InnoDB 为每一个缓存页都创建了一个**控制块**，控制块信息包括「缓存页的表空间、页号、缓存页地址、链表节点」等等。

控制块也是占有内存空间的，它是放在 Buffer Pool 的最前面，接着才是缓存页，如下图：

<img src="https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost4@main/mysql/innodb/%E7%BC%93%E5%AD%98%E9%A1%B5.drawio.png" alt="img" style="zoom:67%;" />



##### 如何管理空闲页？

Buffer Pool 是一片连续的内存空间，当 MySQL 运行一段时间后，这片连续的内存空间中的缓存页既有空闲的，也有被使用的。

所以，为了能够快速找到空闲的缓存页，可以使用链表结构，将空闲缓存页的「控制块」作为链表的节点，这个链表称为 **Free 链表**（空闲链表）。

<img src="https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost4@main/mysql/innodb/freelist.drawio.png" alt="img" style="zoom: 50%;" />

Free 链表上除了有控制块，还有一个头节点，该头节点包含链表的头节点地址，尾节点地址，以及当前链表中节点的数量等信息。

Free 链表节点是一个一个的控制块，而每个控制块包含着对应缓存页的地址，所以相当于 Free 链表节点都对应一个空闲的缓存页。

##### 如何管理脏页？

buffer pool为了提高磁盘IO效率，并不是立即将修改后的页刷新到磁盘，而是将其标记为**脏页**，然后再由后台线程将脏页写入到磁盘。

那为了能快速知道哪些缓存页是脏的，于是就设计出 **Flush 链表**，它跟 Free 链表类似的，链表的节点也是控制块，区别在于 Flush 链表的元素都是脏页。

<img src="https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost4@main/mysql/innodb/Flush.drawio.png" alt="img" style="zoom:50%;" />

因为存在redo log，所有buffer pool中的脏页不急着立即将脏页刷新至磁盘中。

但下面几种情况会触发脏页的刷新：

- 当 redo log 日志满了的情况下，会主动触发脏页刷新到磁盘；
- Buffer Pool 空间不足时，需要将一部分数据页淘汰掉，如果淘汰的是脏页，需要先将脏页同步到磁盘；
- MySQL 认为空闲时，后台线程会定期将适量的脏页刷入到磁盘；
- MySQL 正常关闭之前，会把所有的脏页刷入到磁盘；



##### 如何提高缓存命中率？

Buffer Pool 的大小是有限的，对于一些频繁访问的数据我们希望可以一直留在 Buffer Pool 中，而一些很少访问的数据希望可以在某些时机可以淘汰掉，从而保证 Buffer Pool 不会因为满了而导致无法再缓存新的数据，同时还能保证常用数据留在 Buffer Pool 中。

要实现这个，最容易想到的就是 LRU（Least recently used）算法。

该算法的思路是，链表头部的节点是最近使用的，而链表末尾的节点是最久没被使用的。那么，当空间不够了，就淘汰最久没被使用的节点，从而腾出空间。

简单的 LRU 算法的实现思路是这样的：

- 当访问的页在 Buffer Pool 里，就直接把该页对应的 LRU 链表节点移动到链表的头部。
- 当访问的页不在 Buffer Pool 里，除了要把页放入到 LRU 链表的头部，还要淘汰 LRU 链表末尾的节点。



但是，简单的 LRU 算法并没有被 MySQL 使用，因为简单的 LRU 算法无法避免下面这两个问题：

- 预读失效；

  MySQL 的预读机制，在加载数据页时，会提前把它相邻的数据页一并加载进来，目的是为了减少磁盘 IO。但是可能这些**被提前加载进来的数据页，并没有被访问**，相当于这个预读是白做了，这个就是**预读失效**。

  如果这些预读页如果一直不会被访问到，就会出现一个很奇怪的问题，不会被访问的预读页却占用了 LRU 链表前排的位置，而末尾淘汰的页，可能是频繁访问的页，这样就大大降低了缓存命中率。

- Buffer Pool 污染；

  当某一个 SQL 语句**扫描了大量的数据**时，在 Buffer Pool 空间比较有限的情况下，可能会将 **Buffer Pool 里的所有页都替换出去，导致大量热数据被淘汰了**，等这些热数据又被再次访问的时候，由于缓存未命中，就会产生大量的磁盘 IO，MySQL 性能就会急剧下降，这个过程被称为 **Buffer Pool 污染**。

  比如，在一个数据量非常大的表，执行了这条语句：

  ```sql
  select * from t_user where name like "%xiaolin%";
  ```

  可能这个查询出来的结果就几条记录，但是由于这条语句会发生索引失效，所以这个查询过程是全表扫描的，接着会发生如下的过程：

  - 从磁盘读到的页加入到 LRU 链表的 old 区域头部；
  - 当从页里读取行记录时，也就是页被访问的时候，就要将该页放到 young 区域头部；
  - 接下来拿行记录的 name 字段和字符串 xiaolin 进行模糊匹配，如果符合条件，就加入到结果集里；
  - 如此往复，直到扫描完表中的所有记录。

  经过这一番折腾，原本 young 区域的热点数据都会被替换掉。

  

MySQL改进的LRU算法

- 将 LRU 划分了 2 个区域：**old 区域 和 young 区域**。young 区域在 LRU 链表的前半部分，old 区域则是在后半部分，如下图：

  <img src="https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost4@main/mysql/innodb/young%2Bold.png" alt="img" style="zoom:67%;" />

  **划分这两个区域后，预读的页就只需要加入到 old 区域的头部，当页被真正访问的时候，才将页插入 young 区域的头部**。如果预读的页一直没有被访问，就会从 old 区域移除，这样就不会影响 young 区域中的热点数据。若old区域中的被加载了，那么会将其添加到young区域，young 区域末尾的页（7号），会被挤到 old 区域。

- 针对buffer pool污染

  很多缓冲页其实只会被访问一次，但是它却只因为被访问了一次而进入到 young 区域，从而导致热点数据被替换了。

  MySQL 是这样做的，进入到 young 区域条件增加了一个**停留在 old 区域的时间判断**。

  - 如果后续的访问时间与第一次访问的时间**在某个时间间隔内**，那么**该缓存页就不会被从 old 区域移动到 young 区域的头部**；
  - 如果后续的访问时间与第一次访问的时间**不在某个时间间隔内**，那么**该缓存页移动到 young 区域的头部**；

  这个间隔时间是由 `innodb_old_blocks_time` 控制的，默认是 1000 ms。

  也就说，**只有同时满足「被访问」与「在 old 区域停留时间超过 1 秒」两个条件，才会被插入到 young 区域头部**，这样就解决了 Buffer Pool 污染的问题 。

  

所以总结来说，

Innodb 通过三种链表来管理缓页：

- Free List （空闲页链表），管理空闲页；
- Flush List （脏页链表），管理脏页；
- LRU List，管理脏页+干净页，将最近且经常查询的数据缓存在其中，而不常查询的数据就淘汰出去。

nnoDB 对 LRU 做了一些优化，我们熟悉的 LRU 算法通常是将最近查询的数据放到 LRU 链表的头部，而 InnoDB 做 2 点优化：

- 将 LRU 链表 分为**young 和 old 两个区域**，加入缓冲池的页，优先插入 old 区域；页被访问时，才进入 young 区域，目的是为了解决预读失效的问题。
- 当**「页被访问」且「 old 区域停留时间超过 `innodb_old_blocks_time` 阈值（默认为1秒）」**时，才会将页插入到 young 区域，否则还是插入到 old 区域，目的是为了解决批量数据访问，大量热数据淘汰的问题。

