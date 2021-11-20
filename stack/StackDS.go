package stack

type Stack struct {
	Values []string
	Top    int
}

func (s *Stack) Push(a string) {
	s.Values = append(s.Values, a)
	s.Top += 1
}

func (s *Stack) Pop() string {
	temp := s.Values[len(s.Values)-1]
	s.Values = s.Values[:len(s.Values)-1]
	s.Top -= 1
	return temp
}
