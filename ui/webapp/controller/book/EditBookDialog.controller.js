sap.ui.define([
    "../BaseController",
    "sap/m/MessageToast",
    "sap/ui/model/json/JSONModel",
    "sap/ui/core/Core",
    "sap/ui/core/mvc/XMLView",
], function (BaseController, MessageToast, JSONModel, Core) {
    "use strict";

    return BaseController.extend("library-app.controller.book.EditBookDialog", {
        onInit: async function () {
            this.oDialogBookModel = new JSONModel(this.initBookModel());
            this.getView().setModel(this.oDialogBookModel, "dialogBook");

            this.oTotalCopiesModel = new JSONModel({value: 0});
            this.getView().setModel(this.oTotalCopiesModel, "totalCopies");
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
                saveResponse.edit_book = true;

                const notAvailableCopies = this.oTotalCopiesModel.getData().value - bookData.available_copies;
                saveResponse.available_copies = saveResponse.total_copies - notAvailableCopies;

                Core.getEventBus().publish("library-app", "booksUpdated", saveResponse);

                MessageToast.show("Successfully updated book");
            } catch (error) {
                MessageToast.show(error.error || "Error updating book");
                return;
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

    });
});
