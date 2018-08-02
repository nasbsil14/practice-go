package main

import (
	"fmt"
	"sync"
	_"github.com/go-sql-driver/mysql"
	"database/sql"
	"log"
)

func fetchDB(wg *sync.WaitGroup, con *sql.DB, req <-chan int, res chan []string) {
	defer wg.Done()

	for id := range req {
		rows, err := con.Query("SELECT Id, Title FROM Category WHERE Id = ?", id)
		if err != nil {
			log.Fatal(err)
		}

		results := make([]string, 0, 100)
		for rows.Next() {
			var id int
			var title string
			if err := rows.Scan(&id, &title); err != nil {
				log.Fatal(err)
			}
			fmt.Println(id, title)
			results = append(results, title)
		}
		res <- results
	}
}

func main() {
	fmt.Println("START")

	con, _ := sql.Open("mysql", "root:test@tcp(127.0.0.1:3306)/TestDB")
	defer con.Close()

	var wg sync.WaitGroup

	req := make(chan int)
	defer close(req)
	res := make(chan []string)
	defer close(res)

	// ワーカーを3つ作る
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go fetchDB(&wg, con, req, res)
	}

	// 同時には3つまでしか処理できない
	for i := 0; i < 10; i++ {
		req <- i
		res <- req
	}
	//for title := range res {
	//	fmt.Println(title)
	//}


	//	close(q) // これ大事

	wg.Wait() // すべてのgoroutineが終了するのをまつ

	fmt.Println("END")
}