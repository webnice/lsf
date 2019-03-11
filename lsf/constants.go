package lsf // import "gopkg.in/webnice/lsf.v1/lsf"

//import "gopkg.in/webnice/debug.v1"
//import "gopkg.in/webnice/log.v2"
//import ()

const (
	// SortAscending Сортировка от меньшего к большему, прямой порядок
	SortAscending = SortOrder(`asc`)

	// SortDescending Сортировка от большего к меньшему, обратный порядок
	SortDescending = SortOrder(`desc`)
)

const (
	// EqEqual Сравнение 'точное' (=)
	EqEqual = Equate(`eq`)

	// EqLessThan Сравнение 'меньше' (<)
	EqLessThan = Equate(`lt`)

	// EqLessEqual Сравнение 'меньше или равно' (<=)
	EqLessEqual = Equate(`le`)

	// EqGreaterThan Сравнение 'больше' (>)
	EqGreaterThan = Equate(`gt`)

	// EqGreaterEqual Сравнение 'больше или равно' (>=)
	EqGreaterEqual = Equate(`ge`)

	// EqNotEqual Сравнение 'не равно' (!=)
	EqNotEqual = Equate(`ne`)

	// EqLikeThan Сравнение 'похоже' - SQL эквивалент LIKE (LIKE)
	EqLikeThan = Equate(`ke`)

	// EqNotLikeThan Сравнение 'не похоже' - SQL эквивалент NOT LIKE (NOT LIKE)
	EqNotLikeThan = Equate(`kn`)
)
