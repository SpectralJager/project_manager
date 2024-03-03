/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./templates/*.html"],
  theme: {
    extend: {},
    colors: {
      'white': '#ffffff',
      'black': '#000000',
      'pm-1': '#EE7EBF',
      'pm-2': '#36396F',
      'pm-3': '#D7DDE8',
      'pm-4': '#F0DAD0',
      'pm-5': '#323467',
    },
    fontFamily: {
      main: ['Roboto', 'sans-serif'],
      title: ['Oswald', 'sans-serif'],
    },
  },
  plugins: [
    require('@tailwindcss/forms'),
    require('@tailwindcss/typography'),
  ],
}

