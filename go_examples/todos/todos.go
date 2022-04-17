// ref:: https://blog.logrocket.com/building-simple-app-go-postgresql/
package todos

import (
	"context"
	// "database/sql"

	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4"
)

func Server() {
	// TODO: use envs.
	// https://stackoverflow.com/questions/14588212/postgresql-resetting-password-of-postgresql-on-ubuntu
	// https://ubiq.co/database-blog/create-user-postgresql/
	connStr := "postgresql://developer:developer@localhost:5432/test_database?sslmode=disable"

	// conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	conn, err := pgx.Connect(context.Background(), connStr)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	defer conn.Close(context.Background())

	// fmt.Println(name, weight)

	// // Connect to database
	// db, err := sql.Open("postgres", connStr)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return indexHandler(c, conn, err)
	})

	app.Post("/", func(c *fiber.Ctx) error {
		return postHandler(c, conn)
	})

	app.Put("/update", func(c *fiber.Ctx) error {
		return putHandler(c, conn)
	})

	app.Delete("/delete", func(c *fiber.Ctx) error {
		return deleteHandler(c, conn)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Fatalln(app.Listen(fmt.Sprintf(":%v", port)))
}

func indexHandler(c *fiber.Ctx, conn *pgx.Conn, err error) error {
	// return c.SendString("Hello")
	// var res interface{}

	// var todos []string

	// rows, err := db.Query("SELECT * FROM todos")

	// defer rows.Close()

	// if err != nil {
	// 	log.Fatalln(err)
	// 	c.JSON("An error occured")
	// }

	// for rows.Next() {
	// 	rows.Scan(&res)

	// 	fmt.Println("data::")
	// 	fmt.Println(res)

	// 	// todos = append(todos, res)
	// }

	var id int64
	var item string

	// err = conn.QueryRow(context.Background(), "SELECT * FROM todos").Scan(&id, &item)

	err = conn.QueryRow(context.Background(), "select id, item from todos where id=$1", 1).Scan(&id, &item)

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(id, item)

	return c.SendString("Hello")

	// return c.Render("index", fiber.Map{
	// 	"Todos": todos,
	// })
}

func postHandler(c *fiber.Ctx, conn *pgx.Conn) error {
	return c.SendString("Hello")
}

func putHandler(c *fiber.Ctx, conn *pgx.Conn) error {
	return c.SendString("Hello")
}

func deleteHandler(c *fiber.Ctx, conn *pgx.Conn) error {
	return c.SendString("Hello")
}
