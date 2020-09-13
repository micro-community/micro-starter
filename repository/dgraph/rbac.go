package dgraph

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/dgraph-io/dgo/v200"
	"github.com/dgraph-io/dgo/v200/protos/api"
	"github.com/micro/go-micro/v3/logger"

	rbac "github.com/micro-community/auth/protos/rbac"
)

type UID struct {
	UID string `json:"uid"`
}

type Rbac struct {
	dg *dgo.Dgraph
}

func New(dg *dgo.Dgraph) *Rbac {
	return &Rbac{
		dg: dg,
	}
}

// AddUser is a single request handler called via client.AddUser or the generated client code
func (e *Rbac) AddUser(ctx context.Context, req *rbac.User, rsp *rbac.Response) error {
	logger.Infof("Received Rbac.AddUser request, ID: %s, Name: %s", req.Id, req.Name)
	// 首先查询数据库中是否已有该ID
	variables := map[string]string{"$id1": req.Id}
	q := `query Me($id1: string){
		count(func: type(User)) @filter(eq(person.id, $id1)) {
			count(uid)
		}
	}`
	drsp, err := e.dg.NewTxn().QueryWithVars(ctx, q, variables)
	if err != nil {
		return fmt.Errorf("query err: %v", err)
	}

	type Count struct {
		Count int `json:"count"`
	}

	type Root struct {
		Count []Count `json:"count"`
	}
	var r Root
	err = json.Unmarshal(drsp.Json, &r)
	if err != nil || len(r.Count) < 1 {
		return fmt.Errorf("json unmarshal drsp error: %v", err)
	}

	if r.Count[0].Count > 0 {
		return fmt.Errorf("User %s already exists", req.Id)
	}

	// 创建新User
	p := User{
		Uid:    "_:" + req.Id,
		Type:   "User",
		ID:     req.Id,
		Name:   req.Name,
		Age:    req.Age,
		Gender: req.Gender,
	}

	mu := &api.Mutation{
		CommitNow: true,
	}
	pb, err := json.Marshal(p)
	if err != nil {
		logger.Fatal(err)
		return fmt.Errorf("json Marshal error: %v", err)
	}

	mu.SetJson = pb
	result, err := e.dg.NewTxn().Mutate(ctx, mu)
	if err != nil {
		return fmt.Errorf("dgraph Mutate error: %v", err)
	}

	rsp.Msg = fmt.Sprintf("person created, id: %s,  uid: %s", req.Id, result.Uids[req.Id])
	return nil
}

// RemoveUser is a single request handler called via client.RemoveUser or the generated client code
func (e *Rbac) RemoveUser(ctx context.Context, req *rbac.Request, rsp *rbac.Response) error {
	logger.Infof("Received Rbac.RemoveUser request, ID: %s", req.Id)
	// 首先查询数据库中是否已有该ID
	variables := map[string]string{"$id1": req.Id}
	q := `query Me($id1: string){
		find(func: type(User)) @filter(eq(person.id, $id1)) {
			uid
		}
	}`
	drsp, err := e.dg.NewTxn().QueryWithVars(ctx, q, variables)
	if err != nil {
		return fmt.Errorf("query err: %v", err)
	}
	logger.Info(string(drsp.Json))

	type Root struct {
		UID []UID `json:"find"`
	}

	var r Root
	err = json.Unmarshal(drsp.Json, &r)
	if err != nil {
		return fmt.Errorf("json unmarshal Root error: %v", err)
	}

	if len(r.UID) == 0 {
		rsp.Msg = fmt.Sprintf("%s not exists", req.Id)
		return nil
	}

	// mutate multiple items, then commit
	txn := e.dg.NewTxn()
	for _, uid := range r.UID {
		d := map[string]string{"uid": uid.UID}
		logger.Info(d)
		pb, err := json.Marshal(d)
		if err != nil {
			return err
		}
		mu := &api.Mutation{
			DeleteJson: pb,
		}
		drsp, err = txn.Mutate(ctx, mu)
		if err != nil {
			return fmt.Errorf("txn Mutate error: %v", err)
		}
	}
	err = txn.Commit(ctx)
	if err != nil {
		return fmt.Errorf("RemoveUser commit error: %v", err)
	}
	rsp.Msg = "OK"
	return nil
}

