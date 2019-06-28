package pg

import (
	"fmt"

	"github.com/go-pg/pg"
)

// Params 测试用途
type Params struct {
	X int
	Y int
}

// Sum 返回相加值
func (p *Params) Sum() int {
	return p.X + p.Y
}

// go-pg recognizes `?` in queries as placeholders and replaces them
// with parameters when queries are executed. `?` can be escaped with backslash.
// Parameters are escaped before replacing according to PostgreSQL rules.
// Specifically:
//   - all parameters are properly quoted against SQL injections;
//   - null byte is removed;
//   - JSON/JSONB gets `\u0000` escaped as `\\u0000`.
func main() {

	db := pg.Connect(&pg.Options{
		User:     "postgres",
		Password: "123456",
		Database: "mydb",
	})
	defer db.Close()

	var n int
	_, err := db.QueryOne(pg.Scan(&n), "insert into my_table values (4,'apple2') RETURNING id")
	//panicIf(err)
	fmt.Print(err)
	//panic(err)
	fmt.Println(n)

	/*
		var num int

		// Simple params.
		_, err := pgdb.Query(pg.Scan(&num), "SELECT ?", 42)
		if err != nil {
			panic(err)
		}
		fmt.Println("simple:", num)

		// Indexed params.
		_, err = pgdb.Query(pg.Scan(&num), "SELECT ?0 + ?0", 1)
		if err != nil {
			panic(err)
		}
		fmt.Println("indexed:", num)

		// Named params.
		params := &Params{
			X:  1,
			Y:  1,
		}
		_, err = pgdb.Query(pg.Scan(&num), "SELECT ?x + ?y + ?Sum", params)
		if err != nil {
			panic(err)
		}
		fmt.Println("named:", num)

		// Global params.
		_, err = pgdb.WithParam("z", 1).Query(pg.Scan(&num), "SELECT ?x + ?y + ?z", params)
		if err != nil {
			panic(err)
		}
		fmt.Println("global:", num)

		// Model params.
		var tableName, tableAlias, tableColumns, columns string
		_, err = pgdb.Model(&Params{}).Query(
			pg.Scan(&tableName, &tableAlias, &tableColumns, &columns),
			"SELECT '?TableName', '?TableAlias', '?TableColumns', '?Columns'",
		)
		if err != nil {
			panic(err)
		}
		fmt.Println("table name:", tableName)
		fmt.Println("table alias:", tableAlias)
		fmt.Println("table columns:", tableColumns)
		fmt.Println("columns:", columns)

	*/

}
