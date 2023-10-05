package helpers


type PaginationResultUser struct {
	//Data       []domainUser.User
	Total      int64
	Limit      int64
	Current    int64
	NextCursor uint
	PrevCursor uint
	NumPages   int64
}
func PaginationValues(limit int64, page int64, total int64) (numPages int64, nextCursor int64, prevCursor int64) {
	numPages = (total + limit - 1) / limit
	if page < numPages {
		nextCursor = page + 1
	}
	if page > 1 {
		prevCursor = page - 1
	}
	return
}