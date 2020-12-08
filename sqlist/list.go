package sqlist

import (
	"errors"
	"fmt"
)

type elemType []int

type SqList struct {
	maxSize       int
	currentLength int
	data          elemType
}

func New(size int) *SqList {
	return &SqList{
		data:          make(elemType, size),
		maxSize:       size,
		currentLength: 0,
	}
}

func (s *SqList) Insert(i int, elem int) error {
	if s.currentLength+1 > s.maxSize {
		fmt.Println(i, "顺序线性表已满")
		return errors.New("顺序线性表已满")
	}
	if i < 1 || i > s.currentLength+1 {
		fmt.Println(i, "不在范围内")
		return errors.New("不在范围内")
	}

	if i <= s.currentLength {
		for k := s.currentLength; k >= i; k-- {
			s.data[k] = s.data[k-1]
		}
	}

	s.data[i-1] = elem
	s.currentLength++

	fmt.Println(s.data)
	return nil
}

func (s *SqList) Get(i int) (int, error) {
	if i < 1 || s.currentLength == 0 || i > s.currentLength {
		return 0, errors.New("下标不存在")
	}

	fmt.Println(s.data)
	return s.data[i-1], nil
}

func (s *SqList) Del(i int) error {
	if s.currentLength == 0 {
		return errors.New("线性表为空")
	}
	if i < 0 || i > s.currentLength {
		return errors.New("下标不存在")
	}

	if i < s.currentLength {
		for k := i; k < s.currentLength; k++ {
			s.data[k-1] = s.data[k]
		}
	}
	s.data[s.currentLength-1] = 0

	s.currentLength--

	fmt.Println(s.data)
	return nil
}
func (s *SqList) Add(elem int) error {
	if s.currentLength >= s.maxSize {
		return errors.New("顺序线性表已满")
	}

	s.data[s.currentLength] = elem

	s.currentLength++

	fmt.Println(s.data)
	return nil
}
