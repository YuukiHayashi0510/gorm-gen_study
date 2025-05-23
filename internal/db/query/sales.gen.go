// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"study/internal/db/model"
)

func newSale(db *gorm.DB, opts ...gen.DOOption) sale {
	_sale := sale{}

	_sale.saleDo.UseDB(db, opts...)
	_sale.saleDo.UseModel(&model.Sale{})

	tableName := _sale.saleDo.TableName()
	_sale.ALL = field.NewAsterisk(tableName)
	_sale.SalesID = field.NewString(tableName, "sales_id")
	_sale.SalesDate = field.NewTime(tableName, "sales_date")
	_sale.TotalSales = field.NewInt64(tableName, "total_sales")
	_sale.TotalPrice = field.NewInt64(tableName, "total_price")
	_sale.TotalTax = field.NewInt64(tableName, "total_tax")
	_sale.TotalReducedTax = field.NewInt64(tableName, "total_reduced_tax")
	_sale.IsReturned = field.NewInt32(tableName, "is_returned")
	_sale.ReturnedAt = field.NewTime(tableName, "returned_at")
	_sale.SalesProduct = saleHasManySalesProduct{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("SalesProduct", "model.SalesProduct"),
		Product: struct {
			field.RelationField
			SubCategory struct {
				field.RelationField
				Category struct {
					field.RelationField
				}
			}
		}{
			RelationField: field.NewRelation("SalesProduct.Product", "model.Product"),
			SubCategory: struct {
				field.RelationField
				Category struct {
					field.RelationField
				}
			}{
				RelationField: field.NewRelation("SalesProduct.Product.SubCategory", "model.SubCategory"),
				Category: struct {
					field.RelationField
				}{
					RelationField: field.NewRelation("SalesProduct.Product.SubCategory.Category", "model.Category"),
				},
			},
		},
	}

	_sale.fillFieldMap()

	return _sale
}

type sale struct {
	saleDo

	ALL             field.Asterisk
	SalesID         field.String
	SalesDate       field.Time
	TotalSales      field.Int64
	TotalPrice      field.Int64
	TotalTax        field.Int64
	TotalReducedTax field.Int64
	IsReturned      field.Int32
	ReturnedAt      field.Time
	SalesProduct    saleHasManySalesProduct

	fieldMap map[string]field.Expr
}

func (s sale) Table(newTableName string) *sale {
	s.saleDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s sale) As(alias string) *sale {
	s.saleDo.DO = *(s.saleDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *sale) updateTableName(table string) *sale {
	s.ALL = field.NewAsterisk(table)
	s.SalesID = field.NewString(table, "sales_id")
	s.SalesDate = field.NewTime(table, "sales_date")
	s.TotalSales = field.NewInt64(table, "total_sales")
	s.TotalPrice = field.NewInt64(table, "total_price")
	s.TotalTax = field.NewInt64(table, "total_tax")
	s.TotalReducedTax = field.NewInt64(table, "total_reduced_tax")
	s.IsReturned = field.NewInt32(table, "is_returned")
	s.ReturnedAt = field.NewTime(table, "returned_at")

	s.fillFieldMap()

	return s
}

func (s *sale) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *sale) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 9)
	s.fieldMap["sales_id"] = s.SalesID
	s.fieldMap["sales_date"] = s.SalesDate
	s.fieldMap["total_sales"] = s.TotalSales
	s.fieldMap["total_price"] = s.TotalPrice
	s.fieldMap["total_tax"] = s.TotalTax
	s.fieldMap["total_reduced_tax"] = s.TotalReducedTax
	s.fieldMap["is_returned"] = s.IsReturned
	s.fieldMap["returned_at"] = s.ReturnedAt

}

func (s sale) clone(db *gorm.DB) sale {
	s.saleDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s sale) replaceDB(db *gorm.DB) sale {
	s.saleDo.ReplaceDB(db)
	return s
}

type saleHasManySalesProduct struct {
	db *gorm.DB

	field.RelationField

	Product struct {
		field.RelationField
		SubCategory struct {
			field.RelationField
			Category struct {
				field.RelationField
			}
		}
	}
}

func (a saleHasManySalesProduct) Where(conds ...field.Expr) *saleHasManySalesProduct {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a saleHasManySalesProduct) WithContext(ctx context.Context) *saleHasManySalesProduct {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a saleHasManySalesProduct) Session(session *gorm.Session) *saleHasManySalesProduct {
	a.db = a.db.Session(session)
	return &a
}

func (a saleHasManySalesProduct) Model(m *model.Sale) *saleHasManySalesProductTx {
	return &saleHasManySalesProductTx{a.db.Model(m).Association(a.Name())}
}

type saleHasManySalesProductTx struct{ tx *gorm.Association }

func (a saleHasManySalesProductTx) Find() (result []*model.SalesProduct, err error) {
	return result, a.tx.Find(&result)
}

func (a saleHasManySalesProductTx) Append(values ...*model.SalesProduct) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a saleHasManySalesProductTx) Replace(values ...*model.SalesProduct) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a saleHasManySalesProductTx) Delete(values ...*model.SalesProduct) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a saleHasManySalesProductTx) Clear() error {
	return a.tx.Clear()
}

