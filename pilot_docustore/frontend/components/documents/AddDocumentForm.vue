<script setup lang="ts">
const emit = defineEmits<{
  (e: 'submitted', id: string): any
}>()

const newDocumentId = ref<string | null>(null)
const newDocumentButtonDisabled = computed(() => !newDocumentId.value)

function handleNewDocumentFormSubmit() {
  if (!newDocumentId.value)
    return

  emit('submitted', newDocumentId.value)
}
</script>

<template>
  <form class="mt-10 flex items-center gap-4" @submit.prevent="handleNewDocumentFormSubmit">
    <label for="new-document">Add a new document</label>
    <Input id="new-document" v-model="newDocumentId" type="text" name="new-document" />
    <Button type="submit" :prevent="false" :disabled="newDocumentButtonDisabled">
      Add
    </Button>
  </form>
</template>
