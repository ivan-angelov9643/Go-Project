sap.ui.define([
	"./BaseController",
	"sap/ui/model/json/JSONModel",
	"library-app/model/formatter",
	"sap/ui/core/mvc/XMLView",
	"sap/ui/core/Core",
	"sap/m/MessageToast",
], function (BaseController, JSONModel, formatter, XMLView, Core, MessageToast) {
	"use strict";

	return BaseController.extend("library-app.controller.Profile", {
		formatter: formatter,

		onInit: async function () {
			Core.getEventBus().subscribe("library-app", "userUpdated", this.handleUserUpdated, this);

			this.oUserModel = new JSONModel(this.initUserModel());
			this.getView().setModel(this.oUserModel, "user");

			await this.setupUserModel();
		},

		onEditProfile: async function () {
			if (!this._oEditProfileDialog) {
				const oOwnerComponent = this.getOwnerComponent();
				oOwnerComponent.runAsOwner(() => {
					this._oEditProfileDialog = new XMLView({
						id: "editProfileDialogView",
						viewName: "library-app.view.EditProfileDialog",
					});
					this.getView().addDependent(this._oEditProfileDialog);
				});
			}
			const oUserModel = this.getView().getModel("user");
			const oUserData = oUserModel.getData();
			const oDialogUserModel = this._oEditProfileDialog.getModel("dialogUser");

			this.fillUserModel(oDialogUserModel, oUserData);
			this._oEditProfileDialog.byId("editProfileDialog").open();
		},

		onExit: function () {
			Core.getEventBus().unsubscribe("library-app", "userUpdated", this.handleUserUpdated, this);
			if (this._oEditProfileDialog) {
				this._oEditProfileDialog.destroy();
			}
		},

		handleUserUpdated: async function (ns, ev, eventData) {
			this.fillUserModel(this.oUserModel, eventData);
		},

		setupUserModel: async function () {
			const token = await this.getOwnerComponent().getToken();
			const userID = this.getUserID(token);
			const userData = await this.sendRequest(`http://localhost:8080/api/users/${userID}`, "GET", token);

			this.fillUserModel(this.oUserModel, userData);
		},

		fillUserModel: function (userModel, data) {
			userModel.setProperty("/preferred_username", data.preferred_username);
			userModel.setProperty("/given_name", data.given_name);
			userModel.setProperty("/family_name", data.family_name);
			userModel.setProperty("/email", data.email);
		}
	});
});
