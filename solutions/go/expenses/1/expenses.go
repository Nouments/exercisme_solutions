package expenses

import (
	"fmt"
)

// Record represents an expense record.
type Record struct {
	Day      int
	Amount   float64
	Category string
}

// DaysPeriod represents a period of days for expenses.
type DaysPeriod struct {
	From int
	To   int
}

// Filter returns the records for which the predicate function returns true.
func Filter(in []Record, predicate func(Record) bool) []Record {
	tab:=[]Record{}
	for _,val := range in{
		if predicate(val){
			tab = append(tab, val)
		}
	}
	return tab
}

// ByDaysPeriod returns predicate function that returns true when
// the day of the record is inside the period of day and false otherwise.
func ByDaysPeriod(p DaysPeriod) func(Record) bool {
	 return func(a Record) bool {
        return p.From <= a.Day && a.Day <= p.To
    }
}

// ByCategory returns predicate function that returns true when
// the category of the record is the same as the provided category
// and false otherwise.
func ByCategory(c string) func(Record) bool {
	return func(a Record)bool{
		return c == a.Category
	}
}

// TotalByPeriod returns total amount of expenses for records
// inside the period p.
func TotalByPeriod(in []Record, p DaysPeriod) float64 {
	amount := float64(0)
	for _,val := range in {
		for i:=p.From;i<p.To+1;i++{
			if i == val.Day{
				amount += val.Amount
			}
		}
	}
	return float64(amount)
}

// CategoryExpenses returns total amount of expenses for records
// in category c that are also inside the period p.
// An error must be returned only if there are no records in the list that belong
// to the given category, regardless of period of time.
func CategoryExpenses(in []Record, p DaysPeriod, c string) (float64, error) {
	exist := false
	amount := float64(0)
	for i := range in{
		if in[i].Category == c{
			exist = true
			break
		}
	}
	if !exist {
		return 0, fmt.Errorf("unknown category entertainment")
	}

	a := Filter(in,ByCategory(c))
	if len(a) == 0 {
		return 0,nil
	}
	for _,val := range a {
		if p.From<=val.Day && val.Day<=p.To{
			amount += val.Amount
		}
	}
	return amount,nil
}
