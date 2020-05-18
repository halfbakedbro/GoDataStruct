package array

type MyArray struct {
	length int32
	data   []int64
}

func (ma MyArray) Get(index int32) (int64, bool) {
	if index > ma.length {
		return 0, 0
	}
	return ma.data[index], 0
}

func (ma *MyArray) Push(value int64) {
	ma.data = append(ma.data, value)
	ma.length++
}
