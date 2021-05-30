package main

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm/clause"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	USERNAME = "root"
	PASSWORD = "root123"
	HOST     = "10.86.124.122"
	PORT     = 3306
	DBNAME   = "zioyi"
)

type Product struct {
	ID        uint
	Code      string
	Price     uint `gorm:"default:1"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (Product) TableName() string {
	return "product"
}

func (p *Product) BeforeUpdate(tx *gorm.DB) (err error) {
	fmt.Println("before update called")
	return nil
}

func main() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", USERNAME, PASSWORD, HOST, PORT, DBNAME)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		fmt.Printf("open mysql connection failed, err: %v", err)
	}

	fetchProduct(db)

}

func debugResult(a ...interface{}) {
	a = append([]interface{}{"\nstart debug info:"}, a...)
	for _, i := range a {
		fmt.Printf("%+v\n", i)
	}
	fmt.Print()
}

func createProduct(db *gorm.DB) {
	var p = Product{Code: "book", Price: 10}
	result := db.Create(&p)
	fmt.Printf("%+v \n %+v", result, p)
}

func batchCreateProduct(db *gorm.DB) {
	var products = []Product{{Code: "pencil", Price: 5}, {Code: "erase", Price: 2}, {Code: "rule", Price: 2}}
	result := db.Create(&products)
	debugResult(result, products)
}

func batchCreateProductWay2(db *gorm.DB) {
	var products []Product
	for i := 1; i <= 1500; i++ {
		code := fmt.Sprintf("pencil-seris-%d", i)
		products = append(products, Product{Code: code, Price: uint(i)})
	}
	r := db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.Assignments(map[string]interface{}{"price": 111115}),
	}).Create(&products)
	debugResult(r, "")
}

func updateProduct(db *gorm.DB) {
	p := Product{Code: "banana12"}
	r := db.Clauses(clause.OnConflict{DoNothing: true}).Create(&p)
	debugResult(r, p)
	// 这种方式无法更新零值
	r = db.Model(&p).Updates(Product{Code: "banana13", Price: 0})
	debugResult(r, p)
	p.Price = 0
	r = db.Select("*").Updates(&p)
	debugResult(r, p)
	/*
		&{Config:0xc000166360 Error:<nil> RowsAffected:1 Statement:0xc0001bc1c0 clone:0}
		{ID:4519 Code:banana12 Price:1 CreatedAt:2021-05-24 14:54:10.594 +0800 CST UpdatedAt:2021-05-24 14:54:10.594 +0800 CST}
		&{Config:0xc000166360 Error:<nil> RowsAffected:1 Statement:0xc0001bce00 clone:0}
		{ID:4519 Code:banana13 Price:1 CreatedAt:2021-05-24 14:54:10.594 +0800 CST UpdatedAt:2021-05-24 14:54:10.6 +0800 CST}
		&{Config:0xc000166360 Error:<nil> RowsAffected:1 Statement:0xc0001bd180 clone:0}
		{ID:4519 Code:banana13 Price:0 CreatedAt:2021-05-24 14:54:10.594 +0800 CST UpdatedAt:2021-05-24 14:54:10.606 +0800 CST}
	*/

	//p := Product{Code: "banana-update-zero-value-4"}
	//r = db.Clauses(clause.OnConflict{DoNothing: true}).Create(&p)
	//debugResult(r, p)
	//p.Price = 0
	//p.Code = "banana-update-zero-value-4444"
	//r = db.Model(&p).Where("id = ?", p.ID).Updates(&p)
	productId := 4525
	values := map[string]interface{}{}
	values["code"] = "banana-updated-zero-1"
	values["price"] = 0
	r = db.Model(&Product{}).Where("id = ?", productId).Updates(values)
}

func fetchProduct(db *gorm.DB) {
	var products []Product
	r := db.Unscoped().Where("code = banana12").Find(&products)
	debugResult(r, products)
	//r = db.Delete(&products)
	//debugResult(r, products)
	r = db.Unscoped().Last(&products, "code = ?", "banana")
	fmt.Printf("%t\n", errors.Is(r.Error, gorm.ErrRecordNotFound))
	debugResult(r, products)

	r = db.Where("code IN ?", []string{"book", "rule", "e"}).Find(&products)
	debugResult(r, products, len(products))

	r = db.Where([]int32{1, 2, 3}).Find(&products)
	debugResult(r, products, len(products))

	// 零值查询
	r = db.Where(&Product{Code: "pencil-seris-1500", Price: 0}).Find(&products)
	debugResult(r, products)
	/*
			start debug info:
			&{Config:0xc000166360 Error:<nil> RowsAffected:1 Statement:0xc000189340 clone:0}
			[{ID:2013 Code:pencil-seris-1500 Price:111115 CreatedAt:2021-05-24 14:33:58.573 +0800 CST UpdatedAt:2021-05-24 14:33:58.573 +0800 CST DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}}]

		等价SQL： SELECT * FROM product WHERE code = "pencil-seris-1500";
	*/

	// 正确用法
	r = db.Where("code = ? AND price = ?", "pencil-seris-1500", 0).Find(&products)
	debugResult(r, products)

	r = db.Where("code = ? AND price = ?", "pencil-seris-1500", "111115\" select * from products").Find(&products)
	debugResult(r, products)
}
