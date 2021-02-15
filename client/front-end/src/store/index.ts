import { createStore, createLogger } from 'vuex';
import userDate from './user-date';

const debug = process.env.NODE_ENV !== 'production';

export default createStore({
  modules: { userDate },
  strict: debug,
  plugins: debug ? [createLogger()] : [],
});
