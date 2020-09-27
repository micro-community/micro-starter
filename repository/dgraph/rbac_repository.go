package dgraph

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/dgraph-io/dgo/v200/protos/api"
	"github.com/micro-community/auth/db"
	"github.com/micro-community/auth/models"
	"github.com/micro/go-micro/v3/logger"
)

type RbacRepository struct {
}

func NewRBACRepository() *RbacRepository {
	return &RbacRepository{}
}

//QueryUserExist check user
func (e *RbacRepository) QueryUserExist(targetID int64) (*api.Response, error) {
	queryString := `query Me($id1: string){
		find(func: type(User)) @filter(eq(person.id, $id1)) {
			uid
		}
	}`
	target := fmt.Sprintf("%d", targetID)
	// Assigned uids for nodes which were created would be returned in the resp.AssignedUids map.
	//variables := map[string]string{"$id1": target}
	rsp, err := db.DDB().QueryID(target, queryString)
	if err != nil {
		return nil, fmt.Errorf("query role err: %v", err)
	}
	return rsp, nil

}

// AddUser is a single request handler called via client.AddUser or the generated client code
func (e *RbacRepository) AddUser(ctx context.Context, user *models.User) error {
	logger.Infof("Received RbacRepository.AddUser request, ID: %d, Name: %s", user.ID, user.Name)
	//首先查询数据库中是否已有该ID

	drsp, err := e.QueryUserExist(user.ID)
	if err != nil {
		return fmt.Errorf("query err: %v", err)
	}

	var r Root
	err = json.Unmarshal(drsp.Json, &r)
	if err != nil || len(r.Count) < 1 {
		return fmt.Errorf("json unmarshal drsp error: %v", err)
	}

	if r.Count[0].Count > 0 {
		return fmt.Errorf("User %d already exists", user.ID)
	}

	// 创建新User
	p := models.User{
		//Uid:    "_:" + target,
		Type:   "User",
		ID:     user.ID,
		Name:   user.Name,
		Age:    user.Age,
		Gender: user.Gender,
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
	_, err = db.DDB().Mutate(pb)
	if err != nil {
		return fmt.Errorf("dgraph Mutate error: %v", err)
	}

	return nil
}

// RemoveUser is a single request handler called via client.RemoveUser or the generated client code
func (e *RbacRepository) RemoveUser(ctx context.Context, user *models.User) error {
	logger.Infof("Received RbacRepository.RemoveUser request, ID: %d", user.ID)
	// 首先查询数据库中是否已有该ID

	drsp, err := e.QueryUserExist(user.ID)

	var r Root
	err = json.Unmarshal(drsp.Json, &r)
	if err != nil {
		return fmt.Errorf("json unmarshal Root error: %v", err)
	}
	err = db.DDB().BatchDelete(r.UID)
	if err != nil {
		return fmt.Errorf("RemoveUser commit error: %v", err)
	}
	return nil
}

// QueryUserRoles is a single request handler called via client.QueryUserRoles or the generated client code
func (e *RbacRepository) QueryUserRoles(ctx context.Context, role *models.Role) ([]*models.Role, error) {
	logger.Infof("Received RbacRepository.QueryUserRoles request, ID: %d", role.ID)

	targetID := fmt.Sprintf("%d", role.ID)
	//	variables := map[string]string{"$id": targetID}
	q := `query Me($id: string){
		roles(func: type(User)) @filter(eq(person.id, $id)) @normalize {
			role {
				id
			  name
			}
		}
	}`

	drsp, err := db.DDB().QueryWithVar(targetID, q)
	if err != nil {
		return nil, fmt.Errorf("query user role err: %v", err)
	}
	var roles []models.Role
	err = json.Unmarshal(drsp.Json, &roles)
	if err != nil {
		return nil, fmt.Errorf("json unmarshal roles error: %v", err)
	}
	var resRoles []*models.Role
	for _, role := range roles {
		resRoles = append(resRoles, &models.Role{ID: role.ID, Name: role.Name})
	}
	return resRoles, nil
}

// QueryUserResources is a single request handler called via client.QueryUserResources or the generated client code
func (e *RbacRepository) QueryUserResources(ctx context.Context, user *models.User) ([]*models.Resource, error) {
	logger.Infof("Received RbacRepository.QueryUserResources request, ID: %d", user.ID)

	targetID := fmt.Sprintf("%d", user.ID)
	//variables := map[string]string{"$id1": user.ID}

	q := `query Me($id1: string){
		resources(func: type(User)) @filter(eq(person.id, $id1)) @normalize {
			role {
				resource {
					id
					name
				}
			}
		}
	}`

	drsp, err := db.DDB().QueryWithVar(targetID, q)
	if err != nil {
		return nil, fmt.Errorf("query err: %v", err)
	}
	type Root struct {
		Resources []models.Resource `json:"Resource"`
	}

	var r Root
	err = json.Unmarshal(drsp.Json, &r)
	if err != nil {
		return nil, fmt.Errorf("json unmarshal Root error: %v", err)
	}

	//过滤重复
	seen := map[int]bool{}
	resRoles := []*models.Resource{}
	for _, res := range r.Resources {
		if !seen[res.ID] {
			seen[res.ID] = true
			resRoles = append(resRoles, &models.Resource{ID: res.ID, Name: res.Name})
		}
	}
	return resRoles, nil
}

// LinkUserRole is a single request handler called via client.LinkUserRole or the generated client code
func (e *RbacRepository) LinkUserRole(ctx context.Context, user *models.User, role *models.Role) error {
	logger.Info("Received RbacRepository.LinkUserRole request: user: %d, role: %d", user.ID, role.ID)

	userid := fmt.Sprintf("%d", user.ID)
	roleid := fmt.Sprintf("%d", role.ID)

	// 首先查询user id 和 role 对应的 id
	//variables := map[string]string{"$id1": userid, "$id2": roleid}
	q := `query Me($id1: string, $rid2id: string){
		user(func: type(User)) @filter(eq(person.id, $id1)) {
			uid
		}
		role(func: type(Role)) @filter(eq(role.id, $id2)) {
			uid
		}
	}`
	drsp, err := db.DDB().Query2ID(userid, roleid, q)
	if err != nil {
		return fmt.Errorf("query err: %v", err)
	}
	type Root struct {
		UID1 []models.UID `json:"user"`
		UID2 []models.UID `json:"role"`
	}

	var r Root
	err = json.Unmarshal(drsp.Json, &r)
	if err != nil {
		return fmt.Errorf("json unmarshal Root error: %v", err)
	}
	if len(r.UID1) == 0 {
		return fmt.Errorf("user id <%d> not found", user.ID)
	}
	if len(r.UID2) == 0 {
		return fmt.Errorf("role id <%d> not found", role.ID)
	}
	_, err = db.DDB().UpdateRelationShip(r.UID1[0].UID, "role", r.UID2[0].UID, true)
	if err != nil {
		return fmt.Errorf("LinkUserRole Mutate error: %v", err)
	}

	//	rsp.Msg = "OK"
	return nil
}

// UnlinkUserRole is a single request handler called via client.UnlinkUserRole or the generated client code
func (e *RbacRepository) UnlinkUserRole(ctx context.Context, user *models.User, role *models.Role) error {
	logger.Info("Received RbacRepository.UnlinkUserRole request: user: %d, role: %d", user.ID, role.ID)

	userid := fmt.Sprintf("%d", user.ID)
	roleid := fmt.Sprintf("%d", role.ID)

	// 首先查询user id 和 role 对应的 id
	//variables := map[string]string{"$id1": userid, "$id2": roleid}
	q := `query Me($id1: string, $id2: string){
		find_id1(func: type(User)) @filter(eq(person.id, $id1)) {
			uid
		}
		find_id2(func: type(Role)) @filter(eq(role.id, $id2)) {
			uid
		}
	}`
	drsp, err := db.DDB().Query2ID(userid, roleid, q)
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
		return fmt.Errorf("user id <%d> not found", user.ID)
	}
	if len(r.UID2) == 0 {
		return fmt.Errorf("role id <%d> not found", role.ID)
	}
	_, err = db.DDB().UpdateRelationShip(r.UID1[0].UID, "role", r.UID2[0].UID, false)
	if err != nil {
		return fmt.Errorf("LinkUserRole Mutate error: %v", err)
	}
	//	rsp.Msg = "OK"
	return nil
}

