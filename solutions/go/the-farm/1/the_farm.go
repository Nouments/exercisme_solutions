package thefarm

import "fmt"

type InvalidCowsError struct {
	cows    int
	message string
}

func (e *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", e.cows, e.message)
}


func DivideFood(f FodderCalculator, cows int) (float64, error) {
	fa, err := f.FodderAmount(cows)
	if err != nil {
		return 0, err
	}

	ff, err := f.FatteningFactor()
	if err != nil {
		return 0, err
	}

	return (fa * ff) / float64(cows), nil
}

func ValidateInputAndDivideFood(f FodderCalculator, cows int) (float64, error) {
	if cows < 1 {
		return 0, fmt.Errorf("invalid number of cows")
	}
	return DivideFood(f, cows)
}

func ValidateNumberOfCows(cows int) error {
	if cows < 0 {
		return &InvalidCowsError{
			cows:    cows,
			message: "there are no negative cows",
		}
	}

	if cows == 0 {
		return &InvalidCowsError{
			cows:    cows,
			message: "no cows don't need food",
		}
	}

	return nil
}