package Database

import (
	"fmt"
	"trashgo/Models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "root:@tcp(127.0.0.1:3306)/trashgo_nontps_new?charset=utf8mb4&parseTime=True&loc=Local"
	con, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Gagal Meng koneksikan database")
	}
	fmt.Println("database terkoneksi")
	DB = con
	con.AutoMigrate(&Models.User{})
	con.AutoMigrate(&Models.Laporan{})

}
