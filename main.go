package main

import (
	"fmt"

	"time"

	"./fileUtils"

	"github.com/kataras/iris"
)

type Message struct {
	Number string
	Text   string
}

func main() {
	app := iris.Default()

	// Method:   GET
	// Resource: http://localhost:8080/
	app.Handle("GET", "/", func(ctx iris.Context) {
		ctx.HTML("Hello world!")
	})

	// same as app.Handle("GET", "/ping", [...])
	// Method:   GET
	// Resource: http://localhost:8080/ping
	app.Get("/ping", func(ctx iris.Context) {
		ctx.WriteString("pong")
	})

	// Method:   GET
	// Resource: http://localhost:8080/hello
	app.Get("/hello", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "Hello iris web framework."})
	})

	// Method:   POST
	// Resource: http://localhost:8080/sms
	app.Post("/sms", func(ctx iris.Context) {
		key := ctx.GetHeader("X-Bondo-Key")
		if key != "Bondo1337" {
			panic("wrong Api Key")
		}

		message := Message{}
		err := ctx.ReadJSON(&message)

		if err != nil {
			panic(err.Error())
		}

		string := fmt.Sprintf("To: %s\nflash: true\n\n%s", message.Number, message.Text)
		fileUtils.WriteFile(message.Number+time.Now().Format("2006.01.02-15:04:05"), string)
	})

	// http://localhost:8080
	// http://localhost:8080/ping
	// http://localhost:8080/hello
	app.Run(iris.Addr(":8080"))
}
