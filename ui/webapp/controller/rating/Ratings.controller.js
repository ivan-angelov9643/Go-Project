sap.ui.define([
	'../BaseController',
	'sap/ui/model/json/JSONModel',
	'sap/ui/Device',
	'library-app/model/formatter',
	"sap/ui/core/Core",
	"sap/ui/core/mvc/XMLView",
	"sap/ui/model/Filter",
	"sap/ui/model/FilterOperator",
], function (BaseController, JSONModel, Device, formatter, Core) {
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

			this._setDefaultSearchFields()
		},

		loadData: async function() {
			await this.loadRatings(this.oRatingModel, this.oRatingModel.getData().page, null,
				this.sUsernameSearch, this.sTitleSearch);
		},

		onExit: function () {
			Core.getEventBus().unsubscribe("library-app", "ratingsUpdated", this.handleRatingsUpdated, this);
		},

		handleRatingsUpdated: async function (ns, ev, eventData) {
			await this.loadData();

			eventData.from_ratings = true;
			Core.getEventBus().publish("library-app", "booksUpdated", eventData);
		},

		onPreviousPage: async function () {
			await this.loadRatings(this.oRatingModel, this.oRatingModel.getData().page - 1, null,
				this.sUsernameSearch, this.sTitleSearch);
		},

		onNextPage: async function () {
			await this.loadRatings(this.oRatingModel, this.oRatingModel.getData().page + 1, null,
				this.sUsernameSearch, this.sTitleSearch);
		},

		onSearch: async function () {
			if (!this._searchFieldsChanged()) {
				return;
			}

			this.sUsernameSearch = this.byId("usernameSearch").getValue();
			this.sTitleSearch = this.byId("titleSearch").getValue();

			await this.loadRatings(this.oRatingModel, 1, null,
				this.sUsernameSearch, this.sTitleSearch);
		},

		onClearSearch: async function () {
			if (this._searchFieldsEmpty()) {
				return;
			}

			this._setDefaultSearchFields()

			this.byId("usernameSearch").setValue("");
			this.byId("titleSearch").setValue("");

			await this.loadRatings(this.oRatingModel, 1);
		},

		_setDefaultSearchFields() {
			this.sUsernameSearch = "";
			this.sTitleSearch = "";
		},

		_searchFieldsChanged() {
			return this.sUsernameSearch !== this.byId("usernameSearch").getValue() ||
				this.sTitleSearch !== this.byId("titleSearch").getValue();
		},

		_searchFieldsEmpty() {
			return this.sUsernameSearch === "" &&
				this.sTitleSearch === "";
		},
	});
});