const appDescription = process.env.NUXT_PUBLIC_APP_DESCRIPTION

export default defineNuxtConfig({
  modules: [
    '@vueuse/nuxt',
    '@unocss/nuxt',
    '@nuxtjs/color-mode',
  ],

  publicRuntimeConfig: { baseURL: process.env.BASE_URL || 'https://localhost:8080' },

  runtimeConfig: {
    // from .env file (only in dev)
    apiSecret: '',
    public: {
      appName: '',
      appDescription: '',
      apiBase: process.env.NUXT_PUBLIC_API_BASE,
    },
  },

  components: [
    { path: '~/components/documents', prefix: '' },
    { path: '~/components/form', prefix: '' },
    { path: '~/components/interaction', prefix: '' },
    { path: '~/components/layout', prefix: '' },
    { path: '~/components/podinfo', prefix: '' },
    '~/components',
  ],

  experimental: {
    // when using generate, payload js assets included in sw precache manifest
    // but missing on offline, disabling extraction it until fixed
    payloadExtraction: false,
    inlineSSRStyles: false,
    renderJsonPayloads: true,
    typedPages: true,
  },

  css: [
    '@unocss/reset/tailwind.css',
    '~/assets/css/main.css',
  ],

  colorMode: {
    classSuffix: '',
  },

  nitro: {
    esbuild: {
      options: {
        target: 'esnext',
      },
    },
    prerender: false,
  },

  app: {
    head: {
      viewport: 'width=device-width,initial-scale=1',
      link: [
        { rel: 'icon', href: '/favicon.ico', sizes: 'any' },
        { rel: 'icon', type: 'image/svg+xml', href: '/nuxt.svg' },
        { rel: 'apple-touch-icon', href: '/apple-touch-icon.png' },
      ],
      meta: [
        { name: 'viewport', content: 'width=device-width, initial-scale=1' },
        { name: 'description', content: appDescription },
        { name: 'apple-mobile-web-app-status-bar-style', content: 'black-translucent' },
      ],
    },
  },

  devtools: {
    enabled: true,
  },
})
