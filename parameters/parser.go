package parameters

type ColumnParser interface {
	ParseColumn() (string, error)
}

var parserMap = map[string]ColumnParser{
	"INT": &IntColumnParser{},
}

func getColumnParser(dataType string) ColumnParser {
	return parserMap[dataType]
}
