package db

import (
	"xorm.io/xorm"
)

func Create(engine *xorm.Engine, table interface{}) error {
	_, err := engine.Insert(table)
	return err
}

func Read(engine *xorm.Engine, table interface{}, condition interface{}) error {
	_, err := engine.Where(condition).Get(table)
	return err
}

func Update(engine *xorm.Engine, table interface{}, condition interface{}) error {
	_, err := engine.Where(condition).Update(table)
	return err
}

func Delete(engine *xorm.Engine, table interface{}, condition interface{}) error {
	_, err := engine.Where(condition).Delete(table)
	return err
}

func GetByID(engine *xorm.Engine, table interface{}, id int64) (bool, error) {
	found, err := engine.ID(id).Get(table)
	return found, err
}
func GetByName(engine *xorm.Engine, table interface{}, name string) (bool, error) {
	found, err := engine.Where("name = ?", name).Get(table)
	return found, err
}

func GetAll(engine *xorm.Engine, tableSlice interface{}) error {
	err := engine.Find(tableSlice)
	return err
}

func SearchSector(engine *xorm.Engine, table interface{}, name string, companyID uint64, number string) (bool, error) {
	found, err := engine.Where("name = ? AND company_id = ? AND number = ?", name, companyID, number).Get(table)
	return found, err
}

func SearchAllSector(engine *xorm.Engine, tableSlice interface{}, company string) error {
	err := engine.Where("company = ?", company).Find(tableSlice)
	return err
}
