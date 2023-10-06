<script setup lang="ts">
const props = withDefaults(defineProps<{
  variant?: 'Error' | 'Warning' | 'Success' | 'Info' | 'Default'
  disabled?: boolean
  type?: 'submit' | 'button'
  prevent?: boolean
}>(), {
  variant: 'Default',
  disabled: false,
  type: 'button',
  prevent: true,
})

const emit = defineEmits<{
  (e: 'click'): void
}>()

function handleClick(e: Event) {
  if (props.prevent)
    e.preventDefault()

  emit('click')
}

const isDefault = computed(() => props.variant === 'Default')
const isError = computed(() => props.variant === 'Error')
const isWarning = computed(() => props.variant === 'Warning')
const isSuccess = computed(() => props.variant === 'Success')
const isInfo = computed(() => props.variant === 'Info')
</script>

<template>
  <button
    class="block rounded-md px-3 py-2 text-white transition"
    :class="{
      'bg-indigo-2/15 hover:bg-indigo-2/35': isDefault,
      'bg-red-5 hover:bg-red-7': isError,
      'bg-yellow-5 hover:bg-yellow-7': isWarning,
      'bg-green-6 hover:bg-green-7': isSuccess,
      'bg-sky-6 hover:bg-sky-7': isInfo,
      'cursor-not-allowed text-gray-5!': disabled,
    }"
    :disabled="disabled"
    :type="type"
    @click="handleClick"
  >
    <slot />
  </button>
</template>
