import { defineConfig } from 'vite'
import { svelte } from '@sveltejs/vite-plugin-svelte'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [svelte()],
  resolve: {
    alias: {
      'wailsjs': '/wailsjs'
    }
  },
  build: {
    rollupOptions: {
      external: [
        // '/wailsjs/go/models',
        // '/wailsjs/runtime'
      ]
    }
  }
})
