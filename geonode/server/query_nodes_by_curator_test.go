package server

import (
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	geonodev1 "github.com/choraio/mods/geonode/api/v1"
	v1 "github.com/choraio/mods/geonode/types/v1"
)

type queryNodesByCurator struct {
	*baseSuite
	res *v1.QueryNodesByCuratorResponse
	err error
}

func TestQueryNodesByCurator(t *testing.T) {
	gocuke.NewRunner(t, &queryNodesByCurator{}).
		Path("./features/query_nodes_by_curator.feature").
		Run()
}

func (s *queryNodesByCurator) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
}

func (s *queryNodesByCurator) Node(a gocuke.DocString) {
	var node geonodev1.Node
	err := jsonpb.UnmarshalString(a.Content, &node)
	require.NoError(s.t, err)

	id, err := s.srv.ss.NodeTable().InsertReturningID(s.ctx, &geonodev1.Node{
		Curator:  node.Curator,
		Metadata: node.Metadata,
	})
	require.NoError(s.t, err)
	require.Equal(s.t, node.Id, id)
}

func (s *queryNodesByCurator) QueryNodesByCurator(a gocuke.DocString) {
	var req v1.QueryNodesByCuratorRequest
	err := jsonpb.UnmarshalString(a.Content, &req)
	require.NoError(s.t, err)

	s.res, s.err = s.srv.NodesByCurator(s.ctx, &req)
}

func (s *queryNodesByCurator) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *queryNodesByCurator) ExpectTheError(a gocuke.DocString) {
	require.EqualError(s.t, s.err, a.Content)
}

func (s *queryNodesByCurator) ExpectResponse(a gocuke.DocString) {
	var expected v1.QueryNodesByCuratorResponse
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	require.Equal(s.t, &expected, s.res)
}