// QueryUserRoles is a single request handler called via client.QueryUserRoles or the generated client code
func (e *Rbac) QueryUserRoles(ctx context.Context, req *rbac.Request, rsp *rbac.Roles) error {
	logger.Infof("Received Rbac.QueryUserRoles request, ID: %s", req.Id)
	variables := map[string]string{"$id1": req.Id}
	q := `query Me($id1: string){
		find(func: type(User)) @filter(eq(person.id, $id1)) @normalize {
			role {
				role.id: role.id
				role.name: role.name
			}
		}
	}`
	drsp, err := e.dg.NewTxn().QueryWithVars(ctx, q, variables)
	if err != nil {
		return fmt.Errorf("query err: %v", err)
	}
	type Root struct {
		Role []Role `json:"find"`
	}

	var r Root
	err = json.Unmarshal(drsp.Json, &r)
	if err != nil {
		return fmt.Errorf("json unmarshal Root error: %v", err)
	}

	for _, role := range r.Role {
		rsp.Roles = append(rsp.Roles, &rbac.Role{Id: role.ID, Name: role.Name})
	}
	return nil
}

// QueryUserResources is a single request handler called via client.QueryUserResources or the generated client code
func (e *Rbac) QueryUserResources(ctx context.Context, req *rbac.Request, rsp *rbac.Resources) error {
	logger.Infof("Received Rbac.QueryUserResources request, ID: %s", req.Id)
	variables := map[string]string{"$id1": req.Id}
	q := `query Me($id1: string){
		find(func: type(User)) @filter(eq(person.id, $id1)) @normalize {
			role {
				resource {
					resource.id: resource.id
					resource.name: resource.name
				}
			}
		}
	}`
	drsp, err := e.dg.NewTxn().QueryWithVars(ctx, q, variables)
	if err != nil {
		return fmt.Errorf("query err: %v", err)
	}
	type Root struct {
		Resource []Resource `json:"find"`
	}

	var r Root
	err = json.Unmarshal(drsp.Json, &r)
	if err != nil {
		return fmt.Errorf("json unmarshal Root error: %v", err)
	}

	seen := map[string]bool{}
	for _, res := range r.Resource {
		if !seen[res.ID] {
			seen[res.ID] = true
			rsp.Resources = append(rsp.Resources, &rbac.Resource{Id: res.ID, Name: res.Name})
		}
	}
	return nil
}

// LinkUserRole is a single request handler called via client.LinkUserRole or the generated client code
func (e *Rbac) LinkUserRole(ctx context.Context, req *rbac.LinkRequest, rsp *rbac.Response) error {
	logger.Info("Received Rbac.LinkUserRole request: id1: %s, id2: %s", req.Id1, req.Id2)
	// 首先查询id1 和 id2 对应的 uid
	variables := map[string]string{"$id1": req.Id1, "$id2": req.Id2}
	q := `query Me($id1: string, $id2: string){
		find_id1(func: type(User)) @filter(eq(person.id, $id1)) {
			uid
		}
		find_id2(func: type(Role)) @filter(eq(role.id, $id2)) {
			uid
		}
	}`
	drsp, err := e.dg.NewTxn().QueryWithVars(ctx, q, variables)
	if err != nil {
		return fmt.Errorf("query err: %v", err)
	}
	type Root struct {
		UID1 []UID `json:"find_id1"`
		UID2 []UID `json:"find_id2"`
	}

	var r Root
	err = json.Unmarshal(drsp.Json, &r)
	if err != nil {
		return fmt.Errorf("json unmarshal Root error: %v", err)
	}
	if len(r.UID1) == 0 {
		return fmt.Errorf("id1 <%s> not found", req.Id1)
	}
	if len(r.UID2) == 0 {
		return fmt.Errorf("id2 <%s> not found", req.Id2)
	}

	// link
	mu := &api.Mutation{
		CommitNow: true,
	}

	nq := &api.NQuad{
		Subject:   r.UID1[0].UID,
		Predicate: "role",
		ObjectId:  r.UID2[0].UID,
	}
	mu.Set = []*api.NQuad{nq}
	_, err = e.dg.NewTxn().Mutate(ctx, mu)
	if err != nil {
		return fmt.Errorf("LinkUserRole Mutate error: %v", err)
	}

	rsp.Msg = "OK"
	return nil
}

