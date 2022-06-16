package gokit

import (
	"context"
	"encoding/json"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

type Service interface {
	Hash(ctx context.Context, password string) (string, error)
	Validate(ctx context.Context, password, hash string) (bool, error)
}

// 无状态实现类
type vaultService struct{}

// 返回service实例（动态类型）
func NewService() Service {
	return vaultService{}
}

func (vaultService) Hash(ctx context.Context, password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (vaultService) Validate(ctx context.Context, password, hash string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false, err
	}
	return true, nil
}

/**
Hash方法的请求与响应
*/
type hashRequest struct {
	Password string `json:"password"`
}

type hashResponse struct {
	Hash  string `json:"hash"`
	Error string `json:"error"`
}

//辅助方法，Go kit的http.DecodeRequestFunc类型，将http.Request的JSON主体解码到service.go
func decodeHashRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req hashRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

/**
validate方法的请求与响应
*/
type validateRequest struct {
	Password string `json:"password"`
	Hash     string `json:"hash"`
}

type validateResponse struct {
	Valid bool   `json:"valid"`
	Error string `json:"error"`
}

func decodeValidateRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req validateRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// 对hashRequest和validateRequest对象进行编码
func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

/**
端点（endpoint）是go kit中地一种特殊功能类型，代表单个RPC方法，作为代码跳转到代码的通用方法。
Go kit中主要有两种类型的错误：传输错误和业务逻辑错误。应当避免从endpoint函数返回业务逻辑错误信息，端点应该返回传输错误。
*/
func MakeHashEndpoint(srv Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(hashRequest)
		v, err := srv.Hash(ctx, req.Password)
		if err != nil {
			return hashResponse{v, err.Error()}, nil
		}
		return hashResponse{v, ""}, nil
	}
}

func MakeValidateEndpoint(srv Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(validateRequest)
		v, err := srv.Validate(ctx, req.Password, req.Hash)
		if err != nil {
			return validateResponse{v, err.Error()}, nil
		}
		return validateResponse{v, ""}, nil
	}
}

/**
将端点包装到服务实现中
*/
type Endpoints struct {
	HashEndpoint     endpoint.Endpoint
	ValidateEndpoint endpoint.Endpoint
}

func (e Endpoints) Hash(ctx context.Context, password string) (string, error) {
	req := hashRequest{
		Password: password,
	}
	resp, err := e.HashEndpoint(ctx, req)
	if err != nil {
		return "", err
	}
	hashResp := resp.(hashResponse)
	if hashResp.Error != "" {
		return "", errors.New(hashResp.Error)
	}
	return hashResp.Hash, nil
}

func (e Endpoints) Validate(ctx context.Context, password, hash string) (bool, error) {
	req := validateRequest{Password: password, Hash: hash}
	resp, err := e.ValidateEndpoint(ctx, req)
	if err != nil {
		return false, err
	}
	validateResp := resp.(validateResponse)
	if validateResp.Error != "" {
		return false, errors.New(validateResp.Error)
	}
	return validateResp.Valid, nil
}
