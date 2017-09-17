package format

import "testing"

const (
	expectedOutput = "First  Last   Age\n" +
		"John   Smith  18\n" +
		"Darth  Vader  67"
)

type Person struct {
	First string
	Last  string
	Age   string
}

func columnHeaders() []string {
	return []string{
		"First",
		"Last",
		"Age",
	}
}

func columnData(persons []Person) [][]string {
	var data [][]string
	for _, d := range persons {
		data = append(data, []string{
			d.First,
			d.Last,
			string(d.Age),
		})
	}
	return data
}

func TestColumnFormatter(t *testing.T) {
	persons := []Person{
		Person{
			First: "John",
			Last:  "Smith",
			Age:   "18",
		},
		Person{
			First: "Darth",
			Last:  "Vader",
			Age:   "67",
		},
	}

	headers := columnHeaders()
	data := columnData(persons)
	formatter := NewColumnFormatter()
	cols := formatter.Format(headers, data)

	if cols != expectedOutput {
		t.Fail()
	}

}
