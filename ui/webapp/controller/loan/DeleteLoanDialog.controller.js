sap.ui.define([
    "../BaseController",
    "sap/m/MessageToast",
    "sap/ui/model/json/JSONModel",
    "sap/ui/core/Core"
], function (BaseController, MessageToast, JSONModel, Core) {
    "use strict";

    return BaseController.extend("library-app.controller.loan.DeleteLoanDialog", {
        onInit: function () {
            this.oDialogLoanModel = new JSONModel(this.initLoanModel());
            this.getView().setModel(this.oDialogLoanModel, "dialogLoan");
        },

        onConfirmDelete: async function () {
            const loanData = this.getView().getModel("dialogLoan").getData();

            try {
                const token = await this.getOwnerComponent().getToken();
                const deleteResponse = await this.sendRequest(
                    `http://localhost:8080/api/loans/${loanData.id}`,
                    "DELETE",
                    token
                );

                Core.getEventBus().publish("library-app", "loansUpdated", {delete_loan: true});
                MessageToast.show("Successfully deleted loan");
            } catch (error) {
                MessageToast.show(error.error || "Error deleting loan");
            }

            this.onDialogClose();
        },

        onCancelDelete: function () {
            this.onDialogClose();
        },

        onDialogClose: function () {
            const dialog = this.byId("deleteLoanDialog");
            if (dialog) {
                dialog.close();
            }
        },
    });
});
