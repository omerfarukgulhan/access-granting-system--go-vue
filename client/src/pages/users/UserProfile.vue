<template>
  <div class="user-profile container">
    <h1 class="title has-text-centered m-4">User Profile</h1>
    <UserInfoUpdate :user="user" @user-updated="updateUserInfo"/>
    <UserPasswordUpdate @password-updated="updatePassword"/>
    <div class="m-4 has-text-centered">
      <button class="button is-danger" @click="deleteAccount" :class="{ 'is-loading': isDeletingAccount }"
              :disabled="isDeletingAccount">
        Delete Account
      </button>
      <p v-if="deleteError" class="has-text-danger">{{ deleteError }}</p>
    </div>
  </div>
</template>

<script>
import UserInfoUpdate from '../../components/user/UserInfoUpdate.vue';
import UserPasswordUpdate from '../../components/user/UserPasswordUpdate.vue';
import axios from 'axios';

const serverUrl = import.meta.env.VITE_SERVER_URL + "/users";
const profileImageUrl = import.meta.env.VITE_PROFILE_IMAGE_URL;

export default {
  components: {
    UserInfoUpdate,
    UserPasswordUpdate
  },
  data() {
    return {
      user: {
        id: this.$store.getters['auth/getUser'].id,
        username: this.$store.getters['auth/getUser'].username,
      },
      isDeletingAccount: false,
      deleteError: null,
    };
  },
  methods: {
    async updateUserInfo(formData) {
      try {
        const response = await axios.put(`${serverUrl}/${this.user.id}`, formData, {
          headers: {
            Authorization: `${this.$store.getters['auth/getPrefix']} ${this.$store.getters['auth/getToken']}`,
            'Content-Type': 'multipart/form-data',
          },
        });
        response.data.data.profileImage = `${profileImageUrl}/${response.data.data.profileImage}`;
        this.$store.commit('auth/SET_USER', response.data.data);
      } catch (error) {
        console.error(error);
      }
    },
    async updatePassword(passwordData) {
      try {
        await axios.put(`${serverUrl}/update-password`, passwordData, {
          headers: {
            Authorization: `${this.$store.getters['auth/getPrefix']} ${this.$store.getters['auth/getToken']}`,
          },
        });
      } catch (error) {
        console.error(error);
      }
    },
    async deleteAccount() {
      this.isDeletingAccount = true;
      this.deleteError = null;
      try {
        await axios.delete(`${serverUrl}/${this.user.id}`, {
          headers: {
            Authorization: `${this.$store.getters['auth/getPrefix']} ${this.$store.getters['auth/getToken']}`,
          },
        });
        this.$store.dispatch('auth/logout');
        this.$router.push('/');
      } catch (error) {
        this.deleteError = error.response?.data?.message || 'Failed to delete account';
      } finally {
        this.isDeletingAccount = false;
      }
    },
  },
};
</script>

<style scoped>
.user-profile {
  max-width: 600px;
  margin: auto;
}
</style>
