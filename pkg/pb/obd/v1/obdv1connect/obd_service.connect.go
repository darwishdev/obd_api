// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: obd/v1/obd_service.proto

package obdv1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	_ "github.com/darwishdev/obd_api/pkg/pb/obd/v1"
	car "github.com/darwishdev/obd_api/pkg/pb/obd/v1/car"
	center "github.com/darwishdev/obd_api/pkg/pb/obd/v1/center"
	review "github.com/darwishdev/obd_api/pkg/pb/obd/v1/review"
	user "github.com/darwishdev/obd_api/pkg/pb/obd/v1/user"
	winch "github.com/darwishdev/obd_api/pkg/pb/obd/v1/winch"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// ObdName is the fully-qualified name of the Obd service.
	ObdName = "obd.v1.Obd"
)

// ObdClient is a client for the obd.v1.Obd service.
type ObdClient interface {
	UserCreate(context.Context, *connect_go.Request[user.UserCreateRequest]) (*connect_go.Response[user.UserCreateResponse], error)
	UserLogin(context.Context, *connect_go.Request[user.UserLoginRequest]) (*connect_go.Response[user.UserLoginResponse], error)
	UserUpdate(context.Context, *connect_go.Request[user.UserUpdateRequest]) (*connect_go.Response[user.UserUpdateResponse], error)
	UserAuthorize(context.Context, *connect_go.Request[emptypb.Empty]) (*connect_go.Response[user.UserAuthorizeResponse], error)
	CarBrandsList(context.Context, *connect_go.Request[car.BrandsListRequest]) (*connect_go.Response[car.BrandsListResponse], error)
	CarCreate(context.Context, *connect_go.Request[car.CarCreateRequest]) (*connect_go.Response[car.CarCreateResponse], error)
	CarUpdate(context.Context, *connect_go.Request[car.CarUpdateRequest]) (*connect_go.Response[car.CarUpdateResponse], error)
	CentersList(context.Context, *connect_go.Request[center.CentersListRequest]) (*connect_go.Response[center.CentersListResponse], error)
	WinchList(context.Context, *connect_go.Request[winch.WinchListRequest]) (*connect_go.Response[winch.WinchListResponse], error)
	ReviewsList(context.Context, *connect_go.Request[review.ReviewsListRequest]) (*connect_go.Response[review.ReviewsListResponse], error)
	ReviewCreate(context.Context, *connect_go.Request[review.ReviewCreateRequest]) (*connect_go.Response[review.ReviewCreateResponse], error)
}

// NewObdClient constructs a client for the obd.v1.Obd service. By default, it uses the Connect
// protocol with the binary Protobuf Codec, asks for gzipped responses, and sends uncompressed
// requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewObdClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) ObdClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &obdClient{
		userCreate: connect_go.NewClient[user.UserCreateRequest, user.UserCreateResponse](
			httpClient,
			baseURL+"/obd.v1.Obd/UserCreate",
			opts...,
		),
		userLogin: connect_go.NewClient[user.UserLoginRequest, user.UserLoginResponse](
			httpClient,
			baseURL+"/obd.v1.Obd/UserLogin",
			opts...,
		),
		userUpdate: connect_go.NewClient[user.UserUpdateRequest, user.UserUpdateResponse](
			httpClient,
			baseURL+"/obd.v1.Obd/UserUpdate",
			opts...,
		),
		userAuthorize: connect_go.NewClient[emptypb.Empty, user.UserAuthorizeResponse](
			httpClient,
			baseURL+"/obd.v1.Obd/UserAuthorize",
			opts...,
		),
		carBrandsList: connect_go.NewClient[car.BrandsListRequest, car.BrandsListResponse](
			httpClient,
			baseURL+"/obd.v1.Obd/CarBrandsList",
			opts...,
		),
		carCreate: connect_go.NewClient[car.CarCreateRequest, car.CarCreateResponse](
			httpClient,
			baseURL+"/obd.v1.Obd/CarCreate",
			opts...,
		),
		carUpdate: connect_go.NewClient[car.CarUpdateRequest, car.CarUpdateResponse](
			httpClient,
			baseURL+"/obd.v1.Obd/CarUpdate",
			opts...,
		),
		centersList: connect_go.NewClient[center.CentersListRequest, center.CentersListResponse](
			httpClient,
			baseURL+"/obd.v1.Obd/CentersList",
			opts...,
		),
		winchList: connect_go.NewClient[winch.WinchListRequest, winch.WinchListResponse](
			httpClient,
			baseURL+"/obd.v1.Obd/WinchList",
			opts...,
		),
		reviewsList: connect_go.NewClient[review.ReviewsListRequest, review.ReviewsListResponse](
			httpClient,
			baseURL+"/obd.v1.Obd/ReviewsList",
			opts...,
		),
		reviewCreate: connect_go.NewClient[review.ReviewCreateRequest, review.ReviewCreateResponse](
			httpClient,
			baseURL+"/obd.v1.Obd/ReviewCreate",
			opts...,
		),
	}
}

