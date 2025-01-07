sap.ui.define([
    "./BaseController",
    "sap/m/MessageToast",
    "sap/ui/model/json/JSONModel",
    "sap/ui/core/Core",
], function (BaseController, MessageToast, JSONModel, Core) {
    "use strict";

    return BaseController.extend("myApp.controller.EditProfileDialog", {
        onInit: function () {
            this.oDialogUserModel = new JSONModel(this.initUserModel());
            this.getView().setModel(this.oDialogUserModel, "dialogUser");
        },

        onSaveProfile: async function () {
            const userData = this.oDialogUserModel.getData();

            try {
                const token = await this.getOwnerComponent().getToken();
                const userID = JSON.parse(atob(token.split(".")[1])).sub
                const updateResponse = await this.sendRequest(
                    `http://localhost:8080/api/users/${userID}`,
                    "PUT",
                    token,
                    userData
                );

                Core.getEventBus().publish("library-app", "userUpdated", updateResponse);

                MessageToast.show("Successfully saved profile details");
            } catch (error) {
                MessageToast.show(error.error || "Error updating profile details");
                return;
            }
            this.onDialogClose();
        },

        onCancelEdit: function () {
            this.onDialogClose();
        },

        onDialogClose: function () {
            const dialog = this.byId("editProfileDialog");
            if (dialog) {
                dialog.close();
            }
        }
    });
});
