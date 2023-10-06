import Toast from 'vue-toastification'

import 'vue-toastification/dist/index.css' // if needed

export default defineNuxtPlugin((nuxtApp) => {
  nuxtApp.vueApp.use(Toast, {
    position: 'bottom-right',
    hideProgressBar: true,
    timeout: 3000,
  })
})
