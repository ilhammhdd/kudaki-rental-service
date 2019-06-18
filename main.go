package main

import (
	"os"
	"strings"

	"github.com/ilhammhdd/kudaki-externals/mysql"

	"github.com/ilhammhdd/go-toolkit/safekit"

	"github.com/ilhammhdd/kudaki-rental-service/externals/eventdriven"
)

func init() {
	if len(os.Args) > 1 {
		for _, val := range os.Args[1:] {
			f := strings.Split(val, " ")
			os.Setenv(f[1], f[2])
		}
	}

	mysql.OpenDB(os.Getenv("DB_PATH"), os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
}

func main() {
	wp := safekit.NewWorkerPool()

	wp.Worker <- new(eventdriven.AddCartItem)
	wp.Worker <- new(eventdriven.DeleteCartItem)

	wp.PoolWG.Wait()
}
