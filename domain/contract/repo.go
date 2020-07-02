package contract

//RepoManager defines the repository aggregator interface
type RepoManager interface {
	Ping() PingRepo
}

// PingRepo defines the data set for ping
type PingRepo interface{}
