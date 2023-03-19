import { defineConfig } from 'vite';
// import react from '@vitejs/plugin-react';
import VuetifyPlugin from 'vite-plugin-vuetify';
import VuePlugin from '@vitejs/plugin-vue';

// https://vitejs.dev/config/
export default defineConfig({
  resolve: {
    alias: {
      '@tagioalisi': '/@tagioalisi',
    },
    extensions: ['.mjs', '.js', '.ts', '.jsx', '.tsx', '.json', '.css'],
  },
  build: {
    target: 'esnext',
    manifest: true,
  },

  define: {
    __HTTP_ENDPOINT__: JSON.stringify(process.env.BOT_ENDPOINT || ''),
    __GRPC_ENDPOINT__: JSON.stringify(process.env.BOT_GRPC_ENDPOINT || ''),
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

  plugins: [VuePlugin(), VuetifyPlugin()],
});
