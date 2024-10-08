package main

import (
	f "github.com/foundation-go/foundation"
	fg "github.com/foundation-go/foundation/gateway"
	pbBots "github.com/quinta-nails/protobuf/gen/go/bots"
	pbCompanies "github.com/quinta-nails/protobuf/gen/go/companies"
	pbUsers "github.com/quinta-nails/protobuf/gen/go/users"
)

var (
	app = f.InitGateway("gateway")

	authenticationExcept = []string{
		pbCompanies.CompaniesService_Reserve_FullMethodName,
	}

	services = []*fg.Service{
		{Name: "users", Register: pbUsers.RegisterUsersServiceHandlerFromEndpoint},
		{Name: "companies", Register: pbCompanies.RegisterCompaniesServiceHandlerFromEndpoint},
		{Name: "bots", Register: pbBots.RegisterTelegramBotsServiceHandlerFromEndpoint},
	}
)

func main() {
	opts := f.NewGatewayOptions()

	opts.Services = services
	opts.WithAuthentication = true
	opts.AuthenticationExcept = authenticationExcept
	opts.CORSOptions = fg.NewCORSOptions()

	app.Start(opts)
}
