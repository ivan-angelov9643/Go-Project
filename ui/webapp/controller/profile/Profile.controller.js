sap.ui.define([
	"../BaseController",
	"sap/ui/model/json/JSONModel",
	"library-app/model/formatter",
	"sap/ui/core/mvc/XMLView",
	"sap/ui/core/Core",
], function (BaseController, JSONModel, formatter, XMLView, Core) {
	"use strict";

	return BaseController.extend("library-app.controller.profile.Profile", {
		formatter: formatter,

		onInit: async function () {
			Core.getEventBus().subscribe("library-app", "userUpdated", this.handleUserUpdated, this);

			this.oUserModel = new JSONModel(this.initUserModel());
			this.getView().setModel(this.oUserModel, "user");

			await this.loadCurrentUser(this.oUserModel);
		},

		onEditProfile: async function () {
			if (!this._oEditProfileDialog) {
				const oOwnerComponent = this.getOwnerComponent();
				oOwnerComponent.runAsOwner(() => {
					this._oEditProfileDialog = new XMLView({
						id: "editProfileDialogView",
						viewName: "library-app.view.profile.EditProfileDialog",
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
	});
});
