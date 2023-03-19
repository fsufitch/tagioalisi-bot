<script setup lang="ts">
import { ref, watch, computed } from 'vue';
import { getGrpcService } from '@tagioalisi/services/grpc';

const grpc = getGrpcService();

const $loading = ref(false);
const $statusError = ref<string>('');
const $queryTimeMs = ref<number>(0);
const statusCheckTrigger = ref<number>(0);
const checkStatus = () => statusCheckTrigger.value++;

watch([grpc.$error, grpc.$greeter, statusCheckTrigger], async ([error, greeter]) => {
  if (error) {
    $statusError.value = error;
    return;
  }
  if (!greeter) {
    $statusError.value = 'GREETER_MISSING';
    return;
  }
  $loading.value = true;
  const timeStartMs = new Date().getTime();
  try {
    const reply = await greeter.sayHello({ name: 'service-status-check' });
    if (!reply.message.includes('service-status-check')) {
      throw `unexpected reply message: ${reply.message}`;
    }
    $statusError.value = '';
  } catch (err) {
    $statusError.value = 'QUERY_ERROR';
    console.error('query error', err);
  } finally {
    const timeEndMs = new Date().getTime();
    $queryTimeMs.value = timeEndMs - timeStartMs;
    $loading.value = false;
  }
});

checkStatus();
</script>

<template>
  <VCard class="pa-3">
    <VCardTitle>
      <h3>
        Service Status
        <VProgressCircular v-if="$loading" indeterminate />
      </h3>
    </VCardTitle>
    <VCardItem>
      <VList>
        <VListItem>
          <template #prepend>
            <VIcon icon="mdi-link-variant"></VIcon>
          </template>
          <VListItemTitle>
            <code v-if="grpc.$endpoint.value">{{ grpc.$endpoint.value }}</code>
            <VChip v-else color="error" variant="elevated">gRPC endpoint unset</VChip>
          </VListItemTitle>
          <VListItemSubtitle> gRPC endpoint </VListItemSubtitle>
        </VListItem>
        <VListItem>
          <template #prepend>
            <VIcon icon="mdi-timer-outline"></VIcon>
          </template>
          <VListItemTitle>
            <code>{{ $queryTimeMs }} ms</code>
          </VListItemTitle>
          <VListItemSubtitle> Query time </VListItemSubtitle>
        </VListItem>
        <VListItem v-if="$statusError" color="error">
          <template #prepend>
            <VIcon icon="mdi-alert"></VIcon>
          </template>
          <VListItemTitle>
            <VChip color="error"> {{ $statusError }}</VChip>
          </VListItemTitle>
        </VListItem>
        <VListItem v-else>
          <template #prepend>
            <VIcon icon="mdi-check-circle-outline"></VIcon>
          </template>
          <VListItemTitle>
            <VChip type="success" icon=""> Test query successful! </VChip>
          </VListItemTitle>
        </VListItem>
      </VList>
    </VCardItem>
    <VCardActions>
      <VSpacer />
      <VBtn color="primary" icon="mdi-reload" @click="checkStatus()"> </VBtn>
    </VCardActions>
  </VCard>
</template>
