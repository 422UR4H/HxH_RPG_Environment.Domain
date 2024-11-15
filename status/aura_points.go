package status

type AuraPoints struct {
	Status
}

// TODO: rerify if should I private status and its fields
func NewAuraPoints() *AuraPoints {
	points := 0
	return &AuraPoints{
		Status: Status{
			Min:     points,
			Current: points,
			Max:     points,
		},
	}
}

func (ap *AuraPoints) StatusUpgrade(level int) {
	ap.Status.StatusUpgrade(level)
}

func (ap *AuraPoints) IncreaseAt(value int) int {
	return ap.Status.IncreaseAt(value)
}

func (ap *AuraPoints) DecreaseAt(value int) int {
	return ap.Status.DecreaseAt(value)
}

func (ap *AuraPoints) GetMin() int {
	return ap.Status.Min
}

func (ap *AuraPoints) GetCurrent() int {
	return ap.Status.Current
}

func (ap *AuraPoints) GetMax() int {
	return ap.Status.Max
}
