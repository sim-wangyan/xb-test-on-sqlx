package main

import (
	"fmt"
	"testing"

	"github.com/fndome/xb"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

func InitSqlxDB() *sqlx.DB {

	var err interface{}
	Db, err = sqlx.Connect("mysql",
		"root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True")
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		if Db != nil {
			Db.Close()
		}
	}
	Db.SetMaxOpenConns(16)
	Db.SetMaxIdleConns(8)

	return Db
}

func TestBulderX(t *testing.T) {

	//"COUNT(DISTINCT d.id) AS `d.id_count`"
	arr := []interface{}{3000, 4000, 5000, 6000}

	subP := func(sb *xb.BuilderX) {
		sb.Select("id").From("t_pet")
	}

	builder := xb.Of(&Cat{}).As("c").
		Select("distinct c.color", "COUNT(DISTINCT d.id) AS `d.id_count`"). //"COUNT(DISTINCT d.id) AS `d.id_count`"
		FromX(func(fb *xb.FromBuilder) {
			fb.JOIN(xb.LEFT).Sub(subP).As("p").
				On("p.id = c.pet_id").
				Cond(func(on *xb.ON) {
					on.Gt("p.weight", 10)
				}).
				JOIN(xb.INNER).Of("t_dog").As("d").
				On("d.id = c.pet_id")
		}).
		Eq("p.id", 1).
		In("weight", arr...).
		GroupBy("c.color").
		Having(func(cb *xb.CondBuilderX) {
			cb.Gt("id", 1000)
		}).
		Sort("p.id", xb.DESC).
		Paged(func(pb *xb.PageBuilder) {
			pb.Rows(10).Last(101)
		})

	countSql, dataSql, vs, metaMap := builder.WithoutOptimization().Build().SqlOfPage()
	fmt.Println(dataSql)
	fmt.Println(vs)
	fmt.Println(metaMap)
	fmt.Println(countSql)

	InitSqlxDB()

	catList := []Cat{}
	err := Db.Select(&catList, dataSql, vs...)
	if err != nil {
		fmt.Println(err)
	}
	s := fmt.Sprintf("price : %v", *(catList[0].Price))
	fmt.Println(s)
}
