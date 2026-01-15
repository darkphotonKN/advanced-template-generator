"use client";

import { useState } from "react";
import { useItemList } from "@/features/item/hooks/use-item";
import { ItemList } from "@/features/item/components/item-list";
import { ItemCreateDialog } from "@/features/item/components/item-create-dialog";
import { Button } from "@/components/ui/button";
import { PlusCircle } from "lucide-react";

export default function ItemsPage() {
  const [page, setPage] = useState(1);
  const [createOpen, setCreateOpen] = useState(false);
  const { data, isLoading, error } = useItemList(page, 10);

  if (error) {
    return (
      <div className="container mx-auto py-10">
        <div className="text-center text-red-500">
          Error loading items: {error.message}
        </div>
      </div>
    );
  }

  return (
    <div className="container mx-auto py-10">
      <div className="flex justify-between items-center mb-6">
        <h1 className="text-3xl font-bold">Items</h1>
        <Button onClick={() => setCreateOpen(true)}>
          <PlusCircle className="mr-2 h-4 w-4" />
          Add Item
        </Button>
      </div>

      <ItemList
        items={data?.data || []}
        isLoading={isLoading}
        page={page}
        totalPages={Math.ceil((data?.total || 0) / 10)}
        onPageChange={setPage}
      />

      <ItemCreateDialog
        open={createOpen}
        onOpenChange={setCreateOpen}
      />
    </div>
  );
}