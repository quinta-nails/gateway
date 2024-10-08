package config

import "net/url"

type ServicesConfig struct {
	UsersEndpoint     url.URL `env:"GRPC_USERS_ENDPOINT,required" envDefault:"127.0.0.1:51052"`
	CompaniesEndpoint url.URL `env:"GRPC_COMPANIES_ENDPOINT,required" envDefault:"127.0.0.1:51053"`
	BotsEndpoint      url.URL `env:"GRPC_BOTS_ENDPOINT,required" envDefault:"127.0.0.1:51054"`
}
