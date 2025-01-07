sap.ui.define([
	'./BaseController',
	'sap/ui/model/json/JSONModel',
	'sap/ui/Device',
	'library-app/model/formatter'
], function (BaseController, JSONModel, Device, formatter) {
	"use strict";
	return BaseController.extend("library-app.controller.Reservations", {formatter: formatter,

		onInit: function () {
			// sap.ui.getCore().getEventBus().subscribe("library-app", "RouteChanged", this.handleRouteChanged, this);

			this.reservationModel = new JSONModel({
				reservations: null,
			});
			this.reservationModel.setSizeLimit(Number.MAX_VALUE);
			this.getView().setModel(this.reservationModel, "reservation");
			this.loadData();
			this.getView().setModel(this.reservationModel, "reservation");
		},

		onExit: function () {
			// sap.ui.getCore().getEventBus().unsubscribe("library-app", "RouteChanged", this.handleRouteChanged, this);
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
				reservation.book_tittle = book ? book.title : 'Unknown Book';
			});
			reservationModel.setProperty("/reservations", reservationsData);
		},
	});
});