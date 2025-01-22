sap.ui.define([
    "../BaseController",
    "sap/m/MessageToast",
    "sap/ui/model/json/JSONModel",
    "sap/ui/core/Core",
    "sap/ui/core/mvc/XMLView",
], function (BaseController, MessageToast, JSONModel, Core, XMLView) {
    "use strict";

    return BaseController.extend("library-app.controller.book.EditBookDialog", {
        onInit: async function () {
            Core.getEventBus().subscribe("library-app", "authorsUpdated", this.handleAuthorsUpdated, this);
            Core.getEventBus().subscribe("library-app", "categoriesUpdated", this.handleCategoriesUpdated, this);

            this.oDialogBookModel = new JSONModel(this.initBookModel());
            this.getView().setModel(this.oDialogBookModel, "dialogBook");

            this.oCategoryModel = new JSONModel({
                count: null,
                page_size: null,
                page: null,
                data: null,
                total_pages: null,
            });
            this.oCategoryModel.setSizeLimit(Number.MAX_VALUE);
            this.getView().setModel(this.oCategoryModel, "category");
            // TODO
            await this.loadCategories(this.oCategoryModel);
        },

        onExit: function () {
            Core.getEventBus().unsubscribe("library-app", "authorsUpdated", this.handleAuthorsUpdated, this);
            Core.getEventBus().unsubscribe("library-app", "categoriesUpdated", this.handleCategoriesUpdated, this);
        },

        handleAuthorsUpdated: async function (ns, ev, eventData) {
            // TODO
            await this.loadAuthors(this.oAuthorModel)
        },

        handleCategoriesUpdated: async function (ns, ev, eventData) {
            // TODO
            await this.loadCategories(this.oCategoryModel)
        },

        onSaveBook: async function () {
            const bookData = this.oDialogBookModel.getData();

            bookData.year = parseInt(bookData.year, 10);
            bookData.total_copies = parseInt(bookData.total_copies, 10);

            try {
                const token = await this.getOwnerComponent().getToken();
                const saveResponse = await this.sendRequest(
                    `http://localhost:8080/api/books/${bookData.id}`,
                    "PUT",
                    token,
                    bookData
                );
//TODO
                saveResponse.edit_book = true;
                Core.getEventBus().publish("library-app", "booksUpdated", saveResponse);

                MessageToast.show("Successfully updated book");
            } catch (error) {
                MessageToast.show(error.error || "Error updating book");
                return;
            }

            if (this._oAuthorSelectDialog && !this._oAuthorSelectDialog.bIsDestroyed) {
                this._oAuthorSelectDialog.destroy();
            }

            this.onDialogClose();
        },

        onCancelEdit: function () {
            this.onDialogClose();
        },

        onDialogClose: function () {
            const dialog = this.byId("editBookDialog");
            if (dialog) {
                dialog.close();
            }
        },

        onOpenAuthorDialog: async function () {
            if (!this._oAuthorSelectDialog || this._oAuthorSelectDialog.bIsDestroyed) {
                const oOwnerComponent = this.getOwnerComponent();
                oOwnerComponent.runAsOwner(() => {
                    this._oAuthorSelectDialog= new XMLView({
                        id: "authorSelectDialogView",
                        viewName: "library-app.view.book.AuthorSelectDialog",
                    });
                    this.getView().addDependent(this._oAuthorSelectDialog);
                });
            }
            // const oData = this.oDialogBookModel.getData();
            // const oDialogBookModel = this._oAuthorSelectDialog.getModel("dialogBook");

            // this.fillBookModel(oDialogBookModel, oData);

            console.log(this._oAuthorSelectDialog)
            this._oAuthorSelectDialog.byId("authorSelectDialog").open();

        },
    });
});
