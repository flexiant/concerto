package api

import (
	"github.com/flexiant/concerto/testdata"

	"testing"
)

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
