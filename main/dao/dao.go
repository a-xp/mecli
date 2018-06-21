package dao

import (
	"MecomoApiMock/main/mecomo"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"fmt"
)

type IotDAO struct {
	dsn string
	db* sql.DB
}

func (dao *IotDAO) open() error{
	if dao.db==nil {
		db,err := sql.Open("sqlite3", dao.dsn)
		if err!=nil {
			return err
		}
		if err := checkSchema(db); err != nil {
			return err
		}
		dao.db = db
	}
	return nil
}

func checkSchema(db *sql.DB) error {
	_,err := db.Exec("CREATE TABLE IF NOT EXISTS devices (id INTEGER NOT NULL PRIMARY KEY, name TEXT NOT NULL, reg_date TEXT)")
	if err != nil {
		return err
	}
	return nil
}

func (dao *IotDAO) Close() error{
	if dao.db!=nil {
		err := dao.db.Close()
		dao.db = nil
		return err
	}
	return nil
}

func (dao* IotDAO) StoreDevices(devices []mecomo.Device)  {
	for _,device := range devices {
		dao.db.Exec("DELETE FROM devices WHERE id = ?", device.Id)
		dao.db.Exec("INSERT INTO devices (id, name, reg_date) VALUES (?,?,?)", device.Id, device.Name, device.RegDate)
	}
}

func (dao* IotDAO) GetDevices() (result []mecomo.Device)  {
	rows,err := dao.db.Query("SELECT id,name,reg_date FROM devices")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		device := mecomo.Device{}
		rows.Scan(&device.Id, &device.Name, &device.RegDate)
		result = append(result, device)
	}
	return result
}

func (dao* IotDAO) GetDevice(id int) mecomo.Device {
	row := dao.db.QueryRow("SELECT id,name,reg_date FROM devices WHERE id=?", id)
	device := mecomo.Device{}
	row.Scan(&device.Id, &device.Name, &device.RegDate)
	return device
}

func CreateDAO(dsn string) (IotDAO, error) {
	dao := IotDAO{dsn, nil}
	return dao, dao.open()
}