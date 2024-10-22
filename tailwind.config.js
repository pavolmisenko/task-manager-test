/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./src/web/templates/**/*.html"],
  theme: {
    extend: {},
  },
  plugins: [require("daisyui")],
  daisyui: {
    themes: ["dark", "garden"],
  },
};
