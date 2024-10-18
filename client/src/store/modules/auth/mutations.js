export default {
  SET_TOKEN(state, token) {
    state.token = token;
  },
  SET_USER(state, user) {
    state.user = user;
  },
  SET_PREFIX(state, prefix) {
    state.prefix = prefix;
  },
  CLEAR_AUTH(state) {
    state.token = null;
    state.user = null;
    state.prefix = null;
  },
};