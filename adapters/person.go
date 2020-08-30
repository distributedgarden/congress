package adapters

import (
	"github.com/distributedgarden/propublicago/congress"
	rg "github.com/redislabs/redisgraph-go"
)

type PersonRepository interface {
	BatchCreate(members []congress.Member) (bool, error)
	PrintAll() (bool, error)
}

type PersonRepo struct {
	db *rg.Graph
}

func NewPersonRepository(db *rg.Graph) *PersonRepo {
	return &PersonRepo{db}
}

func (p *PersonRepo) BatchCreate(members []congress.Member) (bool, error) {
	for _, member := range members {
		node := rg.Node{
			Label: "Person",
			Properties: map[string]interface{}{
				"title":         member.Title,
				"first_name":    member.FirstName,
				"last_name":     member.LastName,
				"party":         member.Party,
				"state":         member.State,
				"next_election": member.NextElection,
			},
		}
		p.db.AddNode(&node)
	}
	_, err := p.db.Commit()
	if err != nil {
		return false, nil
	}

	return true, nil
}

func (p *PersonRepo) PrintAll() (bool, error) {
	query := `MATCH (p:Person) RETURN p.title, p.first_name, p.last_name, p.party, p.state, p.next_election`
	results, _ := p.db.Query(query)
	results.PrettyPrint()

	return true, nil
}
