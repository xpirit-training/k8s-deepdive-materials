<script setup lang="ts">
import { useToast } from 'vue-toastification'
import { type Document, useDocustoreService } from '#imports'

const props = defineProps<{
  documentId: string
}>()

const toast = useToast()

const { getDocumentById, updateDocument, deleteDocument } = useDocustoreService()

const response = await getDocumentById(props.documentId)

async function handleUpdated(doc: Document) {
  const response = await updateDocument(doc)
  if (response.isSuccessful)
    toast.success('Document updated')
  else
    toast.error('An error occured')
}

async function handleDeleted(id: string) {
  const response = await deleteDocument(id)
  if (response.isSuccessful) {
    toast.success('Document removed')
    await navigateTo('/documents')
  }
  else {
    toast.error('An error occured')
  }
}
</script>

<template>
  <SectionHeading>
    <template #headline>
      Document: <span class="font-mono">{{ documentId }}</span>
    </template>
  </SectionHeading>

  <div class="mb-4">
    <Back to="/documents" />
  </div>

  <Suspense>
    <div v-if="response.isSuccessful && response.data?.value">
      <EditDocumentForm
        :document="response.data.value"
        @updated="handleUpdated"
        @deleted="handleDeleted"
      />
    </div>
    <div v-else>
      <StatusBox variant="Error">
        <template #headline>
          Not Found
        </template>
        <template #description>
          Couldn't find a document with ID '{{ documentId }}'
        </template>
      </StatusBox>
    </div>
    <template #fallback>
      <p>Loading</p>
    </template>
  </Suspense>
</template>
