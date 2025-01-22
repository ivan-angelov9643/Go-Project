sap.ui.define([
	'../BaseController',
	'sap/ui/model/json/JSONModel',
	'library-app/model/formatter',
	"sap/ui/core/Core",
	"sap/ui/core/mvc/XMLView",
	"sap/f/LayoutType",
	"sap/m/MessageToast",
	"sap/ui/model/Filter",
	"sap/ui/model/FilterOperator"
], function (BaseController, JSONModel, formatter, Core, XMLView, LayoutType, MessageToast, Filter, FilterOperator) {
	"use strict";
	return BaseController.extend("library-app.controller.book.Books", {
		formatter: formatter,

		onInit: async function () {
			const oRouter = this.getOwnerComponent().getRouter();
			oRouter.attachRoutePatternMatched(this.loadData, this);

			Core.getEventBus().subscribe("library-app", "booksUpdated", this.handleBooksUpdated, this);

			this.oBookModel = new JSONModel({
				count: null,
				page_size: null,
				page: null,
				data: null,
				total_pages: null,
			});
			this.oBookModel.setSizeLimit(Number.MAX_VALUE);
			this.getView().setModel(this.oBookModel, "book");

			this.oSelectedBookModel = new JSONModel(this.initBookModel());
			this.getView().setModel(this.oSelectedBookModel, "selectedBook");

			this.oRatingModel = new JSONModel({
				ratings: null,
				page_size: null,
				page: null,
				data: null,
				total_pages: null,
			});
			this.oRatingModel.setSizeLimit(Number.MAX_VALUE);
			this.getView().setModel(this.oRatingModel, "rating");

			await this.loadBooks(this.oBookModel, 1);

			this.getView().byId("titleSearch").setFilterFunction(function(sTerm, oItem) {
				return oItem.getText().match(new RegExp(sTerm, "i")) || oItem.getKey().match(new RegExp(sTerm, "i"));
			});
		},

		loadData: async function() {
			await this.loadBooks(this.oBookModel, this.oBookModel.getData().page);
		},

		onCreateBook: async function () {
			if (!this._oCreateBookDialog) {
				const oOwnerComponent = this.getOwnerComponent();
				oOwnerComponent.runAsOwner(() => {
					this._oCreateBookDialog = new XMLView({
						id: "createBookDialogView",
						viewName: "library-app.view.book.CreateBookDialog",
					});
					this.getView().addDependent(this._oCreateBookDialog);
				});
			}

			this._oCreateBookDialog.byId("createBookDialog").open();
		},

		onEditBook: async function (oEvent) {
			if (!this._oEditBookDialog) {
				const oOwnerComponent = this.getOwnerComponent();
				oOwnerComponent.runAsOwner(() => {
					this._oEditBookDialog = new XMLView({
						id: "editBookDialogView",
						viewName: "library-app.view.book.EditBookDialog",
					});
					this.getView().addDependent(this._oEditBookDialog);
				});
			}
			const oData = this.oSelectedBookModel.getData();
			const oDialogBookModel = this._oEditBookDialog.getModel("dialogBook");

			this.fillBookModel(oDialogBookModel, oData);
			this._oEditBookDialog.byId("editBookDialog").open();
		},

		onDeleteBook: async function (oEvent) {
			if (!this._oDeleteBookDialog) {
				const oOwnerComponent = this.getOwnerComponent();
				oOwnerComponent.runAsOwner(() => {
					this._oDeleteBookDialog = new XMLView({
						id: "deleteBookDialogView",
						viewName: "library-app.view.book.DeleteBookDialog",
					});
					this.getView().addDependent(this._oDeleteBookDialog);
				});
			}
			const oData = this.oSelectedBookModel.getData();
			const oDialogBookModel = this._oDeleteBookDialog.getModel("dialogBook");

			this.fillBookModel(oDialogBookModel, oData);
			this._oDeleteBookDialog.byId("deleteBookDialog").open();
		},

		onSelectBook: async function(oEvent) {
			const oSelectedBookData = oEvent.getParameter("listItem").getBindingContext("book").getObject();

			const token = await this.getOwnerComponent().getToken();

			const oBookData = await this.getBookData(token ,oSelectedBookData.id)
			oBookData.author_name = oSelectedBookData.author_name;
			oBookData.category_name = oSelectedBookData.category_name
			this.fillBookModel(this.oSelectedBookModel, oBookData)

			const user_id = this.getUserID(token);

			await this.reserveButtonUpdateVisible(user_id, oSelectedBookData.id);
			await this.rateButtonUpdateVisible(user_id, oSelectedBookData.id);

			await this.loadRatings(this.oRatingModel, 1, oSelectedBookData.id);

			const oFlexibleColumnLayout = this.getView().byId("flexibleColumnLayout");
			oFlexibleColumnLayout.setLayout(LayoutType.TwoColumnsBeginExpanded);

			const oTable = this.getView().byId("booksTable");
			if (oTable) {
				oTable.removeSelections(true);
			}
		},

		onNavBack: function() {
			const oFlexibleColumnLayout = this.getView().byId("flexibleColumnLayout");
			oFlexibleColumnLayout.setLayout(LayoutType.OneColumn);

			const oTable = this.getView().byId("booksTable");
			if (oTable) {
				oTable.removeSelections(true);
			}
		},

		onReserveBook: async function () {
			const token = await this.getOwnerComponent().getToken();
			const oBook = this.oSelectedBookModel.getData();
			const user_id = this.getUserID(token);

			if (await this.userHasActiveLoanOnBook(user_id, oBook.id)) {
				MessageToast.show("You already have an active loan on this book");
				return;
			}
			if (await this.userHasReservedBook(user_id, oBook.id)) {
				MessageToast.show("You already have a reservation for this book");
				return;
			}
			if (oBook.available_copies < 1) {
				MessageToast.show("There aren't any available copies at the moment");
				return;
			}

			const body = {
				book_id: oBook.id,
				user_id: user_id,
			};

			try {
				const createResponse = await this.sendRequest(
					`http://localhost:8080/api/reservations`,
					"POST",
					token,
					body
				);

				Core.getEventBus().publish("library-app", "reservationsUpdated", {make_reservation: true});

				MessageToast.show("Reservation successful!");

			} catch (error) {
				MessageToast.show(error.error || "Error reserving book");
			}
			this.oSelectedBookModel.setProperty("/available_copies", oBook.available_copies - 1);
		},

		onRateBook: async function () {
			if (!this._oCreateRatingDialog) {
				const oOwnerComponent = this.getOwnerComponent();
				oOwnerComponent.runAsOwner(() => {
					this._oCreateRatingDialog = new XMLView({
						id: "createRatingDialogView",
						viewName: "library-app.view.book.CreateRatingDialog",
					});
					this.getView().addDependent(this._oCreateRatingDialog);
				});
			}
			const oBookData = this.oSelectedBookModel.getData();
			const oDialogRatingModel = this._oCreateRatingDialog.getModel("dialogRating");

			this.fillRatingModel(oDialogRatingModel, {book_id: oBookData.id, book_title: oBookData.title} );
			this._oCreateRatingDialog.byId("createRatingDialog").open();
		},

		onExit: function () {
			Core.getEventBus().unsubscribe("library-app", "booksUpdated", this.handleBooksUpdated, this);
		},

		handleBooksUpdated: async function (ns, ev, eventData) {
			await this.loadData()

			const selectedBookID = this.oSelectedBookModel.getData().id;

			if (eventData.from_reservations || eventData.from_loans || eventData.from_ratings) {
				const token = await this.getOwnerComponent().getToken();
				const user_id = this.getUserID(token);
				await this.reserveButtonUpdateVisible(user_id, selectedBookID);
				await this.rateButtonUpdateVisible(user_id, selectedBookID);
			}

			if (eventData.from_ratings && eventData.book_id === selectedBookID) {
				await this.loadRatings(this.oRatingModel, this.oRatingModel.getData().page, selectedBookID);
			}

			if (!eventData.from_reservations) {
				Core.getEventBus().publish("library-app", "reservationsUpdated", {from_books: true});
			}

			if (!eventData.from_loans) {
				Core.getEventBus().publish("library-app", "loansUpdated", {from_books: true});
			}

			if (eventData.delete) {
				this.onNavBack();
				return;
			}

			if (eventData.edit_book) {
				this.fillBookModel(this.oSelectedBookModel, eventData);
			}
		},

		reserveButtonUpdateVisible: async function (user_id, book_id) {
			var oButton = this.byId("reserveBookButton");
			oButton.setVisible(
				!await this.userHasReservedBook(user_id, book_id) &&
				!await this.userHasActiveLoanOnBook(user_id, book_id)
			);
		},

		rateButtonUpdateVisible: async function (user_id, book_id) {
			var oButton = this.byId("rateBookButton");
			oButton.setVisible(
				await this.userHasLoanOnBook(user_id, book_id) &&
				!await this.userHasRatedBook(user_id, book_id)
			);
		},

		onTitleSearchChange: function(oEvent) {
			this.sTitleSearch = oEvent.getParameter("value");
			this._applyCombinedFilters();
		},

		onAuthorNameSearchChange: function(oEvent) {
			this.sAuthorNameSearch = oEvent.getParameter("value");
			this._applyCombinedFilters();
		},

		onCategoryNameSearchChange: function(oEvent) {
			this.sCategoryNameSearch = oEvent.getParameter("value");
			this._applyCombinedFilters();
		},

		onLanguageSearchChange: function(oEvent) {
			this.sLanguageNameSearch = oEvent.getParameter("value");
			this._applyCombinedFilters();
		},

		_applyCombinedFilters: function() {
			let aFilters = [];

			if (this.sTitleSearch && this.sTitleSearch.trim() !== "") {
				aFilters.push(
					new Filter("title", FilterOperator.Contains, this.sTitleSearch)
				);
			}

			if (this.sAuthorNameSearch && this.sAuthorNameSearch.trim() !== "") {
				aFilters.push(
					new Filter("author_name", FilterOperator.Contains, this.sAuthorNameSearch)
				);
			}

			if (this.sCategoryNameSearch && this.sCategoryNameSearch.trim() !== "") {
				aFilters.push(
					new Filter("category_name", FilterOperator.Contains, this.sCategoryNameSearch)
				);
			}

			if (this.sLanguageNameSearch && this.sLanguageNameSearch.trim() !== "") {
				aFilters.push(
					new Filter("language", FilterOperator.Contains, this.sLanguageNameSearch)
				);
			}

			let oTable = this.getView().byId("booksTable");
			let oBinding = oTable.getBinding("items");

			oBinding.filter(aFilters);
		},

		onPreviousPage: async function () {
			await this.loadBooks(this.oBookModel, this.oBookModel.getData().page - 1);
		},

		onNextPage: async function () {
			await this.loadBooks(this.oBookModel, this.oBookModel.getData().page + 1);
		},

		onPreviousPageRatings: async function () {
			await this.loadRatings(this.oRatingModel, this.oRatingModel.getData().page - 1, this.oSelectedBookModel.getData().id);
		},

		onNextPageRatings: async function () {
			await this.loadRatings(this.oRatingModel, this.oRatingModel.getData().page + 1, this.oSelectedBookModel.getData().id);
		},
	});
});