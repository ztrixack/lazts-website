/** @type {import('tailwindcss').Config} */

module.exports = {
  mode: 'jit',
  darkMode: "class",
  content: ["./templates/**/*.html", "./static/js/**/*.js"],
  theme: {
    extend: {
      aspectRatio: {
        '4/3': '4 / 3',
        '3/4': '3 / 4',
        '9/16': '9 / 16',
      },
      animation: {
        blackhole: 'blackhole 9s linear infinite',
      },
      keyframes: {
        blackhole: {
          '0%': { transform: 'rotate(0turn)' },
          '100%': { transform: 'rotate(1turn)' },
        },
      },
      fontFamily: {
        'kanit': ['Kanit', 'sans-serif'],
      }
    },
  },
  plugins: [
    require('@tailwindcss/typography'),
  ],
}
