package custom_types


type score float64

func (s score ) inside(start, end float64) bool {
	return float64(s) >= start && float64(s) <= end
}



