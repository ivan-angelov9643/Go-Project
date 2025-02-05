sap.ui.define([
    "../BaseController",
    "sap/m/MessageToast",
    "sap/ui/model/json/JSONModel",
    "sap/ui/core/Core"
], function (BaseController, MessageToast, JSONModel, Core) {
    "use strict";

    return BaseController.extend("library-app.controller.book.CreateBookDialog", {
        onInit: async function () {
            this.oDialogBookModel = new JSONModel(this.initBookModel());
            this.getView().setModel(this.oDialogBookModel, "dialogBook");
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
        },

        onCancelCreate: function () {
            this.onDialogClose();
        },

        onDialogClose: function () {
            this.fillBookModel(this.oDialogBookModel, this.initBookModel());
            const dialog = this.byId("createBookDialog");
            if (dialog) {
                dialog.close();
            }
        }
    });
});
