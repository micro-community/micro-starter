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

//QueryExist is under writing
func (d *DormDB) QueryExist(targetID int64) (*api.Response, error) {

	queryString := `query Me($id1: string){
		find(func: type(User)) @filter(eq(person.id, $id1)) {
			uid
		}
	}`

	target := fmt.Sprintf("%d", targetID)
	// Assigned uids for nodes which were created would be returned in the resp.AssignedUids map.
	variables := map[string]string{"$id1": target}

	resp, err := d.txn().QueryWithVars(context.Background(), queryString, variables)
	if err != nil {
		logger.Fatal(err)
	}

	return resp, err

}

//QueryWithVar  ..
func (d *DormDB) QueryWithVar(targetID, queryString string) (*api.Response, error) {

	// Assigned uids for nodes which were created would be returned in the resp.AssignedUids map.
	variables := map[string]string{"$id": targetID}

	resp, err := d.txn().QueryWithVars(context.Background(), queryString, variables)
	if err != nil {
		logger.Fatal(err)
	}

	return resp, err

}

//MutateObject is under writing
func (d *DormDB) MutateObject(typestruct interface{}, target, queryString string) (*api.Response, error) {

	pb, err := json.Marshal(typestruct)
	if err != nil {
		logger.Errorf("json Marshal error: %v", err)
	}
	return d.Mutate(pb)
	// 	if err != nil {
	// 		logger.Fatalf("dgraph Mutate error: %v", err)
	// 	}
}

func (d *DormDB) Mutate(b []byte) (*api.Response, error) {

	mu := &api.Mutation{
		CommitNow: true,
	}

	mu.SetJson = b
	return d.txn().Mutate(context.Background(), mu)
}

func (d *DormDB) BatchDelete(uids []string) error {

	
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
