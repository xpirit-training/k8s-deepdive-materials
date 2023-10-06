<script setup lang="ts">
const emit = defineEmits<{
  (e: 'idSubmitted', id: string): void
}>()

const docId = ref<string | null>(null)

function handleFormSubmit() {
  if (!docId.value)
    throw new Error('id provided is null')

  emit('idSubmitted', docId.value!)
}
</script>

<template>
  <h2 class="mb-4 text-2xl">
    Load or Create Document
  </h2>
  <form class="flex flex-col gap-4 text-left" @submit.prevent="handleFormSubmit">
    <div>
      <Input id="id" v-model="docId" type="text" name="id" label="Document ID" />
    </div>
    <div>
      <Button type="submit" :prevent="false">
        Load
      </Button>
    </div>
  </form>
</template>
