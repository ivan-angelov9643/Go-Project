sap.ui.define([
    "../BaseController",
    "sap/m/MessageToast",
    "sap/ui/model/json/JSONModel",
    "sap/ui/core/Core"
], function (BaseController, MessageToast, JSONModel, Core) {
    "use strict";

    return BaseController.extend("library-app.controller.reservation.DeleteReservationDialog", {
        onInit: function () {
            this.oDialogReservationModel = new JSONModel(this.initReservationModel());
            this.getView().setModel(this.oDialogReservationModel, "dialogReservation");
        },

        onConfirmDelete: async function () {
            const reservationData = this.getView().getModel("dialogReservation").getData();

            try {
                const token = await this.getOwnerComponent().getToken();
                const deleteResponse = await this.sendRequest(
                    `http://localhost:8080/api/reservations/${reservationData.id}`,
                    "DELETE",
                    token
                );

                Core.getEventBus().publish("library-app", "reservationsUpdated", {delete_reservation: true});
                MessageToast.show("Successfully deleted reservation");
            } catch (error) {
                MessageToast.show(error.error || "Error deleting reservation");
            }

            this.onDialogClose();
        },

        onCancelDelete: function () {
            this.onDialogClose();
        },

        onDialogClose: function () {
            const dialog = this.byId("deleteReservationDialog");
            if (dialog) {
                dialog.close();
            }
        },
    });
});
