package config

import (
	"errors"
	"fmt"
	"time"

	"github.com/beego/beego/orm"
	"github.com/go-sql-driver/mysql"
)

//struct fields related to the MySQL connection configuration,
// such as connection Alias, RetryCount, RetryDelay, TimeZone, DebugFlag, and a pointer to a mysql.Config object.
type mysqlConnectionConfig struct {
	cf         *mysql.Config
	Alias      string
	RetryCount int
	RetryDelay time.Duration
	TimeZone   *time.Location
	DebugFlag  bool
}

func ConnectToDatabase(keys ...string) error {
	fmt.Println("Connection to DB")
	//root-username,p@=password,12.23.23=(ip address),databasename
	conString := "root:P@mmount847@tcp(10.75.60.225:3307)/yp?interpolateParams=true&charset=utf8mb4"
	fmt.Println("ConString:", conString)
	//deriving cf as type myqlcoontconfig interface
	cf := mysqlConnectionConfig{}
	//parseDSN is used to parse the constring into the struct
	err := cf.ParseDSN(conString)
	if err != nil {
		return errors.New("failed to parse con string into struct")
	}
	//sets all the values of the struct
	ISTLocation, _ := time.LoadLocation("Asia/Kolkata")
	cf.TimeZone = ISTLocation
	cf.Alias = "default"
	cf.RetryDelay = 500 * time.Millisecond
	cf.RetryCount = 5
	cf.DebugFlag = true
	err = beegoRegisterDB(cf)
	if err != nil {
		return err
	}
	fmt.Println("DB last")
	return nil
}

//this func register the DB using ORM frame work
func beegoRegisterDB(cf mysqlConnectionConfig) error {
	var err error
	for breaker := cf.RetryCount; breaker > 0; breaker-- {
		if breaker < cf.RetryCount {
			time.Sleep(cf.RetryDelay)
		}
		//check if database exsits
		_, err := orm.GetDB(cf.Alias)
		if err != nil {
			err = orm.RegisterDataBase(cf.Alias, "mysql", cf.FormatDSN())
		}
		if err != nil {
			fmt.Println("failed to register db")
			continue
		}
		//test newly registered database
		err = MysqlTest(cf.Alias)
		if err != nil {
			fmt.Println("failed to query")
			continue
		}
		orm.DefaultTimeLoc = cf.TimeZone
		orm.Debug = cf.DebugFlag
		break
	}
	//checks if any error occur,prints the value of alias ,dsn,error
	if err != nil {
		fmt.Println(fmt.Sprintf("beegoRegisterDB(alias: %s, dsn: %s) failed, error: %v", cf.Alias, cf.String(), err))
	}
	return err
}

//func with 'c' of type 'mysqlConnectionConfig' can be called as method
func (c *mysqlConnectionConfig) ParseDSN(s string) error {
	var err error
	//invoking PARSEDSN method wiht string argument init
	c.cf, err = mysql.ParseDSN(s)
	//return *config struct
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil

}
func MysqlTest(alias string) error {
	o := orm.NewOrm()
	o.Using(alias)
	a, err := o.Raw("SELECT 1").Exec()
	fmt.Println(a)
	return err
}
func (c *mysqlConnectionConfig) String() string {
	t := c.cf.Clone()
	t.Passwd = "root"
	return t.FormatDSN()
}
func (c *mysqlConnectionConfig) FormatDSN() string {
	return c.cf.FormatDSN()
}
