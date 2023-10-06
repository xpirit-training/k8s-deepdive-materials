<script setup lang="ts">
interface Props {
  documentIds: string[]
}
defineProps<Props>()
</script>

<template>
  <table class="my-8 min-w-full border-collapse text-left">
    <thead>
      <tr>
        <th class="sticky top-0 z-10 border-b border-slate-400/20 p-0 py-3 pr-2 text-xl font-semibold leading-6">
          ID
        </th>
        <th class="sticky top-0 z-10 border-b border-slate-400/20 p-0 py-3 pr-2 text-xl font-semibold leading-6">
          &nbsp;
        </th>
      </tr>
    </thead>
    <tbody class="align-baseline">
      <tr v-for="documentId in documentIds" :key="documentId">
        <td class="whitespace-nowrap border-b py-3 pr-2 font-medium leading-6 dark:border-slate-400/10">
          <NuxtLink :to="`/documents/${documentId}`" class="underline underline-offset-4" :title="`Open document with id '${documentId}'`">
            {{ documentId }}
          </NuxtLink>
        </td>
        <td class="whitespace-nowrap border-b py-3 pl-2 font-medium leading-6 dark:border-slate-400/10">
             
        </td>
      </tr>
    </tbody>
  </table>
</template>
