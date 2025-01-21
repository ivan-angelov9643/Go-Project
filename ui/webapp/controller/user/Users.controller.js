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

		onInit: async function () {
			const oRouter = this.getOwnerComponent().getRouter();
			oRouter.attachRoutePatternMatched(this.loadData, this);

			Core.getEventBus().subscribe("library-app", "usersUpdated", this.handleUsersUpdated, this);

			this.oUserModel = new JSONModel({
				count: null,
				page_size: null,
				page: null,
				data: null,
				total_pages: null,
			});
			this.oUserModel.setSizeLimit(Number.MAX_VALUE);
			this.getView().setModel(this.oUserModel, "user");
			await this.loadUsers(this.oUserModel,1);
		},

		loadData: async function() {
			await this.loadUsers(this.oUserModel, this.oUserModel.getData().page);
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
			Core.getEventBus().subscribe("library-app", "usersUpdated", this.handleUsersUpdated, this);
		},

		handleUsersUpdated: async function (ns, ev, eventData) {
			await this.loadData();
		},

		onPreviousPage: async function () {
			await this.loadUsers(this.oUserModel, this.oUserModel.getData().page - 1);
		},

		onNextPage: async function () {
			await this.loadUsers(this.oUserModel, this.oUserModel.getData().page + 1);
		},
	});
});