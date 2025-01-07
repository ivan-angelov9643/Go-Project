sap.ui.define([
    "../BaseController",
    "sap/m/MessageToast",
    "sap/ui/model/json/JSONModel",
    "sap/ui/core/Core",
    'library-app/model/formatter',
], function (BaseController, MessageToast, JSONModel, Core, formatter) {
    "use strict";

    return BaseController.extend("library-app.controller.loan.ExtendLoanDialog", {
        formatter: formatter,

        onInit: function () {
            this.oDialogLoanModel = new JSONModel(this.initLoanModel());
            this.getView().setModel(this.oDialogLoanModel, "dialogLoan");
        },

        onExtendLoan: async function () {
            const loanData = this.oDialogLoanModel.getData();
            const daysToExtend = parseInt(loanData.days_to_extend, 10);

            if (isNaN(daysToExtend) || daysToExtend <= 0) {
                MessageToast.show("Please enter a valid number of days to extend.");
                return;
            }

            try {
                const token = await this.getOwnerComponent().getToken();

                const currentDueDate = new Date(loanData.due_date);
                const newDueDate = new Date(currentDueDate);
                newDueDate.setDate(currentDueDate.getDate() + daysToExtend);

                loanData.due_date = newDueDate.toISOString();

                const updateResponse = await this.sendRequest(
                    `http://localhost:8080/api/loans/${loanData.id}`,
                    "PUT",
                    token,
                    loanData
                );

                Core.getEventBus().publish("library-app", "loansUpdated", updateResponse);

                MessageToast.show("Successfully extended the loan.");
            } catch (error) {
                MessageToast.show(error.error || "Error extending the loan.");
                return;
            }

            this.onDialogClose();
            this.fillLoanModel(this.oDialogLoanModel, this.initLoanModel());
        },

        onCancelExtend: function () {
            this.onDialogClose();
        },

        onDialogClose: function () {
            const dialog = this.byId("extendLoanDialog");
            if (dialog) {
                dialog.close();
            }
        },
    });
});
