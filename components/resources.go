package components

import (
	"fmt"

	"github.com/fatih/color"
)

type ResourceBarComponentInterface interface {
	GetResourceBarComponent() *ResourceBarComponent
}

type ResourceBarComponent struct {
	Total int
	Max   int
	Regen float64
}

func (hc *ResourceBarComponent) ResourceBarComponent() *ResourceBarComponent {
	return hc
}

func (hc ResourceBarComponent) String() string {
	d := float64(hc.Total) / float64(hc.Max)
	s := fmt.Sprintf("%d/%d", hc.Total, hc.Max)
	switch {
	case d < .25:
		return color.RedString(s)
	case d < .50:
		return color.YellowString(s)
	default:
		return color.GreenString(s)
	}
}
