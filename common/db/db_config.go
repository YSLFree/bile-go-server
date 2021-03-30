package db

const (
	//Driver 驱动
	Driver string = "mysql"
	//Access str
	Access string = "root:123456@tcp(127.0.0.1:3306)/"
	//DBName dbname
	DBName string = "bile"
	//DbSet .
	DbSet string = "?charset=utf8mb4"
)

//SQL sql语句接口
type SQL struct{}

//SearchSQL search  user whether exist
func (sql *SQL) SearchSQL() string {
	return `SELECT * FROM logininfo WHERE account==? AND password==?`
}

//RegisterSQL register
func (sql *SQL) RegisterSQL() string {
	return `INSERT INTO logininfo (account, password) values (?, ?)`
}

//LoginSQL login
func (sql *SQL) LoginSQL() string {
	return `SELECT * FROM logininfo WHERE account==? AND password==?`
}
