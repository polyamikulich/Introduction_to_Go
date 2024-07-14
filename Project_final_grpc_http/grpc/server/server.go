package server

import (
	"Go_Project/accounts/models"
	"Go_Project/proto"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"sync"
)

func New() *Server {
	return &Server{
		accounts: make(map[string]*models.Account),
		guard:    &sync.RWMutex{},
	}
}

type Server struct {
	accounts map[string]*models.Account
	guard    *sync.RWMutex
	proto    proto.UnimplementedAccountsServer
}

func (s *Server) CreateAccount(ctx context.Context, in *proto.CreateAccountReq) (*proto.Empty, error) {

	s.guard.Lock()

	if _, ok := s.accounts[in.Username]; ok {
		s.guard.Unlock()

		return nil, status.Errorf(codes.AlreadyExists, "account %s already exist", in.Username)
	}

	s.accounts[in.Username] = &models.Account{
		Name:   in.Username,
		Amount: int(in.Amount),
	}

	s.guard.Unlock()

	return nil, nil
}

func (s *Server) GetAccount(ctx context.Context, in *proto.GetAccountReq) (*proto.GetAccountResp, error) {

	s.guard.RLock()

	account, ok := s.accounts[in.Username]

	s.guard.RUnlock()

	if !ok {
		return nil, status.Errorf(codes.NotFound, "account %s not found", in.Username)
	}

	response := proto.GetAccountResp{
		Username: account.Name,
		Amount:   int32(account.Amount),
	}

	return &response, status.Errorf(codes.OK, "")
}

func (s *Server) DeleteAccount(ctx context.Context, in *proto.DeleteAccountReq) (*proto.Empty, error) {
	s.guard.RLock()

	if _, ok := s.accounts[in.Username]; !ok {
		return nil, status.Errorf(codes.NotFound, "account %s not found", in.Username)
	}

	s.guard.RUnlock()

	delete(s.accounts, in.Username)

	return nil, status.Errorf(codes.OK, "account %s is deleted", in.Username)
}

func (s *Server) PatchAccount(ctx context.Context, in *proto.PatchAccountReq) (*proto.Empty, error) {
	if len(in.Username) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "empty name")
	}
	if in.Amount == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "empty amount")
	}

	s.guard.Lock()

	account, ok := s.accounts[in.Username]

	s.guard.Unlock()

	if !ok {
		return nil, status.Errorf(codes.NotFound, "account %s not found", in.Username)
	}

	account.Amount += int(in.Amount)

	return nil, status.Errorf(codes.OK, "account %s is patched", in.Username)
}

func (s *Server) ChangeAccount(ctx context.Context, in *proto.ChangeAccountReq) (*proto.Empty, error) {
	if len(in.LastName) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "empty name")
	}
	if len(in.NewName) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "empty new name")
	}

	s.guard.Lock()

	account, ok := s.accounts[in.LastName]

	s.guard.Unlock()

	if !ok {
		return nil, status.Errorf(codes.NotFound, "account %s not found", in.LastName)
	}

	if _, ok := s.accounts[in.NewName]; ok {
		return nil, status.Errorf(codes.AlreadyExists, "account with name %s is already exist", in.NewName)
	}

	account.Name = in.NewName
	delete(s.accounts, in.LastName)
	s.accounts[in.NewName] = account

	return nil, status.Errorf(codes.OK, "account with name %s is changed on %s", in.LastName, in.NewName)
}
