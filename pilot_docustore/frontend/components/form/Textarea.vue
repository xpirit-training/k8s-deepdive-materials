<script setup lang="ts">
withDefaults(defineProps<{
  modelValue: string | number | string[] | undefined
  name: string
  id: string
  required?: boolean
  label?: string
  rows?: number
}>(), {
  required: false,
  rows: 6,
})

defineEmits<{
  (e: 'update:modelValue', value: string | number | string[] | undefined): void
}>()
</script>

<template>
  <label
    v-if="label"
    :for="id"
    class="block text-sm font-medium leading-6"
  >{{ label }}</label>
  <textarea
    :id="id"
    :name="name"
    :required="required"
    :value="modelValue"
    :rows="rows"
    class="block w-full rounded-md bg-white/5 text-white shadow-sm ring-1 ring-white/10 ring-inset sm:text-sm sm:leading-6 focus:ring-2 focus:ring-indigo-500 focus:ring-inset"
    @input="$emit('update:modelValue', $event.target.value)"
  />
</template>
