export interface Announcement {
    id: string;
    title: string;
    body: string;
    author: string;
    pinned: boolean;
    created_at: string
}

export interface CreateAnnouncement {
    title: string;
    body: string;
    author?: string;
    pinned: boolean;
}