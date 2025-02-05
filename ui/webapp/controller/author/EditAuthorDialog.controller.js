sap.ui.define([
    "../BaseController",
    "sap/m/MessageToast",
    "sap/ui/model/json/JSONModel",
    "sap/ui/core/Core"
], function (BaseController, MessageToast, JSONModel, Core) {
    "use strict";

    return BaseController.extend("library-app.controller.author.EditAuthorDialog", {
        onInit: function () {
            this.oDialogAuthorModel = new JSONModel(this.initAuthorModel());
            this.getView().setModel(this.oDialogAuthorModel, "dialogAuthor");
        },

        onSaveAuthor: async function () {
            const authorData = this.oDialogAuthorModel.getData();

            try {
                if (authorData.birth_date) {
                    authorData.birth_date = this.toISO8601(authorData.birth_date);
                }
                if (authorData.death_date) {
                    authorData.death_date = this.toISO8601(authorData.death_date);
                }

                const token = await this.getOwnerComponent().getToken();
                const updateResponse = await this.sendRequest(
                    `http://localhost:8080/api/authors/${authorData.id}`,
                    "PUT",
                    token,
                    authorData
                );

                Core.getEventBus().publish("library-app", "authorsUpdated", updateResponse);

                MessageToast.show("Successfully saved author details");
            } catch (error) {
                MessageToast.show(error.error || "Error updating author details");
                return;
            }

            this.onDialogClose();
        },

        onCancelEdit: function () {
            this.onDialogClose();
        },

        onDialogClose: function () {
            const dialog = this.byId("editAuthorDialog");
            if (dialog) {
                dialog.close();
            }
        },
    });
});
