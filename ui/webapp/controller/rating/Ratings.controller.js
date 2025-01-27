sap.ui.define([
	'../BaseController',
	'sap/ui/model/json/JSONModel',
	'sap/ui/Device',
	'library-app/model/formatter',
	"sap/ui/core/Core",
	"sap/ui/core/mvc/XMLView",
	"sap/ui/model/Filter",
	"sap/ui/model/FilterOperator",
], function (BaseController, JSONModel, Device, formatter, Core, XMLView, Filter, FilterOperator) {
	"use strict";
	return BaseController.extend("library-app.controller.rating.Ratings", {
		formatter: formatter,

		onInit: async function () {
			const oRouter = this.getOwnerComponent().getRouter();
			oRouter.attachRoutePatternMatched(this.loadData, this);

			Core.getEventBus().subscribe("library-app", "ratingsUpdated", this.handleRatingsUpdated, this);

			this.oRatingModel = new JSONModel({
				count: null,
				page_size: null,
				page: null,
				data: null,
				total_pages: null,
			});
			this.oRatingModel.setSizeLimit(Number.MAX_VALUE);
			this.getView().setModel(this.oRatingModel, "rating");
			await this.loadRatings(this.oRatingModel, 1);
		},

		loadData: async function() {
			await this.loadRatings(this.oRatingModel, this.oRatingModel.getData().page);
		},

		onExit: function () {
			Core.getEventBus().unsubscribe("library-app", "ratingsUpdated", this.handleRatingsUpdated, this);
		},

		handleRatingsUpdated: async function (ns, ev, eventData) {
			await this.loadData();

			eventData.from_ratings = true;
			Core.getEventBus().publish("library-app", "booksUpdated", eventData);
		},

		onTitleSearchChange: function(oEvent) {
			this.sTitleSearch = oEvent.getParameter("value");
			this._applyCombinedFilters();
		},

		onUsernameSearchChange: function(oEvent) {
			this.sUsernameSearch = oEvent.getParameter("value");
			this._applyCombinedFilters();
		},

		_applyCombinedFilters: function() {
			let aFilters = [];

			if (this.sTitleSearch && this.sTitleSearch.trim() !== "") {
				aFilters.push(
					new Filter("book_title", FilterOperator.Contains, this.sTitleSearch)
				);
			}

			if (this.sUsernameSearch && this.sUsernameSearch.trim() !== "") {
				aFilters.push(
					new Filter("user_name", FilterOperator.Contains, this.sUsernameSearch)
				);
			}

			let oTable = this.getView().byId("ratingsTable");
			let oBinding = oTable.getBinding("items");

			oBinding.filter(aFilters);
		},

		onPreviousPage: async function () {
			await this.loadRatings(this.oRatingModel, this.oRatingModel.getData().page - 1);
		},

		onNextPage: async function () {
			await this.loadRatings(this.oRatingModel, this.oRatingModel.getData().page + 1);
		},
	});
});