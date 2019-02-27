package main

import (
	"fmt"
	"sort"
)

type Student struct {
	name string
	score float32
}

type By func(s1, s2 *Student) bool //By s1,s2를 파라미터로 받고 bool을 리턴하는 함수 타입

type studentSorter struct{
	students []Student
	by func(s1, s2 *Student) bool
}


func (by By) Sort(students []Student){
	sorter := &studentSorter{
		students : students,
		by : by,
	}
	sort.Sort(sorter);
}

func (s *studentSorter) Len() int{
	return len(s.students)
}

func (s *studentSorter) Less(i, j int) bool {
	return s.by(&s.students[i], &s.students[j])
}

func (s *studentSorter) Swap(i, j int){
	s.students[i], s.students[j] = s.students[j], s.students[i]
}

func main(){
	s := []Student{
		{"Maria", 89.3},
		{"Andrew", 72.6},
		{"John", 93.1},
	}

	name := func(p1, p2 *Student) bool { // 이름 오름차순 정렬 함수
		return p1.name < p2.name
	}
	score := func(p1, p2 *Student) bool { // 점수 오름차순 정렬 함수
		return p1.score < p2.score
	}
	reverseScore := func(p1, p2 *Student) bool { // 점수 내림차순 정렬 함수
		return !score(p1, p2)
	}

	By(name).Sort(s) // name 함수를 사용하여 이름 오름차순 정렬
	fmt.Println(s)

	By(reverseScore).Sort(s) // reverseScore 함수를 사용하여 점수 내림차순 정렬
	fmt.Println(s)
}