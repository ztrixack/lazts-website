/** @type {import('tailwindcss').Config} */
module.exports = {
  mode: 'jit',
  darkMode: "class",
  content: ["./templates/**/*.html", "./static/js/**/*.js"],
  theme: {
    extend: {
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
  ],
}
