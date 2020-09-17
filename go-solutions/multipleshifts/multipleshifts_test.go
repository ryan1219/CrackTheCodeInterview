package multipleshifts

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test1(t *testing.T) {

	// UserShifts are the users current shifts.
	var userShifts = []Shift{
		{Start: "0600", End: "1000"},
		{Start: "1600", End: "2000"},
	}

	// AvailableCompanyShifts represent the available shifts that a user can work for a company.
	var availableCompanyShifts = []Shift{
		{Start: "0000", End: "2359"},
		{Start: "0600", End: "1800"},
		{Start: "0000", End: "1200"},
		{Start: "0600", End: "1200"},
		{Start: "1800", End: "2359"},
		{Start: "0000", End: "0600"},
		{Start: "1200", End: "2359"},
		{Start: "1200", End: "1800"},
	}

	var expected = []Shift{
		{Start: "0000", End: "0600"},
	}
	// assert equality
	assert.ElementsMatch(t, findShifts(userShifts, availableCompanyShifts), expected, "they should be equal")
}

func Test2(t *testing.T) {

	// UserShifts are the users current shifts.
	var userShifts = []Shift{
		{Start: "0600", End: "1000"},
	}

	// AvailableCompanyShifts represent the available shifts that a user can work for a company.
	var availableCompanyShifts = []Shift{
		{Start: "0000", End: "2359"},
		{Start: "0600", End: "1800"},
		{Start: "0000", End: "1200"},
		{Start: "0600", End: "1200"},
		{Start: "1800", End: "2359"},
		{Start: "0000", End: "0600"},
		{Start: "1200", End: "2359"},
		{Start: "1200", End: "1800"},
	}

	var expected = []Shift{
		{Start: "0000", End: "0600"},
		{Start: "1800", End: "2359"},
		{Start: "1200", End: "2359"},
		{Start: "1200", End: "1800"},
	}
	// assert equality
	assert.ElementsMatch(t, findShifts(userShifts, availableCompanyShifts), expected, "they should be equal")
}
