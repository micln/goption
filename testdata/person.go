package testdata

type Person struct {
	Name      string
	Age       int
	Interests []string

	Parent   *Person
	Children [5]*Person

	Friends []*Person
}

type Car struct {
	Name string
}
