package spiritual

import (
	"github.com/422UR4H/HxH_RPG_System/internal/domain/entity/enum"
	"github.com/422UR4H/HxH_RPG_System/internal/domain/entity/experience"
)

type NenCategory struct {
	exp   experience.Exp
	name  enum.CategoryName
	hatsu IHatsu
}

func NewNenCategory(
	exp experience.Exp,
	name enum.CategoryName,
	hatsu IHatsu,
) *NenCategory {
	return &NenCategory{exp: exp, name: name, hatsu: hatsu}
}

func (nc *NenCategory) CascadeUpgradeTrigger(exp int) int {
	diff := nc.exp.IncreasePoints(exp)
	nc.hatsu.CascadeUpgrade(exp)
	return diff
}

func (nc *NenCategory) GetExpPoints() int {
	return nc.exp.GetPoints()
}

func (nc *NenCategory) GetLevel() int {
	return int(float64(nc.exp.GetLevel()) * nc.hatsu.GetPercentOf(nc.name) / 100.0)
}

func (nc *NenCategory) Clone(name enum.CategoryName) *NenCategory {
	return NewNenCategory(*nc.exp.Clone(), name, nc.hatsu)
}
