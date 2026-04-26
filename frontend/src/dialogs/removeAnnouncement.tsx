import { Button } from "@/components/ui/button";
import { Dialog, DialogContent, DialogDescription, DialogHeader, DialogTitle } from "@/components/ui/dialog";
import { Trash2 } from "lucide-react";

interface Props {
  title: string;
  open: boolean;
  onOpenChange: (open: boolean) => void;
  onConfirm: () => void;
}

export function RemoveAnnouncementDialog({ title, open, onOpenChange, onConfirm }: Props) {
  return (
    <Dialog open={open} onOpenChange={onOpenChange}>
      <DialogContent className="w-[400px] flex flex-col items-center gap-6 py-8">
        <DialogHeader className="flex flex-col items-center gap-3 text-center">
          <div className="text-red-500 bg-red-100 rounded-full p-5">
            <Trash2 className="w-8 h-8" />
          </div>
          <DialogTitle>ยืนยันการลบประกาศ</DialogTitle>
          <DialogDescription>
            คุณต้องการลบประกาศ{" "}
            <span className="font-medium text-gray-900">"{title}"</span> ใช่ไหม?
          </DialogDescription>
        </DialogHeader>

        <div className="flex gap-3 w-full">
          <Button variant="outline" className="flex-1 h-10" onClick={() => onOpenChange(false)}>
            ยกเลิก
          </Button>
          <Button variant="destructive" className="flex-1 h-10" onClick={onConfirm}>
            ยืนยัน
          </Button>
        </div>
      </DialogContent>
    </Dialog>
  );
}
