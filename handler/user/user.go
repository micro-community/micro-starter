package user

import (
	"context"

	pb "github.com/micro-community/micro-starter/protos"
	"github.com/micro-community/micro-starter/repository"
	"github.com/micro/micro/v3/service"
)

//UserHandler implements the user proto interface,User : people、tenant(orgs、company)
type UserHandler struct {
	Name    string
	user    repository.IUser // instance of the user service
}

// New returns an initUser handler
func NewUser(service *service.Service, user repository.IUser) *UserHandler {
	return &UserHandler{
		Name:    service.Name(),
		user:    user,
	}
}

//GetUser return User By ID
func (u *UserHandler) GetUser(ctx context.Context, req *pb.GetUserRequest, resp *pb.UserInfo) error {
	//	var user models.User
	// pb.ID = req.UserId

	// _, err := u.srv.Login()
	// if err != nil {
	// 	logger.Error(err)
	// 	return err
	// }

	return nil
}

func (u *UserHandler) InsertUser(ctx context.Context, req *pb.InsertUserRequest, resp *pb.InsertUserResponse) error {
	//	var user models.User

	// id, err := pb.Insert()
	// if err != nil {
	// 	logger.Error(err)
	// 	return err
	// }

	// resp.UserId = id

	return nil
}

func (u *UserHandler) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest, resp *pb.DeleteUserResponse) error {
	panic("implement me")

}

func (u *UserHandler) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest, resp *pb.UserInfo) error {
	//	var data models.User

	// result, err := data.Update(req.UserId)
	// if err != nil {
	// 	logger.Error(err)
	// 	return err
	// }
	// resp = result.ToView()

	return nil
}
