package lsf // import "gopkg.in/webnice/lsf.v1/lsf"

//import "gopkg.in/webnice/debug.v1"
//import "gopkg.in/webnice/log.v2"
import (
	"fmt"
	"net/url"
	"strings"
)

// New creates a new object and return interface
func New() Interface {
	var qfo = &impl{
		order:  make([]*orderElement, 0),
		filter: make([]*filterElement, 0),
	}
	return qfo
}

// Limit Лимит результата
func (qfo *impl) Limit(from uint64, count uint64) Interface {
	qfo.from, qfo.count = from, count
	return qfo
}

// OrderReset Очистка параметров сортировки результата
func (qfo *impl) OrderReset() Interface {
	qfo.order = qfo.order[:0]
	return qfo
}

// Order Добавление сортировки результата по указанному полю в указанном порядке
func (qfo *impl) Order(fieldName string, sorting SortOrder) Interface {
	qfo.order = append(qfo.order, &orderElement{
		Field:   fieldName,
		Sorting: sorting,
	})
	return qfo
}

// FilterReset Очистка настроек фильтрации
func (qfo *impl) FilterReset() Interface {
	qfo.filter = qfo.filter[:0]
	return qfo
}

// Filter Добавление фильтрации по указанному полю указанным способом сравнения и значением сравнения
func (qfo *impl) Filter(fieldName string, ete Equate, value string) Interface {
	qfo.filter = append(qfo.filter, &filterElement{
		Field: fieldName,
		Ete:   ete,
		Value: value,
	})
	return qfo
}

// Reset Сброс всех настроек лимита, сортировки и фильтрации
func (qfo *impl) Reset() Interface {
	qfo.from, qfo.count = 0, 0
	return qfo.OrderReset().FilterReset()
}

// URL Возвращает настройки лимита, сортировки и фильтрации в виде объекта URL
// Если в функцию передан URL не nil, то настройки лимита, сортировки и фильтрации добавляются к переданному URL
func (qfo *impl) URL(srcURL *url.URL) (ret *url.URL, err error) {
	const urlDelimiter = `&`
	var buf []byte

	ret = &url.URL{}
	// Копирование, если передан URL
	if srcURL != nil {
		if buf, err = srcURL.MarshalBinary(); err != nil {
			err = fmt.Errorf("source net/url.URL failed: %s", err)
			return
		}
		if err = ret.UnmarshalBinary(buf); err != nil {
			err = fmt.Errorf("net/url.URL implementation error: %s", err)
			return
		}
	}
	// Добавление собственных значений
	if ret.RawQuery != "" {
		ret.RawQuery += urlDelimiter + qfo.rawQueryValues().Encode()
	} else {
		ret.RawQuery = qfo.rawQueryValues().Encode()
	}
	// Если собственный объект пустой, то в конце останется & - зачистка
	ret.RawQuery = strings.TrimRight(ret.RawQuery, urlDelimiter)

	return
}

// Создание значения запроса urlQueryParam на основе настроек лимита, сортировки и фильтрации
func (qfo *impl) rawQueryValues() (ret url.Values) {
	const (
		qLimiting  = `limit`
		qSorting   = `order`
		qFiltering = `filter`
	)
	var i int

	ret = make(url.Values)
	// Лимит
	if qfo.from > 0 || qfo.count > 0 {
		ret.Set(
			qLimiting,
			fmt.Sprintf("%d:%d", qfo.from, qfo.count),
		)
	}
	// Сортировка
	if len(qfo.order) > 0 {
		for i = range qfo.order {
			ret.Add(
				qSorting,
				fmt.Sprintf("%s:%s", qfo.order[i].Field, qfo.order[i].Sorting.String()),
			)
		}
	}
	// Фильтрация
	if len(qfo.filter) > 0 {
		for i = range qfo.filter {
			ret.Add(
				qFiltering,
				fmt.Sprintf("%s:%s:%s", qfo.filter[i].Field, qfo.filter[i].Ete.String(), qfo.filter[i].Value),
			)
		}
	}

	return
}
