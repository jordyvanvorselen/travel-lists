/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["../view/**/*.templ"],
  theme: {
    extend: {},
  },
  plugins: [require("@tailwindcss/typography"), require("daisyui")],
  daisyui: {
    themes: ["night"],
    darkTheme: "night",
  }
}

