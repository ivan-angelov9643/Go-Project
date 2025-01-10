sap.ui.define([
    "../BaseController",
    "sap/m/MessageToast",
    "sap/ui/model/json/JSONModel",
    "sap/ui/core/Core",
    "library-app/model/formatter",
], function (BaseController, MessageToast, JSONModel, Core, formatter) {
    "use strict";

    return BaseController.extend("library-app.controller.loan.MakeLoanDialog", {
        formatter: formatter,

        onInit: function () {
            this.oDialogLoanModel = new JSONModel(this.initLoanModel());
            this.getView().setModel(this.oDialogLoanModel, "dialogLoan");

            this.oDialogReservationModel = new JSONModel(this.initReservationModel());
            this.getView().setModel(this.oDialogReservationModel, "dialogReservation");
        },

        onCreateLoan: async function () {
            const loanData = this.oDialogLoanModel.getData();
            const reservationData = this.oDialogReservationModel.getData();

            const loanDurationDays = parseInt(loanData.loan_duration, 10) || 0;

            if (loanDurationDays <= 0) {
                MessageToast.show("Please enter a valid loan duration (days).");
                return;
            }

            try {
                const token = await this.getOwnerComponent().getToken();

                const startDate = new Date();
                const dueDate = new Date(startDate);
                dueDate.setDate(startDate.getDate() + loanDurationDays);

                const body = {
                    user_id: reservationData.user_id,
                    book_id: reservationData.book_id,
                    start_date: startDate.toISOString(),
                    due_date: dueDate.toISOString(),
                    status: "active"
                };

                await this.sendRequest(
                    `http://localhost:8080/api/loans`,
                    "POST",
                    token,
                    body
                );

                await this.sendRequest(
                    `http://localhost:8080/api/reservations/${reservationData.id}`,
                    "DELETE",
                    token
                );

                Core.getEventBus().publish("library-app", "reservationsUpdated", {make_loan: true});

                MessageToast.show("Loan created successfully.");
            } catch (error) {
                MessageToast.show(error.error || "Error creating loan.");
                return;
            }

            this.onDialogClose();
            this.fillLoanModel(this.oDialogLoanModel, this.initLoanModel());
        },

        onCancelMakeLoan: function () {
            this.onDialogClose();
        },

        onDialogClose: function () {
            const dialog = this.byId("makeLoanDialog");
            if (dialog) {
                dialog.close();
            }
        },
    });
});
