<script setup lang="ts">
const route = useRoute<'documents-id'>()
const documentId = route.params.id
</script>

<template>
  <DocumentDetails :document-id="documentId" />
</template>
