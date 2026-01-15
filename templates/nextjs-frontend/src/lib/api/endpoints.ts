export const API_ENDPOINTS = {
  // Auth endpoints
  AUTH: {
    LOGIN: "/auth/login",
    REGISTER: "/auth/register",
    LOGOUT: "/auth/logout",
    REFRESH: "/auth/refresh",
    ME: "/auth/me",
  },

  // Item endpoints
  ITEM: {
    LIST: "/items",
    CREATE: "/items",
    GET: (id: string | number) => `/items/${id}`,
    UPDATE: (id: string | number) => `/items/${id}`,
    DELETE: (id: string | number) => `/items/${id}`,
  },

  // Upload endpoints
  UPLOAD: {
    IMAGE: "/upload/image",
    FILE: "/upload/file",
    PRESIGNED_URL: "/upload/presigned-url",
  },
} as const;