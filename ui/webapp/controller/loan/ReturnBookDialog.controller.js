sap.ui.define([
    "../BaseController",
    "sap/m/MessageToast",
    "sap/ui/model/json/JSONModel",
    "sap/ui/core/Core"
], function (BaseController, MessageToast, JSONModel, Core) {
    "use strict";

    return BaseController.extend("library-app.controller.loan.ReturnBookDialog", {
        onInit: function () {
            this.oDialogLoanModel = new JSONModel(this.initLoanModel());
            this.getView().setModel(this.oDialogLoanModel, "dialogLoan");
        },

        onConfirmReturn: async function () {
            const loanData = this.oDialogLoanModel.getData();
            loanData.return_date = new Date().toISOString();
            loanData.status = "completed";

            try {
                const token = await this.getOwnerComponent().getToken();
                const returnResponse = await this.sendRequest(
                    `http://localhost:8080/api/loans/${loanData.id}`,
                    "PUT",
                    token,
                    loanData
                );

                Core.getEventBus().publish("library-app", "loansUpdated", {return_book: true});

                MessageToast.show("Successfully returned the book");
            } catch (error) {
                MessageToast.show(error.error || "Error processing the return");
                return;
            }

            this.onDialogClose();
        },

        onCancelReturn: function () {
            this.onDialogClose();
        },

        onDialogClose: function () {
            const dialog = this.byId("returnBookDialog");
            if (dialog) {
                dialog.close();
            }
        },
    });
});
