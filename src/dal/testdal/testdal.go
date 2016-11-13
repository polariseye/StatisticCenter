package testdal

import (
	"github.com/polariseye/statisticcenter/src/dal"
	"github.com/polariseye/statisticcenter/src/model/testmodel"
)

type testDalStruct struct {
}

var TestDal *testDalStruct

const (
	con_Test_Insert string = ""
)

func (this *testDalStruct) ByTran() {

}

func (this *testDalStruct) InsertTest(model testmodel.TestModel) {
	dal.GetConnection(dal.Con_Db_GameModel).Execute(con_Test_Insert,
		model.Name)
}

func (this *testDalStruct) GetList() {
	rows, errMsg := dal.GetConnection(dal.Con_Db_Game).Query(con_Test_Insert)
	if errMsg != nil {

	}

	rows.Close()
}
