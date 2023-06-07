package underscore

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func Test_Count_Can_Count_Numbers(t *testing.T) {
	numbers := Range(1, 100)
	count := Count(numbers, func(n int) bool {
		return n%2 == 0
	})

	assert.Equal(t, 50, count)
}

type People struct {
	Name   string
	Age    int
	Gender string
}

func Test_Count_Can_Count__People(t *testing.T) {
	people := []People{
		{Name: "Andy", Age: 43, Gender: "M"},
		{Name: "Fred", Age: 33, Gender: "M"},
		{Name: "Jack", Age: 23, Gender: "M"},
		{Name: "Jill", Age: 43, Gender: "F"},
		{Name: "Anna", Age: 33, Gender: "F"},
		{Name: "Arya", Age: 23, Gender: "F"},
		{Name: "Jane", Age: 13, Gender: "F"},
	}

	a := Count(people, func(p People) bool {
		return strings.HasPrefix(p.Name, "A")
	})
	assert.Equal(t, 3, a)

	females := Count(people, func(p People) bool {
		return p.Gender == "F"
	})
	assert.Equal(t, 4, females)

	males := Count(people, func(p People) bool {
		return p.Gender == "M"
	})
	assert.Equal(t, 3, males)

	over30 := Count(people, func(p People) bool {
		return p.Age > 30
	})
	assert.Equal(t, 4, over30)

	under30 := Count(people, func(p People) bool {
		return p.Age < 30
	})
	assert.Equal(t, 3, under30)

	under20 := Count(people, func(p People) bool {
		return p.Age < 20
	})
	assert.Equal(t, 1, under20)
}
