import { useFetch, useRuntimeConfig } from 'nuxt/app'

export interface Document {
  id: string
  data: string
}

interface CreateDocumentData extends Document {
}
interface UpdateDocumentData extends Document {
}

export function useDocustoreService() {
  const runtimeConfig = useRuntimeConfig()
  const baseUrl = runtimeConfig.public.apiBase
  const documentsUrl = new URL('api/v1/document/', baseUrl) // TODO: pluralize me

  async function listDocuments(): Promise<ServiceDataResponse<string[]>> {
    const url = new URL('api/v1/documents/', baseUrl)

    const { data, status, error } = await useFetch<string[]>(url.toString(), {
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

  async function getDocumentById(id: string): Promise<ServiceDataResponse<Document>> {
    const url = new URL(id, documentsUrl)

    const { data, status, error } = await useFetch<Document>(url.toString(), {
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

  async function createDocument(id: string, documentData?: string): Promise<ServiceDataResponse<Document>> {
    // const url = documentsUrl
    const url = new URL(id, documentsUrl)

    const { data, status, error } = await useFetch<Document>(url.toString(), {
      method: 'POST',
      body: {
        id,
        data: documentData,
      } as CreateDocumentData,
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

  async function updateDocument(doc: Document): Promise<ServiceDataResponse<Document>> {
    const url = new URL(doc.id, documentsUrl)

    const { data, status, error } = await useFetch<Document>(url.toString(), {
      method: 'POST',
      body: {
        id: doc.id,
        data: doc.data,
      } as UpdateDocumentData,
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

  async function deleteDocument(id: string): Promise<ServiceResponse> {
    const url = new URL(id, documentsUrl)

    const { status, error } = await useFetch(url.toString(), {
      method: 'DELETE',
    })

    if (status.value === 'error') {
      return {
        isSuccessful: false,
        error: statusToErrorCodes(error.value?.statusCode),
      }
    }

    return {
      isSuccessful: true,
    }
  }

  return {
    listDocuments,
    getDocumentById,
    createDocument,
    updateDocument,
    deleteDocument,
  }
}
