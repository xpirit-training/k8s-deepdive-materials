import { useFetch, useRuntimeConfig } from 'nuxt/app'

export interface PodInfo {
  'pod_name': string
  'namespace_name': string
}

export function usePodInfoService() {
  const runtimeConfig = useRuntimeConfig()
  const baseUrl = runtimeConfig.public.apiBase
  const podInfoUrl = new URL('api/v1/podinfo/', baseUrl)

  async function getPodInfo(): Promise<ServiceDataResponse<PodInfo>> {
    const url = podInfoUrl
    console.log('🚀 ~ file: usePodInfoService.ts:15 ~ getPodInfo ~ podInfoUrl:', podInfoUrl)

    const { data, status, error } = await useFetch<PodInfo>(url.toString(), {
      method: 'GET',
    })

    if (status.value === 'error') {
      return {
        isSuccessful: false,
        error: statusToErrorCodes(error.value?.statusCode),
      }
    }

    return {
      data,
      isSuccessful: true,
    }
  }

  return { getPodInfo }
}
