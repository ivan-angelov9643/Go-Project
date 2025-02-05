sap.ui.define([
    "../BaseController",
    "sap/m/MessageToast",
    "sap/ui/model/json/JSONModel",
    "sap/ui/core/Core"
], function (BaseController, MessageToast, JSONModel, Core) {
    "use strict";

    return BaseController.extend("library-app.controller.user.EditUserDialog", {
        onInit: function () {
            this.oDialogUserModel = new JSONModel(this.initUserModel());
            this.getView().setModel(this.oDialogUserModel, "dialogUser");
        },

        onSaveUser: async function () {
            const userData = this.oDialogUserModel.getData();

            try {
                const token = await this.getOwnerComponent().getToken();
                const userID = userData.id;
                const updateResponse = await this.sendRequest(
                    `http://localhost:8080/api/users/${userID}`,
                    "PUT",
                    token,
                    userData
                );

                Core.getEventBus().publish("library-app", "usersUpdated", updateResponse);

                MessageToast.show("Successfully saved user details");
            } catch (error) {
                MessageToast.show(error.error || "Error updating user details");
                return;
            }

            this.onDialogClose();
        },

        onCancelEdit: function () {
            this.onDialogClose();
        },

        onDialogClose: function () {
            const dialog = this.byId("editUserDialog");
            if (dialog) {
                dialog.close();
            }
        },
    });
});
