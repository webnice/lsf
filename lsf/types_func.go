package lsf // import "gopkg.in/webnice/lsf.v1/lsf"

//import "gopkg.in/webnice/debug.v1"
//import "gopkg.in/webnice/log.v2"
//import ()

// SortOrder Тип порядка сортировки результата
type SortOrder string

// Equate Тип способа сравнения значения для фильтрации
type Equate string

// String Возвращает значение типа как строку
func (so SortOrder) String() string { return string(so) }

// String Возвращает значение типа как строку
func (et Equate) String() string { return string(et) }
