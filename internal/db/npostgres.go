package db

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
	"todo/internal/models"
)

type NPostgreSQLRepository struct {
	DatabaseDialect string
	DatabaseURL     string
}

var pgOrm orm.Ormer = nil

func (c NPostgreSQLRepository) ConnectDB() interface{} {
	if pgOrm == nil {
		err := orm.RegisterDriver(c.DatabaseDialect, orm.DRPostgres)
		if err != nil {
			return nil
		}
		err = orm.RegisterDataBase("default", c.DatabaseDialect, c.DatabaseURL)
		if err != nil {
			panic("failed to connect database")
		}

		// Create database tables from struct models
		// Initially keep force true inorder to create new table but after that make it false
		// The default behavior for Beego is to add additional columns when the model is updated.
		// But need to manually handle dropping the columns if they are removed from the model.
		err = orm.RunSyncdb("default", false, true)
		if err != nil {
			return nil
		}

		pgOrm = orm.NewOrm()
	}
	return nil
}

func (c NPostgreSQLRepository) CloseDB(conn interface{}) error {
	return nil
}

func (c NPostgreSQLRepository) Create(model interface{}, value interface{}) error {
	c.ConnectDB()
	_, err := pgOrm.Insert(value)
	return err
}

// FindUser this returns the user with the password hash. Hence, do not expose the return value
// without deleting or masking the password field in the result
func (c NPostgreSQLRepository) FindUser(value interface{}) interface{} {
	c.ConnectDB()
	var user models.User
	err := pgOrm.QueryTable(new(models.User)).Filter("email", value).One(&user)
	if err != nil {
		return user
	}
	return user
}

func (c NPostgreSQLRepository) GetAllTodo() interface{} {
	c.ConnectDB()
	var res []orm.Params
	//todos := make([]*models.Todo, 0)
	_, err := pgOrm.QueryTable(new(models.Todo)).RelatedSel().Values(&res)
	if err != nil {
		return nil
	}
	//for _, val := range todos {
	//	_, err = pgOrm.LoadRelated(val, "Editors")
	//}
	return res
}
