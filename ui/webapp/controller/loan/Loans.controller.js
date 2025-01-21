sap.ui.define([
	'../BaseController',
	'sap/ui/model/json/JSONModel',
	'library-app/model/formatter',
	"sap/ui/core/Core",
	"sap/ui/core/mvc/XMLView",
	"sap/ui/model/Filter",
	"sap/ui/model/FilterOperator",
	"sap/ui/model/Sorter"
], function (BaseController, JSONModel, formatter, Core, XMLView, Filter, FilterOperator, Sorter) {
	"use strict";
	return BaseController.extend("library-app.controller.loan.Loans", {
		formatter: formatter,

		onInit: async function () {
			const oRouter = this.getOwnerComponent().getRouter();
			oRouter.attachRoutePatternMatched(this.loadData, this);

			Core.getEventBus().subscribe("library-app", "loansUpdated", this.handleLoansUpdated, this);

			this.oLoanModel = new JSONModel({
				count: null,
				page_size: null,
				page: null,
				data: null,
				total_pages: null,
			});
			this.oLoanModel.setSizeLimit(Number.MAX_VALUE);
			this.getView().setModel(this.oLoanModel, "loan");
			await this.loadLoans(this.oLoanModel, 1);

			this.sSortBy = this.getView().byId("sortBySelect").getSelectedKey();
			this.sSortOrder = this.getView().byId("sortOrderSelect").getSelectedKey();
			this._applySorting()
		},

		loadData: async function() {
			await this.loadLoans(this.oLoanModel, this.oLoanModel.getData().page);
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

		handleLoansUpdated: async function (ns, ev, eventData) {
			await this.loadData();

			if (!eventData.from_books && (eventData.delete_loan || eventData.return_book)) {
				Core.getEventBus().publish("library-app", "booksUpdated", {from_loans: true});
			}
		},

		onTitleSearchChange: function(oEvent) {
			this.sTitleSearch = oEvent.getParameter("value");
			this._applyCombinedFilters();
		},

		onUsernameSearchChange: function(oEvent) {
			this.sUsernameSearch = oEvent.getParameter("value");
			this._applyCombinedFilters();
		},

		onFilterStatusChange: function(oEvent) {
			this._sSelectedStatus = oEvent.getParameter("selectedItem").getKey();
			this._applyCombinedFilters();
		},

		_applyCombinedFilters: function() {
			let aFilters = [];

			if (this.sTitleSearch && this.sTitleSearch.trim() !== "") {
				aFilters.push(
					new Filter("book_title", FilterOperator.Contains, this.sTitleSearch)
				);
			}

			if (this.sUsernameSearch && this.sUsernameSearch.trim() !== "") {
				aFilters.push(
					new Filter("user_name", FilterOperator.Contains, this.sUsernameSearch)
				);
			}

			if (this._sSelectedStatus && this._sSelectedStatus !== "all") {
				aFilters.push(
					new Filter("status", FilterOperator.EQ, this._sSelectedStatus)
				);
			}

			let oTable = this.getView().byId("loansTable");
			let oBinding = oTable.getBinding("items");

			oBinding.filter(aFilters);
		},

		onSortByChange: function (oEvent) {
			this.sSortBy = oEvent.getParameter("selectedItem").getKey();
			this._applySorting();
		},

		onSortOrderChange: function (oEvent) {
			this.sSortOrder = oEvent.getParameter("selectedItem").getKey();
			this._applySorting();
		},

		_applySorting: function () {
			let oTable = this.getView().byId("loansTable");
			let oBinding = oTable.getBinding("items");

			if (this.sSortBy && this.sSortOrder) {
				let bDescending = this.sSortOrder === "desc";
				let oSorter = new Sorter(this.sSortBy, bDescending);

				oBinding.sort(oSorter);
			}
		},

		onPreviousPage: async function () {
			await this.loadLoans(this.oLoanModel, this.oLoanModel.getData().page - 1);
		},

		onNextPage: async function () {
			await this.loadLoans(this.oLoanModel, this.oLoanModel.getData().page + 1);
		},
	});
});