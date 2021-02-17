import { Entry } from '@/shared/models';

const debugMode = process.env.NODE_ENV !== 'production';
const host = debugMode ? 'http://localhost:5050' : '';
const options = {
  headers: {
    accept: 'application/json, text/plain, */*',
  },
  mode: 'cors',
  credentials: 'omit',
} as RequestInit;

export default {
  async getAll() {
    try {
      const res = await fetch(`${host}/api/`, options);
      return res.json();
    } catch (err) {
      if (debugMode) {
        console.error('<><><><><><><><><><><><><><><><><><><><><><><><>');
        console.error(err);
      }
      return [];
    }
  },
  async createEntries(
    yyyy: number|string,
    mm: number|string,
    dd: number|string,
    entryData: Entry[],
  ) {
    try {
      const res = await fetch(`${host}/api/${yyyy}/${mm}/${dd}/entries`, {
        ...options,
        method: 'POST',
        body: entryData as unknown as FormData,
      });
      return res.json();
    } catch (err) {
      if (debugMode) {
        console.error('<><><><><><><><><><><><><><><><><><><><><><><><>');
        console.error(err);
      }
      return {};
    }
  },
};
