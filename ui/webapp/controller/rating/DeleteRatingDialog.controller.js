sap.ui.define([
    "../BaseController",
    "sap/m/MessageToast",
    "sap/ui/model/json/JSONModel",
    "sap/ui/core/Core"
], function (BaseController, MessageToast, JSONModel, Core) {
    "use strict";

    return BaseController.extend("library-app.controller.rating.DeleteRatingDialog", {
        onInit: function () {
            this.oDialogRatingModel = new JSONModel(this.initRatingModel());
            this.getView().setModel(this.oDialogRatingModel, "dialogRating");
        },

        onConfirmDelete: async function () {
            const ratingData = this.getView().getModel("dialogRating").getData();

            try {
                const token = await this.getOwnerComponent().getToken();
                const deleteResponse = await this.sendRequest(
                    `http://localhost:8080/api/ratings/${ratingData.id}`,
                    "DELETE",
                    token
                );

                Core.getEventBus().publish("library-app", "ratingsUpdated", {book_id: ratingData.book_id});
                MessageToast.show("Successfully deleted rating");
            } catch (error) {
                MessageToast.show(error.error || "Error deleting rating");
            }

            this.onDialogClose();
        },

        onCancelDelete: function () {
            this.onDialogClose();
        },

        onDialogClose: function () {
            const dialog = this.byId("deleteRatingDialog");
            if (dialog) {
                dialog.close();
            }
        },
    });
});
