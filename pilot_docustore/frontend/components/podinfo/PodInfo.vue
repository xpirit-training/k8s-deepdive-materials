<script setup lang="ts">
import { usePodInfoService } from '#imports'

const { getPodInfo } = usePodInfoService()

const podInfoResponse = await getPodInfo()
</script>

<template>
  <SectionHeading>
    <template #headline>
      Info about the current Pod
    </template>
  </SectionHeading>

  <Suspense>
    <div
      v-if="podInfoResponse.isSuccessful && podInfoResponse.data?.value"
      class="mt-8 border border-slate/16 rounded-md bg-slate/15 p-8 shadow shadow-black/30 shadow-xl"
    >
      <h2 class="text-2xl">
        Namespace: <span class="ml-4">{{ podInfoResponse.data?.value.namespace_name }}</span>
      </h2>
      <h3 class="text-xl">
        Pod name: <span class="ml-4">{{ podInfoResponse.data?.value.pod_name }}</span>
      </h3>
    </div>

    <template #fallback>
      <p>Loading</p>
    </template>
  </Suspense>
</template>
