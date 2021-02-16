import { Entry } from '@/shared/models';
import axios from 'axios';

const host = process.env.NODE_ENV === 'production'
  ? 'https://go--rest--api.herokuapp.com'
  : 'http://localhost:5050';
const headers = {
  'Access-Control-Allow-Origin': '*',
  'Referrer-Policy': 'origin',
};

export default {
  async getAll() {
    const res = await axios.get(`${host}/api/`, { headers });
    return res.data;
  },
  async createEntries(
    yyyy: number | string,
    mm: number | string,
    dd: number | string,
    entryData: Entry[],
  ) {
    const res = await axios.post(`${host}/api/${yyyy}/${mm}/${dd}/entries`, entryData);
    return res.data;
  },
};
