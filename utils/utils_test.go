package utils

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strconv"
	"strings"
	"testing"
)

func TestMap(t *testing.T) {
	mapped_int := Map([]int{0, 1, 2}, func(x int) int {
		return x * 2
	})
	for i, x := range []int{0, 2, 4} {
		assert.Equal(t, x, mapped_int[i])
	}

	mapped_str := Map([]string{"HELLO", "WORLD", "HOW", "ARE", "YOU"}, func(s string) string {
		return strings.ToLower(s)
	})
	for i, x := range []string{"hello", "world", "how", "are", "you"} {
		assert.Equal(t, x, mapped_str[i])
	}

	mapped_stoi := Map([]string{"-42", "4", "32", "42069"}, func(s string) int {
		res, _ := strconv.Atoi(s)
		return res
	})
	for i, x := range []int{-42, 4, 32, 42069} {
		assert.Equal(t, x, mapped_stoi[i])
	}
}

func TestMapSafe(t *testing.T) {
	mapped_noErrs, errors := MapSafe([]int{0, 1, 2}, func(x int) (string, error) {
		return fmt.Sprintf("LETS FREAKING GO: %d!", x), nil
	})
	assert.Empty(t, errors)
	for i, strVal := range []string{
		"LETS FREAKING GO: 0!",
		"LETS FREAKING GO: 1!",
		"LETS FREAKING GO: 2!",
	} {
		assert.Equal(t, strVal, mapped_noErrs[i])
	}

	mapped_errs, errors := MapSafe([]int{0, 1, 2}, func(x int) (float64, error) {
		if x == 0 {
			return 0.0, fmt.Errorf("NO ZERO!")
		}
		return float64(x), nil
	})
	assert.NotEmpty(t, errors)
	for _, err := range errors {
		assert.EqualError(t, err, "NO ZERO!")
	}
	assert.Equal(t, 1.0, mapped_errs[0])
	assert.Equal(t, 2.0, mapped_errs[1])
}
