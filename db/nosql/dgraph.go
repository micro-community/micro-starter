package nosql

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/dgraph-io/dgo/v200"
	"github.com/dgraph-io/dgo/v200/protos/api"
	"github.com/micro/go-micro/v3/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/encoding/gzip"
)

//DgraphCfg for DgraphDB
type DgraphCfg struct {
	User     string
	Password string
	Host     string
	Port     int
	DBName   string
	Url      string
}

var (
	dg             *dgo.Dgraph
	dgraphConnHost = "127.0.0.1:9080"
)

func newDGraphClient() *dgo.Dgraph {

	dialOpts := append([]grpc.DialOption{},
		grpc.WithInsecure(),
		grpc.WithDefaultCallOptions(grpc.UseCompressor(gzip.Name)))

	conn, err := grpc.Dial(dgraphConnHost, dialOpts...)
	if err != nil {
		logger.Fatal("While trying to dial gRPC")
	}

	return dgo.NewDgraphClient(
		api.NewDgraphClient(conn),
	)
}

func QueryReadOnly(q string) (*api.Response, error) {

	if dg == nil {
		dg = newDGraphClient()
	}
	return dg.NewReadOnlyTxn().Query(context.Background(), q)
}

func Query(q string) (*api.Response, error) {

	if dg == nil {
		dg = newDGraphClient()
	}
	return dg.NewTxn().Query(context.Background(), q)
}

//QueryWithVar is under writing
func QueryWithVar(typestruct interface{}, target, queryString string) (*api.Response, error) {

	if dg == nil {
		dg = newDGraphClient()
	}

	mu := &api.Mutation{
		CommitNow: true,
	}
	pb, err := json.Marshal(typestruct)
	if err != nil {
		logger.Fatal(err)
	}

	mu.SetJson = pb

	ctx := context.Background()

	assigned, err := dg.NewTxn().Mutate(ctx, mu)
	if err != nil {
		logger.Fatal(err)
	}

	// Assigned uids for nodes which were created would be returned in the resp.AssignedUids map.
	variables := map[string]string{"$id": assigned.Uids[target]}

	resp, err := dg.NewTxn().QueryWithVars(ctx, queryString, variables)
	if err != nil {
		logger.Fatal(err)
	}

	return resp, err

}

func Mutate(b []byte) (*api.Response, error) {
	if dg == nil {
		dg = newDGraphClient()
	}

	mu := &api.Mutation{
		CommitNow: true,
	}

	mu.SetJson = b
	return dg.NewTxn().Mutate(context.Background(), mu)
}

func Delete(b []byte) error {
	if dg == nil {
		dg = newDGraphClient()
	}

	fmt.Println(string(b))

	mu := &api.Mutation{
		CommitNow:  true,
		DeleteJson: b,
	}
	resp, err := dg.NewTxn().Mutate(context.Background(), mu)

	fmt.Println(string(resp.Json))

	return err
}

func Update(set string) error {
	if dg == nil {
		dg = newDGraphClient()
	}
	mu := &api.Mutation{
		CommitNow: true,
		SetNquads: []byte(set),
	}
	_, err := dg.NewTxn().Mutate(context.Background(), mu)
	//fmt.Println(string(resp.Json))
	return err
}

func UpdateWithQuery(query, set string) error {
	if dg == nil {
		dg = newDGraphClient()
	}
	mu := &api.Mutation{
		CommitNow: true,
	}

	req := &api.Request{CommitNow: true}
	req.Query = query

	mu.SetNquads = []byte(set)
	req.Mutations = []*api.Mutation{mu}
	resp, err := dg.NewTxn().Do(context.Background(), req)

	fmt.Println(string(resp.Json))

	return err
}
