package migrations

import (
	"database/sql"
	"gohub/app/models"
	"gohub/pkg/migrate"
	"gorm.io/gorm"
)

func init() {

	type Link struct {
		models.BaseModel

		Name string `gorm:"type:varchar(255);not null;index;comment:名称"`
		URL  string `gorm:"type:varchar(255);index;default:null;comment:URL"`

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&Link{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&Link{})
	}

	migrate.Add("2022_09_19_150905_add_link_table", up, down)
}