//QueryRoleExist is under writing
func (e *RbacRepository) QueryRoleExist(targetID int) ([]string, error) {
	queryString := `query Me($id1: string){
		count(func: type(Role)) @filter(eq(role.id, $id1)) {
			count(uid)
		}
	}`
	target := fmt.Sprintf("%d", targetID)
	// Assigned uids for nodes which were created would be returned in the resp.AssignedUids map.
	//variables := map[string]string{"$id1": target}
	rsp, err := db.DDB().QueryID(target, queryString)
	if err != nil {
		return nil, fmt.Errorf("query role err: %v", err)
	}
	type Root struct {
		UID []models.UID `json:"role"`
	}
	var r Root

	err = json.Unmarshal(rsp.Json, &r)
	if err != nil {
		return nil, fmt.Errorf("json unmarshal drsp error: %v", err)
	}
	if len(r.UID) > 0 {
		return nil, fmt.Errorf("Role %d already exists", targetID)
	}
	var resultUIDs []string
	for _, uid := range r.UID {
		resultUIDs = append(resultUIDs, uid.UID)
	}
	return resultUIDs, nil
}

// AddRole is a single request handler called via client.AddRole or the generated client code
func (e *RbacRepository) AddRole(ctx context.Context, role *models.Role) error {
	logger.Infof("Received RbacRepository.AddRole request, ID: %d, Name: %d", role.ID, role.Name)
	_, err := e.QueryRoleExist(role.ID)
	if err != nil {
		return fmt.Errorf("query err: %v", err)
	}
	// 创建新Role
	newRole := &models.Role{
		//	Uid:  "_:" + role.ID,
		Type: "Role",
		ID:   role.ID,
		Name: role.Name,
	}

	_, err = db.DDB().MutateObject(newRole)
	if err != nil {
		return fmt.Errorf("dgraph Mutate error: %v", err)
	}

	//rsp.Msg = fmt.Sprintf("role created, id: %s,  uid: %s", req.Id, result.Uids[req.Id])
	return nil
}

