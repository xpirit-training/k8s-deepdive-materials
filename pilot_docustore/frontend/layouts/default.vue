<script setup lang="ts">
import { useAppData } from '#imports'

const { appName, appDescription } = useAppData()

const links = [
  {
    url: '/documents',
    name: 'Documents',
  },
  {
    url: '/podinfo',
    name: 'Pod Info',
  },
]
</script>

<template>
  <header class="w-full">
    <nav class="grid grid-cols-[auto_1fr_auto] mb-8 items-center px-6 py-4 md:py-8">
      <NuxtLink to="/" class="flex items-center gap-2">
        <div class="i-carbon:rule-filled text-4xl color-indigo-5" />
        <div>
          <h1 class="text-2xl font-bold leading-6 tracking-wider font-serif">
            {{ appName }}
          </h1>
          <p class="italic color-slate-5">
            {{ appDescription }}
          </p>
        </div>
      </NuxtLink>
      <Navigation :links="links" />
      <div>
        <NuxtLink to="https://xpirit.com/" target="_blank">
          <div class="i-carbon:arrow-right" />
        </NuxtLink>
      </div>
    </nav>
  </header>
  <div class="mx-auto max-w-2xl px-6 font-sans md:px-0">
    <main class="p">
      <slot />
    </main>
  </div>
</template>
