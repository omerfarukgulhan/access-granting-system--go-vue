import axios from 'axios';

const serverUrl = import.meta.env.VITE_SERVER_URL + "/users";
const profileImageUrl = import.meta.env.VITE_PROFILE_IMAGE_URL

export default {
  async fetchUsers(context) {
    try {
      const response = await axios.get(serverUrl);
      const users = response.data.data.map(user => ({
        ...user,
        profileImage: `${profileImageUrl}/${user.profileImage}`
      }));
      context.commit('SET_USERS', users);
    } catch (error) {
      console.error(error);
    }
  },
  async fetchUserById(context, payload) {
    try {
      const response = await axios.get(`${serverUrl}/${payload.id}`);
      const user = {
        ...response.data.data,
        profileImage: `${profileImageUrl}/${ response.data.data.profileImage}`
      };
      context.commit('SET_CURRENT_USER', user);
    } catch (error) {
      console.error(error);
    }
  }
};