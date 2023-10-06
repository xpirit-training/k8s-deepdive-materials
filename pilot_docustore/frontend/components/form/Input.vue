<script setup lang="ts">
withDefaults(defineProps<{
  modelValue: string | number | null
  name: string
  id: string
  type: string
  required?: boolean
  readonly?: boolean
  label?: string
}>(), {
  required: false,
})

defineEmits<{
  (e: 'update:modelValue', value: string | number): void
}>()
</script>

<template>
  <label
    v-if="label"
    :for="id"
    class="block text-sm font-medium leading-6"
  >{{ label }}</label>
  <input
    :id="id"
    :name="name"
    :type="type"
    :required="required"
    :readonly="readonly"
    :value="modelValue"
    class="block w-full rounded-md bg-white/5 text-white shadow-sm ring-1 ring-white/10 ring-inset sm:text-sm sm:leading-6 focus:ring-2 focus:ring-indigo-500 focus:ring-inset"
    :class="{ 'cursor-not-allowed text-gray-5!': readonly }"
    @input="$emit('update:modelValue', $event.target.value)"
  >
</template>
