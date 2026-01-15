import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";
import { itemService } from "../services/api";
import type { CreateItemRequest, UpdateItemRequest } from "../types";
import { useToast } from "@/components/ui/use-toast";

const QUERY_KEY = "items";

// Hook to fetch all items
export const useItemList = (page = 1, pageSize = 10) => {
  return useQuery({
    queryKey: [QUERY_KEY, "list", page, pageSize],
    queryFn: () => itemService.getAll(page, pageSize),
  });
};

// Hook to fetch a single item
export const useItem = (id: string) => {
  return useQuery({
    queryKey: [QUERY_KEY, id],
    queryFn: () => itemService.getById(id),
    enabled: !!id,
  });
};

// Hook to create an item
export const useCreateItem = () => {
  const queryClient = useQueryClient();
  const { toast } = useToast();

  return useMutation({
    mutationFn: (data: CreateItemRequest) => itemService.create(data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: [QUERY_KEY] });
      toast({
        title: "Success",
        description: "Item created successfully",
      });
    },
    onError: (error: Error) => {
      toast({
        title: "Error",
        description: error.message || "Failed to create item",
        variant: "destructive",
      });
    },
  });
};

// Hook to update an item
export const useUpdateItem = () => {
  const queryClient = useQueryClient();
  const { toast } = useToast();

  return useMutation({
    mutationFn: ({ id, data }: { id: string; data: UpdateItemRequest }) =>
      itemService.update(id, data),
    onSuccess: (_, { id }) => {
      queryClient.invalidateQueries({ queryKey: [QUERY_KEY] });
      queryClient.invalidateQueries({ queryKey: [QUERY_KEY, id] });
      toast({
        title: "Success",
        description: "Item updated successfully",
      });
    },
    onError: (error: Error) => {
      toast({
        title: "Error",
        description: error.message || "Failed to update item",
        variant: "destructive",
      });
    },
  });
};

// Hook to delete an item
export const useDeleteItem = () => {
  const queryClient = useQueryClient();
  const { toast } = useToast();

  return useMutation({
    mutationFn: (id: string) => itemService.delete(id),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: [QUERY_KEY] });
      toast({
        title: "Success",
        description: "Item deleted successfully",
      });
    },
    onError: (error: Error) => {
      toast({
        title: "Error",
        description: error.message || "Failed to delete item",
        variant: "destructive",
      });
    },
  });
};