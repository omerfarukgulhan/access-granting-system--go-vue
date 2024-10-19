<template>
  <div class="user-profile container">
    <h1 class="title has-text-centered m-4">User Profile</h1>
    <div class="box has-background-white">
      <h2 class="subtitle has-text-grey-dark">Update Information</h2>
      <form @submit.prevent="updateUser">
        <div class="field">
          <label class="label has-text-black">Username</label>
          <div class="control">
            <input class="input" type="text" v-model="username" required :disabled="isUpdatingUser"/>
          </div>
        </div>
        <div class="field">
          <label class="label has-text-black">Profile Image</label>
          <div class="control">
            <input class="input" type="file" @change="handleFileUpload" :disabled="isUpdatingUser"/>
          </div>
        </div>
        <div class="field">
          <div class="control">
            <button type="submit" class="button is-primary is-fullwidth" :class="{ 'is-loading': isUpdatingUser }"
                    :disabled="isUpdatingUser">
              Update Info
            </button>
          </div>
        </div>
        <p v-if="userError" class="has-text-danger">{{ userError }}</p>
      </form>
    </div>
    <div class="box has-background-white">
      <h2 class="subtitle has-text-grey-dark">Change Password</h2>
      <form @submit.prevent="updatePassword">
        <div class="field">
          <label class="label has-text-black">Old Password</label>
          <div class="control">
            <input class="input" type="password" v-model="oldPassword" required :disabled="isUpdatingPassword"/>
          </div>
        </div>
        <div class="field">
          <label class="label has-text-black">New Password</label>
          <div class="control">
            <input class="input" type="password" v-model="newPassword" required :disabled="isUpdatingPassword"/>
          </div>
        </div>
        <div class="field">
          <div class="control">
            <button type="submit" class="button is-primary is-fullwidth" :class="{ 'is-loading': isUpdatingPassword }"
                    :disabled="isUpdatingPassword">
              Change Password
            </button>
          </div>
        </div>
        <p v-if="passwordError" class="has-text-danger">{{ passwordError }}</p>
      </form>
    </div>
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
import axios from 'axios';

const serverUrl = import.meta.env.VITE_SERVER_URL + "/users";
const profileImageUrl = import.meta.env.VITE_PROFILE_IMAGE_URL;

export default {
  data() {
    return {
      username: '',
      profileImage: '',
      oldPassword: '',
      newPassword: '',
      isUpdatingUser: false,
      isUpdatingPassword: false,
      isDeletingAccount: false,
      userError: null,
      passwordError: null,
      deleteError: null,
    };
  },
  mounted() {
    this.username = this.$store.getters['auth/getUser'].username;
  },
  methods: {
    handleFileUpload(event) {
      this.profileImage = event.target.files[0];
    },
    async updateUser() {
      this.isUpdatingUser = true;
      this.userError = null;
      const formData = new FormData();
      formData.append('username', this.username);
      if (this.profileImage) {
        formData.append('profileImage', this.profileImage);
      }

      try {
        const response = await axios.put(`${serverUrl}/${this.$store.getters['auth/getUser'].id}`, formData, {
          headers: {
            Authorization: `${this.$store.getters['auth/getPrefix']} ${this.$store.getters['auth/getToken']}`,
            'Content-Type': 'multipart/form-data',
          },
        });
        response.data.data.profileImage = `${profileImageUrl}/${response.data.data.profileImage}`;
        this.$store.commit('auth/SET_USER', response.data.data);
      } catch (error) {
        this.userError = error.response?.data?.message || 'Failed to update user information';
      } finally {
        this.isUpdatingUser = false;
      }
    },
    async updatePassword() {
      this.isUpdatingPassword = true;
      this.passwordError = null;

      try {
        await axios.put(`${serverUrl}/update-password`, {
          oldPassword: this.oldPassword,
          newPassword: this.newPassword,
        }, {
          headers: {
            Authorization: `${this.$store.getters['auth/getPrefix']} ${this.$store.getters['auth/getToken']}`,
          },
        });
      } catch (error) {
        this.passwordError = error.response?.data?.message || 'Failed to update password';
      } finally {
        this.isUpdatingPassword = false;
      }
    },
    async deleteAccount() {
      this.isDeletingAccount = true;
      this.deleteError = null;
      try {
        await axios.delete(`${serverUrl}/${this.$store.getters['auth/getUser'].id}`, {
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
