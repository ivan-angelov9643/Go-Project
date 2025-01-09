sap.ui.define([
	"sap/ui/core/mvc/Controller",
	"sap/ui/core/UIComponent"
], function(Controller, UIComponent) {
	"use strict";

	return Controller.extend("library-app.controller.BaseController", {

		getRouter : function () {
			return UIComponent.getRouterFor(this);
		},

		getBundleTextByModel: function(sI18nKey, oResourceModel, aPlaceholderValues){
			return oResourceModel.getResourceBundle().then(function(oBundle){
				return oBundle.getText(sI18nKey, aPlaceholderValues);
			});
		},

		getBundleText: function (sI18nKey, aPlaceholderValues) {
			let i18nModel = this.getView().getModel("i18n");
			if(!i18nModel){
				i18nModel = this.getOwnerComponent().getModel("i18n")
			}
			return this.getBundleTextByModel(sI18nKey, i18nModel, aPlaceholderValues);
		},

		sendRequest: function (url, method, token, body = null) {
			return new Promise((resolve, reject) => {
				const options = {
					url: url,
					type: method,
					contentType: "application/json",
					beforeSend: function (xhr) {
						if (token) {
							xhr.setRequestHeader("Authorization", `Bearer ${token}`);
						}
					},
					success: function (response) {
						resolve(response);
					},
					error: function (xhr, status, error) {
						console.error(`Error in request to ${url}:`, xhr.responseText || status || error);
						const parsedError = JSON.parse(xhr.responseText);
						reject(parsedError);
					}
				};

				if (body) {
					options.data = JSON.stringify(body);
				}

				jQuery.ajax(options);
			});
		},

		getAvailableCopies: async function (book_id) {
			const token = await this.getOwnerComponent().getToken();

			const bookResponse = await this.sendRequest(`http://localhost:8080/api/books/${book_id}`, "GET", token);
			const totalCopies = bookResponse.total_copies;

			const loansResponse = await this.sendRequest(`http://localhost:8080/api/loans`, "GET", token);
			const activeLoans = loansResponse.filter(
				(loan) => loan.book_id === book_id && loan.status === "active"
			);

			const reservationsResponse = await this.sendRequest(`http://localhost:8080/api/reservations`, "GET", token);
			const activeReservations = reservationsResponse.filter(
				(reservation) => reservation.book_id === book_id
			);

			return totalCopies - activeLoans.length - activeReservations.length
		},

		userHasReservedBook: async function (user_id, book_id) {
			const token = await this.getOwnerComponent().getToken();

			const reservationsResponse = await this.sendRequest(
				`http://localhost:8080/api/reservations`,
				"GET",
				token
			);

			const reservation = reservationsResponse.find(
				(reservation) =>
					reservation.user_id === user_id && reservation.book_id === book_id
			);

			return reservation !== undefined;
		},

		userHasActiveLoanOnBook: async function (user_id, book_id) {
			const token = await this.getOwnerComponent().getToken();

			const loansResponse = await this.sendRequest(
				`http://localhost:8080/api/loans`,
				"GET",
				token
			);

			const activeLoan = loansResponse.find(
				(loan) =>
					loan.user_id === user_id &&
					loan.book_id === book_id &&
					loan.status === "active"
			);

			return activeLoan !== undefined;
		},

		getUserID: function (token) {
			return JSON.parse(atob(token.split(".")[1])).sub;
		},

		loadAuthors: async function (model) {
			const token = await this.getOwnerComponent().getToken();
			const authorsData = await this.sendRequest("http://localhost:8080/api/authors", "GET", token);

			model.setProperty("/authors", authorsData);
		},

		loadCategories: async function (model) {
			const token = await this.getOwnerComponent().getToken();
			const categoriesData = await this.sendRequest("http://localhost:8080/api/categories", "GET", token);

			model.setProperty("/categories", categoriesData);
		},

		loadReviews: async function (model, book_id = null) {
			const token = await this.getOwnerComponent().getToken();

			let [reviewData, usersData, booksData] = await Promise.all([
				this.sendRequest('http://localhost:8080/api/reviews', "GET", token),
				this.sendRequest('http://localhost:8080/api/users', "GET", token),
				this.sendRequest('http://localhost:8080/api/books', "GET", token)
			]);

			if (book_id) {
				reviewData = reviewData.filter(review => review.book_id === book_id);
			}

			reviewData.forEach(review => {
				const user = usersData.find(u => u.id === review.user_id);
				const book = booksData.find(b => b.id === review.book_id);

				review.user_name = user ? user.preferred_username : 'Unknown User';
				review.book_title = book ? book.title : 'Unknown Book';
			});

			model.setProperty("/reviews", reviewData);
		},

		toISO8601: function (dateString) {
			const date = new Date(dateString);
			if (isNaN(date.getTime())) {
				throw new Error(`Invalid date format: ${dateString}`);
			}
			return date.toISOString(); // Formats as ISO 8601 (YYYY-MM-DDTHH:mm:ss.sssZ)
		},

		initUserModel: function () {
			return {
				id: null,
				preferred_username: null,
				given_name: null,
				family_name: null,
				email: null
			};
		},

		fillUserModel: function (userModel, data) {
			userModel.setProperty("/id", data.id);
			userModel.setProperty("/preferred_username", data.preferred_username);
			userModel.setProperty("/given_name", data.given_name);
			userModel.setProperty("/family_name", data.family_name);
			userModel.setProperty("/email", data.email);
		},

		initAuthorModel: function () {
			return {
				id: null,
				first_name: null,
				last_name: null,
				nationality: null,
				birth_date: null,
				death_date: null,
				bio: null,
				website: null
			};
		},

		fillAuthorModel: function (authorModel, data) {
			authorModel.setProperty("/id", data.id)
			authorModel.setProperty("/first_name", data.first_name);
			authorModel.setProperty("/last_name", data.last_name);
			authorModel.setProperty("/nationality", data.nationality);
			authorModel.setProperty("/birth_date", data.birth_date);
			authorModel.setProperty("/death_date", data.death_date);
			authorModel.setProperty("/bio", data.bio);
			authorModel.setProperty("/website", data.website);
		},

		initCategoryModel: function () {
			return {
				id: null,
				name: null,
				description: null,
			};
		},

		fillCategoryModel: function (categoryModel, data) {
			categoryModel.setProperty("/id", data.id);
			categoryModel.setProperty("/name", data.name);
			categoryModel.setProperty("/description", data.description);
		},

		initBookModel: function () {
			return {
				id: null,
				title: null,
				year: null,
				author_id: null,
				author_name: null,
				category_id: null,
				category_name: null,
				total_copies: null,
				available_copies: null,
				language: null,
			};
		},

		fillBookModel: function (bookModel, data) {
			bookModel.setProperty("/id", data.id);
			bookModel.setProperty("/title", data.title);
			bookModel.setProperty("/year", data.year);
			bookModel.setProperty("/author_id", data.author_id);
			bookModel.setProperty("/author_name", data.author_name);
			bookModel.setProperty("/category_id", data.category_id);
			bookModel.setProperty("/category_name", data.category_name);
			bookModel.setProperty("/total_copies", data.total_copies);
			bookModel.setProperty("/available_copies", data.available_copies);
			bookModel.setProperty("/language", data.language);
		},

		initLoanModel: function () {
			return {
				id: null,
				user_id: null,
				book_id: null,
				start_date: null,
				due_date: null,
				return_date: null,
				status: null,
				days_to_extend: null
			};
		},

		fillLoanModel: function (loanModel, data) {
			loanModel.setProperty("/id", data.id);
			loanModel.setProperty("/user_id", data.user_id);
			loanModel.setProperty("/user_name", data.user_name);
			loanModel.setProperty("/book_id", data.book_id);
			loanModel.setProperty("/book_title", data.book_title);
			loanModel.setProperty("/start_date", data.start_date);
			loanModel.setProperty("/due_date", data.due_date);
			loanModel.setProperty("/return_date", data.return_date);
			loanModel.setProperty("/status", data.status);
			loanModel.setProperty("/days_to_extend", data.days_to_extend);
		},

		initReservationModel: function () {
			return {
				id: null,
				book_id: null,
				book_title: null,
				user_id: null,
				user_name: null,
				created_at: null,
				expiry_date: null,
				days_to_extend: null,
				hours_to_extend: null,
				minutes_to_extend: null,
			};
		},

		fillReservationModel: function (reservationModel, data) {
			reservationModel.setProperty("/id", data.id);
			reservationModel.setProperty("/book_id", data.book_id);
			reservationModel.setProperty("/book_title", data.book_title);
			reservationModel.setProperty("/user_id", data.user_id);
			reservationModel.setProperty("/user_name", data.user_name);
			reservationModel.setProperty("/created_at", data.created_at);
			reservationModel.setProperty("/expiry_date", data.expiry_date);
			reservationModel.setProperty("/days_to_extend", data.days_to_extend);
			reservationModel.setProperty("/hours_to_extend", data.hours_to_extend);
			reservationModel.setProperty("/minutes_to_extend", data.minutes_to_extend);
		}
	});
});