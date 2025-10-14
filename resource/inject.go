package resource

import (
	"github.com/relaunch-cot/service-chat/handler"
	"github.com/relaunch-cot/service-chat/repositories"
	"github.com/relaunch-cot/service-chat/server"
)

var Repositories repositories.Repositories
var Handler handler.Handlers
var Server server.Servers

func Inject() {
	mysqlClient := OpenMysqlConn()

	Repositories.Inject(mysqlClient)
	Handler.Inject(&Repositories)
	Server.Inject(&Handler)
}
