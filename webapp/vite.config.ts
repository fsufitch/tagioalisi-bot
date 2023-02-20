import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react';

// https://vitejs.dev/config/
export default defineConfig({
  resolve: {
    alias: {
      '@tagioalisi': '/@tagioalisi',
    },
    extensions: ['.mjs', '.js', '.ts', '.jsx', '.tsx', '.json', '.css'],
  },
  define: {
    __BOT_BASE_URL__: JSON.stringify(process.env.BOT_ENDPOINT || ''),
    __BOT_GRPC_BASE_URL__: JSON.stringify(process.env.BOT_GRPC_ENDPOINT || ''),
  },
  server: {
    host: '0.0.0.0',
    port: 8080,
    hmr: true,
    https: {
      key: './keys/default.key',
      cert: './keys/default.pem',
    },
  },

  plugins: [react()],
});
