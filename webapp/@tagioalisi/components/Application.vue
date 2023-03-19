<script setup lang="ts">
import WelcomeView from './WelcomeView.vue';
import { getGrpcService } from '@tagioalisi/services/grpc';

// Initialize the grpc connection
const grpc = getGrpcService();
grpc.reset();
</script>

<template>
  <VApp theme="discordDark">
    <VMain>
      <VContainer class="fill-height align-center justify-center">
        <WelcomeView />
      </VContainer>
    </VMain>
  </VApp>
</template>
