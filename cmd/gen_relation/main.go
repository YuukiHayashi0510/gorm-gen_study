package main

import (
	"study/internal/db"
	"study/internal/db/model"

	"gorm.io/gen"
	"gorm.io/gen/field"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath:           "./internal/db/query", // 出力パス
		Mode:              gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
		FieldWithIndexTag: true,
		FieldWithTypeTag:  true,
		FieldNullable:     true,
	})

	sqlHandler, err := db.NewSqlHandler("test:password@tcp(127.0.0.1:3306)/database")
	if err != nil {
		panic(err)
	}

	g.UseDB(sqlHandler.Conn)

	g.Execute()

	allModels := []interface{}{
		g.GenerateModel(
			model.TableNameSubCategory,
			gen.FieldRelateModel(field.BelongsTo, "Category", model.Category{}, &field.RelateConfig{
				RelatePointer: true,
			}),
		),
		g.GenerateModel(
			model.TableNameProduct,
			gen.FieldRelateModel(field.BelongsTo, "SubCategory", model.SubCategory{}, &field.RelateConfig{
				RelatePointer: true,
				GORMTag: field.GormTag{
					"foreignKey": []string{"SubCategoryID"}, "references": []string{"SubCategoryID"},
				},
			}),
		),
		g.GenerateModel(
			model.TableNameSalesProduct,
			gen.FieldRelateModel(field.BelongsTo, "Product", model.Product{}, &field.RelateConfig{
				RelatePointer: true,
				GORMTag: field.GormTag{
					"foreignKey": []string{"ProductID"}, "references": []string{"ProductID"},
				},
			}),
		),
		g.GenerateModel(
			model.TableNameSale,
			gen.FieldRelateModel(field.HasMany, "SalesProduct", model.SalesProduct{}, &field.RelateConfig{
				RelateSlicePointer: true,
				GORMTag: field.GormTag{
					"foreignKey": []string{"SalesID"}, "references": []string{"SalesID"},
				},
			}),
		),
	}
	g.ApplyBasic(allModels...)

	g.Execute()
}
