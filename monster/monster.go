package monster

import (
	"strconv"

	"github.com/dreamsofcode-io/nethttp/entity"
)

func monsters() []entity.Monster {
	return []entity.Monster{
		{
			ID:     1,
			Name:   "Dracula",
			Powers: []string{"Immortality", "Shape-shifting", "Mind Control"},
		},
		{
			ID:     2,
			Name:   "Frankenstein",
			Powers: []string{"Superhuman Strength", "Endurance"},
		},
		{
			ID:     3,
			Name:   "Werewolf",
			Powers: []string{"Shape-shifting", "Enhanced Senses", "Regeneration"},
		},
		{
			ID:     4,
			Name:   "Zombie",
			Powers: []string{"Undead Physiology", "Immunity to Pain"},
		},
		{
			ID:     5,
			Name:   "Mummy",
			Powers: []string{"Immortality", "Control over Sand"},
		},
	}
}

func loadMonsters() map[string]entity.Monster {
	monsters := monsters()
	res := make(map[string]entity.Monster, len(monsters))

	for _, x := range monsters {
		res[strconv.Itoa(int(x.ID))] = x
	}

	return res
}