// RemoveRole is a single request handler called via client.RemoveRole or the generated client code
func (e *RbacRepository) RemoveRole(ctx context.Context, role *models.Role) error {

	logger.Infof("Received RbacRepository.RemoveRole request, ID: %d", role.ID)
	// 首先查询数据库中是否已有该ID
	uids, err := e.QueryRoleExist(role.ID)
	if err != nil {
		return fmt.Errorf("query err: %v", err)
	}
	// mutate multiple items, then commit
	db.DDB().BatchDelete(uids)
	return nil
}

// QueryRoleResources is a single request handler called via client.QueryRoleResources or the generated client code
func (e *RbacRepository) QueryRoleResources(ctx context.Context, role *models.Role, resource *models.Resource) ([]*models.Resource, error) {
	logger.Infof("Received RbacRepository.QueryRoleResources request, ID: %s", resource.ID)

	roleID := fmt.Sprintf("%d", role.ID)
	//variables := map[string]string{"$id1": roleID}
	q := `query Me($id: string){
		role(func: type(Role)) @filter(eq(role.id, $id)) @normalize {
			resource {
				resource.id
				resource.name
			}
		}
	}`
	drsp, err := db.DDB().QueryWithVar(roleID, q)
	//	drsp, err := db.DDB().QueryWithVars(ctx, q, variables)
	if err != nil {
		return nil, fmt.Errorf("query err: %v", err)
	}
	type Root struct {
		Resource []models.Resource `json:"find"`
	}

	var r Root
	err = json.Unmarshal(drsp.Json, &r)
	if err != nil {
		return nil, fmt.Errorf("json unmarshal Root error: %v", err)
	}

	var resResource []*models.Resource

	for _, res := range r.Resource {
		resResource = append(resResource, &models.Resource{ID: res.ID, Name: res.Name})
	}
	return resResource, nil
}

// LinkRoleResource is a single request handler called via client.LinkRoleResource or the generated client code
func (e *RbacRepository) LinkRoleResource(ctx context.Context, role *models.Role, resource *models.Resource) error {
	logger.Info("Received RbacRepository.LinkRoleResource request: id1: %d, id2: %d", role.ID, resource.ID)
	// 首先查询id1 和 id2 对应的 uid

	roleid := fmt.Sprintf("%d", role.ID)
	resourceid := fmt.Sprintf("%d", resource.ID)
	//	variables := map[string]string{"$id1": role.ID, "$id2": req.ID}

	q := `query Me($id1: string, $id2: string){
		roles(func: type(Role)) @filter(eq(role.id, $id1)) {
			uid
		}
		resources(func: type(Resource)) @filter(eq(resource.id, $id2)) {
			uid
		}
	}`
	drsp, err := db.DDB().Query2ID(roleid, resourceid, q)
	if err != nil {
		return fmt.Errorf("query err: %v", err)
	}
	type Root struct {
		UID1 []models.UID `json:"roles"`
		UID2 []UID        `json:"resources"`
	}

	var r Root
	err = json.Unmarshal(drsp.Json, &r)
	if err != nil {
		return fmt.Errorf("json unmarshal Root error: %v", err)
	}
	if len(r.UID1) == 0 {
		return fmt.Errorf("id1 <%s> not found", roleid)
	}
	if len(r.UID2) == 0 {
		return fmt.Errorf("id2 <%s> not found", resourceid)
	}

	_, err = db.DDB().UpdateRelationShip(r.UID1[0].UID, "resource", r.UID2[0].UID, true)
	if err != nil {
		return fmt.Errorf("LinkRoleResource Mutate error: %v", err)
	}

	//rsp.Msg = "OK"
	return nil
}

