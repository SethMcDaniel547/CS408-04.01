/** @type {import('tailwindcss').Config} */
module.exports = {
  content: {
    files: ["./views/**/*.templ"],
  },
  theme: {
    extend: {
      backgroundColor: {
        'background': 'hsl(var(--background))',
      },
      textColor: {
        'foreground': 'hsl(var(--foreground))',
      },
      borderColor: {
        'border': 'hsl(var(--border))',
      },
    },
  },
  plugins: [],
}