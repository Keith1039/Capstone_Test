package parameters

const DEFAULTDATECODE = RANDOM
const DEFAULTDATERANGE = "2001-01-01,2024-12-31"

type DateParser struct {
	Column column
}

func (p *DateParser) ParseColumn() (string, error) {

	return "", nil
}
