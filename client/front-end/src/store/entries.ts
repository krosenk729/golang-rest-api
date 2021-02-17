/* eslint-disable no-param-reassign, no-shadow */
import api from '@/services/api';
import { Entry } from '@/shared/models';
import {
  AddEntriesPayload,
  EntryState,
  HashValue,
  ListReduceParams,
} from './entries.types';

const state = (): EntryState => ({
  entries: [],
});

// getters
const getters = {
  entryWeeks(state: EntryState) {
    const weeksHash = state.entries.reduce((hash, { date }) => {
      // TODO @katherine - aggreagte weeks
      hash[date as string] = date;
      return hash;
    }, {} as { [index: string]: string|Date|undefined });
    return Object.keys(weeksHash);
  },
  entryDays(state: EntryState) {
    const dateHash = state.entries.reduce((hash, { date }) => {
      hash[date as string] = date;
      return hash;
    }, {} as { [index: string]: string|Date|undefined });
    return Object.keys(dateHash);
  },
  lists(state: EntryState) {
    const {
      tagsHash,
      verbHash,
    } = state.entries.reduce(({ tagsHash, verbHash }: ListReduceParams, { tags, verb }) => {
      // eslint-disable-next-line no-unused-expressions
      tags?.forEach((t) => {
        tagsHash[t] = tagsHash[t] || { tag: t, count: 0 };
        tagsHash[t].count += 1;
      });
      verbHash[verb] = verbHash[verb] || { verb, count: 0 };
      verbHash[verb].count += 1;
      return { tagsHash, verbHash };
    }, { tagsHash: {}, verbHash: {} });

    return {
      tagsList: Object.values(tagsHash).sort((a: HashValue, b: HashValue) => b.count - a.count),
      verbList: Object.values(verbHash).sort((a: HashValue, b: HashValue) => b.count - a.count),
    };
  },
};

// actions
const actions = {
  async refreshData({ commit }: { commit: Function }) {
    const entries = await api.getAll();
    commit('updateEntries', entries);
  },
  clearData({ commit }: { commit: Function }) {
    commit('updateEntries', []);
  },
  async addEntries({ commit }: { commit: Function }, payload: AddEntriesPayload) {
    const [mm, dd, yyyy] = payload.date.toLocaleDateString().split('/');
    const entries = await api.createEntries(yyyy, mm, dd, payload.newEntries);
    commit('addEntries', entries);
  },
};

// mutations
const mutations = {
  updateEntries(state: EntryState, entries: Entry[]) {
    state.entries = entries;
  },
  addEntries(state: EntryState, entries: Entry[]) {
    state.entries = entries.concat(state.entries);
  },
};

export default {
  namespaced: false,
  state,
  getters,
  actions,
  mutations,
};
