package adapters

import (
	"github.com/distributedgarden/congress/clients/propublica/congress"
	rg "github.com/redislabs/redisgraph-go"
)

type CommitteeRepository interface {
	BatchCreate(committees []congress.Committee) (bool, error)
	PrintAll() (bool, error)
}

type CommitteeRepo struct {
	db *rg.Graph
}

func NewCommitteeRepository(db *rg.Graph) *CommitteeRepo {
	return &CommitteeRepo{db}
}

func (c *CommitteeRepo) BatchCreate(committees []congress.Committee) (bool, error) {
	for _, committee := range committees {
		node := rg.Node{
			Label: "Committee",
			Properties: map[string]interface{}{
				"name":           committee.Name,
				"chamber":        committee.Chamber,
				"chair":          committee.Chair,
				"chair_party":    committee.ChairParty,
				"chair_state":    committee.ChairState,
				"ranking_member": committee.RankingMember,
			},
		}
		c.db.AddNode(&node)
	}
	_, err := c.db.Commit()
	if err != nil {
		return false, nil
	}

	return true, nil
}

func (c *CommitteeRepo) PrintAll() (bool, error) {
	query := `MATCH (c:Committee) RETURN c.name, c.chamber, c.ranking_member, c.chair_party`
	results, _ := c.db.Query(query)
	results.PrettyPrint()

	return true, nil
}
