/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./templates/*.html"],
  theme: {
    extend: {},
    colors: {
      'pm-1': '#E63946',
      'pm-2': '#EC9A9A',
      'pm-3': '#F1FAEE',
      'pm-4': '#CDEAE5',
      'pm-5': '#A8DADC',
      'pm-6': '#77ABBD',
      'pm-7': '#457B9D',
      'pm-8': '#31587A',
      'pm-9': '#1D3557'
    },
    fontFamily: {
      sans: ['Roboto', 'sans-serif'],
      serif: ['Merriweather', 'serif'],
    },
  },
  plugins: [
    require('@tailwindcss/forms'),
    require('@tailwindcss/typography'),
  ],
}

