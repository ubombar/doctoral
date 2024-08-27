package doctoral

type IdentifierType string

const (
	UNKNOWN IdentifierType = "unknown"
	PDF     IdentifierType = "pdf"
)

func GetTypeOfIdentifier(identifier string) IdentifierType {
	return PDF
}
