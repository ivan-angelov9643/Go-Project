sap.ui.define([
	'./BaseController',
	'sap/m/ActionSheet',
	'sap/m/Button',
	'sap/m/MessageToast',
	'sap/ui/Device',
	'sap/m/library',
	'library-app/model/formatter',
	"sap/ui/core/mvc/XMLView",
	"sap/ui/core/Core",
	// '../libs/keycloak-js/dist/keycloak'
], function (
	BaseController,
	ActionSheet,
	Button,
	MessageToast,
	Device,
	mobileLibrary,
	formatter,
	XMLView,
	Core
	// keycloakLibrary
) {
	"use strict";
	return BaseController.extend("library-app.controller.App", {
		formatter: formatter,

		onInit: async function () {
			if (Device.resize.width <= 1024) {
				this.onSideNavButtonPress();
			}
			Device.media.attachHandler(this._handleWindowResize, this);

			this.getRouter().attachRouteMatched(this.onRouteChange.bind(this));
			this.getRouter().navTo("home")

			await this._initializeMainViews()
		},
		
		logOut: function () {
			const authModel = this.getView().getModel("authModel")
			const keycloak = authModel.getProperty("/keycloak");
			authModel.setProperty("/keycloak", null);
			authModel.setProperty("/username", "Unknown");
			if(keycloak){
				keycloak.logout();
			}
		},

		onExit: function () {
			Device.media.detachHandler(this._handleWindowResize, this);
		},

		onRouteChange: function (oEvent) {
			const selectedPageKey = oEvent.getParameter('name')
			this.getView().getModel('side').setProperty('/selectedKey', selectedPageKey);

			sap.ui.getCore().getEventBus().publish("library-app", "RouteChanged", { selectedPageKey });
		},

		onUserNamePress: function (oEvent) {
			var oSource = oEvent.getSource();
			this.getView().getModel("i18n").getResourceBundle().then(function (oBundle) {
				var fnHandleUserMenuLogoutPress = function (oEvent) {
					this.logOut()
				}.bind(this);

				var oActionSheet = new ActionSheet(this.getView().createId("userMessageActionSheet"), {
					title: oBundle.getText("userHeaderTitle"),
					showCancelButton: false,
					buttons: [
						new Button({
							text: '{i18n>userAccountLogout}',
							type: mobileLibrary.ButtonType.Transparent,
							press: fnHandleUserMenuLogoutPress
						})
					],
					afterClose: function () {
						oActionSheet.destroy();
					}
				});
				this.getView().addDependent(oActionSheet);
				oActionSheet.openBy(oSource);
			}.bind(this));
		},

		onSideNavButtonPress: function () {
			var oToolPage = this.byId("app");
			oToolPage.setSideExpanded(!oToolPage.getSideExpanded());
		},

		_handleWindowResize: function (oDevice) {
			var oToolPage = this.byId("app");
			if (Device.resize.width < 1024) {
				oToolPage.setSideExpanded(false);
			}
			if (Device.resize.width >= 1024) {
				oToolPage.setSideExpanded(true);
			}
		},

		_initializeMainViews: async function () {
			const oView = this.getView();
			const aMainViews = [
				{ id: "authorsView", name: "library-app.view.author.Authors" },
				{ id: "booksView", name: "library-app.view.book.Books" },
				{ id: "categoriesView", name: "library-app.view.category.Categories" },
				{ id: "loansView", name: "library-app.view.loan.Loans" },
				{ id: "ratingsView", name: "library-app.view.rating.Ratings" },
				{ id: "reservationsView", name: "library-app.view.reservation.Reservations" },
				{ id: "usersView", name: "library-app.view.user.Users" }
			];

			for (let i = 0; i < aMainViews.length; i++) {
				const viewInfo = aMainViews[i];

				let oExistingView = oView.byId(viewInfo.id);
				if (!oExistingView) {
					const oOwnerComponent = this.getOwnerComponent();
					oOwnerComponent.runAsOwner(() => {
						XMLView.create({
							id: viewInfo.id,
							viewName: viewInfo.name
						}).then((oViewInstance) => {
							oView.addDependent(oViewInstance);
						});
					});
				}
			}
		}
	});
});