package adapters

import (
	"github.com/gomodule/redigo/redis"
	rg "github.com/redislabs/redisgraph-go"
)

func NewRepositories() (*Repositories, error) {
	pool := newPool()
	graph, _ := SetupCongressGraph(pool)
	return &Repositories{
		Person:    NewPersonRepository(&graph),
		Committee: NewCommitteeRepository(&graph),
	}, nil

}

func SetupCongressGraph(db *redis.Pool) (rg.Graph, error) {
	graph := rg.GraphNew("congress", db.Get())
	graph.Delete()

	return graph, nil
}

type Repositories struct {
	Person    PersonRepository
	Committee CommitteeRepository
}
