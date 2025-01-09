sap.ui.define([
    "../BaseController",
    "sap/m/MessageToast",
    "sap/ui/model/json/JSONModel",
    "sap/ui/core/Core",
    "library-app/model/formatter",
], function (BaseController, MessageToast, JSONModel, Core, formatter) {
    "use strict";

    return BaseController.extend("library-app.controller.reservation.ExtendReservationDialog", {
        formatter: formatter,

        onInit: function () {
            this.oDialogReservationModel = new JSONModel(this.initReservationModel());
            this.getView().setModel(this.oDialogReservationModel, "dialogReservation");
        },

        onExtendReservation: async function () {
            const reservationData = this.oDialogReservationModel.getData();
            const daysToExtend = parseInt(reservationData.days_to_extend, 10) || 0;
            const hoursToExtend = parseInt(reservationData.hours_to_extend, 10) || 0;
            const minutesToExtend = parseInt(reservationData.minutes_to_extend, 10) || 0;

            if (daysToExtend <= 0 && hoursToExtend <= 0 && minutesToExtend <= 0) {
                MessageToast.show("Please enter a valid time to extend (days, hours, or minutes).");
                return;
            }

            try {
                const token = await this.getOwnerComponent().getToken();

                const currentExpiryDate = new Date(reservationData.expiry_date);
                const newExpiryDate = new Date(currentExpiryDate);

                newExpiryDate.setDate(currentExpiryDate.getDate() + daysToExtend);
                newExpiryDate.setHours(currentExpiryDate.getHours() + hoursToExtend);
                newExpiryDate.setMinutes(currentExpiryDate.getMinutes() + minutesToExtend);

                reservationData.expiry_date = newExpiryDate.toISOString();

                const updateResponse = await this.sendRequest(
                    `http://localhost:8080/api/reservations/${reservationData.id}`,
                    "PUT",
                    token,
                    reservationData
                );

                Core.getEventBus().publish("library-app", "reservationsUpdated", updateResponse);

                MessageToast.show("Successfully extended the reservation.");
            } catch (error) {
                MessageToast.show(error.error || "Error extending the reservation.");
                return;
            }

            this.onDialogClose();
            this.fillReservationModel(this.oDialogReservationModel, this.initReservationModel());
        },

        onCancelExtend: function () {
            this.onDialogClose();
        },

        onDialogClose: function () {
            const dialog = this.byId("extendReservationDialog");
            if (dialog) {
                dialog.close();
            }
        },
    });
});