// UnlinkUserRole is a single request handler called via client.UnlinkUserRole or the generated client code
func (e *Rbac) UnlinkUserRole(ctx context.Context, req *rbac.LinkRequest, rsp *rbac.Response) error {
	logger.Info("Received Rbac.UnlinkUserRole request: id1: %s, id2: %s", req.Id1, req.Id2)
	// 首先查询id1 和 id2 对应的 uid
	variables := map[string]string{"$id1": req.Id1, "$id2": req.Id2}
	q := `query Me($id1: string, $id2: string){
		find_id1(func: type(User)) @filter(eq(person.id, $id1)) {
			uid
		}
		find_id2(func: type(Role)) @filter(eq(role.id, $id2)) {
			uid
		}
	}`
	drsp, err := e.dg.NewTxn().QueryWithVars(ctx, q, variables)
	if err != nil {
		return fmt.Errorf("query err: %v", err)
	}
	type Root struct {
		UID1 []UID `json:"find_id1"`
		UID2 []UID `json:"find_id2"`
	}

	var r Root
	err = json.Unmarshal(drsp.Json, &r)
	if err != nil {
		return fmt.Errorf("json unmarshal Root error: %v", err)
	}
	if len(r.UID1) == 0 {
		return fmt.Errorf("id1 <%s> not found", req.Id1)
	}
	if len(r.UID2) == 0 {
		return fmt.Errorf("id2 <%s> not found", req.Id2)
	}

	// unlink
	mu := &api.Mutation{
		CommitNow: true,
	}

	nq := &api.NQuad{
		Subject:   r.UID1[0].UID,
		Predicate: "role",
		ObjectId:  r.UID2[0].UID,
	}
	mu.Del = []*api.NQuad{nq}
	_, err = e.dg.NewTxn().Mutate(ctx, mu)
	if err != nil {
		return fmt.Errorf("UnlinkUserRole Mutate error: %v", err)
	}

	rsp.Msg = "OK"
	return nil
}

// AddRole is a single request handler called via client.AddRole or the generated client code
func (e *Rbac) AddRole(ctx context.Context, req *rbac.Role, rsp *rbac.Response) error {
	logger.Infof("Received Rbac.AddRole request, ID: %s, Name: %s", req.Id, req.Name)
	// 首先查询数据库中是否已有该ID
	variables := map[string]string{"$id1": req.Id}
	q := `query Me($id1: string){
		count(func: type(Role)) @filter(eq(role.id, $id1)) {
			count(uid)
		}
	}`
	drsp, err := e.dg.NewTxn().QueryWithVars(ctx, q, variables)
	if err != nil {
		return fmt.Errorf("query err: %v", err)
	}

	type Count struct {
		Count int `json:"count"`
	}

	type Root struct {
		Count []Count `json:"count"`
	}
	var r Root
	err = json.Unmarshal(drsp.Json, &r)
	if err != nil || len(r.Count) < 1 {
		return fmt.Errorf("json unmarshal drsp error: %v", err)
	}

	if r.Count[0].Count > 0 {
		return fmt.Errorf("Role %s already exists", req.Id)
	}

	// 创建新Role
	role := Role{
		Uid:  "_:" + req.Id,
		Type: "Role",
		ID:   req.Id,
		Name: req.Name,
	}

	mu := &api.Mutation{
		CommitNow: true,
	}
	pb, err := json.Marshal(role)
	if err != nil {
		logger.Fatal(err)
		return fmt.Errorf("json Marshal error: %v", err)
	}

	mu.SetJson = pb
	result, err := e.dg.NewTxn().Mutate(ctx, mu)
	if err != nil {
		return fmt.Errorf("dgraph Mutate error: %v", err)
	}

	rsp.Msg = fmt.Sprintf("role created, id: %s,  uid: %s", req.Id, result.Uids[req.Id])
	return nil
}

// RemoveRole is a single request handler called via client.RemoveRole or the generated client code
func (e *Rbac) RemoveRole(ctx context.Context, req *rbac.Request, rsp *rbac.Response) error {
	logger.Infof("Received Rbac.RemoveRole request, ID: %s", req.Id)
	// 首先查询数据库中是否已有该ID
	variables := map[string]string{"$id1": req.Id}
	q := `query Me($id1: string){
		find(func: type(Role)) @filter(eq(role.id, $id1)) {
			uid
		}
	}`
	drsp, err := e.dg.NewTxn().QueryWithVars(ctx, q, variables)
	if err != nil {
		return fmt.Errorf("query err: %v", err)
	}
	logger.Info(string(drsp.Json))

	type Root struct {
		UID []UID `json:"find"`
	}

	var r Root
	err = json.Unmarshal(drsp.Json, &r)
	if err != nil {
		return fmt.Errorf("json unmarshal drsp error: %v", err)
	}

	if len(r.UID) == 0 {
		rsp.Msg = fmt.Sprintf("%s not exists", req.Id)
		return nil
	}

	// mutate multiple items, then commit
	txn := e.dg.NewTxn()
	for _, uid := range r.UID {
		d := map[string]string{"uid": uid.UID}
		logger.Info(d)
		pb, err := json.Marshal(d)
		if err != nil {
			return err
		}
		mu := &api.Mutation{
			DeleteJson: pb,
		}
		drsp, err = txn.Mutate(ctx, mu)
		if err != nil {
			return fmt.Errorf("txn Mutate error: %v", err)
		}
	}
	err = txn.Commit(ctx)
	if err != nil {
		return fmt.Errorf("RemoveRole commit error: %v", err)
	}
	rsp.Msg = "OK"
	return nil
}

