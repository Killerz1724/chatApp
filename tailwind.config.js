/** @type {import('tailwindcss').Config} */
import tailwindScrollbar from "tailwind-scrollbar";
module.exports = {
  content: [
    "./src/pages/**/*.{js,ts,jsx,tsx}",
    "./src/components/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {},
    fontFamily: {
      jakartaSans: ["Plus Jakarta Sans", "sans-serif"],
    },
  },
  plugins: [tailwindScrollbar],
};
