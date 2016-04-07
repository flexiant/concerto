package testdata

import "github.com/flexiant/concerto/api/types"

// GetLoadBalancerData loads test data
func GetLoadBalancerData() *[]types.LoadBalancer {

	testLoadBalancers := []types.LoadBalancer{
		{
			Id:                          "fakeId0",
			Name:                        "fakeName0",
			Fqdn:                        "fakeFqdn0",
			Protocol:                    "fakeProtocol0",
			Port:                        1234,
			Algorithm:                   "fakeAlgorithm0",
			SslCertificate:              "fakeSslCert0",
			Ssl_certificate_private_key: "fakeSslCertPK0",
			Domain_id:                   "fakeDomId0",
			Cloud_provider_id:           "fakeCloudProvId0",
			Traffic_in:                  1024,
			Traffic_out:                 2048,
		},
		{
			Id:                          "fakeId1",
			Name:                        "fakeName1",
			Fqdn:                        "fakeFqdn1",
			Protocol:                    "fakeProtocol1",
			Port:                        1235,
			Algorithm:                   "fakeAlgorithm1",
			SslCertificate:              "fakeSslCert1",
			Ssl_certificate_private_key: "fakeSslCertPK1",
			Domain_id:                   "fakeDomId1",
			Cloud_provider_id:           "fakeCloudProvId1",
			Traffic_in:                  10240,
			Traffic_out:                 20480,
		},
	}

	return &testLoadBalancers
}

// GetLoadBalancerRecordData loads test data
func GetLBNodeData() *[]types.LBNode {

	testLBNode := []types.LBNode{
		{
			Id:       "fakeId0.0",
			Name:     "fakeName0.0",
			PublicIp: "fakePubIp0",
			State:    "fakeState0",
			ServerId: "fakeServerId0",
			Port:     1234,
		},
		{
			Id:       "fakeId0.1",
			Name:     "fakeName0.1",
			PublicIp: "fakePubIp1",
			State:    "fakeState1",
			ServerId: "fakeServerId1",
			Port:     1235,
		},
		{
			Id:       "fakeId0.2",
			Name:     "fakeName0.2",
			PublicIp: "fakePubIp2",
			State:    "fakeState2",
			ServerId: "fakeServerId2",
			Port:     1236,
		},
	}

	return &testLBNode
}
