<template>
  <section class="section">
    <div class="container has-text-centered is-flex is-justify-content-center is-align-items-center">
      <div v-if="loading" class="has-text-centered">
        <p>Loading...</p>
      </div>
      <div v-else-if="error" class="has-text-centered has-text-danger">
        <p>{{ error }}</p>
      </div>
      <div v-else class="card has-background-white-bis" style="max-width: 400px; margin: 0 auto;">
        <div class="card-image">
          <figure class="image is-4by3">
            <img
                v-if="user && user.profileImage"
                :src="user.profileImage"
                alt="User Image"
            />
            <img
                v-else
                :src="defaultProfileImage"
                alt="Default User Image"
            />
          </figure>
        </div>
        <div class="card-content">
          <p class="title is-4 has-text-black">{{ user?.username || 'Unknown User' }}</p>
          <div class="content">
            <h6 class="has-text-black">Email: {{ user?.email || 'No Email' }}</h6>
            <UserRoles :roles="user?.roles" :userId="user.id" :isAdmin="isAdmin" @remove-role="deleteUserRole"/>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script>
import axios from 'axios';
import UserRoles from '../../components/user/UserRoles.vue';

const serverUrl = import.meta.env.VITE_SERVER_URL + "/user-roles";
const profileImageUrl = import.meta.env.VITE_PROFILE_IMAGE_URL;

export default {
  components: {
    UserRoles,
  },
  props: ['id'],
  data() {
    return {
      loading: true,
      error: null,
    };
  },
  computed: {
    defaultProfileImage() {
      return `${profileImageUrl}/default.png`;
    },
    user() {
      return this.$store.getters['users/currentUser'];
    },
    isAdmin() {
      return this.$store.getters['auth/isAdmin'];
    }
  },
  methods: {
    async fetchUser() {
      this.loading = true;
      this.error = null;
      try {
        await this.$store.dispatch('users/fetchUserById', {id: this.id});
      } catch (err) {
        this.error = err.response?.data?.message || 'Failed to fetch user. Please try again.';
      } finally {
        this.loading = false;
      }
    },
    async deleteUserRole(roleId) {
      this.loading = true;
      this.error = null;
      try {
        await axios.delete(`${serverUrl}/${this.id}/${roleId}`, {
          headers: {
            Authorization: `${this.$store.getters['auth/getPrefix']} ${this.$store.getters['auth/getToken']}`,
          },
        });
        await this.fetchUser();
      } catch (err) {
        this.error = err.response?.data?.message || 'Failed to delete role.';
      } finally {
        this.loading = false;
      }
    }
  },
  created() {
    this.fetchUser();
  }
};
</script>
