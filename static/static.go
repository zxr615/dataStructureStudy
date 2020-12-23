package static

type Node struct {
	Data interface{}
	Cur int
}

type Static struct {
	staticList []Node
}

// 我们通常把未被使用的数组元素称为备用链表。
// 而数组第一个元素，即下标为0的元素的cur就存放备用链表的第一个结点的下标
// 而数组的最后一个元素的cur则存放第一个有数值的元素的下标，
// 相当于单链表中的头结点作用，当整个链表为空时，则为0
func NewStatic() *Static {
	return &Static{}
}

func (s *Static) InitList(size int) *Static {
	s.staticList = make([]Node, size)

	for i := 0; i < size - 1; i++ {
		s.staticList[i].Cur = i+1
	}

	s.staticList[size - 1].Cur = 0

	return s
}

func (s *Static) Add(value interface{}) *Static {
	i := s.malloc()

	s.staticList[i].Data = value

	 return s
}

func (s *Static) malloc() int {
	i := s.staticList[0].Cur

	if s.staticList[0].Cur != 0 {
		s.staticList[0].Cur = s.staticList[i].Cur
	}

	return i
}
