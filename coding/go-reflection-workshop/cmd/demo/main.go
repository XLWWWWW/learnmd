package main

import (
	"fmt"
	"log"

	"example.com/go-reflection-workshop/practice"
)

func main() {
	type Config struct {
		Host string `default:"localhost" map:"host" validate:"required"`
		Port int16  `default:"8080" map:"port" validate:"min=1"`
	}
	cfg := Config{}
	if err := practice.ApplyDefaults(&cfg); err != nil {
		log.Fatal(err)
	}
	issues, err := practice.Validate(cfg)
	if err != nil {
		log.Fatal(err)
	}
	view, err := practice.ToMap(cfg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("config: host=%v port=%v issues=%d\n", view["host"], view["port"], len(issues))

	type User struct {
		ID   int64  `db:"user_id"`
		Name string `db:"name"`
	}
	var user User
	if err := practice.ScanRow(&user, []string{"user_id", "name"}, []any{int64(7), "小明"}); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("row: ID=%d Name=%s\n", user.ID, user.Name)
	result, err := practice.Invoke(func(a, b int) int { return a + b }, 2, 3)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("invoke:", result[0])
}
