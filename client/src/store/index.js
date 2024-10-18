import {createStore} from 'vuex';
import usersModule from './modules/users/index.js'
import authModule from './modules/auth/index.js'

const store = createStore({
  modules: {
    users: usersModule,
    auth: authModule
  },
});

export default store;
