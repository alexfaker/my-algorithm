package data_structure

type DoubleDirectQueue interface {
	LeftIn(a int)
	LeftOut() int
	GetLeft() int
	RightIn(a int)
	RightOut() int
	GetRight() int
	IsEmpty() bool
}

type DoubleDirectQueueImp struct {
	arr []int
}

func NewDoubleDirectQueueImp() *DoubleDirectQueueImp {
	return &DoubleDirectQueueImp{arr: []int{}}
}

func (d *DoubleDirectQueueImp) LeftIn(a int) {
	d.arr = append([]int{a}, d.arr...)

}

func (d *DoubleDirectQueueImp) RightIn(a int) {
	d.arr = append(d.arr, a)
}

func (d *DoubleDirectQueueImp) IsEmpty() bool {
	return len(d.arr) == 0
}

func (d *DoubleDirectQueueImp) GetLeft() int {
	if len(d.arr) <= 0 {
		return -1
	}
	return d.arr[0]
}

func (d *DoubleDirectQueueImp) GetRight() int {
	if len(d.arr) <= 0 {
		return -1
	}
	return d.arr[len(d.arr)-1]
}

func (d *DoubleDirectQueueImp) LeftOut() int {
	l := len(d.arr)
	if l <= 0 {
		return -1
	}
	t := d.arr[0]
	if l == 1 {
		d.arr = []int{}
	} else {
		d.arr = d.arr[1:]
	}
	return t
}

func (d *DoubleDirectQueueImp) RightOut() int {
	l := len(d.arr)
	if l <= 0 {
		return -1
	}
	t := d.arr[l-1]
	if l == 1 {
		d.arr = []int{}
	} else {
		// ???
		d.arr = d.arr[:l-1]
	}
	return t
}

type DoubleDirectQueueStrImp struct {
	arr []string
}

func NewDoubleDirectQueueStrImp() *DoubleDirectQueueStrImp {
	return &DoubleDirectQueueStrImp{arr: []string{}}
}

func (d *DoubleDirectQueueStrImp) LeftIn(a string) {
	d.arr = append([]string{a}, d.arr...)

}
func (d *DoubleDirectQueueStrImp) RightIn(a string) {
	d.arr = append(d.arr, a)
}

func (d *DoubleDirectQueueStrImp) IsEmpty() bool {
	return len(d.arr) == 0
}

func (d *DoubleDirectQueueStrImp) GetLeft() string {
	if len(d.arr) <= 0 {
		return ""
	}
	return d.arr[0]
}

func (d *DoubleDirectQueueStrImp) GetRight() string {
	if len(d.arr) <= 0 {
		return ""
	}
	return d.arr[len(d.arr)-1]
}

func (d *DoubleDirectQueueStrImp) LeftOut() string {
	l := len(d.arr)
	if l <= 0 {
		return ""
	}
	t := d.arr[0]
	if l == 1 {
		d.arr = []string{}
	} else {
		d.arr = d.arr[1:]
	}
	return t
}

func (d *DoubleDirectQueueStrImp) RightOut() string {
	l := len(d.arr)
	if l <= 0 {
		return ""
	}
	t := d.arr[l-1]
	if l == 1 {
		d.arr = []string{}
	} else {
		// ???
		d.arr = d.arr[:l-1]
	}
	return t
}
