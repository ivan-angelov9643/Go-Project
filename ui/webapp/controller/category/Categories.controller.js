sap.ui.define([
	'../BaseController',
	'sap/ui/model/json/JSONModel',
	'library-app/model/Formatter',
	"sap/ui/core/Core",
	"sap/ui/core/mvc/XMLView",
], function (BaseController, JSONModel, formatter, Core, XMLView) {
	"use strict";
	return BaseController.extend("library-app.controller.category.Categories", {
		formatter: formatter,

		onInit: async function () {
			const oRouter = this.getOwnerComponent().getRouter();
			oRouter.attachRoutePatternMatched(this.loadData, this);

			Core.getEventBus().subscribe("library-app", "categoriesUpdated", this.handleCategoriesUpdated, this);

			this.oCategoryModel = new JSONModel({
				count: null,
				page_size: null,
				page: null,
				data: null,
				total_pages: null,
			});
			this.oCategoryModel.setSizeLimit(Number.MAX_VALUE);
			this.getView().setModel(this.oCategoryModel, "category");
			await this.loadCategories(this.oCategoryModel, 1);
		},

		loadData: async function() {
			await this.loadCategories(this.oCategoryModel, this.oCategoryModel.getData().page);
		},


		onCreateCategory: async function () {
			if (!this._oCreateCategoryDialog) {
				const oOwnerComponent = this.getOwnerComponent();
				oOwnerComponent.runAsOwner(() => {
					this._oCreateCategoryDialog = new XMLView({
						id: "createCategoryDialogView",
						viewName: "library-app.view.category.CreateCategoryDialog",
					});
					this.getView().addDependent(this._oCreateCategoryDialog);
				});
			}

			this._oCreateCategoryDialog.byId("createCategoryDialog").open();
		},

		onEditCategory: async function (oEvent) {
			if (!this._oEditCategoryDialog) {
				const oOwnerComponent = this.getOwnerComponent();
				oOwnerComponent.runAsOwner(() => {
					this._oEditCategoryDialog = new XMLView({
						id: "editCategoryDialogView",
						viewName: "library-app.view.category.EditCategoryDialog",
					});
					this.getView().addDependent(this._oEditCategoryDialog);
				});
			}
			const oContext = oEvent.getSource().getBindingContext("category");
			const oData = oContext.getObject();
			const oDialogCategoryModel = this._oEditCategoryDialog.getModel("dialogCategory");

			this.fillCategoryModel(oDialogCategoryModel, oData);
			this._oEditCategoryDialog.byId("editCategoryDialog").open();
		},

		onDeleteCategory: async function (oEvent) {
			if (!this._oDeleteCategoryDialog) {
				const oOwnerComponent = this.getOwnerComponent();
				oOwnerComponent.runAsOwner(() => {
					this._oDeleteCategoryDialog = new XMLView({
						id: "deleteCategoryDialogView",
						viewName: "library-app.view.category.DeleteCategoryDialog",
					});
					this.getView().addDependent(this._oDeleteCategoryDialog);
				});
			}
			const oContext = oEvent.getSource().getBindingContext("category");
			const oData = oContext.getObject();
			const oDialogCategoryModel = this._oDeleteCategoryDialog.getModel("dialogCategory");

			this.fillCategoryModel(oDialogCategoryModel, oData);
			this._oDeleteCategoryDialog.byId("deleteCategoryDialog").open();
		},

		onExit: function () {
			Core.getEventBus().unsubscribe("library-app", "categoriesUpdated", this.handleCategoriesUpdated, this);
		},

		handleCategoriesUpdated: async function (ns, ev, eventData) {
			await this.loadData();
			Core.getEventBus().publish("library-app", "booksUpdated", eventData);
		},

		onPreviousPage: async function () {
			await this.loadCategories(this.oCategoryModel, this.oCategoryModel.getData().page - 1);
		},

		onNextPage: async function () {
			await this.loadCategories(this.oCategoryModel, this.oCategoryModel.getData().page + 1);
		},
	});
});
