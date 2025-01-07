sap.ui.define([
	'./BaseController',
	'sap/ui/model/json/JSONModel',
	'sap/ui/Device',
	'library-app/model/formatter'
], function (BaseController, JSONModel, Device, formatter) {
	"use strict";
	return BaseController.extend("library-app.controller.Reviews", {
		formatter: formatter,

		onInit: function () {
			// sap.ui.getCore().getEventBus().subscribe("library-app", "RouteChanged", this.handleRouteChanged, this);

			this.oReviewModel = new JSONModel({
				reviews: null,
			});
			this.oReviewModel.setSizeLimit(Number.MAX_VALUE);
			this.getView().setModel(this.oReviewModel, "review");
			this.loadReviews(this.oReviewModel);
		},

		onExit: function () {
			// sap.ui.getCore().getEventBus().unsubscribe("library-app", "RouteChanged", this.handleRouteChanged, this);
		},
	});
});