The Library Management System (LMS) offers a solution for managing books, authors, categories, reservations and loans. The LMS allows Librarians to manage the library's inventory of books, track loans, and handle user reservations, while Members can search for books, review them, reserve them, and check their loan history. The LMS is developed using Golang, gorilla/mux, gorm, and PostgreSQL. It uses a web-based client powered by frontend framework Open UI 5 to create a dynamic and user-friendly experience. 

The system's security and authentication are handled using Keycloak, ensuring robust role-based access control and secure user management. This integration allows for centralized authentication and fine-grained permission control across different user roles. The main roles in the system include:

●	Member – can view books, authors, and categories, make reservations, create ratings, and track loan history.
●	Librarian – can manage books, authors, categories, reservations, loans and ratings.
●	Administrator – has full control over books, authors, categories, loans, reservations, ratings and users.
