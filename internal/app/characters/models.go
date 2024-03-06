package characters

import (
	"time"

	"github.com/uptrace/bun"
)

type Character struct {
	bun.BaseModel `bun:"table:myusers,alias:u"`
	Strength      int
	Dexterity     int
	Constitution  int
	Intelligence  int
	Wisdom        int
	Charisma      int
	Name          string
	Level         int
	CreatedAt     time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt     time.Time `bun:",nullzero,notnull,default:current_timestamp"`
}
