# MySQL

操作数据库的方法：

- [使用第三方驱动，用官方database/sql包进行开发]( https://github.com/ffhelicopter/Go42/blob/master/content/42_39_mysql.md)
- 使用第三方库开发
  - sqlx
  - xorm



如无特殊说明，示例代码均基于表：

```mysql
CREATE TABLE person (
	user_id int primary key auto_increment,
    username varchar(260),
    sex varchar(260),
    email varchar(260)
);
```



## ORM

### 简介

**概念**

ORM，即 Object-Relational-Mapping，对象关系映射。它的作用是在关系型数据库和对象之间作一个映射，这样，我们在具体的操作数据库的时候，就不需要再去和复杂的SQL语句打交道，只要像平时操作对象一样操作它就可以了。

**ORM 把数据库映射成对象**

```
数据库的表（table） --> 类（class）
记录（record，行数据） --> 对象（object）
字段（field） --> 对象的属性（attribute）
```

举例来说，下面是一行 SQL 语句：

```sql
SELECT id, first_name, last_name, phone, birth_date, sex
FROM persons WHERE id = 10;
```

改成 ORM 的写法：

```go
person := Person{}
person.Get(10)
name := person.FirstName
```

一比较就可以发现，ORM 使用对象，封装了数据库操作，因此可以不碰 SQL 语言。开发者只使用面向对象编程，与数据对象直接交互，不用关心底层数据库。

总结起来，ORM 有下面这些**优点**：

- 数据模型都在一个地方定义，更容易更新和维护，也利于重用代码。
- ORM 有现成的工具，很多功能都可以自动完成，比如数据消毒、预处理、事务等等。
- 它迫使你使用 MVC 架构，ORM 就是天然的 Model，最终使代码更清晰。
- 基于 ORM 的业务代码比较简单，代码量少，语义性好，容易理解。
- 你不必编写性能不佳的 SQL。

但是，ORM 也有很突出的**缺点**：

- ORM 库不是轻量级工具，需要花很多精力学习和设置。
- 对于复杂的查询，ORM 要么是无法表达，要么是性能不如原生的 SQL。
- ORM 抽象掉了数据库层，开发者无法了解底层的数据库操作，也无法定制一些特殊的 SQL。

