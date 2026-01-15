import { apiClient } from "@/lib/api/client";
import { API_ENDPOINTS } from "@/lib/api/endpoints";
import type {
  Item,
  CreateItemRequest,
  UpdateItemRequest,
  ItemListResponse,
} from "../types";

export const itemService = {
  // Get all items
  getAll: async (page = 1, pageSize = 10): Promise<ItemListResponse> => {
    const { data } = await apiClient.get(API_ENDPOINTS.ITEM.LIST, {
      params: { page, pageSize },
    });
    return data;
  },

  // Get a single item
  getById: async (id: string): Promise<Item> => {
    const { data } = await apiClient.get(API_ENDPOINTS.ITEM.GET(id));
    return data;
  },

  // Create a new item
  create: async (payload: CreateItemRequest): Promise<Item> => {
    const { data } = await apiClient.post(API_ENDPOINTS.ITEM.CREATE, payload);
    return data;
  },

  // Update an existing item
  update: async (id: string, payload: UpdateItemRequest): Promise<Item> => {
    const { data } = await apiClient.put(API_ENDPOINTS.ITEM.UPDATE(id), payload);
    return data;
  },

  // Delete an item
  delete: async (id: string): Promise<void> => {
    await apiClient.delete(API_ENDPOINTS.ITEM.DELETE(id));
  },
};