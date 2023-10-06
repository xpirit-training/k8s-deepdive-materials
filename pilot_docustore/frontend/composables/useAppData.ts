import { useRuntimeConfig } from "nuxt/app";

export function useAppData() {
  const runtimeConfig = useRuntimeConfig();

  return {
    appName: runtimeConfig.public.appName || "Docustore",
    appDescription: runtimeConfig.public.appDescription || "Document Management Software"
  }
}
