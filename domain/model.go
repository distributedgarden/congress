package domain

import (
	"github.com/google/uuid"
	"time"
)

type Person struct {
	Id           uuid.UUID
	Title        string
	FirstName    string
	LastName     string
	Party        string
	State        string
	NextElection string
}

type Committee struct {
	Id            uuid.UUID
	Name          string
	Chamber       string
	Chair         string
	ChairParty    string
	ChairState    string
	RankingMember string
}

type Proxy struct {
	Id uuid.UUID
}

type Organization struct {
	Id          uuid.UUID
	Name        string
	Description string
}

type Industry struct {
	Id   uuid.UUID
	Name string
}

type Sector struct {
	Id   uuid.UUID
	Name string
}

type Role struct {
	Id          uuid.UUID
	Title       string
	Description string
}

type Timespan struct {
	Id    uuid.UUID
	Start time.Time
	End   time.Time
}

type Alias struct {
	Id   uuid.UUID
	Name string
}

type Contribution struct {
	Id uuid.UUID
}

type Bill struct {
	Id           uuid.UUID
	Summary      string
	BillId       string
	SponsorState string
}

type MemberEdge struct {}

type ProxiesEdge struct {}

type PositionEdge struct {}

type DurationEdge struct {}

type SubsetEdge struct {}

type ContributesEdge struct {}

type StaffAffiliationEdge struct {}

type SponsorsEdge struct {}

type CosponsorsEdge struct {}
