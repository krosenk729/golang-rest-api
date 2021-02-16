import { createStore, createLogger } from 'vuex';
import entries from './entries';
import userDate from './user-date';

const debug = process.env.NODE_ENV !== 'production';

export default createStore({
  modules: {
    entries,
    userDate,
  },
  strict: debug,
  plugins: debug ? [createLogger()] : [],
});
