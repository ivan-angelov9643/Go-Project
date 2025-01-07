sap.ui.define([
  "sap/ui/core/library",
  "sap/ui/core/UIComponent",
  "./model/models",
  "sap/ui/model/json/JSONModel",
  './libs/keycloak-js/dist/keycloak'
], function(library, UIComponent, models, JSONModel) {
  "use strict";

  return UIComponent.extend("library-app.Component", {
    metadata: {
      manifest: "json",
      interfaces: [library.IAsyncContentCreation]
    },

    init: async function () {
      UIComponent.prototype.init.apply(this, arguments);

      this.authModel = new JSONModel({
        keycloak: null,
        username: "Unknown",
      })

      const keycloak = new Keycloak('/libs/keycloak-cfg/keycloak.json');
      const authenticated = await keycloak.init({ onLoad: 'login-required', checkLoginIframe: false })
      if(authenticated){
        this.authModel.setProperty("/keycloak", keycloak);
        this.authModel.setProperty("/username", keycloak.tokenParsed?.preferred_username);
      } else {
        keycloak.logout();
      }
      this.setModel(this.authModel, "authModel");

      if( keycloak.hasRealmRole("mainrole.administrator")) {
        const sideModel = this.getModel('sideAdministrator')
        this.setModel(sideModel, "side");
      }
      if( keycloak.hasRealmRole("mainrole.librarian")) {
        const sideModel = this.getModel('sideLibrarian')
        this.setModel(sideModel, "side");
      }
      if( keycloak.hasRealmRole("mainrole.user")) {
        const sideModel = this.getModel('sideUser')
        this.setModel(sideModel, "side");
      }

      this.getRouter().initialize();
    },

    getToken: async function () {
      const authModel = this.getModel("authModel");
      const keycloak = authModel.getProperty("/keycloak");
      if(keycloak.isTokenExpired()){
        try {
          const tokenUpdated = await keycloak.updateToken(1800)
          if(tokenUpdated) {
            authModel.setProperty("/username", keycloak.tokenParsed?.preferred_username); // ?
          } else {
            keycloak.logout();
          }
        } catch (error) {
          keycloak.logout();
        }
      }
      return keycloak.token;
    },

    hasRole: function (role) {
      const authModel = this.getModel("authModel");
      const keycloak = authModel.getProperty("/keycloak");
      return keycloak.hasRealmRole(role)
    },
  });
});