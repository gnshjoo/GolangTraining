// Package acdc asks if you ar ready to rock
package acdc

func Sum(n ...int) int {
	s := 0
	for _, v := range n {
		s += v
	}

	return s
}
