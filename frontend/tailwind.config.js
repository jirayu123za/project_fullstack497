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
        'R1' : '#E4C1F9',
        'R2' : '#FFD45E',
        'R3' : '#FFAFCC',
        'R4' : '#93B955',
      },
      fontFamily: {
        'poppins': ["Poppins", 'sans-serif'],
      },
    },
  },
  plugins: [],
}

