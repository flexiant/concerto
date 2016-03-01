package api

import (
	"encoding/json"
	"github.com/flexiant/concerto/api/types"
	"github.com/flexiant/concerto/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

// GetAdminReportListMocked test mocked function
func GetAdminReportListMocked(t *testing.T, adminReportsIn *[]types.Report) *[]types.Report {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	rs, err := NewReportService(cs)
	assert.Nil(err, "Couldn't load report service")
	assert.NotNil(rs, "Report service not instanced")

	// to json
	arIn, err := json.Marshal(adminReportsIn)
	assert.Nil(err, "Cluster test data corrupted")

	// call service
	cs.On("Get", "/v1/admin/reports").Return(arIn, 200, nil)
	arOut, err := rs.GetAdminReportList()
	assert.Nil(err, "Error getting admin report")

	// exclude time data from comparison
	assert.Equal(len(*adminReportsIn), len(arOut), "Returned wrong number of elements")

	for i := 0; i < len(*adminReportsIn); i++ {
		(*adminReportsIn)[i].StartTime = arOut[i].StartTime
		(*adminReportsIn)[i].EndTime = arOut[i].EndTime
	}

	assert.Equal(*adminReportsIn, arOut, "GetClusterList returned different clusters")

	return &arOut
}
