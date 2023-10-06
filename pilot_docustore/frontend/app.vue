<script setup lang="ts">
import { useAppData } from '#imports'

const { appName } = useAppData()

useHead({
  title: appName,
})
</script>

<template>
  <NuxtLayout>
    <NuxtPage />
  </NuxtLayout>
</template>
