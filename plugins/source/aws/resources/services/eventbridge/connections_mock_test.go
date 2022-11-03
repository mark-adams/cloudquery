// Code generated by codegen; DO NOT EDIT.

package eventbridge

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/eventbridge"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildEventbridgeConnectionsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEventbridgeClient(ctrl)
	object := types.Connection{}
	err := faker.FakeObject(&object)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListConnections(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&eventbridge.ListConnectionsOutput{
			Connections: []types.Connection{object},
		}, nil)

	return client.Services{
		Eventbridge: m,
	}
}

func TestEventbridgeConnections(t *testing.T) {
	client.AwsMockTestHelper(t, Connections(), buildEventbridgeConnectionsMock, client.TestOptions{})
}