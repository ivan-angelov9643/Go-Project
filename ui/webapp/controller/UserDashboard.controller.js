sap.ui.define([
	'./BaseController',
	'sap/ui/model/json/JSONModel',
	'sap/ui/Device',
	'library-app/model/formatter'
], function (BaseController, JSONModel, Device, formatter) {
	"use strict";
	return BaseController.extend("library-app.controller.UserDashboard", {
		formatter: formatter,

		onInit: function () {
			// sap.ui.getCore().getEventBus().subscribe("library-app", "RouteChanged", this.handleRouteChanged, this);
		},

		onExit: function() {
			// sap.ui.getCore().getEventBus().unsubscribe("library-app", "RouteChanged", this.handleRouteChanged, this);
		},

		handleRouteChanged: function(channel, eventId, pageData) {
			if (pageData.selectedPageKey === "home"){
				this.loadData()
			}
		},

		loadData: async function () {

		}
	});
});