// QueryRoleResources is a single request handler called via client.QueryRoleResources or the generated client code
func (e *Rbac) QueryRoleResources(ctx context.Context, req *rbac.Request, rsp *rbac.Resources) error {
	logger.Infof("Received Rbac.QueryRoleResources request, ID: %s", req.Id)
	variables := map[string]string{"$id1": req.Id}
	q := `query Me($id1: string){
		find(func: type(Role)) @filter(eq(role.id, $id1)) @normalize {
			resource {
				resource.id: resource.id
				resource.name: resource.name
			}
		}
	}`
	drsp, err := e.dg.NewTxn().QueryWithVars(ctx, q, variables)
	if err != nil {
		return fmt.Errorf("query err: %v", err)
	}
	type Root struct {
		Resource []Resource `json:"find"`
	}

	var r Root
	err = json.Unmarshal(drsp.Json, &r)
	if err != nil {
		return fmt.Errorf("json unmarshal Root error: %v", err)
	}

	for _, res := range r.Resource {
		rsp.Resources = append(rsp.Resources, &rbac.Resource{Id: res.ID, Name: res.Name})
	}
	return nil
}

// LinkRoleResource is a single request handler called via client.LinkRoleResource or the generated client code
func (e *Rbac) LinkRoleResource(ctx context.Context, req *rbac.LinkRequest, rsp *rbac.Response) error {
	logger.Info("Received Rbac.LinkRoleResource request: id1: %s, id2: %s", req.Id1, req.Id2)
	// 首先查询id1 和 id2 对应的 uid
	variables := map[string]string{"$id1": req.Id1, "$id2": req.Id2}
	q := `query Me($id1: string, $id2: string){
		find_id1(func: type(Role)) @filter(eq(role.id, $id1)) {
			uid
		}
		find_id2(func: type(Resource)) @filter(eq(resource.id, $id2)) {
			uid
		}
	}`
	drsp, err := e.dg.NewTxn().QueryWithVars(ctx, q, variables)
	if err != nil {
		return fmt.Errorf("query err: %v", err)
	}
	type Root struct {
		UID1 []UID `json:"find_id1"`
		UID2 []UID `json:"find_id2"`
	}

	var r Root
	err = json.Unmarshal(drsp.Json, &r)
	if err != nil {
		return fmt.Errorf("json unmarshal Root error: %v", err)
	}
	if len(r.UID1) == 0 {
		return fmt.Errorf("id1 <%s> not found", req.Id1)
	}
	if len(r.UID2) == 0 {
		return fmt.Errorf("id2 <%s> not found", req.Id2)
	}

	// link
	mu := &api.Mutation{
		CommitNow: true,
	}

	nq := &api.NQuad{
		Subject:   r.UID1[0].UID,
		Predicate: "resource",
		ObjectId:  r.UID2[0].UID,
	}
	mu.Set = []*api.NQuad{nq}
	_, err = e.dg.NewTxn().Mutate(ctx, mu)
	if err != nil {
		return fmt.Errorf("LinkRoleResource Mutate error: %v", err)
	}

	rsp.Msg = "OK"
	return nil
}

// UnlinkRoleResource is a single request handler called via client.UnlinkRoleResource or the generated client code
func (e *Rbac) UnlinkRoleResource(ctx context.Context, req *rbac.LinkRequest, rsp *rbac.Response) error {
	logger.Info("Received Rbac.UnlinkRoleResource request: id1: %s, id2: %s", req.Id1, req.Id2)
	// 首先查询id1 和 id2 对应的 uid
	variables := map[string]string{"$id1": req.Id1, "$id2": req.Id2}
	q := `query Me($id1: string, $id2: string){
		find_id1(func: type(Role)) @filter(eq(role.id, $id1)) {
			uid
		}
		find_id2(func: type(Resource)) @filter(eq(resource.id, $id2)) {
			uid
		}
	}`
	drsp, err := e.dg.NewTxn().QueryWithVars(ctx, q, variables)
	if err != nil {
		return fmt.Errorf("query err: %v", err)
	}
	type Root struct {
		UID1 []UID `json:"find_id1"`
		UID2 []UID `json:"find_id2"`
	}

	var r Root
	err = json.Unmarshal(drsp.Json, &r)
	if err != nil {
		return fmt.Errorf("json unmarshal Root error: %v", err)
	}
	if len(r.UID1) == 0 {
		return fmt.Errorf("id1 <%s> not found", req.Id1)
	}
	if len(r.UID2) == 0 {
		return fmt.Errorf("id2 <%s> not found", req.Id2)
	}

	// unlink
	mu := &api.Mutation{
		CommitNow: true,
	}

	nq := &api.NQuad{
		Subject:   r.UID1[0].UID,
		Predicate: "resource",
		ObjectId:  r.UID2[0].UID,
	}
	mu.Del = []*api.NQuad{nq}
	_, err = e.dg.NewTxn().Mutate(ctx, mu)
	if err != nil {
		return fmt.Errorf("UnlinkRoleResource Mutate error: %v", err)
	}

	rsp.Msg = "OK"
	return nil
}

