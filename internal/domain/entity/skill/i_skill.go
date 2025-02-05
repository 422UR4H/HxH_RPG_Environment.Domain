package skill

import "github.com/422UR4H/HxH_RPG_System/internal/domain/entity/experience"

type ISkill interface {
	experience.ITriggerCascadeExp

	GetLevel() int
	GetValueForTest() int
	GetExpPoints() int
	GetAggregateExpByLvl(lvl int) int
}
