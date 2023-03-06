import 'vuetify/styles';
import * as Vue from 'vue';
import Application from './components/Application.vue';

async function main() {
  const app = Vue.createApp(Application);

  const { createVuetify } = await import('./vuetify-loader');
  const vuetify = await createVuetify();
  app.use(vuetify);

  const { GrpcServicePlugin } = await import('./services/grpc');
  app.use(GrpcServicePlugin);

  app.mount('#app-container');
}

(() => main())();
