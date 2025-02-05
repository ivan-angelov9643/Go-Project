sap.ui.define([
    "../BaseController",
    "sap/m/MessageToast",
    "sap/ui/model/json/JSONModel",
    "sap/ui/core/Core"
], function (BaseController, MessageToast, JSONModel, Core) {
    "use strict";

    return BaseController.extend("library-app.controller.rating.CreateRatingDialog", {
        onInit: function () {
            this.oDialogRatingModel = new JSONModel(this.initRatingModel());
            this.getView().setModel(this.oDialogRatingModel, "dialogRating");
        },

        onCreateRating: async function () {
            const token = await this.getOwnerComponent().getToken();
            const ratingData = this.oDialogRatingModel.getData();
            ratingData.user_id = await this.getUserID(token);
            ratingData.value = parseInt(ratingData.value, 10);

            if (!ratingData.value || ratingData.value < 1 || ratingData.value > 5) {
                MessageToast.show("Please enter a valid rating between 1 and 5.");
                return;
            }

            try {
                await this.sendRequest(
                    `http://localhost:8080/api/ratings`,
                    "POST",
                    token,
                    ratingData
                );

                Core.getEventBus().publish("library-app", "ratingsUpdated",
                    {
                        create_rating: true,
                        book_id: ratingData.book_id
                    }
                );

                MessageToast.show("Successfully created rating.");
            } catch (error) {
                MessageToast.show(error.error || "Error creating rating.");
                return;
            }

            this.onDialogClose();
            this.fillRatingModel(this.oDialogRatingModel, this.initRatingModel());
        },

        onCancelCreate: function () {
            this.onDialogClose();
        },

        onDialogClose: function () {
            const dialog = this.byId("createRatingDialog");
            if (dialog) {
                dialog.close();
            }
        }
    });
});
