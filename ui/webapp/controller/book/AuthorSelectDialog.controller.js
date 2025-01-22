sap.ui.define([
    '../BaseController',
    'sap/ui/model/json/JSONModel',
], function(BaseController, JSONModel) {
    "use strict";

    return BaseController.extend("library-app.controller.book.Books", {

        onInit: async function() {
            // events

            this.oAuthorModel = new JSONModel({
                count: null,
                page_size: null,
                page: null,
                data: null,
                total_pages: null,
            });
            this.oAuthorModel.setSizeLimit(Number.MAX_VALUE);
            this.getView().setModel(this.oAuthorModel, "author");
            await this.loadAuthors(this.oAuthorModel, 1);

            this.oDisplayAuthorModel = new JSONModel({
                count: null,
                page_size: null,
                page: null,
                data: null,
                total_pages: null,
            });
            this.oDisplayAuthorModel.setSizeLimit(Number.MAX_VALUE);
            this.getView().setModel(this.oDisplayAuthorModel, "displayAuthor");
            await this.loadAuthors(this.oDisplayAuthorModel, 1); // can be filled with oauthor data
        },

        onAuthorSearch: async function (oEvent) {
            this.sSearch = oEvent.getParameter("value");

            await this.loadAuthors(this.oAuthorModel, 1, this.sSearch)
            await this.loadAuthors(this.oDisplayAuthorModel, 1, this.sSearch)
        },

        onLoadMoreAuthors: async function (oEvent) {
            const actual = oEvent.getParameters().actual;
            const total = oEvent.getParameters().total;

            if (this.oAuthorModel.getData().page < this.oAuthorModel.getData().total_pages &&
                actual === total) {
                await this.loadAuthors(this.oAuthorModel, this.oAuthorModel.getData().page + 1, this.sSearch);
                this._appendData(this.oDisplayAuthorModel, this.oAuthorModel);
            }
        },

        onAuthorConfirm: function (oEvent) {
            const oSelectedItem = oEvent.getParameter("selectedItem");
            if (oSelectedItem) {
                const oAuthorContext = oSelectedItem.getBindingContext("displayAuthor");
                const sAuthorName = oAuthorContext.getProperty("first_name") + " " + oAuthorContext.getProperty("last_name");
                const sAuthorID = oAuthorContext.getProperty("id");

                const oDialogBookModel = this.getView().getModel("dialogBook");
                oDialogBookModel.setProperty("/author_id", sAuthorID);
                oDialogBookModel.setProperty("/author_name", sAuthorName);

                this.onDialogCancel();
            }
        },

        onDialogCancel: function () {
            this.getView().destroy()
        },

        _appendData: function (displayModel, model) {
            const displayData = displayModel.getData().data;
            const newData = model.getData().data;

            displayData.push(...newData);

            displayModel.setProperty("/data", displayData);
        }
    });
});
