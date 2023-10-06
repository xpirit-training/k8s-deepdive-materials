<script setup lang="ts">
import { useToast } from 'vue-toastification'
import { useDocustoreService } from '#imports'

const toast = useToast()

const { listDocuments, createDocument } = useDocustoreService()
const documentIds = await listDocuments()

async function handleAddDocumentFormSubmit(newDocumentId: string) {
  const response = await createDocument(newDocumentId)
  if (response.isSuccessful) {
    toast.success('Document created')
    await navigateTo(`/documents/${newDocumentId}`)
  }
  else {
    toast.error(`Couldn't create document with id '${newDocumentId}'`)
  }
}
</script>

<template>
  <SectionHeading>
    <template #headline>
      All the documents
    </template>
    <template #description>
      Here is a list of all document IDs. Click on one to view it's data or create a new one.
    </template>
  </SectionHeading>

  <AddDocumentForm @submitted="handleAddDocumentFormSubmit" />

  <Suspense>
    <DocumentList v-if="documentIds.data?.value && documentIds.data.value.length > 0" :document-ids="documentIds.data.value" />
    <template #fallback>
      loading
    </template>
  </Suspense>
</template>
