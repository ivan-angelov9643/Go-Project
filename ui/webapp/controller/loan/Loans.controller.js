sap.ui.define([
	'../BaseController',
	'sap/ui/model/json/JSONModel',
	'library-app/model/formatter',
	"sap/ui/core/Core",
	"sap/ui/core/mvc/XMLView",
], function (BaseController, JSONModel, formatter, Core, XMLView) {
	"use strict";
	return BaseController.extend("library-app.controller.loan.Loans", {
		formatter: formatter,

		onInit: function () {
			Core.getEventBus().subscribe("library-app", "loansUpdated", this.handleLoansUpdated, this);

			this.oLoanModel = new JSONModel({
				loans: null,
			});
			this.oLoanModel.setSizeLimit(Number.MAX_VALUE);
			this.getView().setModel(this.oLoanModel, "loan");
			this.loadLoans();
		},

		onExtendLoan: async function (oEvent) {
			if (!this._oExtendLoanDialog) {
				const oOwnerComponent = this.getOwnerComponent();
				oOwnerComponent.runAsOwner(() => {
					this._oExtendLoanDialog = new XMLView({
						id: "extendLoanDialogView",
						viewName: "library-app.view.loan.ExtendLoanDialog",
					});
					this.getView().addDependent(this._oExtendLoanDialog);
				});
			}
			const oContext = oEvent.getSource().getBindingContext("loan");
			const oData = oContext.getObject();
			const oDialogLoanModel = this._oExtendLoanDialog.getModel("dialogLoan");

			this.fillLoanModel(oDialogLoanModel, oData);
			this._oExtendLoanDialog.byId("extendLoanDialog").open();
		},

		onReturnBook: async function (oEvent) {
			if (!this._oReturnBookDialog) {
				const oOwnerComponent = this.getOwnerComponent();
				oOwnerComponent.runAsOwner(() => {
					this._oReturnBookDialog = new XMLView({
						id: "returnBookDialogView",
						viewName: "library-app.view.loan.ReturnBookDialog",
					});
					this.getView().addDependent(this._oReturnBookDialog);
				});
			}
			const oContext = oEvent.getSource().getBindingContext("loan");
			const oData = oContext.getObject();
			const oDialogLoanModel = this._oReturnBookDialog.getModel("dialogLoan");

			this.fillLoanModel(oDialogLoanModel, oData);
			this._oReturnBookDialog.byId("returnBookDialog").open();
		},

		onDeleteLoan: async function (oEvent) {
			if (!this._oDeleteLoanDialog) {
				const oOwnerComponent = this.getOwnerComponent();
				oOwnerComponent.runAsOwner(() => {
					this._oDeleteLoanDialog = new XMLView({
						id: "deleteLoanDialogView",
						viewName: "library-app.view.loan.DeleteLoanDialog",
					});
					this.getView().addDependent(this._oDeleteLoanDialog);
				});
			}
			const oContext = oEvent.getSource().getBindingContext("loan");
			const oData = oContext.getObject();
			const oDialogLoanModel = this._oDeleteLoanDialog.getModel("dialogLoan");

			this.fillLoanModel(oDialogLoanModel, oData);
			this._oDeleteLoanDialog.byId("deleteLoanDialog").open();
		},

		onExit: function () {
			Core.getEventBus().unsubscribe("library-app", "loansUpdated", this.handleLoansUpdated, this);
		},

		loadLoans: async function () {
			const token = await this.getOwnerComponent().getToken();

			const [loansData, usersData, booksData] = await Promise.all([
				this.sendRequest('http://localhost:8080/api/loans', "GET", token),
				this.sendRequest('http://localhost:8080/api/users', "GET", token),
				this.sendRequest('http://localhost:8080/api/books', "GET", token)
			]);

			loansData.forEach(loan => {
				const user = usersData.find(u => u.id === loan.user_id);
				const book = booksData.find(b => b.id === loan.book_id);

				loan.user_name = user ? user.preferred_username : 'Unknown User';
				loan.book_title = book ? book.title : 'Unknown Book';
			});
			this.oLoanModel.setProperty("/loans", loansData);
		},

		handleLoansUpdated: async function (ns, ev, eventData) {
			this.loadLoans();

			if (!eventData.from_books && (eventData.delete_loan || eventData.return_book)) {
				Core.getEventBus().publish("library-app", "booksUpdated", {from_loans: true});
			}
		},

	});
});