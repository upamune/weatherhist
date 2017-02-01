package weatherhist

import (
	"fmt"
	"strconv"
	"testing"
)

func TestGetStation(t *testing.T) {
	cases := []struct {
		ID           string
		GroupNumber  string
		ExpectedName string
		ExpectedType string
	}{
		{
			"47570",
			"36",
			"若松",
			StationTypeS,
		},
		{
			"0967",
			"63",
			"大屋",
			StationTypeA,
		},
		{
			"967",
			"63",
			"大屋",
			StationTypeA,
		},
		{
			"47639",
			"49",
			"富士山",
			StationTypeS,
		},
		{
			"47639",
			"50",
			"富士山",
			StationTypeS,
		},
		{
			"1234567890",
			"0987654321",
			"",
			"",
		},
	}

	for _, c := range cases {
		station, err := GetStation(c.ID, c.GroupNumber)
		if err != nil {
			if c.ExpectedName == "" {
				// この場合は見つからないのが正常
				continue
			}
			t.Errorf("can't find such a station ID: %s, GroupNumber: %s", c.ID, c.GroupNumber)
		}

		if station.ID != StationID(c.ID) {
			// 4桁の0埋めで試してみる
			n, err := strconv.Atoi(c.ID)
			if err != nil {
				t.Errorf("Expected: %s, Actual: %v", c.ID, station.ID)
			}
			if station.ID != StationID(fmt.Sprintf("%04d", n)) {
				t.Errorf("Expected: %s, Actual: %v", c.ID, station.ID)
			}
		}

		if station.GroupNumber != c.GroupNumber {
			t.Errorf("Expected: %s, Actual: %v", c.GroupNumber, station.GroupNumber)
		}

		if station.Name != c.ExpectedName {
			t.Errorf("Expected: %s, Actual: %v", c.ExpectedName, station.Name)
		}

		if station.Type != c.ExpectedType {
			t.Errorf("Expected: %s, Actual: %v", c.ExpectedType, station.Type)
		}
	}
}
