export interface Item {
  id: string;
  name: string;
  description?: string;
  createdAt: string;
  updatedAt: string;
}

export interface CreateItemRequest {
  name: string;
  description?: string;
}

export interface UpdateItemRequest {
  name?: string;
  description?: string;
}

export interface ItemListResponse {
  data: Item[];
  total: number;
  page: number;
  pageSize: number;
}