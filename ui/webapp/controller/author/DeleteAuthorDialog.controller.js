sap.ui.define([
    "../BaseController",
    "sap/m/MessageToast",
    "sap/ui/model/json/JSONModel",
    "sap/ui/core/Core"
], function (BaseController, MessageToast, JSONModel, Core) {
    "use strict";

    return BaseController.extend("library-app.controller.author.DeleteAuthorDialog", {
        onInit: function () {
            this.oDialogAuthorModel = new JSONModel(this.initAuthorModel());
            this.getView().setModel(this.oDialogAuthorModel, "dialogAuthor");
        },

        onConfirmDelete: async function () {
            const authorData = this.getView().getModel("dialogAuthor").getData();

            try {
                const token = await this.getOwnerComponent().getToken();
                const deleteResponse = await this.sendRequest(
                    `http://localhost:8080/api/authors/${authorData.id}`,
                    "DELETE",
                    token
                );

                Core.getEventBus().publish("library-app", "authorsUpdated", deleteResponse);

                MessageToast.show("Successfully deleted author");
            } catch (error) {
                MessageToast.show(error.error || "Error deleting author");
                return;
            }
            this.onDialogClose();
        },

        onCancelDelete: function () {
            this.onDialogClose();
        },

        onDialogClose: function () {
            const dialog = this.byId("deleteAuthorDialog");
            if (dialog) {
                dialog.close();
            }
        }
    });
});
