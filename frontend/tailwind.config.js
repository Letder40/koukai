const colors = require('tailwindcss/colors')

/** @type {import('tailwindcss').Config} */
module.exports = {
  mode: 'jit',
  theme: {
    extend: {},
    colors: {
      dblue: "#7289da",
      dblack: {
        Light: "#424549",
        Medium: "#36393e",
        Strong: "#282b30",
        Stronger: "#1e2124",
      },
      white: colors.white,
    },
  },
  purge: [
     './src/**/*.{js,jsx,ts,tsx}',
  ],
  plugins: [],
}
