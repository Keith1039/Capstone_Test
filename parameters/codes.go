package parameters

// ENUMS that will be used to generate SQL strings
const (
	RANDOM = iota
	STATIC
	SEQ
	UUID
	BOOL
	FIRSTNAME
	LASTNAME
	FULLNAME
	ADDRESS
	REGEX
	EMAIL
	NULL
)

var defaultRegex = map[string]string{
	"EMAIL": "^[\\w-\\.]+@([\\w-]+\\.)+[\\w-]{2,4}$",
}