// obdClient implements ObdClient.
type obdClient struct {
	userCreate    *connect_go.Client[user.UserCreateRequest, user.UserCreateResponse]
	userLogin     *connect_go.Client[user.UserLoginRequest, user.UserLoginResponse]
	userUpdate    *connect_go.Client[user.UserUpdateRequest, user.UserUpdateResponse]
	userAuthorize *connect_go.Client[emptypb.Empty, user.UserAuthorizeResponse]
	carBrandsList *connect_go.Client[car.BrandsListRequest, car.BrandsListResponse]
	carCreate     *connect_go.Client[car.CarCreateRequest, car.CarCreateResponse]
	carUpdate     *connect_go.Client[car.CarUpdateRequest, car.CarUpdateResponse]
	centersList   *connect_go.Client[center.CentersListRequest, center.CentersListResponse]
	winchList     *connect_go.Client[winch.WinchListRequest, winch.WinchListResponse]
	reviewsList   *connect_go.Client[review.ReviewsListRequest, review.ReviewsListResponse]
	reviewCreate  *connect_go.Client[review.ReviewCreateRequest, review.ReviewCreateResponse]
}

// UserCreate calls obd.v1.Obd.UserCreate.
func (c *obdClient) UserCreate(ctx context.Context, req *connect_go.Request[user.UserCreateRequest]) (*connect_go.Response[user.UserCreateResponse], error) {
	return c.userCreate.CallUnary(ctx, req)
}

// UserLogin calls obd.v1.Obd.UserLogin.
func (c *obdClient) UserLogin(ctx context.Context, req *connect_go.Request[user.UserLoginRequest]) (*connect_go.Response[user.UserLoginResponse], error) {
	return c.userLogin.CallUnary(ctx, req)
}

// UserUpdate calls obd.v1.Obd.UserUpdate.
func (c *obdClient) UserUpdate(ctx context.Context, req *connect_go.Request[user.UserUpdateRequest]) (*connect_go.Response[user.UserUpdateResponse], error) {
	return c.userUpdate.CallUnary(ctx, req)
}

// UserAuthorize calls obd.v1.Obd.UserAuthorize.
func (c *obdClient) UserAuthorize(ctx context.Context, req *connect_go.Request[emptypb.Empty]) (*connect_go.Response[user.UserAuthorizeResponse], error) {
	return c.userAuthorize.CallUnary(ctx, req)
}

// CarBrandsList calls obd.v1.Obd.CarBrandsList.
func (c *obdClient) CarBrandsList(ctx context.Context, req *connect_go.Request[car.BrandsListRequest]) (*connect_go.Response[car.BrandsListResponse], error) {
	return c.carBrandsList.CallUnary(ctx, req)
}

// CarCreate calls obd.v1.Obd.CarCreate.
func (c *obdClient) CarCreate(ctx context.Context, req *connect_go.Request[car.CarCreateRequest]) (*connect_go.Response[car.CarCreateResponse], error) {
	return c.carCreate.CallUnary(ctx, req)
}

// CarUpdate calls obd.v1.Obd.CarUpdate.
func (c *obdClient) CarUpdate(ctx context.Context, req *connect_go.Request[car.CarUpdateRequest]) (*connect_go.Response[car.CarUpdateResponse], error) {
	return c.carUpdate.CallUnary(ctx, req)
}

