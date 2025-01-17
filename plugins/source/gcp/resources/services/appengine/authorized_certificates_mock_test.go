// Code generated by codegen; DO NOT EDIT.

package appengine

import (
	"context"
	"fmt"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/cloudquery/plugins/source/gcp/client"
	"google.golang.org/grpc"
	"testing"

	pb "cloud.google.com/go/appengine/apiv1/appenginepb"
)

func createAuthorizedCertificates(gsrv *grpc.Server) error {
	fakeServer := &fakeAuthorizedCertificatesServer{}
	pb.RegisterAuthorizedCertificatesServer(gsrv, fakeServer)
	return nil
}

type fakeAuthorizedCertificatesServer struct {
	pb.UnimplementedAuthorizedCertificatesServer
}

func (f *fakeAuthorizedCertificatesServer) ListAuthorizedCertificates(context.Context, *pb.ListAuthorizedCertificatesRequest) (*pb.ListAuthorizedCertificatesResponse, error) {
	resp := pb.ListAuthorizedCertificatesResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func TestAuthorizedCertificates(t *testing.T) {
	client.MockTestGrpcHelper(t, AuthorizedCertificates(), createAuthorizedCertificates, client.TestOptions{})
}
