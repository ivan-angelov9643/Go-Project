sap.ui.define([
    '../BaseController',
    'sap/ui/model/json/JSONModel',
], function(BaseController, JSONModel) {
    "use strict";

    return BaseController.extend("library-app.controller.book.Categories", {

        onInit: async function() {
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

            this.oDisplayCategoryModel = new JSONModel({
                data: null,
            });
            this.oDisplayCategoryModel.setSizeLimit(Number.MAX_VALUE);
            this.getView().setModel(this.oDisplayCategoryModel, "displayCategory");
            await this.loadCategories(this.oDisplayCategoryModel, 1); // can be filled with oCategory data
        },

        onCategorySearch: async function(oEvent) {
            this.sSearch = oEvent.getParameter("value");

            await this.loadCategories(this.oCategoryModel, 1, this.sSearch);
            await this.loadCategories(this.oDisplayCategoryModel, 1, this.sSearch);
        },

        onLoadMoreCategories: async function(oEvent) {
            const actual = oEvent.getParameters().actual;
            const total = oEvent.getParameters().total;

            if (this.oCategoryModel.getData().page < this.oCategoryModel.getData().total_pages &&
                actual === total) {
                await this.loadCategories(this.oCategoryModel, this.oCategoryModel.getData().page + 1, this.sSearch);
                this.AppendData(this.oDisplayCategoryModel, this.oCategoryModel);
            }
        },

        onCategoryConfirm: function(oEvent) {
            const oSelectedItem = oEvent.getParameter("selectedItem");
            if (oSelectedItem) {
                const oCategoryContext = oSelectedItem.getBindingContext("displayCategory");
                const sCategoryName = oCategoryContext.getProperty("name");
                const sCategoryID = oCategoryContext.getProperty("id");

                const oDialogBookModel = this.getView().getModel("dialogBook");
                oDialogBookModel.setProperty("/category_id", sCategoryID);
                oDialogBookModel.setProperty("/category_name", sCategoryName);

                this.onDialogCancel();
            }
        },

        onDialogCancel: function() {
            this.getView().destroy();
        },
    });
});
