const { defineConfig } = require("cypress");

module.exports = defineConfig({
  Integration: {
    baseUrl: "http://127.0.0.1:8080",
    supportFile: false,
  },

  e2e: {
    setupNodeEvents(on, config) {
      // implement node event listeners here
    },
  },
});
