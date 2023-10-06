<script setup lang="ts">
const props = defineProps<{
  variant: 'Error' | 'Warning' | 'Success' | 'Info'
}>()

defineSlots<{
  headline(): any
  description(): any
}>()

const isError = computed(() => props.variant === 'Error')
const isWarning = computed(() => props.variant === 'Warning')
const isSuccess = computed(() => props.variant === 'Success')
const isInfo = computed(() => props.variant === 'Info')
</script>

<template>
  <div
    class="rounded p-4"
    :class="{
      'bg-red-5': isError,
      'bg-yellow-5': isWarning,
      'bg-green-6': isSuccess,
      'bg-sky-5': isInfo,
    }"
  >
    <div class="flex gap-2">
      <div class="shrink-0">
        <div
          class="inline-block text-xl"
          :class="{
            'i-carbon:close-filled': isError,
            'i-carbon:warning-alt-filled': isWarning,
            'i-carbon:checkmark-filled': isSuccess,
            'i-carbon:information-filled': isInfo,
          }"
        />
      </div>
      <div>
        <div class="font-bold">
          <slot name="headline" />
        </div>
        <p class="mt-2">
          <slot name="description" />
        </p>
      </div>
    </div>
  </div>
</template>
