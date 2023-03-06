import * as ngrpc from 'nice-grpc-web';
import { App, computed, ref, Ref, watch, InjectionKey, inject } from 'vue';
import { GreeterDefinition } from '@tagioalisi/proto/hello';

export interface GrpcService {
  reset(): void;
  reconnect(): void;
  saveSessionEndpoint(): void;
  clearSessionEndpoint(): void;

  $error: Ref<string>;
  $endpoint: Ref<string>;
  $channel: Ref<ngrpc.Channel | null>;
  $greeter: Ref<ngrpc.Client<GreeterDefinition> | null>;
}

const ENDPOINT_STORAGE = 'tagioalisi.grpc';
export const GRPC_SERVICE_KEY: InjectionKey<GrpcService> = Symbol('tagioalisi.GrpcService');

export const GrpcServicePlugin = (app: App) => {
  const $error = ref('');
  const $endpoint = ref('');
  const $channel = ref<ngrpc.Channel | null>(null);
  const $connectTrigger = ref(0);
  watch([$endpoint, $connectTrigger], ([endpoint]) => {
    if (!endpoint) {
      console.warn('tried to connect but endpoint was empty');
      $error.value = 'EMPTY_ENDPOINT';
      return;
    }
    $error.value = '';
    try {
      $channel.value = ngrpc.createChannel(endpoint);
    } catch (err) {
      console.error('failed channel creation', err);
      $error.value = 'CHANNEL_CREATE_FAILED';
    }
  });

  const $greeter = computed(() =>
    $channel.value ? ngrpc.createClient(GreeterDefinition, $channel.value) : null,
  );

  const reconnect = () => {
    $connectTrigger.value++;
  };

  const reset = () => {
    const sessionEndpoint = sessionStorage.getItem(ENDPOINT_STORAGE);
    if (sessionEndpoint) {
      $endpoint.value = sessionEndpoint;
      return;
    }
    getDefaultEndpoint().then((endpoint) => ($endpoint.value = endpoint));
  };

  const saveSessionEndpoint = () => sessionStorage.setItem(ENDPOINT_STORAGE, $endpoint.value);
  const clearSessionEndpoint = () => sessionStorage.removeItem(ENDPOINT_STORAGE);

  app.provide(GRPC_SERVICE_KEY, {
    reset,
    reconnect,
    saveSessionEndpoint,
    clearSessionEndpoint,
    $error,
    $endpoint,
    $channel,
    $greeter,
  });
};

const getDefaultEndpoint = async () => {
  const response = await fetch(window.location.href);
  const headerValue = response.headers.get('X-Tagioalisi-Grpc-Endpoint');
  if (headerValue && typeof headerValue === 'string') {
    return headerValue;
  }

  if (__GRPC_ENDPOINT__ && typeof __GRPC_ENDPOINT__ === 'string') {
    return __GRPC_ENDPOINT__;
  }

  return '';
};

export const getGrpcService = () => {
  const grpc = inject(GRPC_SERVICE_KEY);
  if (!grpc) {
    throw 'grpc service injection failed';
  }
  return grpc;
};
