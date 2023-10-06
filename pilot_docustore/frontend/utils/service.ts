export interface ServiceResponse {
  isSuccessful: boolean
  error?: ErrorCodes
}

export interface ServiceDataResponse<T> extends ServiceResponse {
  data?: Ref<T | null>
}

export type ErrorCodes = 'NotFound' | 'BadRequest' | 'ApiError' | 'None'

export function statusToErrorCodes(statusCode: number | undefined): ErrorCodes {
  if (!statusCode)
    return 'None'

  switch (statusCode) {
    case 404:
      return 'NotFound'
    case 400:
      return 'BadRequest'
    default:
      return 'ApiError'
  }
}
