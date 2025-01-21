sap.ui.define([
	'../BaseController',
	'sap/ui/model/json/JSONModel',
	'sap/ui/Device',
	'library-app/model/formatter',
	"sap/ui/core/Core",
	"sap/ui/core/mvc/XMLView",
	"sap/ui/model/Filter",
	"sap/ui/model/FilterOperator",
	"sap/ui/model/Sorter"
], function (BaseController, JSONModel, Device, formatter, Core, XMLView, Filter, FilterOperator, Sorter) {
	"use strict";
	return BaseController.extend("library-app.controller.Reservations", {formatter: formatter,

		onInit: async function () {
			const oRouter = this.getOwnerComponent().getRouter();
			oRouter.attachRoutePatternMatched(this.loadData, this);

			Core.getEventBus().subscribe("library-app", "reservationsUpdated", this.handleReservationsUpdated, this);

			this.oReservationModel = new JSONModel({
				count: null,
				page_size: null,
				page: null,
				data: null,
				total_pages: null,
			});
			this.oReservationModel.setSizeLimit(Number.MAX_VALUE);
			this.getView().setModel(this.oReservationModel, "reservation");
			await this.loadReservations(this.oReservationModel, 1);

			this.sSortBy = this.getView().byId("sortBySelect").getSelectedKey();
			this.sSortOrder = this.getView().byId("sortOrderSelect").getSelectedKey();
			this._applySorting()
		},

		loadData: async function() {
			await this.loadReservations(this.oReservationModel, this.oReservationModel.getData().page);
		},

		onExtendReservation: async function (oEvent) {
			if (!this._oExtendReservationDialog) {
				const oOwnerComponent = this.getOwnerComponent();
				oOwnerComponent.runAsOwner(() => {
					this._oExtendReservationDialog = new XMLView({
						id: "extendReservationDialogView",
						viewName: "library-app.view.reservation.ExtendReservationDialog",
					});
					this.getView().addDependent(this._oExtendReservationDialog);
				});
			}
			const oContext = oEvent.getSource().getBindingContext("reservation");
			const oData = oContext.getObject();
			const oDialogReservationModel = this._oExtendReservationDialog.getModel("dialogReservation");

			this.fillReservationModel(oDialogReservationModel, oData);
			this._oExtendReservationDialog.byId("extendReservationDialog").open();
		},

		onMakeLoan: async function (oEvent) {
			if (!this._oMakeLoanDialog) {
				const oOwnerComponent = this.getOwnerComponent();
				oOwnerComponent.runAsOwner(() => {
					this._oMakeLoanDialog = new sap.ui.core.mvc.XMLView({
						id: "makeLoanDialogView",
						viewName: "library-app.view.reservation.MakeLoanDialog",
					});
					this.getView().addDependent(this._oMakeLoanDialog);
				});
			}

			const oContext = oEvent.getSource().getBindingContext("reservation");
			const oData = oContext.getObject();

			const oDialogLoanModel = this._oMakeLoanDialog.getModel("dialogLoan");
			const oDialogReservationModel = this._oMakeLoanDialog.getModel("dialogReservation");

			this.fillReservationModel(oDialogReservationModel, oData);
			this.fillLoanModel(oDialogLoanModel, {
				user_id: oData.user_id,
				book_id: oData.book_id,
			});

			this._oMakeLoanDialog.byId("makeLoanDialog").open();
		},

		onDeleteReservation: async function (oEvent) {
			if (!this._oDeleteReservationDialog) {
				const oOwnerComponent = this.getOwnerComponent();
				oOwnerComponent.runAsOwner(() => {
					this._oDeleteReservationDialog = new XMLView({
						id: "deleteReservationDialogView",
						viewName: "library-app.view.reservation.DeleteReservationDialog",
					});
					this.getView().addDependent(this._oDeleteReservationDialog);
				});
			}
			const oContext = oEvent.getSource().getBindingContext("reservation");
			const oData = oContext.getObject();
			const oDialogReservationModel = this._oDeleteReservationDialog.getModel("dialogReservation");

			this.fillReservationModel(oDialogReservationModel, oData);
			this._oDeleteReservationDialog.byId("deleteReservationDialog").open();
		},

		onExit: function () {
			Core.getEventBus().unsubscribe("library-app", "reservationsUpdated", this.handleReservationsUpdated, this);
		},

		handleReservationsUpdated: async function (ns, ev, eventData) {
			await this.loadData()

			if (eventData.make_loan) {
				Core.getEventBus().publish("library-app", "loansUpdated");
			}

			if (!eventData.from_books &&
				(
					eventData.make_loan ||
					eventData.delete_reservation ||
					eventData.make_reservation
				)
			) {
				Core.getEventBus().publish("library-app", "booksUpdated", {from_reservations: true});
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

			let oTable = this.getView().byId("reservationsTable");
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
			let oTable = this.getView().byId("reservationsTable");
			let oBinding = oTable.getBinding("items");

			if (this.sSortBy && this.sSortOrder) {
				let bDescending = this.sSortOrder === "desc";
				let oSorter = new Sorter(this.sSortBy, bDescending);

				oBinding.sort(oSorter);
			}
		},

		onPreviousPage: async function () {
			await this.loadReservations(this.oReservationModel, this.oReservationModel.getData().page - 1);
		},

		onNextPage: async function () {
			await this.loadReservations(this.oReservationModel, this.oReservationModel.getData().page + 1);
		},
	});
});