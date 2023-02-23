package iteration

const (
	repeatCounter = 5
)

func Repeat(char string, counter int) string {
	var repeated string
	for i := 0; i < counter; i++ {
		repeated += char
	}
	return repeated
}
