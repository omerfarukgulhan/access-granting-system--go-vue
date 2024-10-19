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
            <p class="has-text-black"><strong class="has-text-black">Email:</strong> {{ user?.email || 'No Email' }}</p>
            <h6 class="has-text-black">Roles:</h6>
            <ul v-if="user && user.roles && user.roles.length > 0" class="no-bullet">
              <li v-for="role in user.roles" :key="role.id" class="has-text-black">{{ role.name }}</li>
            </ul>
            <p v-else class="has-text-grey">User has no roles.</p>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script>
const profileImageUrl = import.meta.env.VITE_PROFILE_IMAGE_URL;

export default {
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
    }
  },
  created() {
    this.fetchUser();
  }
};
</script>

<style scoped>
.no-bullet {
  list-style-type: none;
  padding: 0;
  margin: 0;
}
</style>
