package main

type Stack struct {
	Items []int
	Size  int
	Max   int
}

func (s *Stack) Pop() int {
	if s.Size == 0 {
		return -99999999
	}
	item := s.Items[len(s.Items)-1]
	s.Items = s.Items[:len(s.Items)-1]
	s.Size--
	return item
}

func (s *Stack) Push(item int) {
	if s.Size == s.Max {
		tempArr := make([]int, 0, 10)
		s.Items = append(s.Items, tempArr...)
		s.Max += 10
	}
	s.Items = append(s.Items, item)
	s.Size++
}

func (s *Stack) IsEmpty() bool {
	return s.Size == 0
}

func (s *Stack) IsFull() bool {
	return s.Size == s.Max
}

func (s *Stack) Peek() interface{} {
	if s.Size == 0 {
		return nil
	}
	item := s.Items[s.Size-1]
	return item
}
