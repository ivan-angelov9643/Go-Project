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
], function (BaseController, JSONModel, Device, formatter, Core, XMLView) {
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

			this._setDefaultSearchFields()

			await this.loadReservations(this.oReservationModel, 1, this.sSortBy, this.sSortOrder);
		},

		loadData: async function() {
			await this.loadReservations(this.oReservationModel, this.oReservationModel.getData().page,
				this.sSortBy, this.sSortOrder,
				this.sUsernameSearch, this.sTitleSearch);
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

		onPreviousPage: async function () {
			await this.loadReservations(this.oReservationModel, this.oReservationModel.getData().page - 1,
				this.sSortBy, this.sSortOrder,
				this.sUsernameSearch, this.sTitleSearch);
		},

		onNextPage: async function () {
			await this.loadReservations(this.oReservationModel, this.oReservationModel.getData().page + 1,
				this.sSortBy, this.sSortOrder,
				this.sUsernameSearch, this.sTitleSearch);
		},

		onSearch: async function () {
			if (!this._searchFieldsChanged()) {
				return;
			}

			this.sSortBy = this.byId("sortBySelect").getSelectedKey();
			this.sSortOrder = this.byId("sortOrderSelect").getSelectedKey();
			this.sUsernameSearch = this.byId("usernameSearch").getValue();
			this.sTitleSearch = this.byId("titleSearch").getValue();

			await this.loadReservations(this.oReservationModel, 1,
				this.sSortBy, this.sSortOrder,
				this.sUsernameSearch, this.sTitleSearch);
		},

		onClearSearch: async function () {
			if (this._searchFieldsEmpty()) {
				return;
			}

			this._setDefaultSearchFields()

			this.byId("sortBySelect").setSelectedKey("created_at");
			this.byId("sortOrderSelect").setSelectedKey("asc");
			this.byId("usernameSearch").setValue("");
			this.byId("titleSearch").setValue("");

			await this.loadReservations(this.oReservationModel, 1,
				this.sSortBy, this.sSortOrder);
		},

		_setDefaultSearchFields() {
			this.sSortBy = "created_at";
			this.sSortOrder = "asc";
			this.sUsernameSearch = "";
			this.sTitleSearch = "";
		},

		_searchFieldsChanged() {
			return this.sSortBy !== this.byId("sortBySelect").getSelectedKey() ||
				this.sSortOrder !== this.byId("sortOrderSelect").getSelectedKey() ||
				this.sUsernameSearch !== this.byId("usernameSearch").getValue() ||
				this.sTitleSearch !== this.byId("titleSearch").getValue();
		},

		_searchFieldsEmpty() {
			return this.sSortBy === "created_at" &&
			this.sSortOrder === "asc" &&
			this.sUsernameSearch === "" &&
			this.sTitleSearch ==="";
		},
	});
});