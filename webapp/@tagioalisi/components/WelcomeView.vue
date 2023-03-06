<script setup lang="ts">
import ServiceStatusCard from './ServiceStatusCard.vue';
</script>

<template>
  <VRow class="justify-center">
    <VCol cols="8" md="6" lg="5">
      <VCard avatar="/icon.png">
        <VCardTitle>
          <VAvatar image="/icon.png" />
          Hello! My name is Tagioalisi!
        </VCardTitle>
        <VCardText>
          <p> I am a Discord bot. This is my web management interface. Let's have some fun! </p>
        </VCardText>
        <VCardActions>
          <VSpacer />
          <VBtn color="primary">Log In</VBtn>
        </VCardActions>
      </VCard>
    </VCol>
    <VCol cols="8" md="6" lg="5">
      <ServiceStatusCard />
    </VCol>
  </VRow>
</template>
