# fire-admin-boilerplate

cd backend

## 初始化 go 项目

go mod init fire-admin-boilerplate

## 创建 main.go

```
package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
```

## 执行 go mod tidy 加载模块

## 执行 go run main.go

访问 localhost:8080/ping

## json 返回

```
r.GET("json", func(c *gin.Context) {
		// 方法1：使用map
		data := map[string]interface{}{
			"data": map[string]interface{}{
				"name":    "galaxy",
				"message": "hello world",
				"age":     18,
			},
			"msg":  "返回成功",
			"code": 200,
		}

		// 方法2: gin.H
		data := gin.H{
			"data": gin.H{
				"name":    "galaxy",
				"message": "hello world",
				"age":     18,
			},
			"msg":  "返回成功",
			"code": 200,
		}

	// 方法3: 结构体
		type User struct {
			Name string `json:"name"`
			Age  int    `json:"age"`
		}

		data := &User{
			Name: "galaxy",
			Age:  18,
		}

		c.JSON(http.StatusOK, data)
	})
```

## 结构体指针

```
func main(){
    type Person struct {
        name  string
        age int
    }
    p1 := &Person{name: "wbw", age: 18}
    fmt.Println((*p1).name)
    fmt.Println(p1.name) //隐式解引用

    p2 := Person{name: "wbw", age: 18}
    fmt.Println(p2.name)

    p3 := p1
    fmt.Println(p3.name)
    p4 := p2
    fmt.Println(p4.name)

    p3.name = "sdfsd"
    fmt.Println((*p1).name)
    fmt.Println(p1.name)
    fmt.Println(p3.name)

    p4.name = "sdfsdss"
    fmt.Println(p2.name)
    fmt.Println(p4.name)
}

当结构体原型（p1）获取的是结构体指针，那么当它给另一个成员赋值（p3）时，p3的改动会导致p1同时改动。就是说此时p1, p3共同指向一个结构体地址
当结构体原型（p2）获取的只是普通结构体时，那么当它给另一个成员赋值（p4）时，p4的改动不会导致p2的改动。就是说此时p2，p4不是指向同一个结构体地址
当结构体原型（p1）获取的是结构体指针，那么它可以用指针获取成员变量（(*p1).name），也可以使用隐式解引用（p1.name即可获取成员变量）。
```

## 获取参数
### 获取查询参数(Query string parameters)
```
func ping(c *gin.Context) {
	firstname := c.DefaultQuery("firstname", "Guest")
	lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")
	c.JSON(200, gin.H{
		"message":   "pong",
		"firstname": firstname,
		"lastname":  lastname,
	})
}

### 访问
http://localhost:8080/ping?firstname=galaxy&&lastname=guo

```
### 获取路径参数(Parameters in path)
```
func ping(c *gin.Context) {
	id := c.Param("id")
	firstname := c.DefaultQuery("firstname", "Guest")
	lastname := c.Query("lastname") 
	c.JSON(200, gin.H{
		"code":    200,
		"message": "pong",
		"data": gin.H{
			"id":        id,
			"firstname": firstname,
			"lastname":  lastname,
		},
	})
}
```

## 数据库链接
```
go get -u gorm.io/gorm

sqlite数据库
go get -u gorm.io/driver/sqlite
```

## 引入gorm.Model
```
type User struct {
	gorm.Model
	Name string `json:"name"`
	Age  int    `json:"age"`
}

main()中加入
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
```

## 全局数据库
```
// 全局DB,注意不能用 DB,err := (这里为局部变量)
	DB, _ = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
```

## 新增用户
```
func ping(c *gin.Context) {
	id := c.Param("id")
	firstname := c.DefaultQuery("firstname", "Guest")
	lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")
	user := User{ID: id, FirstName: firstname, Age: 18, LastName: lastname}
	DB.Create(&user)

	c.JSON(200, gin.H{
		"code":    200,
		"message": "pong",
		"data": gin.H{
			"id":        id,
			"firstname": firstname,
			"lastname":  lastname,
		},
	})
}
```

## 绑定body,json
```
func pingCreate(c *gin.Context) {
	// id := c.Param("id")
	// firstname := c.DefaultQuery("firstname", "Guest")
	// lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")
	// user := User{ID: id, FirstName: firstname, Age: 18, LastName: lastname}
	// DB.Create(&user)
	var user User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	DB.Create(User{ID: user.ID, FirstName: user.FirstName, Age: user.Age, LastName: user.LastName})

	c.JSON(200, gin.H{
		"code":    200,
		"message": "pong",
		"data": gin.H{
			"id":        user.ID,
			"firstname": user.FirstName,
			"lastname":  user.LastName,
		},
	})
}
```

## swagger
@Param 配置 body
```
// PingExample godoc
// @Summary user create
// @Schemes
// @Description create user list
// @Tags Users
// @Accept json
// @Produce json
// @Param       data	body	string  true  "Request Data" SchemaExample({\r\n"id":"kkkkasdf",\r\n"firstname":"gaalaxy",\r\n"lastname":"fasdfa",\r\n"age":12\r\n})
// @Success 200 {object} model.User
// @Router /admin/users [post]
```
