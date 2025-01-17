sap.ui.define([
	'./BaseController',
	'sap/ui/model/json/JSONModel',
	'sap/ui/Device',
	'library-app/model/formatter',
	"sap/ui/core/Core",
], function (BaseController, JSONModel, Device, formatter, Core) {
	"use strict";
	return BaseController.extend("library-app.controller.Home", {
		formatter: formatter,

		onInit: async function () {
			Core.getEventBus().subscribe("library-app", "userUpdated", this.handleUserUpdated, this);

			this.oUserModel = new JSONModel(this.initUserModel());
			this.getView().setModel(this.oUserModel, "user");

			await this.loadCurrentUser(this.oUserModel);
		},

		onExit: function() {
			Core.getEventBus().unsubscribe("library-app", "userUpdated", this.handleUserUpdated, this);
		},

		handleUserUpdated: async function (ns, ev, eventData) {
			this.fillUserModel(this.oUserModel, eventData);
		},
	});
});