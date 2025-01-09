sap.ui.define([
	'../BaseController',
	'sap/ui/model/json/JSONModel',
	'library-app/model/Formatter',
	"sap/ui/core/Core",
	"sap/ui/core/mvc/XMLView",
], function (BaseController, JSONModel, formatter, Core, XMLView) {
	"use strict";
	return BaseController.extend("library-app.controller.author.Authors", {
		formatter: formatter,

		onInit: function () {
			Core.getEventBus().subscribe("library-app", "authorsUpdated", this.handleAuthorsUpdated, this);
			this.oAuthorModel = new JSONModel({
				authors: null,
			});
			this.oAuthorModel.setSizeLimit(Number.MAX_VALUE);
			this.getView().setModel(this.oAuthorModel, "author");
			this.loadAuthors(this.oAuthorModel);
		},

		onCreateAuthor: async function () {
			if (!this._oCreateAuthorDialog) {
				const oOwnerComponent = this.getOwnerComponent();
				oOwnerComponent.runAsOwner(() => {
					this._oCreateAuthorDialog = new XMLView({
						id: "createAuthorDialogView",
						viewName: "library-app.view.author.CreateAuthorDialog",
					});
					this.getView().addDependent(this._oCreateAuthorDialog);
				});
			}

			this._oCreateAuthorDialog.byId("createAuthorDialog").open();
		},

		onEditAuthor: async function (oEvent) {
			if (!this._oEditAuthorDialog) {
				const oOwnerComponent = this.getOwnerComponent();
				oOwnerComponent.runAsOwner(() => {
					this._oEditAuthorDialog = new XMLView({
						id: "editAuthorDialogView",
						viewName: "library-app.view.author.EditAuthorDialog",
					});
					this.getView().addDependent(this._oEditAuthorDialog);
				});
			}
			const oContext = oEvent.getSource().getBindingContext("author");
			const oData = oContext.getObject();
			const oDialogAuthorModel = this._oEditAuthorDialog.getModel("dialogAuthor");

			this.fillAuthorModel(oDialogAuthorModel, oData);
			this._oEditAuthorDialog.byId("editAuthorDialog").open();
		},

		onDeleteAuthor: async function (oEvent) {
			if (!this._oDeleteAuthorDialog) {
				const oOwnerComponent = this.getOwnerComponent();
				oOwnerComponent.runAsOwner(() => {
					this._oDeleteAuthorDialog = new XMLView({
						id: "deleteAuthorDialogView",
						viewName: "library-app.view.author.DeleteAuthorDialog",
					});
					this.getView().addDependent(this._oDeleteAuthorDialog);
				});
			}
			const oContext = oEvent.getSource().getBindingContext("author");
			const oData = oContext.getObject();
			const oDialogAuthorModel = this._oDeleteAuthorDialog.getModel("dialogAuthor");

			this.fillAuthorModel(oDialogAuthorModel, oData)
			this._oDeleteAuthorDialog.byId("deleteAuthorDialog").open();
		},

		onExit: function () {
			Core.getEventBus().unsubscribe("library-app", "authorsUpdated", this.handleAuthorsUpdated, this);
		},

		handleAuthorsUpdated: async function (ns, ev, eventData) {
			this.loadAuthors(this.oAuthorModel);
			Core.getEventBus().publish("library-app", "booksUpdated", eventData);
		},
	});
});