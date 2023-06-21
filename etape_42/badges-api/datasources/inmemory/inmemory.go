package inmemory

import (
	"badges/types"
	"fmt"
	"math/rand"
)

var (
	Users = map[types.ID]*types.User{
		1 : &types.User{
		},
	}

	Badges = map[types.ID]*types.Badge{
		1 : &types.Badge{
			Name: "go",
			URL: "somewher",
		},
	}
)

type DataSource struct {}

func (ds DataSource) Connect() {}

func (ds DataSource) CreateUser(c types.User) (types.ID, error) {
	id := types.ID(rand.Int31n(100000))
	Users[id] = &c
	return id, nil
}

func (ds DataSource) GetBadges() ([]*types.Badge, error) {
	cc := make([]*types.Badge, 0, len(Badges))
	for _, c := range Badges {
		cc = append(cc, c)
	}
	return cc, nil
}

func (ds DataSource) GetBadge(id types.ID) (*types.Badge, error) {
	c, ok := Badges[id]
	if !ok {
		fmt.Errorf("not found")
	}

	return c, nil
}

func (ds DataSource) GetUserBadges(userID types.ID) ([]types.Badge, error) {
	id := types.ID(rand.Int31n(100000))
	Badges[id] = &c
	return id, nil
}

func (ds DataSource) AddUserBadge(userID types.ID, c types.Badge) (types.ID, error) {
	id := types.ID(rand.Int31n(100000))
	Badges[id] = &c
	return id, nil
}

func (ds DataSource) UpdateBadge(types.ID, *types.Badge) error {
	return nil
}

func (ds DataSource) DeleteBadge(types.ID) error {
	return nil
}
