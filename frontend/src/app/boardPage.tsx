import { announcementService } from "@/services/announcementService";
import type { Announcement, CreateAnnouncement } from "@/types/announcemnet";
import { Badge } from "@/components/ui/badge";
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Textarea } from "@/components/ui/textarea";
import { RemoveAnnouncementDialog } from "@/dialogs/removeAnnouncement";
import { Loader, Pin, PinOff, Trash2 } from "lucide-react";
import { useEffect, useState } from "react";
import { toast } from "sonner";



function AnnouncementCard({
  announcement: a,
  onDelete,
  onTogglePin,
}: {
  announcement: Announcement;
  onDelete: (id: string) => void;
  onTogglePin: (announcement: Announcement) => void;
}) {
  return (
    <Card className={`mb-3 ${a.pinned ? "bg-amber-50" : ""}`}>
      <CardContent className="pt-4">
        <div className="flex items-start justify-between gap-2 mb-2">
          <div className="flex flex-row items-center gap-2">
            {a.pinned && (
              <Badge variant="secondary" className="bg-amber-100 text-amber-700">
                Pinned
              </Badge>
            )}
            <h2 className="text-base font-medium text-gray-900">{a.title}</h2>
          </div>
          <button
            onClick={() => onTogglePin(a)}
            className={`p-1.5 rounded-md transition-colors ${
              a.pinned
                ? "text-amber-500 hover:bg-amber-100"
                : "text-gray-400 hover:bg-gray-100 hover:text-gray-600"
            }`}
          >
            {a.pinned ? <PinOff size={16} /> : <Pin size={16} />}
          </button>
        </div>
        <p className="text-sm text-gray-600 leading-relaxed mb-3">{a.body}</p>
        <div className="flex items-center justify-between">
          <span className="text-xs text-gray-400">
            โดย {a.author} • {new Date(a.created_at).toLocaleString("en-GB", {
              day: "2-digit", month: "2-digit", year: "numeric",
              hour: "2-digit", minute: "2-digit",
              timeZone: "Asia/Bangkok", hour12: false,
            }).replace(",", "")}
          </span>
          <Button
            variant="outline"
            size="sm"
            className="text-red-500 border-red-300 hover:bg-red-50 hover:text-red-600"
            onClick={() => onDelete(a.id)}
          >
            <Trash2 size={14} className="mr-1" />
            ลบ
          </Button>
        </div>
      </CardContent>
    </Card>
  );
}

function CreateForm({ onSubmit }: { onSubmit: (data: CreateAnnouncement) => Promise<void> }) {
  const [title, setTitle] = useState("");
  const [body, setBody] = useState("");
  const [author, setAuthor] = useState("");
  const [submitting, setSubmitting] = useState(false);

  const handleSubmit = async () => {

    // check empty field
    if (!title.trim() || !body.trim() || !author.trim()) {
      toast.error("กรุณากรอกข้อมูลให้ครบทุกช่อง");
      return;
    }

    setSubmitting(true);
    try {
      await onSubmit({ title, body, author, pinned: false });
      setTitle(""); setBody(""); setAuthor("");
    } finally {
      setSubmitting(false);
    }
  };

  return (
    <Card className="mb-6">
      <CardHeader>
        <CardTitle className="text-base">สร้างประกาศใหม่</CardTitle>
      </CardHeader>
      <CardContent className="flex flex-col gap-4">
          <div className="flex flex-col gap-1">
            <Label>หัวข้อ</Label>
            <Input placeholder="หัวข้อประกาศ" value={title} onChange={e => setTitle(e.target.value)}/>
          </div>
          <div className="flex flex-col gap-1">
            <Label>เนื้อหา</Label>
            <Textarea placeholder="รายละเอียดประกาศ" value={body} onChange={e => setBody(e.target.value)} className="min-h-[80px] resize-none"/>
          </div>
          <div className="flex flex-col gap-1">
            <Label>ผู้เขียน</Label>
            <Input placeholder="ชื่อผู้เขียน" value={author} onChange={e => setAuthor(e.target.value)}/>
          </div>
          <Button 
            onClick={() => void 
                handleSubmit()
            } 
            disabled={submitting} 
            className="w-25 h-10">
            สร้างประกาศ
          </Button>
      </CardContent>
    </Card>
  );
}

export function BoardPage() {
  const [announcements, setAnnouncements] = useState<Announcement[]>([]);
  const [loading, setLoading] = useState(true);
  const [deleteTarget, setDeleteTarget] = useState<Announcement | null>(null);

  const fetchAnnouncements = async () => {
    try {
      const data = await announcementService.getAll();
      setAnnouncements(data);
    } catch {
      toast.error("ไม่สามารถโหลดข้อมูลได้ กรุณาลองใหม่");
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    const load = async () => {
      try {
        const data = await announcementService.getAll();
        setAnnouncements(data);
      } catch {
        toast.error("ไม่สามารถโหลดข้อมูลได้ กรุณาลองใหม่");
      } finally {
        setLoading(false);
      }
    };
    void load();
  }, []);

  const handleCreate = async (data: CreateAnnouncement) => {
    try {
      await announcementService.createNewAnnouncement(data);
      await fetchAnnouncements();
      toast.success("สร้างประกาศใหม่สำเร็จ");
    } catch (err) {
      toast.error("สร้างไม่สำเร็จ กรุณาลองใหม่");
      throw err;
    }
  };

  const handleTogglePin = async (pinAnnoun: Announcement) => {
    try {
      await announcementService.togglePinAnnouncement(pinAnnoun);
      await fetchAnnouncements();
    } catch {
      toast.error("ไม่สามารถ pin ได้ กรุณาลองใหม่");
    }
  };

  const confirmDelete = async () => {
    if (!deleteTarget) return;
    try {
      await announcementService.removeAnnouncement(deleteTarget.id);
      await fetchAnnouncements();
      toast.success("ลบประกาศสำเร็จ");
    } catch {
      toast.error("ลบไม่สำเร็จ กรุณาลองใหม่");
    } finally {
      setDeleteTarget(null);
    }
  };

  return (
    <div className="flex flex-row gap-6 items-start p-10">
      <div className="w-[700px] h-[700px] sticky top-10">
        <h1 className="text-2xl font-medium text-gray-900 mb-6">
          📢 Announcement Board
        </h1>
        <CreateForm onSubmit={handleCreate} />
      </div>

      <div className="flex-1">
        <p className="text-xs font-medium text-gray-400 uppercase tracking-wide mb-3">
          ประกาศทั้งหมด
        </p>
        {loading ? (
            <div className="flex flex-col justify-center items-center py-10">
                <Loader className="w-25 h-25 animate-spin text-gray-400" />
                <span className="text-base text-gray-400 mt-2">กำลังโหลด...</span>
            </div>
        ) : announcements.length === 0 ? (
          <p className="text-sm text-gray-400">ยังไม่มีประกาศ</p>
        ) : (
          announcements.map(annouce => (
            <AnnouncementCard
              key={annouce.id}
              announcement={annouce}
              onDelete={() => setDeleteTarget(annouce)}
              onTogglePin={handleTogglePin}
            />
          ))
        )}
      </div>

      <RemoveAnnouncementDialog
        title={deleteTarget?.title ?? ""}
        open={deleteTarget !== null}
        onOpenChange={(open) => { if (!open) setDeleteTarget(null); }}
        onConfirm={confirmDelete}
      />
    </div>
  );
}
