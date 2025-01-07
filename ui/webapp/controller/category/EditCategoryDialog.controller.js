sap.ui.define([
    "../BaseController",
    "sap/m/MessageToast",
    "sap/ui/model/json/JSONModel",
    "sap/ui/core/Core"
], function (BaseController, MessageToast, JSONModel, Core) {
    "use strict";

    return BaseController.extend("library-app.controller.category.EditCategoryDialog", {
        onInit: function () {
            this.oDialogCategoryModel = new JSONModel(this.initCategoryModel());
            this.getView().setModel(this.oDialogCategoryModel, "dialogCategory");
        },

        onSaveCategory: async function () {
            const categoryData = this.oDialogCategoryModel.getData();

            try {
                const token = await this.getOwnerComponent().getToken();
                const updateResponse = await this.sendRequest(
                    `http://localhost:8080/api/categories/${categoryData.id}`,
                    "PUT",
                    token,
                    categoryData
                );

                Core.getEventBus().publish("library-app", "categoriesUpdated", updateResponse);

                MessageToast.show("Successfully saved category details");
            } catch (error) {
                MessageToast.show(error.error || "Error updating category details");
                return;
            }

            this.onDialogClose();
        },

        onCancelEdit: function () {
            this.onDialogClose();
        },

        onDialogClose: function () {
            const dialog = this.byId("editCategoryDialog");
            if (dialog) {
                dialog.close();
            }
        },
    });
});
