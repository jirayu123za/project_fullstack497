/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      colors : {
        'E1' : '#344B59',
        'M1' : '#5888A6',
        'B1' : '#A7D5F2',
        'G1' : '#D9D9D9',
      },
      fontFamily: {
        'poppins': ["Poppins", 'sans-serif'],
      },
    },
  },
  plugins: [],
}

