sap.ui.define([
	'../BaseController',
	'sap/ui/model/json/JSONModel',
	'library-app/model/formatter',
	"sap/ui/core/Core",
	"sap/ui/core/mvc/XMLView",
], function (BaseController, JSONModel, formatter, Core, XMLView) {
	"use strict";
	return BaseController.extend("library-app.controller.user.Users", {
		formatter: formatter,

		onInit: function () {
			// sap.ui.getCore().getEventBus().subscribe("library-app", "RouteChanged", this.handleRouteChanged, this);
			Core.getEventBus().subscribe("library-app", "usersUpdated", this.handleUsersUpdated, this);

			this.oUserModel = new JSONModel({
				users: null,
			});
			this.oUserModel.setSizeLimit(Number.MAX_VALUE);
			this.getView().setModel(this.oUserModel, "user");
			this.loadUsers();
		},

		onEditUser: async function (oEvent) {
			if (!this._oEditUserDialog) {
				const oOwnerComponent = this.getOwnerComponent();
				oOwnerComponent.runAsOwner(() => {
					this._oEditUserDialog = new XMLView({
						id: "editUserDialogView",
						viewName: "library-app.view.user.EditUserDialog",
					});
					this.getView().addDependent(this._oEditUserDialog);
				});
			}
			const oContext = oEvent.getSource().getBindingContext("user");
			const oData = oContext.getObject();
			const oDialogUserModel = this._oEditUserDialog.getModel("dialogUser");

			this.fillUserModel(oDialogUserModel, oData);
			this._oEditUserDialog.byId("editUserDialog").open();
		},

		onDeleteUser: async function (oEvent) {
			if (!this._oDeleteUserDialog) {
				const oOwnerComponent = this.getOwnerComponent();
				oOwnerComponent.runAsOwner(() => {
					this._oDeleteUserDialog = new XMLView({
						id: "deleteUserDialogView",
						viewName: "library-app.view.user.DeleteUserDialog",
					});
					this.getView().addDependent(this._oDeleteUserDialog);
				});
			}
			const oContext = oEvent.getSource().getBindingContext("user");
			const oData = oContext.getObject();
			const oDialogUserModel = this._oDeleteUserDialog.getModel("dialogUser");

			this.fillUserModel(oDialogUserModel, oData);
			this._oDeleteUserDialog.byId("deleteUserDialog").open();
		},

		onExit: function () {
			// sap.ui.getCore().getEventBus().unsubscribe("library-app", "RouteChanged", this.handleRouteChanged, this);
			Core.getEventBus().subscribe("library-app", "usersUpdated", this.handleUsersUpdated, this);
		},

		loadUsers: async function () {
			const token = await this.getOwnerComponent().getToken();
			const userData = await this.sendRequest('http://localhost:8080/api/users', "GET", token);

			this.oUserModel.setProperty("/users", userData);
		},

		handleUsersUpdated: async function (ns, ev, eventData) {
			this.loadUsers();
		},
	});
});