// CentersList calls obd.v1.Obd.CentersList.
func (c *obdClient) CentersList(ctx context.Context, req *connect_go.Request[center.CentersListRequest]) (*connect_go.Response[center.CentersListResponse], error) {
	return c.centersList.CallUnary(ctx, req)
}

// WinchList calls obd.v1.Obd.WinchList.
func (c *obdClient) WinchList(ctx context.Context, req *connect_go.Request[winch.WinchListRequest]) (*connect_go.Response[winch.WinchListResponse], error) {
	return c.winchList.CallUnary(ctx, req)
}

// ReviewsList calls obd.v1.Obd.ReviewsList.
func (c *obdClient) ReviewsList(ctx context.Context, req *connect_go.Request[review.ReviewsListRequest]) (*connect_go.Response[review.ReviewsListResponse], error) {
	return c.reviewsList.CallUnary(ctx, req)
}

// ReviewCreate calls obd.v1.Obd.ReviewCreate.
func (c *obdClient) ReviewCreate(ctx context.Context, req *connect_go.Request[review.ReviewCreateRequest]) (*connect_go.Response[review.ReviewCreateResponse], error) {
	return c.reviewCreate.CallUnary(ctx, req)
}

// ObdHandler is an implementation of the obd.v1.Obd service.
type ObdHandler interface {
	UserCreate(context.Context, *connect_go.Request[user.UserCreateRequest]) (*connect_go.Response[user.UserCreateResponse], error)
	UserLogin(context.Context, *connect_go.Request[user.UserLoginRequest]) (*connect_go.Response[user.UserLoginResponse], error)
	UserUpdate(context.Context, *connect_go.Request[user.UserUpdateRequest]) (*connect_go.Response[user.UserUpdateResponse], error)
	UserAuthorize(context.Context, *connect_go.Request[emptypb.Empty]) (*connect_go.Response[user.UserAuthorizeResponse], error)
	CarBrandsList(context.Context, *connect_go.Request[car.BrandsListRequest]) (*connect_go.Response[car.BrandsListResponse], error)
	CarCreate(context.Context, *connect_go.Request[car.CarCreateRequest]) (*connect_go.Response[car.CarCreateResponse], error)
	CarUpdate(context.Context, *connect_go.Request[car.CarUpdateRequest]) (*connect_go.Response[car.CarUpdateResponse], error)
	CentersList(context.Context, *connect_go.Request[center.CentersListRequest]) (*connect_go.Response[center.CentersListResponse], error)
	WinchList(context.Context, *connect_go.Request[winch.WinchListRequest]) (*connect_go.Response[winch.WinchListResponse], error)
	ReviewsList(context.Context, *connect_go.Request[review.ReviewsListRequest]) (*connect_go.Response[review.ReviewsListResponse], error)
	ReviewCreate(context.Context, *connect_go.Request[review.ReviewCreateRequest]) (*connect_go.Response[review.ReviewCreateResponse], error)
}

