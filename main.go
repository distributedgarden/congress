package main

import (
	"github.com/distributedgarden/congress/adapters"
	"github.com/distributedgarden/propublicago/congress"
	"os"
)

func main() {
	/*
		add message bus and emit commands from here
		move batch create to UoW
	*/
	repo, _ := adapters.NewRepositories()

	c := congress.NewClient(os.Getenv("PROPUBLICA_API_KEY"))

	senateMembers := congress.MembersQueryParameters{Congress: "116", Chamber: "senate"}
	senate, _ := c.GetMembers(&senateMembers)
	repo.Person.BatchCreate(senate.Results[0].Members)

	houseMembers := congress.MembersQueryParameters{Congress: "116", Chamber: "house"}
	house, _ := c.GetMembers(&houseMembers)
	repo.Person.BatchCreate(house.Results[0].Members)

	repo.Person.PrintAll()

	senateCommittees := congress.CommitteeQueryParameters{Congress: "116", Chamber: "senate"}
	sCommittees, _ := c.GetCommittees(&senateCommittees)
	repo.Committee.BatchCreate(sCommittees.Results[0].Committees)

	houseCommittees := congress.CommitteeQueryParameters{Congress: "116", Chamber: "house"}
	hCommittees, _ := c.GetCommittees(&houseCommittees)
	repo.Committee.BatchCreate(hCommittees.Results[0].Committees)

	repo.Committee.PrintAll()
}
