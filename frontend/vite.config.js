import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

const backend =
  process.env.VITE_BACKEND_URL || "http://app-qa:8080"; // valor por defecto para desarrollo

export default defineConfig({
  plugins: [react()],
  server: {
    proxy: {
      "/api": {
        target: backend,
        changeOrigin: true,
        rewrite: (path) => path.replace(/^\/api/, ""),
      },
    },
    host: true,
  },
})