const { tailwindTransform } = require("postcss-lit");

module.exports = {
  content: {
    files: ["./dist/*.js"],
    transform: {
      ts: tailwindTransform,
    },
  },
  theme: {
    extend: {},
  },
  plugins: [require("daisyui")],
  daisyui: {
    themes: ["dark", "garden"],
  },
};
