export type EventType = "music" | "conference";

export interface Event {
  uuid: string;
  location: string;
  type: EventType;
  date_start: string;
  date_end: string;
  description: string;
  title: string;
  tags: string[];
  organizer_id: string;
  number_of_tickets: number;
}
