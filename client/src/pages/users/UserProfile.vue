<template>
  <div class="user-profile container">
    <h1 class="title has-text-centered">User Profile</h1>
    <div class="box">
      <h2 class="subtitle">Update Information</h2>
      <form @submit.prevent="updateUser">
        <div class="field">
          <label class="label">Username</label>
          <div class="control">
            <input class="input" type="text" v-model="username" required/>
          </div>
        </div>

        <div class="field">
          <label class="label">Profile Image</label>
          <div class="control">
            <input class="input" type="file" @change="handleFileUpload"/>
          </div>
        </div>

        <div class="field">
          <div class="control">
            <button type="submit" class="button is-primary is-fullwidth">Update Info</button>
          </div>
        </div>
      </form>
    </div>

    <div class="box">
      <h2 class="subtitle">Change Password</h2>
      <form @submit.prevent="updatePassword">
        <div class="field">
          <label class="label">Old Password</label>
          <div class="control">
            <input class="input" type="password" v-model="oldPassword" required/>
          </div>
        </div>

        <div class="field">
          <label class="label">New Password</label>
          <div class="control">
            <input class="input" type="password" v-model="newPassword" required/>
          </div>
        </div>

        <div class="field">
          <div class="control">
            <button type="submit" class="button is-primary is-fullwidth">Change Password</button>
          </div>
        </div>
      </form>
    </div>

    <div class="box has-text-centered">
      <button class="button is-danger" @click="deleteAccount">Delete Account</button>
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
        response.data.data.profileImage = `${profileImageUrl}/${response.data.data.profileImage}`
        this.$store.commit('auth/SET_USER', response.data.data);
      } catch (error) {
        console.error('Failed to update user:', error);
      }
    },
    async updatePassword() {
      try {
        const response = await axios.put(`${serverUrl}/update-password`, {
          oldPassword: this.oldPassword,
          newPassword: this.newPassword,
        }, {
          headers: {
            Authorization: `${this.$store.getters['auth/getPrefix']} ${this.$store.getters['auth/getToken']}`,
          },
        });
      } catch (error) {
        console.error('Failed to update password:', error);
      }
    },
    async deleteAccount() {
      try {
        const response = await axios.delete(serverUrl, {
          headers: {
            Authorization: `${this.$store.getters['auth/getPrefix']} ${this.$store.getters['auth/getToken']}`,
          },
        });
        console.log('Account deleted:', response.data);
        this.$router.push('/login');
      } catch (error) {
        console.error('Failed to delete account:', error);
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
