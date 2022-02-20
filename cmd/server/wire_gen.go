// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/goxiaoy/go-saas-kit/pkg/api"
	"github.com/goxiaoy/go-saas-kit/pkg/authn/jwt"
	"github.com/goxiaoy/go-saas-kit/pkg/authz/authz"
	"github.com/goxiaoy/go-saas-kit/pkg/conf"
	server2 "github.com/goxiaoy/go-saas-kit/pkg/server"
	uow2 "github.com/goxiaoy/go-saas-kit/pkg/uow"
	api2 "github.com/goxiaoy/go-saas-kit/saas/api"
	"github.com/goxiaoy/go-saas-kit/saas/remote"
	api3 "github.com/goxiaoy/go-saas-kit/user/api"
	remote2 "github.com/goxiaoy/go-saas-kit/user/remote"
	"github.com/goxiaoy/go-saas/common/http"
	"github.com/goxiaoy/go-saas/gorm"
	"github.com/goxiaoy/kit-saas-layout/private/biz"
	conf2 "github.com/goxiaoy/kit-saas-layout/private/conf"
	"github.com/goxiaoy/kit-saas-layout/private/data"
	"github.com/goxiaoy/kit-saas-layout/private/server"
	"github.com/goxiaoy/kit-saas-layout/private/service"
	"github.com/goxiaoy/uow"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(services *conf.Services, security *conf.Security, config *uow.Config, gormConfig *gorm.Config, webMultiTenancyOption *http.WebMultiTenancyOption, confData *conf2.Data, logger log.Logger, arg ...grpc.ClientOption) (*kratos.App, func(), error) {
	tokenizerConfig := jwt.NewTokenizerConfig(security)
	tokenizer := jwt.NewTokenizer(tokenizerConfig)
	clientName := _wireClientNameValue
	saasContributor := api.NewSaasContributor(webMultiTenancyOption)
	userContributor := api.NewUserContributor()
	clientContributor := api.NewClientContributor()
	option := api.NewDefaultOption(saasContributor, userContributor, clientContributor)
	inMemoryTokenManager := api.NewInMemoryTokenManager(tokenizer)
	grpcConn, cleanup := api2.NewGrpcConn(clientName, services, option, inMemoryTokenManager, arg...)
	tenantServiceClient := api2.NewTenantGrpcClient(grpcConn)
	tenantStore := remote.NewRemoteGrpcTenantStore(tenantServiceClient)
	dbOpener, cleanup2 := gorm.NewDbOpener()
	manager := uow2.NewUowManager(gormConfig, config, dbOpener)
	decodeRequestFunc := _wireDecodeRequestFuncValue
	encodeResponseFunc := _wireEncodeResponseFuncValue
	encodeErrorFunc := _wireEncodeErrorFuncValue
	factory := data.NewBlobFactory(confData)
	dbProvider := data.NewProvider(confData, gormConfig, dbOpener, tenantStore, logger)
	greeterRepo := data.NewGreeterRepo(dbProvider, logger)
	greeterUsecase := biz.NewGreeterUsecase(greeterRepo, logger)
	apiGrpcConn, cleanup3 := api3.NewGrpcConn(clientName, services, option, inMemoryTokenManager, arg...)
	permissionServiceClient := api3.NewPermissionGrpcClient(apiGrpcConn)
	permissionChecker := remote2.NewRemotePermissionChecker(permissionServiceClient)
	authzOption := service.NewAuthorizationOption()
	subjectResolverImpl := authz.NewSubjectResolver(authzOption)
	defaultAuthorizationService := authz.NewDefaultAuthorizationService(permissionChecker, subjectResolverImpl, logger)
	greeterService := service.NewGreeterService(greeterUsecase, defaultAuthorizationService, logger)
	httpServer := server.NewHTTPServer(services, security, tokenizer, tenantStore, manager, decodeRequestFunc, encodeResponseFunc, encodeErrorFunc, factory, confData, webMultiTenancyOption, option, greeterService, logger)
	grpcServer := server.NewGRPCServer(services, tokenizer, tenantStore, manager, webMultiTenancyOption, option, greeterService, logger)
	seeder := server.NewSeeder(manager)
	app := newApp(logger, httpServer, grpcServer, seeder)
	return app, func() {
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}

var (
	_wireClientNameValue         = server.ClientName
	_wireDecodeRequestFuncValue  = server2.ReqDecode
	_wireEncodeResponseFuncValue = server2.ResEncoder
	_wireEncodeErrorFuncValue    = server2.ErrEncoder
)
