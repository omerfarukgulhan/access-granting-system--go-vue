export default {
  isAuthenticated(state) {
    return !!state.token;
  },
  getToken(state) {
    return state.token;
  },
  getPrefix(state) {
    return state.prefix;
  },
  getUser(state) {
    return state.user;
  },
  isAdmin(state) {
    return state.user && state.user.roles.some(role => role.name === 'Admin');
  },
};