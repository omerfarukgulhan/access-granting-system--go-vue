import axios from 'axios';

const serverUrl = import.meta.env.VITE_SERVER_URL + "/roles";

export default {
  async fetchRoles(context) {
    try {
      const response = await axios.get(serverUrl,{
        headers: {
          Authorization: `${context.rootGetters['auth/getPrefix']} ${context.rootGetters['auth/getToken']}`,
        },
      });
      console.log(response);
      context.commit('SET_ROLES', response.data.data);
    } catch (error) {
      console.error(error);
    }
  },
  async fetchRoleById(context, payload) {
    try {
      const response = await axios.get(`${serverUrl}/${payload.id}`);
      context.commit('SET_CURRENT_role', response.data.data);
    } catch (error) {
      console.error(error);
    }
  }
};