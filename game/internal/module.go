package internal

import (
	"gameserver/base"
	"gameserver/db"

	"github.com/name5566/leaf/module"
	"gorm.io/gorm"
)

var (
	skeleton = base.NewSkeleton()
	ChanRPC  = skeleton.ChanRPCServer
	mysqlDB  = &gorm.DB{}
)

type Module struct {
	*module.Skeleton
}

func (m *Module) OnInit() {
	m.Skeleton = skeleton
	mysqlDB = db.ConnectDB()

}

func (m *Module) OnDestroy() {
	newDB, _ := mysqlDB.DB()
	newDB.Close()
}
