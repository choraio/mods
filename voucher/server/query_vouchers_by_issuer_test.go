package server

import (
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	voucherv1 "github.com/choraio/mods/voucher/api/v1"
	v1 "github.com/choraio/mods/voucher/types/v1"
)

type queryVouchersByIssuer struct {
	*baseSuite
	res *v1.QueryVouchersByIssuerResponse
	err error
}

func TestQueryVouchersByIssuer(t *testing.T) {
	gocuke.NewRunner(t, &queryVouchersByIssuer{}).
		Path("./features/query_vouchers_by_issuer.feature").
		Run()
}

func (s *queryVouchersByIssuer) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
}

func (s *queryVouchersByIssuer) Voucher(a gocuke.DocString) {
	var voucher voucherv1.Voucher
	err := jsonpb.UnmarshalString(a.Content, &voucher)
	require.NoError(s.t, err)

	id, err := s.srv.ss.VoucherTable().InsertReturningID(s.ctx, &voucherv1.Voucher{
		Issuer:   voucher.Issuer,
		Metadata: voucher.Metadata,
	})
	require.NoError(s.t, err)
	require.Equal(s.t, voucher.Id, id)
}

func (s *queryVouchersByIssuer) QueryVouchersByIssuer(a gocuke.DocString) {
	var req v1.QueryVouchersByIssuerRequest
	err := jsonpb.UnmarshalString(a.Content, &req)
	require.NoError(s.t, err)

	s.res, s.err = s.srv.VouchersByIssuer(s.ctx, &req)
}

func (s *queryVouchersByIssuer) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *queryVouchersByIssuer) ExpectTheError(a gocuke.DocString) {
	require.EqualError(s.t, s.err, a.Content)
}

func (s *queryVouchersByIssuer) ExpectResponse(a gocuke.DocString) {
	var expected v1.QueryVouchersByIssuerResponse
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	require.Equal(s.t, &expected, s.res)
}
