package test_utils

import (
	"github.com/atlas-dns/common"
	"github.com/golang/mock/gomock"
)

var TestContext = func() *common.Context {
	corelationId := "co-relationId"

	common.SetUpLogging()

	return &common.Context{
		Logger:        common.Log,
		CorrelationID: &corelationId,
	}
}()

type ContextMatcher struct {
	ctx *common.Context
}

func EqContextMatcher(ctx *common.Context) gomock.Matcher {
	return ContextMatcher{ctx: ctx}
}

func (m ContextMatcher) Matches(x interface{}) bool {

	_, ok := x.(*common.Context)
	return ok
}

func (m ContextMatcher) String() string {
	return "ContextMatcher"
}
