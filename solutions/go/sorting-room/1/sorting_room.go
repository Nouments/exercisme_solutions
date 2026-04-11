package sorting

import (
    "fmt"
    "math"
    "strconv"
)

func DescribeNumber(f float64) string {
    return fmt.Sprintf("This is the number %.1f", (math.Round(f*10) / 10))
}

type NumberBox interface {
    Number() int
}

func DescribeNumberBox(nb NumberBox) string {
    return fmt.Sprintf("This is a box containing the number %.1f", float64(nb.Number()))
}

type FancyNumber struct {
    n string
}

func (i FancyNumber) Value() string {
    return i.n
}

type FancyNumberBox interface {
    Value() string
}

func ExtractFancyNumber(fnb FancyNumberBox) int {
    fn, ok := fnb.(FancyNumber)
    if !ok {
        return 0
    }

    n, err := strconv.Atoi(fn.Value())
    if err != nil {
        return 0
    }

    return n
}

func DescribeFancyNumberBox(fnb FancyNumberBox) string {
    return fmt.Sprintf("This is a fancy box containing the number %.1f", float64(ExtractFancyNumber(fnb)))
}

func DescribeAnything(i interface{}) string {
    switch v := i.(type) {
    case float64:
        return DescribeNumber(v)
    case int:
        return DescribeNumber(float64(v))
    case NumberBox:
        return DescribeNumberBox(v)
    case FancyNumberBox:
        return DescribeFancyNumberBox(v)
    default:
        return "Return to sender"
    }
}

