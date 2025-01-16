sap.ui.define([
    "../BaseController",
    "sap/m/MessageToast",
    "sap/ui/model/json/JSONModel",
    "sap/ui/core/Core"
], function (BaseController, MessageToast, JSONModel, Core) {
    "use strict";

    return BaseController.extend("library-app.controller.book.CreateBookDialog", {
        onInit: async function () {
            Core.getEventBus().subscribe("library-app", "authorsUpdated", this.handleAuthorsUpdated, this);
            Core.getEventBus().subscribe("library-app", "categoriesUpdated", this.handleCategoriesUpdated, this);

            this.oDialogBookModel = new JSONModel(this.initBookModel());
            this.getView().setModel(this.oDialogBookModel, "dialogBook");

            this.oAuthorModel = new JSONModel({
                authors: null,
            });
            this.oAuthorModel.setSizeLimit(Number.MAX_VALUE);
            this.getView().setModel(this.oAuthorModel, "author");
            await this.loadAuthors(this.oAuthorModel);

            this.oCategoryModel = new JSONModel({
                categories: null,
            });
            this.oCategoryModel.setSizeLimit(Number.MAX_VALUE);
            this.getView().setModel(this.oCategoryModel, "category");
            await this.loadCategories(this.oCategoryModel);
        },

        onExit: function () {
            Core.getEventBus().unsubscribe("library-app", "authorsUpdated", this.handleAuthorsUpdated, this);
            Core.getEventBus().unsubscribe("library-app", "categoriesUpdated", this.handleCategoriesUpdated, this);
        },

        handleAuthorsUpdated: async function (ns, ev, eventData) {
            await this.loadAuthors(this.oAuthorModel)
        },

        handleCategoriesUpdated: async function (ns, ev, eventData) {
            await this.loadCategories(this.oCategoryModel)
        },

        onCreateBook: async function () {
            const bookData = this.oDialogBookModel.getData();

            bookData.year = parseInt(bookData.year, 10);
            bookData.total_copies = parseInt(bookData.total_copies, 10);

            try {
                const token = await this.getOwnerComponent().getToken();
                const createResponse = await this.sendRequest(
                    `http://localhost:8080/api/books`,
                    "POST",
                    token,
                    bookData
                );

                Core.getEventBus().publish("library-app", "booksUpdated", createResponse);

                MessageToast.show("Successfully created book");
            } catch (error) {
                MessageToast.show(error.error || "Error creating book");
                return;
            }

            this.onDialogClose();
            this.fillBookModel(this.oDialogBookModel, this.initBookModel());
        },

        onCancelCreate: function () {
            this.onDialogClose();
        },

        onDialogClose: function () {
            const dialog = this.byId("createBookDialog");
            if (dialog) {
                dialog.close();
            }
        }
    });
});
