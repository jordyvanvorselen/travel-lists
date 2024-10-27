/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["../templates/**/*.templ"],
  theme: {
    extend: {},
  },
  plugins: [require("@tailwindcss/typography"), require("daisyui")],
  daisyui: {
    themes: ["night"],
    darkTheme: "night",
  }
}

