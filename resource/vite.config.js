import { fileURLToPath, URL } from 'node:url'
import { defineConfig,loadEnv } from 'vite'
import vue from '@vitejs/plugin-vue'


// https://vitejs.dev/config/
export default defineConfig(({ mode }) => {
  const env = loadEnv(mode, __dirname)
  return { 
    // base: env.VITE_APP_BASE_API,
    plugins: [
      vue(),
    ],
    server: {
      host: '0.0.0.0',
      port: env.VITE_PORT || 10070,
    },
    resolve: {
      alias: {
        '@': fileURLToPath(new URL('./src', import.meta.url))
      }
    }
}
});




