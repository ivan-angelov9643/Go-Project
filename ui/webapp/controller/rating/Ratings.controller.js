sap.ui.define([
	'../BaseController',
	'sap/ui/model/json/JSONModel',
	'sap/ui/Device',
	'library-app/model/formatter',
	"sap/ui/core/Core",
	"sap/ui/core/mvc/XMLView",
], function (BaseController, JSONModel, Device, formatter, Core, XMLView) {
	"use strict";
	return BaseController.extend("library-app.controller.rating.Ratings", {
		formatter: formatter,

		onInit: async function () {
			const oRouter = this.getOwnerComponent().getRouter();
			oRouter.attachRoutePatternMatched(this.loadData, this);

			Core.getEventBus().subscribe("library-app", "ratingsUpdated", this.handleRatingsUpdated, this);

			this.oRatingModel = new JSONModel({
				ratings: null,
			});
			this.oRatingModel.setSizeLimit(Number.MAX_VALUE);
			this.getView().setModel(this.oRatingModel, "rating");
			await this.loadRatings(this.oRatingModel);
		},

		loadData: async function() {
			await this.loadRatings(this.oRatingModel);
		},

		onDeleteRating: async function (oEvent) {
			if (!this._oDeleteRatingDialog) {
				const oOwnerComponent = this.getOwnerComponent();
				oOwnerComponent.runAsOwner(() => {
					this._oDeleteRatingDialog = new XMLView({
						id: "deleteRatingDialogView",
						viewName: "library-app.view.rating.DeleteRatingDialog",
					});
					this.getView().addDependent(this._oDeleteRatingDialog);
				});
			}

			const oContext = oEvent.getSource().getBindingContext("rating");
			const oData = oContext.getObject();

			const oDialogRatingModel = this._oDeleteRatingDialog.getModel("dialogRating");
			this.fillRatingModel(oDialogRatingModel, oData);

			this._oDeleteRatingDialog.byId("deleteRatingDialog").open();
		},

		onExit: function () {
			Core.getEventBus().unsubscribe("library-app", "ratingsUpdated", this.handleRatingsUpdated, this);
		},

		handleRatingsUpdated: async function (ns, ev, eventData) {
			await this.loadRatings(this.oRatingModel);

			eventData.from_ratings = true;
			Core.getEventBus().publish("library-app", "booksUpdated", eventData);
		}
	});
});