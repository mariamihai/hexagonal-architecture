package ports

// DBPort stores every operation performed in a History table
type DBPort interface {
	CloseDBConnection() error
	AddToHistory(answer int, operation string) error
}
