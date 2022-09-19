package migrations

import (
	"database/sql"
	"gohub/app/models"
	"gohub/pkg/migrate"
	"gorm.io/gorm"
)

func init() {
	//type User struct {
	//	models.BaseModel
	//}
	//
	//type Category struct {
	//	models.BaseModel
	//}

	type Topic struct {
		models.BaseModel

		Title      string `gorm:"type:varchar(255);not null;index;comment:标题"`
		Body       string `gorm:"type:varchar(255);default:null;comment:内容"`
		UserId     string `gorm:"type:bigint;not null;index;comment:用户ID"`
		CategoryId string `gorm:"type:bigint;not null;index;comment:类ID"`

		// 会创建user_id 与category_id 外键的约束
		//User     User
		//Category Category

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&Topic{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&Topic{})
	}

	migrate.Add("2022_09_19_100801_add_topic_table", up, down)
}
