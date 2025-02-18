package graph

import "awesomeProject/library-app/managers"

type Resolver struct {
	//AuthorManager managers.AuthorManagerInterface
	BookManager     *managers.BookManagerInterface
	CategoryManager *managers.CategoryManagerInterface
	//LoanManager        *managers.LoanManagerInterface
	//RatingManager      *managers.RatingManagerInterface
	//ReservationManager *managers.ReservationManagerInterface
	//UserManager        *managers.UserManagerInterface
}
