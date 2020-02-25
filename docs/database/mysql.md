MYSQL
===

* [Sử dụng connection pool để tăng hiệu suất](http://go-database-sql.org/connection-pool.html) 

```gotemplate
db, err := sql.Open("mysql", "username:password@tcp("127.0.0.1:3306)/hello")
if err != nil {
    log.Fatal(err)
}

defer db.Close()

// mặc định là 0 (không giới hạn) 
// nếu giá trị truyền vào <= 0 cũng là không giới hạn 
db.SetMaxOpenConn(16)

// mặc định là 2 connection rảnh 
// nếu giá trị truyền vào <= 0 là không có connection rảnh nào được giữ lại 
db.setMaxIdleConns(10)

// Mặc định các connection không bị đóng 
// Nếu giá trị truyền vào <= o là các connection sẽ được sử dụng lại mãi và không bị đóng
db.SetConnMaxLifetime(10 * time.Second)
```


