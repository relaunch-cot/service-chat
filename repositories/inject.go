package repositories

import (
	"github.com/relaunch-cot/lib-relaunch-cot/repositories/mysql"
	MysqlRepository "github.com/relaunch-cot/service-chat/repositories/mysql"
)

type Repositories struct {
	Mysql MysqlRepository.IMySqlChat
}

func (r *Repositories) Inject(mysqlClient *mysql.Client) {
	r.Mysql = MysqlRepository.NewMysqlRepository(mysqlClient)
}
