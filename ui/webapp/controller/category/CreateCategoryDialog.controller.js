sap.ui.define([
    "../BaseController",
    "sap/m/MessageToast",
    "sap/ui/model/json/JSONModel",
    "sap/ui/core/Core"
], function (BaseController, MessageToast, JSONModel, Core) {
    "use strict";

    return BaseController.extend("library-app.controller.category.CreateCategoryDialog", {
        onInit: function () {
            this.oDialogCategoryModel = new JSONModel(this.initCategoryModel());
            this.getView().setModel(this.oDialogCategoryModel, "dialogCategory");
        },

        onCreateCategory: async function () {
            const categoryData = this.oDialogCategoryModel.getData();

            try {
                const token = await this.getOwnerComponent().getToken();
                const createResponse = await this.sendRequest(
                    `http://localhost:8080/api/categories`,
                    "POST",
                    token,
                    categoryData
                );

                Core.getEventBus().publish("library-app", "categoriesUpdated", createResponse);

                MessageToast.show("Successfully created category");
            } catch (error) {
                MessageToast.show(error.error || "Error creating category");
                return;
            }

            this.onDialogClose();
            this.fillCategoryModel(this.oDialogCategoryModel, this.initCategoryModel());
        },

        onCancelCreate: function () {
            this.onDialogClose();
        },

        onDialogClose: function () {
            const dialog = this.byId("createCategoryDialog");
            if (dialog) {
                dialog.close();
            }
        },
    });
});
