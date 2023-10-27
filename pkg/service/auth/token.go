package auth

import (
  "x-user-service/protos/auth"
  "x-user-service/pkg/database"
)

type TokenService struct {
  DBContext database.DBContext
}

func (handler *TokenService) GenerateToken() (*auth.AuthenticationResponse, error) {
  // get data from database
  db := handler.DBContext
  var user, _ = db.User.GetById(1)
  return &auth.AuthenticationResponse{
    Token:  user.IsActive,
    Iat:    user.CreatedAt.Unix(),
    Exp:    user.UpdatedAt.Unix(),
    Role:   "admin",
    Email:  user.Email,
    UserId: "id",
  }, nil
}
