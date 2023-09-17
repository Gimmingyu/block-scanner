package repository

import (
	"gorm.io/gorm"
)

func GormFindByID[T Table](tx *gorm.DB, id uint) (*T, error) {
	var (
		result = new(T)
		table  T
		err    error
	)

	if err = tx.Table(table.Table()).Where("id = ?", id).First(&result).Error; err != nil {
		return nil, tx.Rollback().Error
	}

	return result, nil
}

func GormFindOne[T Table](tx *gorm.DB, where string, args ...interface{}) (*T, error) {
	var (
		result = new(T)
		table  T
		err    error
	)

	if err = tx.Table(table.Table()).Where(where, args).First(&result).Error; err != nil {
		return nil, tx.Rollback().Error
	}

	return result, nil
}

func GormFindMany[T Table](tx *gorm.DB, where string, args ...interface{}) ([]*T, error) {
	var (
		result []*T
		table  T
		err    error
	)

	if err = tx.Table(table.Table()).Where(where, args).Find(&result).Error; err != nil {
		return nil, tx.Rollback().Error
	}

	return result, nil
}

func GormInsert[T Table](tx *gorm.DB, model *T) error {
	var (
		table T
		err   error
	)

	if err = tx.Table(table.Table()).Create(&model).Error; err != nil {
		return err
	}

	return nil
}

func GormInsertMany[T Table](tx *gorm.DB, models []*T, batchSize int) error {
	var (
		table T
		err   error
	)

	if err = tx.Table(table.Table()).CreateInBatches(&models, batchSize).Error; err != nil {
		return err
	}

	return nil
}

func GormUpdate[T Table](tx *gorm.DB, filter, value map[string]interface{}) error {
	var (
		table T
		err   error
	)

	err = tx.Table(table.Table()).Where(filter).Updates(value).Error
	return err
}

func GormUpdateColumn[T Table](tx *gorm.DB, filter map[string]interface{}, column string, value interface{}) error {
	var (
		table T
		err   error
	)

	err = tx.Table(table.Table()).Where(filter).UpdateColumn(column, value).Error
	return err
}

func GormDelete[T Table](tx *gorm.DB, filter map[string]interface{}) error {
	var (
		table = new(T)
		err   error
	)

	err = tx.Where(filter).Delete(&table).Error
	return err
}

func GormCount[T Table](tx *gorm.DB, where string, args ...interface{}) (int64, error) {
	var (
		table T
		count int64
		err   error
	)

	if err = tx.Table(table.Table()).Where(where, args).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func GormFindWithJoin[T Table](tx *gorm.DB, join, where string, args ...interface{}) ([]*T, error) {
	var (
		result []*T
		table  T
		err    error
	)

	if err = tx.Table(table.Table()).Joins(join).Where(where, args).Find(&result).Error; err != nil {
		return nil, tx.Rollback().Error
	}

	return result, nil
}

func GormFindWithSubQuery[T Table](tx *gorm.DB, subQuery, where string, args ...interface{}) ([]*T, error) {
	var (
		result []*T
		table  T
		err    error
	)

	if err = tx.Table(table.Table()).Where(where, args).Where(subQuery).Find(&result).Error; err != nil {
		return nil, tx.Rollback().Error
	}

	return result, nil
}

func GormQuery(tx *gorm.DB, query string, args []interface{}, dest interface{}) error {
	return tx.Raw(query, args).Find(&dest).Error
}

func GormExec(tx *gorm.DB, query string, args ...interface{}) error {
	return tx.Exec(query, args).Error
}
