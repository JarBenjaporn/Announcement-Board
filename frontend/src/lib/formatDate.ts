interface FormatDateProps {
    date: string;
}
export function FormatDate({ date }: FormatDateProps) {
    return new Date(date).toLocaleDateString("en-GB", {
        day: "2-digit",
        month: "2-digit",
        year: "numeric",
        hour: "2-digit",
        minute: "2-digit",
        timeZone: "Asia/Bangkok",
        hour12: false,
    }).replace(",", "");
}