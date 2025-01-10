sap.ui.define([
	"sap/base/strings/formatMessage"
], function (formatMessage) {
	"use strict";

	return {
		formatMessage: formatMessage,

		formatDate: function (date) {
			const oDate = new Date(date);

			const months = [
				"January", "February", "March", "April", "May", "June",
				"July", "August", "September", "October", "November", "December"
			];

			const day = oDate.getDate();
			const month = months[oDate.getMonth()];
			const year = oDate.getFullYear();

			return `${day} ${month} ${year}`;
		},

		formatTime: function (date) {
			const oDate = new Date(date);

			const hours = oDate.getHours().toString().padStart(2, "0");
			const minutes = oDate.getMinutes().toString().padStart(2, "0");

			return `${hours}:${minutes}`;
		},

		formatDateTime: function (dateTime) {
			const formattedDate = this.formatter.formatDate(dateTime);
			const formattedTime = this.formatter.formatTime(dateTime);

			return `${formattedDate}, ${formattedTime}`;
		},

		formatOptionalDate: function (optionalDate) {
			if (!optionalDate) {
				return this.getBundleText("nullField")
			}
			return this.formatter.formatDate(optionalDate);
		},

		formatOptionalField: function (field) {
			if (!field) {
				return this.getBundleText("nullField");
			}
			return field;
		},

		isReservationVisible: function (userHasReserved) {

			return !userHasReserved;
		}
	};
});