// UnlinkRoleResource is a single request handler called via client.UnlinkRoleResource or the generated client code
func (e *RbacRepository) UnlinkRoleResource(ctx context.Context, role *models.Role, resource *models.Resource) error {
	logger.Info("Received RbacRepository.UnlinkRoleResource request: id1: %d, id2: %d", role.ID, resource.ID)
	// 首先查询id1 和 id2 对应的 uid

	roleid := fmt.Sprintf("%d", role.ID)
	resourceid := fmt.Sprintf("%d", resource.ID)

	//variables := map[string]string{"$id1": roleid, "$id2": resourceid}
	q := `query Me($id1: string, $id2: string){
		find_id1(func: type(Role)) @filter(eq(role.id, $id1)) {
			uid
		}
		find_id2(func: type(Resource)) @filter(eq(resource.id, $id2)) {
			uid
		}
	}`

	drsp, err := db.DDB().Query2ID(roleid, resourceid, q)
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
		return fmt.Errorf("id1 <%s> not found", roleid)
	}
	if len(r.UID2) == 0 {
		return fmt.Errorf("id2 <%s> not found", resourceid)
	}

	_, err = db.DDB().UpdateRelationShip(r.UID1[0].UID, "resource", r.UID2[0].UID, false)
	if err != nil {
		return fmt.Errorf("UnlinkRoleResource Mutate error: %v", err)
	}

	//rsp.Msg = "OK"
	return nil
}

//QueryResourceExist check resource
func (e *RbacRepository) QueryResourceExist(targetID int) ([]string, error) {
	queryString := `query Me($id1: string){
		count(func: type(Resource)) @filter(eq(resource.id, $id1)) {
			count(uid)
		}
	}`
	target := fmt.Sprintf("%d", targetID)
	// Assigned uids for nodes which were created would be returned in the resp.AssignedUids map.
	//variables := map[string]string{"$id1": target}
	rsp, err := db.DDB().QueryID(target, queryString)
	if err != nil {
		return nil, fmt.Errorf("query role err: %v", err)
	}

	var r Root
	err = json.Unmarshal(rsp.Json, &r)
	if err != nil || len(r.Count) < 1 {
		return nil, fmt.Errorf("json unmarshal drsp error: %v", err)
	}
	if r.Count[0].Count > 0 {
		return nil, fmt.Errorf("Resource %s already exists", target)
	}
	return r.UID, nil

}

// AddResource is a single request handler called via client.AddResource or the generated client code
func (e *RbacRepository) AddResource(ctx context.Context, resource *models.Resource) error {
	logger.Infof("Received RbacRepository.AddResource request, ID: %d, Name: %s", resource.ID, resource.Name)

	// 首先查询数据库中是否已有该ID

	ids, err := e.QueryResourceExist(resource.ID)

	if err != nil {
		return fmt.Errorf("query err: %v", err)
	}
	// 创建新Resource
	res := models.Resource{
		Uid: "_:" + ids[0],
		//	Type: "Resource",
		ID:   resource.ID,
		Name: resource.Name,
	}
	_, err = db.DDB().MutateObject(res)
	if err != nil {
		return fmt.Errorf("dgraph Mutate error: %v", err)
	}

	return nil
}

// RemoveResource is a single request handler called via client.RemoveResource or the generated client code
func (e *RbacRepository) RemoveResource(ctx context.Context, resource *models.Resource) error {
	logger.Infof("Received RbacRepository.RemoveResource request, ID: %d", resource.ID)
	// 首先查询数据库中是否已有该ID

	ids, err := e.QueryResourceExist(resource.ID)
	if err != nil {
		return fmt.Errorf("query err: %v", err)
	}

	err = db.DDB().BatchDelete(ids)
	if err != nil {
		return fmt.Errorf("RemoveResource commit error: %v", err)
	}
	//	rsp.Msg = "OK"
	return nil
}
