/* eslint-disable arrow-body-style, no-shadow */
import * as Util from '../shared/utils';
import { Direction, UserDateState } from './user-date.types';

const state = (): UserDateState => ({
  date: new Date(),
});

// getters
const getters = {
  formattedDate: (state: UserDateState) => {
    return state.date.toDateString();
  },
  weekStart: (state: UserDateState) => {
    return Util.getDayOfWeek(state.date, 0);
  },
  weekEnd: (state: UserDateState) => {
    return Util.getDayOfWeek(state.date, 6);
  },
};

// actions
const actions = {
  increaseWeek({ commit }: { commit: Function }) {
    console.log('************');
    commit('changeWeek', 1);
  },
  decreaseWeek({ commit }: { commit: Function }) {
    commit('changeWeek', -1);
  },
};

// mutations
const mutations = {
  changeWeek(state: UserDateState, dir: Direction) {
    console.log('************', dir, state.date.getDate(), (dir * 7), state.date.getDate() + (dir * 7));
    const n = state.date.setDate(state.date.getDate() + (dir * 7));
    state.date = new Date(n);
  },
};

export default {
  namespaced: false,
  state,
  getters,
  actions,
  mutations,
};
