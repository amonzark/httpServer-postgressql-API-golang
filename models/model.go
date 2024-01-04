package models

type Students struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Grade int    `json:"grade"`
}

func SampleStudent() []Students {
	var student = []Students{
		{ID: "300A", Name: "AVATAAR", Age: 20, Grade: 3},
	}
	return student
}
