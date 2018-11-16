package services

type ICalisApiService interface {
	GetCipByIsbn(isbn string) ([]string, error)
	UpdateLocalBookCip() (map[string][]string, error)
}
