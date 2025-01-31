package display

import (
	"fmt"

	"github.com/sgrumley/deskday/pkg/count"
	"github.com/sgrumley/deskday/pkg/format"
)

type Service struct {
	Store DisplayStore
}

type DisplayStore interface {
	GetWorkDayEntries() (int, error)
}

func New(store DisplayStore) *Service {
	return &Service{
		Store: store,
	}
}

func (s *Service) Out() error {
	// call store get workdays
	// daysRecorded := 8
	daysRecorded, err := s.Store.GetWorkDayEntries()
	if err != nil {
		return fmt.Errorf("sql broke: %w", err)
	}
	daysRecorded = 10

	// TODO: get from config or user input??
	policy := 50

	days := Days{
		totalDays:    count.GetWorkdaysInCurrentMonth(),
		daysRecorded: daysRecorded,
		policy:       count.PercentageToFraction(policy),
	}

	days.PrintProgress()

	return nil
}

type Days struct {
	totalDays    int
	daysRecorded int
	policy       int
}

func (d *Days) GetRequiredDays() int {
	return (d.totalDays + d.policy - 1) / d.policy
}

func (d *Days) GetRemainingWorkDays() int {
	return count.GetRemainingWorkDays()
}

// PrintProgress displays a centered progress display with current status
func (d *Days) PrintProgress() {
	daysRequired := d.GetRequiredDays()
	if daysRequired <= 0 {
		return
	}

	// NOTE: being in a map means the order is not deterministic
	labelConfig := map[string]string{
		"Days in office:":        fmt.Sprintf("%d", d.daysRecorded),
		"Days Required:":         fmt.Sprintf("%d", d.GetRequiredDays()),
		"Office Days Remaining:": fmt.Sprintf("%d", d.GetRemainingWorkDays()),
		"Progress:":              fmt.Sprintf("%d/%d", d.daysRecorded, daysRequired),
	}

	kvs := format.Text(labelConfig)
	for _, line := range kvs {
		fmt.Println(line)
	}

	// determine what color the bar should be based on percentage completion
	barColor := format.GREEN
	progress := (d.daysRecorded * 100) / d.GetRequiredDays()
	switch true {
	case progress <= 33:
		barColor = format.RED
	case progress <= 66:
		barColor = format.ROSEWATER
	default:
		barColor = format.GREEN
	}

	progressBar := format.CreateProgressBarWithColor(d.daysRecorded, daysRequired, barColor)
	fmt.Println(progressBar)
}
