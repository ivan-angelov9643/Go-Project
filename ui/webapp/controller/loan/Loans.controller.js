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

			this._setDefaultSearchFields()

			await this.loadLoans(this.oLoanModel, 1, this.sSortBy, this.sSortOrder, this.sStatusFilter);
		},

		loadData: async function() {
			await this.loadLoans(this.oLoanModel, this.oLoanModel.getData().page,
				this.sSortBy, this.sSortOrder, this.sStatusFilter,
				this.sUsernameSearch, this.sTitleSearch);
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

		onPreviousPage: async function () {
			await this.loadLoans(this.oLoanModel, this.oLoanModel.getData().page - 1,
				this.sSortBy, this.sSortOrder, this.sStatusFilter,
				this.sUsernameSearch, this.sTitleSearch);
		},

		onNextPage: async function () {
			await this.loadLoans(this.oLoanModel, this.oLoanModel.getData().page + 1,
				this.sSortBy, this.sSortOrder, this.sStatusFilter,
				this.sUsernameSearch, this.sTitleSearch);
		},

		onSearch: async function () {
			if (!this._searchFieldsChanged()) {
				return;
			}

			this.sSortBy = this.byId("sortBySelect").getSelectedKey();
			this.sSortOrder = this.byId("sortOrderSelect").getSelectedKey();
			this.sStatusFilter = this.byId("statusFilter").getSelectedKey();
			this.sUsernameSearch = this.byId("usernameSearch").getValue();
			this.sTitleSearch = this.byId("titleSearch").getValue();

			await this.loadLoans(this.oLoanModel, 1,
				this.sSortBy, this.sSortOrder, this.sStatusFilter,
				this.sUsernameSearch, this.sTitleSearch);
		},

		onClearSearch: async function () {
			if (this._searchFieldsEmpty()) {
				return;
			}

			this._setDefaultSearchFields()

			this.byId("sortBySelect").setSelectedKey("start_date");
			this.byId("sortOrderSelect").setSelectedKey("asc");
			this.byId("statusFilter").setSelectedKey("all");
			this.byId("usernameSearch").setValue("");
			this.byId("titleSearch").setValue("");

			await this.loadLoans(this.oLoanModel, 1,
				this.sSortBy, this.sSortOrder, this.sStatusFilter);
		},

		_setDefaultSearchFields() {
			this.sSortBy = "start_date";
			this.sSortOrder = "asc";
			this.sStatusFilter = "all"
			this.sUsernameSearch = "";
			this.sTitleSearch = "";
		},

		_searchFieldsChanged() {
			return this.sSortBy !== this.byId("sortBySelect").getSelectedKey() ||
				this.sSortOrder !== this.byId("sortOrderSelect").getSelectedKey() ||
				this.sStatusFilter !== this.byId("statusFilter").getSelectedKey() ||
				this.sUsernameSearch !== this.byId("usernameSearch").getValue() ||
				this.sTitleSearch !== this.byId("titleSearch").getValue();
		},

		_searchFieldsEmpty() {
			return this.sSortBy === "start_date" &&
				this.sSortOrder === "asc" &&
				this.sStatusFilter === "all" &&
				this.sUsernameSearch === "" &&
				this.sTitleSearch ==="";
		},
	});
});