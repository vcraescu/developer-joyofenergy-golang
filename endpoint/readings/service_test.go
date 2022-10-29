package readings

import (
	"testing"

	"github.com/sirupsen/logrus"

	"joi-energy-golang/domain"
	"joi-energy-golang/repository"
)

func TestStoreReadings(t *testing.T) {
	t.Parallel()

	meterReadings := repository.NewMeterReadings(
		map[string][]domain.ElectricityReading{},
	)
	service := NewService(
		logrus.NewEntry(logrus.StandardLogger()),
		&meterReadings,
	)
	service.StoreReadings("1", []domain.ElectricityReading{})
}
