sap.ui.define([
	'../BaseController',
	'sap/ui/model/json/JSONModel',
	'sap/ui/Device',
	'library-app/model/formatter',
	"sap/ui/core/Core",
	"sap/ui/core/mvc/XMLView",
], function (BaseController, JSONModel, Device, formatter, Core, XMLView) {
	"use strict";
	return BaseController.extend("library-app.controller.Reservations", {formatter: formatter,

		onInit: function () {
			Core.getEventBus().subscribe("library-app", "reservationsUpdated", this.handleReservationsUpdated, this);
			this.reservationModel = new JSONModel({
				reservations: null,
			});
			this.reservationModel.setSizeLimit(Number.MAX_VALUE);
			this.getView().setModel(this.reservationModel, "reservation");
			this.loadData();
			this.getView().setModel(this.reservationModel, "reservation");
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

		loadData: async function () {
			const reservationModel = this.getView().getModel("reservation");
			const token = await this.getOwnerComponent().getToken();

			const [reservationsData, usersData, booksData] = await Promise.all([
				this.sendRequest('http://localhost:8080/api/reservations', "GET", token),
				this.sendRequest('http://localhost:8080/api/users', "GET", token),
				this.sendRequest('http://localhost:8080/api/books', "GET", token)
			]);

			reservationsData.forEach(reservation => {
				const user = usersData.find(u => u.id === reservation.user_id);
				const book = booksData.find(b => b.id === reservation.book_id);

				reservation.user_name = user ? user.preferred_username : 'Unknown User';
				reservation.book_title = book ? book.title : 'Unknown Book';
			});
			reservationModel.setProperty("/reservations", reservationsData);
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
	});
});