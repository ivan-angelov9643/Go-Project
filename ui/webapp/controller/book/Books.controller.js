sap.ui.define([
	'../BaseController',
	'sap/ui/model/json/JSONModel',
	'library-app/model/formatter',
	"sap/ui/core/Core",
	"sap/ui/core/mvc/XMLView",
	"sap/f/LayoutType",
	"sap/m/MessageToast",
], function (BaseController, JSONModel, formatter, Core, XMLView, LayoutType, MessageToast) {
	"use strict";
	return BaseController.extend("library-app.controller.book.Books", {
		formatter: formatter,

		onInit: async function () {
			Core.getEventBus().subscribe("library-app", "booksUpdated", this.handleBooksUpdated, this);

			this.oBookModel = new JSONModel({
				books: null,
			});
			this.oBookModel.setSizeLimit(Number.MAX_VALUE);
			this.getView().setModel(this.oBookModel, "book");

			await this.loadData();
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

		onExit: function () {
			Core.getEventBus().unsubscribe("library-app", "booksUpdated", this.handleBooksUpdated, this);
		},

		onSelectBook: async function(oEvent) {
			const oSelectedBook = oEvent.getParameter("listItem").getBindingContext("book").getObject();
			oSelectedBook.available_copies = await this.getAvailableCopies(oSelectedBook.id)

			this.oSelectedBookModel = new JSONModel(oSelectedBook);
			this.getView().setModel(this.oSelectedBookModel, "selectedBook");

			this.oReviewModel = new JSONModel({
				reviews: null,
			});
			this.oReviewModel.setSizeLimit(Number.MAX_VALUE);
			this.getView().setModel(this.oReviewModel, "review");
			this.loadReviews(this.oReviewModel, oSelectedBook.id);

			const oFlexibleColumnLayout = this.getView().byId("flexibleColumnLayout");
			oFlexibleColumnLayout.setLayout(LayoutType.TwoColumnsBeginExpanded);
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
			const book_id = this.oSelectedBookModel.getData().id
			const user_id = this.getUserID(token);

			if (await this.userHasActiveLoanOnBook(user_id, book_id)) {
				MessageToast.show("You already have an active loan on this book");
				return;
			}
			if (await this.userHasReservedBook(user_id, book_id)) {
				MessageToast.show("You already have a reservation for this book");
				return;
			}
			if (await this.getAvailableCopies(book_id) < 1) {
				MessageToast.show("There aren't any available copies at the moment");
				return;
			}

			const body = {
				book_id: book_id,
				user_id: user_id,
			};

			try {
				const createResponse = await this.sendRequest(
					`http://localhost:8080/api/reservations`,
					"POST",
					token,
					body
				);

				Core.getEventBus().publish("library-app", "reservationsUpdated");

				MessageToast.show("Reservation successful!");

			} catch (error) {
				MessageToast.show(error.error || "Error reserving book");
			}
			this.oSelectedBookModel.setProperty("/available_copies", this.getAvailableCopies(book_id));
		},

		loadData: async function () {
			const token = await this.getOwnerComponent().getToken();

			const [booksData, authorsData, categoriesData] = await Promise.all([
				this.sendRequest('http://localhost:8080/api/books', "GET", token),
				this.sendRequest('http://localhost:8080/api/authors', "GET", token),
				this.sendRequest('http://localhost:8080/api/categories', "GET", token)
			]);

			booksData.forEach(book => {
				const author = authorsData.find(a => a.id === book.author_id);
				const category = categoriesData.find(c => c.id === book.category_id);

				book.author_name = author ? `${author.first_name} ${author.last_name}` : 'Unknown Author';
				book.category_name = category ? category.name : 'Unknown Category';
			});
			this.oBookModel.setProperty("/books", booksData);
		},

		handleBooksUpdated: async function (ns, ev, eventData) {
			await this.loadData()

			if (eventData.delete) {
				this.onNavBack();
				return;
			}

			const selectedBookId = this.oSelectedBookModel.getData().id;
			const selectedBookData = this.oBookModel.getData().books.find(book => book.id === selectedBookId);
			selectedBookData.available_copies = await this.getAvailableCopies(selectedBookId)
			this.fillBookModel(this.oSelectedBookModel, selectedBookData);
		},
	});
});