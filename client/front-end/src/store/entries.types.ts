import { Entry } from '@/shared/models';

export interface HashValue {
  tag?: string;
  verb?: string;
  count: number;
}
export interface CountHash {
  [key: string]: HashValue;
}

export interface EntryState {
  entries: Entry[];
}

export interface AddEntriesPayload {
  date: Date;
  newEntries: Entry[];
}

export interface ListReduceParams {
  tagsHash: CountHash;
  verbHash: CountHash;
}
