package handler

import (
	handlerInterface "Clean-Grpc/pkg/api/handler/interface"
	pb "Clean-Grpc/pkg/api/proto"
	"Clean-Grpc/pkg/domain"
	services "Clean-Grpc/pkg/usecase/interface"
	"context"
	"fmt"
	"net/http"

)

type UserHandler struct {
	userUseCase services.UserUseCase
}

func NewUserHandler(usecase services.UserUseCase) handlerInterface.UserHandler {
	return &UserHandler{
		userUseCase: usecase,
	}
}

func (cr *UserHandler) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.ResgisterResponse, error) {
	user := domain.Users{
		Email:    req.Email,
		Password: req.Password,
	}

	user, err := cr.userUseCase.Save(ctx, user)
	fmt.Println(user)
	if err != nil {
		return nil, err
	}

	return &pb.ResgisterResponse{
		Status: http.StatusCreated,
	}, nil
}



// func (cr *UserHandler) Delete(c *gin.Context) {
// 	paramsId := c.Param("id")
// 	id, err := strconv.Atoi(paramsId)

// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"error": "Cannot parse id",
// 		})
// 		return
// 	}

// 	ctx := c.Request.Context()
// 	user, err := cr.userUseCase.FindByID(ctx, uint(id))

// 	if err != nil {
// 		c.AbortWithStatus(http.StatusNotFound)
// 	}

// 	if user == (domain.Users{}) {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": "User is not booking yet",
// 		})
// 		return
// 	}

// 	cr.userUseCase.Delete(ctx, user)

// 	c.JSON(http.StatusOK, gin.H{"message": "User is deleted successfully"})
// }


func (cr *UserHandler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	// var user domain.Users

	// if result := s.H.DB.Where(&models.User{Email: req.Email}).First(&user); result.Error != nil {
	// 	return &pb.LoginResponse{
	// 		Status: http.StatusNotFound,
	// 		Error:  "User not found",
	// 	}, nil
	// }

	// match := utils.CheckPasswordHash(req.Password, user.Password)

	// if !match {
	// 	return &pb.LoginResponse{
	// 		Status: http.StatusNotFound,
	// 		Error:  "User not found",
	// 	}, nil
	// }

	// token, _ := s.Jwt.GenerateToken(user)

	return &pb.LoginResponse{
		Status: http.StatusOK,
		Token:  "token",
	}, nil
}

