package main

type StringIntMap struct {
	val map[string]int
}

func NewStringIntMap() *StringIntMap {
	return &StringIntMap{
		val: make(map[string]int),
	}
}

func (m *StringIntMap) Add(key string, value int) {
	m.val[key] = value
}

func (m *StringIntMap) Remove(key string) {
	delete(m.val, key)
}

func (m *StringIntMap) Copy() map[string]int {
	newMap := make(map[string]int)
	for k, v := range m.val {
		newMap[k] = v
	}
	return newMap
}

func (m *StringIntMap) Exists(key string) bool {
	_, ok := m.val[key]
	return ok
}

func (m *StringIntMap) Get(key string) (int, bool) {
	val, ok := m.val[key]
	return val, ok
}

func main() {
	//s := NewStringIntMap()
	//
	//s.Add("first", 1)
	//fmt.Println(*s)
	//s.Remove("first")
	//fmt.Println(*s)
	//
	//s.Add("first", 2)
	//newMap := s.Copy()
	//fmt.Println(newMap)
	//
	//fmt.Println(s.Exists("first"))
	//
	//val, ok := s.Get("first")
	//if !ok {
	//	fmt.Println("Нет значения")
	//} else {
	//	fmt.Println(val)
	//}
}
