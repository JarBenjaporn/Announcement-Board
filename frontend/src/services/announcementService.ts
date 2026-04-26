import type { Announcement, CreateAnnouncement } from "@/types/announcemnet";
import axios from "axios";


const API = import.meta.env.VITE_API_URL;

export const announcementService = {
  getAll: async (): Promise<Announcement[]> => {
    const res = await axios.get(`${API}/announcements`);
    return res.data;
  },

  createNewAnnouncement: async (data: CreateAnnouncement): Promise<Announcement> => {
    const res = await axios.post(`${API}/announcements`, data);
    return res.data;
  },

  togglePinAnnouncement: async (announ: Announcement): Promise<Announcement> => {
    const res = await axios.put(`${API}/announcements/${announ.id}`, {
      title: announ.title,
      body: announ.body,
      author: announ.author,
      pinned: !announ.pinned,
    });
    return res.data;
  },

  removeAnnouncement: async (id: string): Promise<void> => {
    await axios.delete(`${API}/announcements/${id}`);
  },
};