/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./internal/app/**/*.{go,css}",
    "./templates/**/*.{js,jsx,ts,tsx,css}",
    "./static/index.css",
  ],
  theme: {
    extend: {},
  },
  plugins: [],
};