// NewObdHandler builds an HTTP handler from the service implementation. It returns the path on
// which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewObdHandler(svc ObdHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/obd.v1.Obd/UserCreate", connect_go.NewUnaryHandler(
		"/obd.v1.Obd/UserCreate",
		svc.UserCreate,
		opts...,
	))
	mux.Handle("/obd.v1.Obd/UserLogin", connect_go.NewUnaryHandler(
		"/obd.v1.Obd/UserLogin",
		svc.UserLogin,
		opts...,
	))
	mux.Handle("/obd.v1.Obd/UserUpdate", connect_go.NewUnaryHandler(
		"/obd.v1.Obd/UserUpdate",
		svc.UserUpdate,
		opts...,
	))
	mux.Handle("/obd.v1.Obd/UserAuthorize", connect_go.NewUnaryHandler(
		"/obd.v1.Obd/UserAuthorize",
		svc.UserAuthorize,
		opts...,
	))
	mux.Handle("/obd.v1.Obd/CarBrandsList", connect_go.NewUnaryHandler(
		"/obd.v1.Obd/CarBrandsList",
		svc.CarBrandsList,
		opts...,
	))
	mux.Handle("/obd.v1.Obd/CarCreate", connect_go.NewUnaryHandler(
		"/obd.v1.Obd/CarCreate",
		svc.CarCreate,
		opts...,
	))
	mux.Handle("/obd.v1.Obd/CarUpdate", connect_go.NewUnaryHandler(
		"/obd.v1.Obd/CarUpdate",
		svc.CarUpdate,
		opts...,
	))
	mux.Handle("/obd.v1.Obd/CentersList", connect_go.NewUnaryHandler(
		"/obd.v1.Obd/CentersList",
		svc.CentersList,
		opts...,
	))
	mux.Handle("/obd.v1.Obd/WinchList", connect_go.NewUnaryHandler(
		"/obd.v1.Obd/WinchList",
		svc.WinchList,
		opts...,
	))
	mux.Handle("/obd.v1.Obd/ReviewsList", connect_go.NewUnaryHandler(
		"/obd.v1.Obd/ReviewsList",
		svc.ReviewsList,
		opts...,
	))
	mux.Handle("/obd.v1.Obd/ReviewCreate", connect_go.NewUnaryHandler(
		"/obd.v1.Obd/ReviewCreate",
		svc.ReviewCreate,
		opts...,
	))
	return "/obd.v1.Obd/", mux
}

// UnimplementedObdHandler returns CodeUnimplemented from all methods.
type UnimplementedObdHandler struct{}

func (UnimplementedObdHandler) UserCreate(context.Context, *connect_go.Request[user.UserCreateRequest]) (*connect_go.Response[user.UserCreateResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("obd.v1.Obd.UserCreate is not implemented"))
}

func (UnimplementedObdHandler) UserLogin(context.Context, *connect_go.Request[user.UserLoginRequest]) (*connect_go.Response[user.UserLoginResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("obd.v1.Obd.UserLogin is not implemented"))
}

func (UnimplementedObdHandler) UserUpdate(context.Context, *connect_go.Request[user.UserUpdateRequest]) (*connect_go.Response[user.UserUpdateResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("obd.v1.Obd.UserUpdate is not implemented"))
}

func (UnimplementedObdHandler) UserAuthorize(context.Context, *connect_go.Request[emptypb.Empty]) (*connect_go.Response[user.UserAuthorizeResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("obd.v1.Obd.UserAuthorize is not implemented"))
}

func (UnimplementedObdHandler) CarBrandsList(context.Context, *connect_go.Request[car.BrandsListRequest]) (*connect_go.Response[car.BrandsListResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("obd.v1.Obd.CarBrandsList is not implemented"))
}

func (UnimplementedObdHandler) CarCreate(context.Context, *connect_go.Request[car.CarCreateRequest]) (*connect_go.Response[car.CarCreateResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("obd.v1.Obd.CarCreate is not implemented"))
}

func (UnimplementedObdHandler) CarUpdate(context.Context, *connect_go.Request[car.CarUpdateRequest]) (*connect_go.Response[car.CarUpdateResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("obd.v1.Obd.CarUpdate is not implemented"))
}

func (UnimplementedObdHandler) CentersList(context.Context, *connect_go.Request[center.CentersListRequest]) (*connect_go.Response[center.CentersListResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("obd.v1.Obd.CentersList is not implemented"))
}

func (UnimplementedObdHandler) WinchList(context.Context, *connect_go.Request[winch.WinchListRequest]) (*connect_go.Response[winch.WinchListResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("obd.v1.Obd.WinchList is not implemented"))
}

func (UnimplementedObdHandler) ReviewsList(context.Context, *connect_go.Request[review.ReviewsListRequest]) (*connect_go.Response[review.ReviewsListResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("obd.v1.Obd.ReviewsList is not implemented"))
}

func (UnimplementedObdHandler) ReviewCreate(context.Context, *connect_go.Request[review.ReviewCreateRequest]) (*connect_go.Response[review.ReviewCreateResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("obd.v1.Obd.ReviewCreate is not implemented"))
}
