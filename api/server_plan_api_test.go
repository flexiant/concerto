package api

import (
	"github.com/flexiant/concerto/testdata"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewServerPlanServiceNil(t *testing.T) {
	assert := assert.New(t)
	rs, err := NewServerPlanService(nil)
	assert.Nil(rs, "Uninitialized service should return nil")
	assert.NotNil(err, "Uninitialized service should return error")
}

func TestGetServerPlanList(t *testing.T) {
	serverPlansIn := testdata.GetServerPlanData()
	for _, serverPlanIn := range *serverPlansIn {
		GetServerPlanListMocked(t, serverPlansIn, serverPlanIn.CloudProviderId)
	}
}

func TestGetServerPlan(t *testing.T) {
	serverPlansIn := testdata.GetServerPlanData()
	for _, serverPlanIn := range *serverPlansIn {
		GetServerPlanMocked(t, &serverPlanIn)
	}
}
