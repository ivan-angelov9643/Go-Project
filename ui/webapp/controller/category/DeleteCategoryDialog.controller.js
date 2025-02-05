sap.ui.define([
    "../BaseController",
    "sap/m/MessageToast",
    "sap/ui/model/json/JSONModel",
    "sap/ui/core/Core"
], function (BaseController, MessageToast, JSONModel, Core) {
    "use strict";

    return BaseController.extend("library-app.controller.category.DeleteCategoryDialog", {
        onInit: function () {
            this.oDialogCategoryModel = new JSONModel(this.initCategoryModel());
            this.getView().setModel(this.oDialogCategoryModel, "dialogCategory");
        },

        onConfirmDelete: async function () {
            const categoryData = this.getView().getModel("dialogCategory").getData();

            try {
                const token = await this.getOwnerComponent().getToken();
                const deleteResponse = await this.sendRequest(
                    `http://localhost:8080/api/categories/${categoryData.id}`,
                    "DELETE",
                    token
                );

                Core.getEventBus().publish("library-app", "categoriesUpdated", deleteResponse);
                MessageToast.show("Successfully deleted category");
            } catch (error) {
                MessageToast.show(error.error || "Error deleting category");
            }

            this.onDialogClose();
        },

        onCancelDelete: function () {
            this.onDialogClose();
        },

        onDialogClose: function () {
            const dialog = this.byId("deleteCategoryDialog");
            if (dialog) {
                dialog.close();
            }
        }
    });
});
