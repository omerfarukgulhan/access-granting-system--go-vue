import axios from 'axios';

const serverUrl = import.meta.env.VITE_SERVER_URL + "/auth";
const profileImageUrl = import.meta.env.VITE_PROFILE_IMAGE_URL;

export default {
  async login(context, payload) {
    try {
      const response = await axios.post(`${serverUrl}/login`, payload);
      if (response.data.success) {
        const {token, prefix, user} = response.data.data;
        user.profileImage = `${profileImageUrl}/${user.profileImage}`;
        context.commit('SET_TOKEN', token);
        context.commit('SET_PREFIX', prefix);
        context.commit('SET_USER', user);
      } else {
        console.error('Login failed:', response.data.message);
        throw new Error(response.data.message);
      }
    } catch (error) {
      console.error('Login error:', error);
      throw error;
    }
  },
  async register(context, payload) {
    try {
      const response = await axios.post(`${serverUrl}/register`, payload);
      if (response.data.success) {
        const prefix = "Bearer";
        context.commit('SET_PREFIX', prefix);
      } else {
        console.error('Registration failed:', response.data.message);
        throw new Error(response.data.message);
      }
    } catch (error) {
      console.error('Registration error:', error);
      throw error;
    }
  },
  logout(context) {
    context.commit('CLEAR_AUTH');
  },
};
