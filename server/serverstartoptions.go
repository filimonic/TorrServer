package server

type ServerStartOptions struct {
	Port       string
	SslPort    string
	SslKey     string
	SslCert    string
	SslEnabled bool
	ReadOnly   bool
	SearchWA   bool
}
