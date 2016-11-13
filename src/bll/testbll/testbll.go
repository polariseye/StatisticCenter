package testbll

import (
	"github.com/polariseye/statisticcenter/src/model"
)

type testBll struct {
}

var (
	TestBll testBll
)

func (this testBll) GetClassName() string {
	return "TestBll"
}

func init() {
	model.GetPlayerInitManager().Register(TestBll.GetClassName(), &TestBll)
}

// 开始初始化
// instance:初始化对象
// 返回值：
// []string:错误信息列表
func (this *testBll) StartInit(instance interface{}) []string {
	return []string{}
}
