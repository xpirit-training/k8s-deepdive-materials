<script setup lang="ts">
interface Props {
  to: string
}
defineProps<Props>()
</script>

<template>
  <NuxtLink :to="to" class="flex items-center gap-1">
    <span class="i-carbon:chevron-left relative top-1px inline-block" /><span>Back</span>
  </NuxtLink>
</template>
