sap.ui.define([
	"sap/ui/core/mvc/Controller",
	"sap/ui/core/UIComponent",
	"sap/ui/core/mvc/XMLView",
], function(Controller, UIComponent, XMLView) {
	"use strict";

	return Controller.extend("library-app.controller.BaseController", {
		page_size: 10,

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

		getBookData: async function (token, book_id) {
			return await this.sendRequest(
				`http://localhost:8080/api/books/${book_id}`,
				"GET",
				token
			)
		},

		userHasReservedBook: async function (user_id, book_id) {
			const token = await this.getOwnerComponent().getToken();
			const reservationsData = await this.sendRequest(
				`http://localhost:8080/api/reservations?user_id=${user_id}&book_id=${book_id}`,
				"GET",
				token
			);

			return reservationsData.count > 0;
		},

		userHasActiveLoanOnBook: async function (user_id, book_id) {
			const token = await this.getOwnerComponent().getToken();
			const loansData = await this.sendRequest(
				`http://localhost:8080/api/loans?user_id=${user_id}&book_id=${book_id}&status=active`,
				"GET",
				token
			);

			return loansData.count > 0;
		},

		userHasLoanOnBook: async function (user_id, book_id) {
			const token = await this.getOwnerComponent().getToken();
			const loansData = await this.sendRequest(
				`http://localhost:8080/api/loans?user_id=${user_id}&book_id=${book_id}`,
				"GET",
				token
			);

			return loansData.count > 0;
		},

		userHasRatedBook: async function (user_id, book_id) {
			const token = await this.getOwnerComponent().getToken();
			const ratingsData = await this.sendRequest(
				`http://localhost:8080/api/ratings?user_id=${user_id}&book_id=${book_id}`,
				"GET",
				token
			);

			return ratingsData.count > 0;
		},

		getUserID: function (token) {
			return JSON.parse(atob(token.split(".")[1])).sub;
		},

		loadBooks: async function (model, page, title = null, authorName = null,
								   categoryName = null, language = null) {
			const token = await this.getOwnerComponent().getToken();
			const booksData = await this.sendRequest(
				`http://localhost:8080/api/books?page_size=${this.page_size}&page=${page}&
				title=${title}&author_name=${authorName}&
				category_name=${categoryName}&language=${language}`,
				"GET",
				token
			);

			booksData.data.forEach(book => {
				book.author_name = book.author_name ? book.author_name : 'Unknown Author';
				book.category_name = book.category_name ? book.category_name : 'Unknown Category';
			});

			model.setProperty("/count", booksData.count);
			model.setProperty("/page_size", booksData.page_size);
			model.setProperty("/page", booksData.page);
			model.setProperty("/data", booksData.data);
			model.setProperty("/total_pages", Math.ceil(booksData.count / booksData.page_size));
		},

		loadAuthors: async function (model, page, authorName = null) {
			const token = await this.getOwnerComponent().getToken();
			const authorsData = await this.sendRequest(
				`http://localhost:8080/api/authors?page_size=${this.page_size}&page=${page}&
				author_name=${authorName}`,
				"GET",
				token
			);

			model.setProperty("/count", authorsData.count);
			model.setProperty("/page_size", authorsData.page_size);
			model.setProperty("/page", authorsData.page);
			model.setProperty("/data", authorsData.data);
			model.setProperty("/total_pages", Math.ceil(authorsData.count / authorsData.page_size));
		},

		loadCategories: async function (model, page, categoryName = null) {
			const token = await this.getOwnerComponent().getToken();
			const categoriesData = await this.sendRequest(
				`http://localhost:8080/api/categories?page_size=${this.page_size}&page=${page}&
				category_name=${categoryName}`,
				"GET",
				token
			);

			model.setProperty("/count", categoriesData.count);
			model.setProperty("/page_size", categoriesData.page_size);
			model.setProperty("/page", categoriesData.page);
			model.setProperty("/data", categoriesData.data);
			model.setProperty("/total_pages", Math.ceil(categoriesData.count / categoriesData.page_size));
		},

		loadReservations: async function (model, page, sortBy = null, sortOrder = null,
										  username = null, title = null) {
			const token = await this.getOwnerComponent().getToken();
			const reservationsData = await this.sendRequest(
				`http://localhost:8080/api/reservations?page_size=${this.page_size}&page=${page}&
				sort_by=${sortBy}&sort_order=${sortOrder}&username=${username}&title=${title}`,
				"GET",
				token
			);

			model.setProperty("/count", reservationsData.count);
			model.setProperty("/page_size", reservationsData.page_size);
			model.setProperty("/page", reservationsData.page);
			model.setProperty("/data", reservationsData.data);
			model.setProperty("/total_pages", Math.ceil(reservationsData.count / reservationsData.page_size));
		},

		loadLoans: async function (model, page, sortBy = null, sortOrder = null, status,
								   username = null, title = null) {
			const token = await this.getOwnerComponent().getToken();
			const loansData = await this.sendRequest(
				`http://localhost:8080/api/loans?page_size=${this.page_size}&page=${page}&
				sort_by=${sortBy}&sort_order=${sortOrder}&status=${status}&username=${username}&title=${title}`,
				"GET",
				token
			);

			model.setProperty("/count", loansData.count);
			model.setProperty("/page_size", loansData.page_size);
			model.setProperty("/page", loansData.page);
			model.setProperty("/data", loansData.data);
			model.setProperty("/total_pages", Math.ceil(loansData.count / loansData.page_size));
		},

		loadRatings: async function (model, page, book_id = null) {
			const token = await this.getOwnerComponent().getToken();
			let ratingData = await this.sendRequest(
				`http://localhost:8080/api/ratings?page_size=${this.page_size}&page=${page}&book_id=${book_id}`,
				"GET",
				token
			);

			model.setProperty("/count", ratingData.count);
			model.setProperty("/page_size", ratingData.page_size);
			model.setProperty("/page", ratingData.page);
			model.setProperty("/data", ratingData.data);
			model.setProperty("/total_pages", Math.ceil(ratingData.count / ratingData.page_size));
		},

		loadUsers: async function (model, page) {
			const token = await this.getOwnerComponent().getToken();
			const userData = await this.sendRequest(
				`http://localhost:8080/api/users?page_size=${this.page_size}&page=${page}`,
				"GET",
				token
			);

			model.setProperty("/count", userData.count);
			model.setProperty("/page_size", userData.page_size);
			model.setProperty("/page", userData.page);
			model.setProperty("/data", userData.data);
			model.setProperty("/total_pages", Math.ceil(userData.count / userData.page_size));
		},

		loadCurrentUser: async function (model) {
			const token = await this.getOwnerComponent().getToken();
			const userID = this.getUserID(token);
			const userData = await this.sendRequest(`http://localhost:8080/api/users/${userID}`, "GET", token);

			this.fillUserModel(model, userData);
		},

		AppendData: function (displayModel, model) {
			const displayData = displayModel.getData().data;
			const newData = model.getData().data;

			displayData.push(...newData);

			displayModel.setProperty("/data", displayData);
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
				loan_duration: null,
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
			loanModel.setProperty("/loan_duration", data.loan_duration);
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
		},

		initRatingModel: function () {
			return {
				id: null,
				user_id: null,
				book_id: null,
				book_title: null,
				content: null,
				value: null
			};
		},

		fillRatingModel: function (ratingModel, data) {
			ratingModel.setProperty("/id", data.id);
			ratingModel.setProperty("/user_id", data.user_id);
			ratingModel.setProperty("/book_id", data.book_id);
			ratingModel.setProperty("/book_title", data.book_title);
			ratingModel.setProperty("/content", data.content);
			ratingModel.setProperty("/value", data.value);
		},

		onOpenAuthorDialog: async function () {
			if (!this._oAuthorSelectDialog || this._oAuthorSelectDialog.bIsDestroyed) {
				const oOwnerComponent = this.getOwnerComponent();
				oOwnerComponent.runAsOwner(() => {
					this._oAuthorSelectDialog= new XMLView({
						id: "authorSelectDialogView",
						viewName: "library-app.view.book.AuthorSelectDialog",
					});
					this.getView().addDependent(this._oAuthorSelectDialog);
				});
			}
			this._oAuthorSelectDialog.byId("authorSelectDialog").open();
		},

		onOpenCategoryDialog: async function () {
			if (!this._oCategorySelectDialog || this._oCategorySelectDialog.bIsDestroyed) {
				const oOwnerComponent = this.getOwnerComponent();
				oOwnerComponent.runAsOwner(() => {
					this._oCategorySelectDialog = new XMLView({
						id: "categorySelectDialogView",
						viewName: "library-app.view.book.CategorySelectDialog",
					});
					this.getView().addDependent(this._oCategorySelectDialog);
				});
			}
			this._oCategorySelectDialog.byId("categorySelectDialog").open();
		},
	});
});