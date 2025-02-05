sap.ui.define([
    "../BaseController",
    "sap/m/MessageToast",
    "sap/ui/model/json/JSONModel",
    "sap/ui/core/Core"
], function (BaseController, MessageToast, JSONModel, Core) {
    "use strict";

    return BaseController.extend("library-app.controller.book.DeleteBookDialog", {
        onInit: function () {
            this.oDialogBookModel = new JSONModel(this.initBookModel());
            this.getView().setModel(this.oDialogBookModel, "dialogBook");
        },

        onConfirmDelete: async function () {
            const bookData = this.getView().getModel("dialogBook").getData();

            try {
                const token = await this.getOwnerComponent().getToken();
                const deleteResponse = await this.sendRequest(
                    `http://localhost:8080/api/books/${bookData.id}`,
                    "DELETE",
                    token
                );

                Core.getEventBus().publish("library-app", "booksUpdated", {delete: true});

                MessageToast.show("Successfully deleted book");
            } catch (error) {
                MessageToast.show(error.error || "Error deleting book");
            }

            this.onDialogClose();
        },

        onCancelDelete: function () {
            this.onDialogClose();
        },

        onDialogClose: function () {
            const dialog = this.byId("deleteBookDialog");
            if (dialog) {
                dialog.close();
            }
        }
    });
});
