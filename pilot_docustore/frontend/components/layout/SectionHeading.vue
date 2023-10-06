<script setup lang="ts">
defineSlots<{
  headline(): any
  description(): any
}>()
</script>

<template>
  <div class="mb-4 border-b border-b-slate pb-4">
    <h3 class="text-xl font-bold leading-6">
      <slot name="headline" />
    </h3>
    <p v-if="$slots.description" class="mt-3 max-w-md">
      <slot name="description" />
    </p>
  </div>
</template>
