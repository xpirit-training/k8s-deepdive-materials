<script setup lang="ts">
import type { Document } from '#imports'

interface Props {
  document: Document
}
const props = defineProps<Props>()

const emit = defineEmits<{
  (e: 'updated', document: Document): void
  (e: 'deleted', id: string): void
}>()

const userDocument = ref(props.document)

function handleUpdateClicked() {
  emit('updated', userDocument.value)
}
function handleDeleteClicked() {
  emit('deleted', props.document.id)
}
</script>

<template>
  <form class="flex flex-col gap-4 text-left">
    <Input id="id" v-model="userDocument.id" type="text" name="id" label="Document ID" :readonly="true" />
    <Textarea id="document_data" v-model="userDocument.data" name="document_data" label="Document content" />
    <div class="mt-10 flex flex-row gap-8">
      <Button type="button" variant="Info" @click="handleUpdateClicked">
        Update
      </Button>
      <Button type="button" variant="Error" @click="handleDeleteClicked">
        Delete
      </Button>
    </div>
  </form>
</template>
