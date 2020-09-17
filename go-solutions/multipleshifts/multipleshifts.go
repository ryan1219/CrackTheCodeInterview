package multipleshifts

/*
Question:
Users can define what shifts they are working on in their profile page but we don't want users choosing multiple shifts that interfere with each other.
**For example:** If a user is assigned a shift of `0000` to `0600`. They can't choose a shift that starts at `0400` because that would interfere with one of their shifts.
**Note:** If a shift is from `0000` to `0600`, a user can also have a shift that starts at `0600`
*/
import (
	"strconv"
)

type Shift struct {
	Start string
	End   string
}

func main() {
	//test1
	// // UserShifts are the users current shifts.
	// var userShifts = []Shift{
	// 	{Start: "0600", End: "1000"},
	// 	{Start: "1600", End: "2000"},
	// }

	// // AvailableCompanyShifts represent the available shifts that a user can work for a company.
	// var availableCompanyShifts = []Shift{
	// 	{Start: "0000", End: "2359"},
	// 	{Start: "0600", End: "1800"},
	// 	{Start: "0000", End: "1200"},
	// 	{Start: "0600", End: "1200"},
	// 	{Start: "1800", End: "2359"},
	// 	{Start: "0000", End: "0600"},
	// 	{Start: "1200", End: "2359"},
	// 	{Start: "1200", End: "1800"},
	// }

	// fmt.Println(findShifts(userShifts, availableCompanyShifts))
}

func findShifts(userShifts []Shift, available []Shift) []Shift {
	selectShifts := make([]Shift, 0)
	for _, shift := range available {
		if !checkOverlap(userShifts, shift) {
			selectShifts = append(selectShifts, shift)
		}
	}

	return selectShifts
}

func checkOverlap(userShifts []Shift, shift Shift) bool {
	for _, user := range userShifts {
		userStart, _ := strconv.Atoi(user.Start)
		userEnd, _ := strconv.Atoi(user.End)
		shiftStart, _ := strconv.Atoi(shift.Start)
		shiftEnd, _ := strconv.Atoi(shift.End)
		if userStart < shiftEnd && shiftStart < userEnd {
			return true
		}
	}

	return false
}
