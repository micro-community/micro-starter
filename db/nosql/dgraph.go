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

type DormDB struct {
	dg *dgo.Dgraph
}

func NewDGraphClient(cfg *DgraphCfg) *DormDB {

	dialOpts := append([]grpc.DialOption{},
		grpc.WithInsecure(),
		grpc.WithDefaultCallOptions(grpc.UseCompressor(gzip.Name)))
	conn, err := grpc.Dial(cfg.Url, dialOpts...)
	if err != nil {
		logger.Fatal("While trying to dial gRPC")
	}

	dClient := dgo.NewDgraphClient(
		api.NewDgraphClient(conn),
	)

	return &DormDB{dg: dClient}
}

func (d *DormDB) txn() *dgo.Txn {
	return d.dg.NewTxn()
}

func (d *DormDB) QueryReadOnly(q string) (*api.Response, error) {
	return d.dg.NewReadOnlyTxn().Query(context.Background(), q)
}

func (d *DormDB) Query(q string) (*api.Response, error) {
	return d.txn().Query(context.Background(), q)
}

//Query2ID  ..
func (d *DormDB) Query2ID(id1, id2, queryString string) (*api.Response, error) {
	// Assigned uids for nodes which were created would be returned in the resp.AssignedUids map.
	variables := map[string]string{"$id1": id2, "$id2": id2}
	resp, err := d.txn().QueryWithVars(context.Background(), queryString, variables)
	if err != nil {
		logger.Fatal("query id1: %s id2: %s with error ", id1, id2, err)
	}

	return resp, err
}

//QueryID  to query a id..
func (d *DormDB) QueryID(targetID, queryString string) (*api.Response, error) {
	// Assigned uids for nodes which were created would be returned in the resp.AssignedUids map.
	variables := map[string]string{"$id": targetID}
	resp, err := d.txn().QueryWithVars(context.Background(), queryString, variables)
	if err != nil {
		logger.Fatal(err)
	}

	return resp, err

}

//MutateObject is under writing
func (d *DormDB) MutateObject(typestruct interface{}) (*api.Response, error) {

	pb, err := json.Marshal(typestruct)
	if err != nil {
		logger.Errorf("json Marshal error: %v", err)
	}
	return d.Mutate(pb)
	// 	if err != nil {
	// 		logger.Fatalf("dgraph Mutate error: %v", err)
	// 	}
}

func (d *DormDB) UpdateRelationShip(subject, predicate, object string, isSetRelationShip bool) (*api.Response, error) {

	mu := &api.Mutation{
		CommitNow: true,
	}

	nq := &api.NQuad{
		Subject:   subject,
		Predicate: predicate,
		ObjectId:  object,
	}

	if isSetRelationShip {
		mu.Set = []*api.NQuad{nq}
	} else {
		mu.Del = []*api.NQuad{nq}
	}

	return d.txn().Mutate(context.Background(), mu)
}

func (d *DormDB) Mutate(b []byte) (*api.Response, error) {

	mu := &api.Mutation{
		CommitNow: true,
	}

	mu.SetJson = b
	return d.txn().Mutate(context.Background(), mu)
}

func (d *DormDB) BatchDelete(uids []string) error {

	ctx := context.Background()

	for _, uid := range uids {
		data := map[string]string{"uid": uid}
		//	logger.Info(data)
		pb, err := json.Marshal(data)
		if err != nil {
			return err
		}
		_, err = d.txn().Mutate(ctx, &api.Mutation{DeleteJson: pb})
		if err != nil {
			return err
		}
	}

	return d.txn().Commit(ctx)

}

func (d *DormDB) Delete(b []byte) error {

	fmt.Println(string(b))

	mu := &api.Mutation{
		CommitNow:  true,
		DeleteJson: b,
	}
	resp, err := d.txn().Mutate(context.Background(), mu)

	fmt.Println(string(resp.Json))

	return err
}

func (d *DormDB) Update(set string) error {

	mu := &api.Mutation{
		CommitNow: true,
		SetNquads: []byte(set),
	}
	_, err := d.txn().Mutate(context.Background(), mu)
	//fmt.Println(string(resp.Json))
	return err
}

func (d *DormDB) UpdateWithQuery(query, set string) error {

	mu := &api.Mutation{
		CommitNow: true,
	}

	req := &api.Request{CommitNow: true}
	req.Query = query

	mu.SetNquads = []byte(set)
	req.Mutations = []*api.Mutation{mu}
	resp, err := d.txn().Do(context.Background(), req)

	fmt.Println(string(resp.Json))

	return err
}
