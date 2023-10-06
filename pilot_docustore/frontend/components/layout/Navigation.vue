<script lang="ts">
export interface PageLink {
  url: string
  name: string
}
</script>

<script setup lang="ts">
defineProps<{
  links: PageLink[]
}>()
</script>

<template>
  <div class="flex items-center justify-center gap-8">
    <template v-for="link in links" :key="link.name">
      <NuxtLink
        :to="link.url"
        class="block p-2 font-bold leading-6 hover:underline"
        exact-active-class="underline"
      >
        {{ link.name }}
      </NuxtLink>
    </template>
  </div>
</template>
