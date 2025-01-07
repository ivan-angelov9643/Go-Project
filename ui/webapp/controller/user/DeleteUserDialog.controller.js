sap.ui.define([
    "../BaseController",
    "sap/m/MessageToast",
    "sap/ui/model/json/JSONModel",
    "sap/ui/core/Core"
], function (BaseController, MessageToast, JSONModel, Core) {
    "use strict";

    return BaseController.extend("library-app.controller.user.DeleteUserDialog", {
        onInit: function () {
            this.oDialogUserModel = new JSONModel(this.initUserModel());
            this.getView().setModel(this.oDialogUserModel, "dialogUser");
        },

        onConfirmDelete: async function () {
            const userData = this.getView().getModel("dialogUser").getData();

            try {
                const token = await this.getOwnerComponent().getToken();
                const deleteResponse = await this.sendRequest(
                    `http://localhost:8080/api/users/${userData.id}`,
                    "DELETE",
                    token
                );

                Core.getEventBus().publish("library-app", "usersUpdated", deleteResponse);
                MessageToast.show("Successfully deleted user");
            } catch (error) {
                MessageToast.show(error.error || "Error deleting user");
            }

            this.onDialogClose();
        },

        onCancelDelete: function () {
            this.onDialogClose();
        },

        onDialogClose: function () {
            const dialog = this.byId("deleteUserDialog");
            if (dialog) {
                dialog.close();
            }
        }
    });
});