// AddResource is a single request handler called via client.AddResource or the generated client code
func (e *Rbac) AddResource(ctx context.Context, req *rbac.Resource, rsp *rbac.Response) error {
	logger.Infof("Received Rbac.AddResource request, ID: %s, Name: %s", req.Id, req.Name)
	// 首先查询数据库中是否已有该ID
	variables := map[string]string{"$id1": req.Id}
	q := `query Me($id1: string){
		count(func: type(Resource)) @filter(eq(resource.id, $id1)) {
			count(uid)
		}
	}`
	drsp, err := e.dg.NewTxn().QueryWithVars(ctx, q, variables)
	if err != nil {
		return fmt.Errorf("query err: %v", err)
	}

	type Count struct {
		Count int `json:"count"`
	}

	type Root struct {
		Count []Count `json:"count"`
	}
	var r Root
	err = json.Unmarshal(drsp.Json, &r)
	if err != nil || len(r.Count) < 1 {
		return fmt.Errorf("json unmarshal drsp error: %v", err)
	}

	if r.Count[0].Count > 0 {
		return fmt.Errorf("Resource %s already exists", req.Id)
	}

	// 创建新Resource
	res := Resource{
		Uid:  "_:" + req.Id,
		Type: "Resource",
		ID:   req.Id,
		Name: req.Name,
	}

	mu := &api.Mutation{
		CommitNow: true,
	}
	pb, err := json.Marshal(res)
	if err != nil {
		logger.Fatal(err)
		return fmt.Errorf("json Marshal error: %v", err)
	}

	mu.SetJson = pb
	result, err := e.dg.NewTxn().Mutate(ctx, mu)
	if err != nil {
		return fmt.Errorf("dgraph Mutate error: %v", err)
	}

	rsp.Msg = fmt.Sprintf("resource created, id: %s,  uid: %s", req.Id, result.Uids[req.Id])
	return nil
}

// RemoveResource is a single request handler called via client.RemoveResource or the generated client code
func (e *Rbac) RemoveResource(ctx context.Context, req *rbac.Request, rsp *rbac.Response) error {
	logger.Infof("Received Rbac.RemoveResource request, ID: %s", req.Id)
	// 首先查询数据库中是否已有该ID
	variables := map[string]string{"$id1": req.Id}
	q := `query Me($id1: string){
		find(func: type(Resource)) @filter(eq(resource.id, $id1)) {
			uid
		}
	}`
	drsp, err := e.dg.NewTxn().QueryWithVars(ctx, q, variables)
	if err != nil {
		return fmt.Errorf("query err: %v", err)
	}
	logger.Info(string(drsp.Json))

	type Root struct {
		UID []UID `json:"find"`
	}

	var r Root
	err = json.Unmarshal(drsp.Json, &r)
	if err != nil {
		return fmt.Errorf("json unmarshal drsp error: %v", err)
	}

	if len(r.UID) == 0 {
		rsp.Msg = fmt.Sprintf("%s not exists", req.Id)
		return nil
	}

	// mutate multiple items, then commit
	txn := e.dg.NewTxn()
	for _, uid := range r.UID {
		d := map[string]string{"uid": uid.UID}
		logger.Info(d)
		pb, err := json.Marshal(d)
		if err != nil {
			return err
		}
		mu := &api.Mutation{
			DeleteJson: pb,
		}
		drsp, err = txn.Mutate(ctx, mu)
		if err != nil {
			return fmt.Errorf("txn Mutate error: %v", err)
		}
	}
	err = txn.Commit(ctx)
	if err != nil {
		return fmt.Errorf("RemoveResource commit error: %v", err)
	}
	rsp.Msg = "OK"
	return nil
}
