package lsf // import "gopkg.in/webnice/lsf.v1/lsf"

//import "gopkg.in/webnice/debug.v1"
//import "gopkg.in/webnice/log.v2"
import (
	"net/url"
)

// Interface is an interface of package
type Interface interface {
	// Limit Лимит результата
	Limit(from uint64, count uint64) Interface

	// OrderReset Очистка параметров сортировки результата
	OrderReset() Interface

	// Order Добавление сортировки результата по указанному полю в указанном порядке
	Order(fieldName string, sorting SortOrder) Interface

	// FilterReset Очистка настроек фильтрации
	FilterReset() Interface

	// Filter Добавление фильтрации по указанному полю указанным способом сравнения и значением сравнения
	Filter(fieldName string, ete Equate, value string) Interface

	// Reset Сброс всех настроек лимита, сортировки и фильтрации
	Reset() Interface

	// URL Возвращает настройки лимита, сортировки и фильтрации в виде объекта URL
	// Если в функцию передан URL не nil, то настройки лимита, сортировки и фильтрации добавляются к переданному URL
	URL(srcURL *url.URL) (ret *url.URL, err error)
}

// impl is an implementation of package
type impl struct {
	from   uint64           // Лимит возвращения результата начиная с указанной позиции
	count  uint64           // Лимит количества возвращаемых значений
	order  []*orderElement  // Сортировка результата
	filter []*filterElement // Фильтрация результата
}

// Структура объекта описания сортировки
type orderElement struct {
	Field   string    // Название поля
	Sorting SortOrder // Порядок сортировки значений поля
}

// Структура объекта описания фильтрации результата
type filterElement struct {
	Field string // Название поля
	Ete   Equate // Способ сравнения значений поля
	Value string // Значение сравнения
}
