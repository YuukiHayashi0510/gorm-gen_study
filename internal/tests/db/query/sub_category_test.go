package query_test

import (
	"database/sql"
	"study/internal/db/query"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestSubCategory(t *testing.T) {
	sqlDB, err := sql.Open("mysql", "test:password@tcp(127.0.0.1:3306)/database")
	if err != nil {
		t.Fatal(err)
	}

	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}))
	if err != nil {
		t.Fatal(err)
	}
	subCategoryQuery := query.Use(db).SubCategory

	t.Run("SubCategoryを取得できる", func(t *testing.T) {
		s, err := subCategoryQuery.Find()
		if err != nil {
			t.Error(err)
		}

		if s == nil {
			t.Error("invalid query")
		}
	})
}
