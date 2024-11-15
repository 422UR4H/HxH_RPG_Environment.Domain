package main

import (
	"fmt"

	experience "github.com/422UR4H/HxH_RPG_Environment.Domain/experience"
	status "github.com/422UR4H/HxH_RPG_Environment.Domain/status"
)

func main() {
	expTable := experience.NewDefaultExpTable()
	fmt.Println(expTable.ToString())
	fmt.Println(expTable.GetLvlByExp(29))

	ap := status.NewAuraPoints()
	fmt.Println(ap.Min)

	fmt.Println(ap.Max)
	fmt.Println(ap.Current)
}
