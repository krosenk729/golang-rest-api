/* eslint-disable arrow-body-style, no-shadow */

type Direction = 1 | -1;

interface UserDateState {
  date: Date;
}

const state = (): UserDateState => ({
  date: new Date(),
});

// getters
const getters = {
  formattedDate: (state: UserDateState) => {
    return state.date.toDateString();
  },
};

// actions
const actions = {
  increaseWeek({ commit }: { commit: Function }) {
    commit('changeWeek', 1);
  },
  decreaseWeek({ commit }: { commit: Function }) {
    commit('changeWeek', -1);
  },
};

// mutations
const mutations = {
  changeWeek(state: UserDateState, dir: Direction) {
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