func (a saleHasManySalesProductTx) Count() int64 {
	return a.tx.Count()
}

type saleDo struct{ gen.DO }

type ISaleDo interface {
	gen.SubQuery
	Debug() ISaleDo
	WithContext(ctx context.Context) ISaleDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISaleDo
	WriteDB() ISaleDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISaleDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISaleDo
	Not(conds ...gen.Condition) ISaleDo
	Or(conds ...gen.Condition) ISaleDo
	Select(conds ...field.Expr) ISaleDo
	Where(conds ...gen.Condition) ISaleDo
	Order(conds ...field.Expr) ISaleDo
	Distinct(cols ...field.Expr) ISaleDo
	Omit(cols ...field.Expr) ISaleDo
	Join(table schema.Tabler, on ...field.Expr) ISaleDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISaleDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISaleDo
	Group(cols ...field.Expr) ISaleDo
	Having(conds ...gen.Condition) ISaleDo
	Limit(limit int) ISaleDo
	Offset(offset int) ISaleDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISaleDo
	Unscoped() ISaleDo
	Create(values ...*model.Sale) error
	CreateInBatches(values []*model.Sale, batchSize int) error
	Save(values ...*model.Sale) error
	First() (*model.Sale, error)
	Take() (*model.Sale, error)
	Last() (*model.Sale, error)
	Find() ([]*model.Sale, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Sale, err error)
	FindInBatches(result *[]*model.Sale, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Sale) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISaleDo
	Assign(attrs ...field.AssignExpr) ISaleDo
	Joins(fields ...field.RelationField) ISaleDo
	Preload(fields ...field.RelationField) ISaleDo
	FirstOrInit() (*model.Sale, error)
	FirstOrCreate() (*model.Sale, error)
	FindByPage(offset int, limit int) (result []*model.Sale, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISaleDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s saleDo) Debug() ISaleDo {
	return s.withDO(s.DO.Debug())
}

func (s saleDo) WithContext(ctx context.Context) ISaleDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s saleDo) ReadDB() ISaleDo {
	return s.Clauses(dbresolver.Read)
}

func (s saleDo) WriteDB() ISaleDo {
	return s.Clauses(dbresolver.Write)
}

func (s saleDo) Session(config *gorm.Session) ISaleDo {
	return s.withDO(s.DO.Session(config))
}

func (s saleDo) Clauses(conds ...clause.Expression) ISaleDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s saleDo) Returning(value interface{}, columns ...string) ISaleDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s saleDo) Not(conds ...gen.Condition) ISaleDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s saleDo) Or(conds ...gen.Condition) ISaleDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s saleDo) Select(conds ...field.Expr) ISaleDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s saleDo) Where(conds ...gen.Condition) ISaleDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s saleDo) Order(conds ...field.Expr) ISaleDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s saleDo) Distinct(cols ...field.Expr) ISaleDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s saleDo) Omit(cols ...field.Expr) ISaleDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s saleDo) Join(table schema.Tabler, on ...field.Expr) ISaleDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s saleDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISaleDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s saleDo) RightJoin(table schema.Tabler, on ...field.Expr) ISaleDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s saleDo) Group(cols ...field.Expr) ISaleDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s saleDo) Having(conds ...gen.Condition) ISaleDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s saleDo) Limit(limit int) ISaleDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s saleDo) Offset(offset int) ISaleDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s saleDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISaleDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s saleDo) Unscoped() ISaleDo {
	return s.withDO(s.DO.Unscoped())
}

func (s saleDo) Create(values ...*model.Sale) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s saleDo) CreateInBatches(values []*model.Sale, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s saleDo) Save(values ...*model.Sale) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s saleDo) First() (*model.Sale, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Sale), nil
	}
}

func (s saleDo) Take() (*model.Sale, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Sale), nil
	}
}

func (s saleDo) Last() (*model.Sale, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Sale), nil
	}
}

func (s saleDo) Find() ([]*model.Sale, error) {
	result, err := s.DO.Find()
	return result.([]*model.Sale), err
}

func (s saleDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Sale, err error) {
	buf := make([]*model.Sale, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s saleDo) FindInBatches(result *[]*model.Sale, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s saleDo) Attrs(attrs ...field.AssignExpr) ISaleDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s saleDo) Assign(attrs ...field.AssignExpr) ISaleDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s saleDo) Joins(fields ...field.RelationField) ISaleDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s saleDo) Preload(fields ...field.RelationField) ISaleDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s saleDo) FirstOrInit() (*model.Sale, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Sale), nil
	}
}

func (s saleDo) FirstOrCreate() (*model.Sale, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Sale), nil
	}
}

func (s saleDo) FindByPage(offset int, limit int) (result []*model.Sale, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s saleDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s saleDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s saleDo) Delete(models ...*model.Sale) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *saleDo) withDO(do gen.Dao) *saleDo {
	s.DO = *do.(*gen.DO)
	return s
}
