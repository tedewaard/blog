/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./src/templates/**/*.templ",
    "./src/templates/**/*.go",
  ],
  safelist: [
    // Technology tag colors - needed because they're generated dynamically
    'bg-blue-100',
    'text-blue-700',
    'bg-green-100',
    'text-green-700',
    'bg-purple-100',
    'text-purple-700',
    'bg-yellow-100',
    'text-yellow-700',
    'bg-red-100',
    'text-red-700',
    'bg-indigo-100',
    'text-indigo-700',
  ],
  theme: {
    extend: {},
  },
  plugins: [],
}
