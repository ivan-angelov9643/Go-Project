sap.ui.define([
    "../BaseController",
    "sap/m/MessageToast",
    "sap/ui/model/json/JSONModel",
    "sap/ui/core/Core"
], function (BaseController, MessageToast, JSONModel, Core) {
    "use strict";

    return BaseController.extend("library-app.controller.author.CreateAuthorDialog", {
        onInit: function () {
            this.oDialogAuthorModel = new JSONModel(this.initAuthorModel());
            this.getView().setModel(this.oDialogAuthorModel, "dialogAuthor");
        },

        onCreateAuthor: async function () {
            const authorData = this.oDialogAuthorModel.getData();

            try {
                if (authorData.birth_date) {
                    authorData.birth_date = this.toISO8601(authorData.birth_date);
                }
                if (authorData.death_date) {
                    authorData.death_date = this.toISO8601(authorData.death_date);
                }

                const token = await this.getOwnerComponent().getToken();
                const createResponse = await this.sendRequest(
                    `http://localhost:8080/api/authors`,
                    "POST",
                    token,
                    authorData
                );

                Core.getEventBus().publish("library-app", "authorsUpdated", createResponse);

                MessageToast.show("Successfully created author");
            } catch (error) {
                MessageToast.show(error.error || "Error creating author");
                return;
            }

            this.onDialogClose();
        },

        onCancelCreate: function () {
            this.onDialogClose();
        },

        onDialogClose: function () {
            this.fillAuthorModel(this.oDialogAuthorModel, this.initAuthorModel());
            const dialog = this.byId("createAuthorDialog");
            if (dialog) {
                dialog.close();
            }
        },
    });
});
