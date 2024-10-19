<template>
  <section class="section">
    <div class="container has-text-centered is-flex is-justify-content-center is-align-items-center">
      <div v-if="loading" class="has-text-centered">
        <p>Loading...</p>
      </div>
      <div v-else-if="error" class="has-text-centered has-text-danger">
        <p>{{ error }}</p>
      </div>
      <div v-else class="card has-background-white-bis role-card" style="max-width: 400px; margin: 0 auto;">
        <div class="card-content">
          <p class="title is-4 has-text-black">{{ role?.name || 'Unknown role' }}</p>
          <div class="content">
            <h6 class="has-text-black">Users:</h6>
            <ul v-if="role && role.users && role.users.length > 0" class="no-bullet">
              <li v-for="user in role.users" :key="user.id" class="has-text-black">{{ user.username }}</li>
            </ul>
            <p v-else class="has-text-grey">role has no users.</p>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script>
export default {
  props: ['id'],
  data() {
    return {
      loading: true,
      error: null,
    };
  },
  computed: {
    role() {
      return this.$store.getters['roles/currentRole'];
    }
  },
  methods: {
    async fetchRole() {
      this.loading = true;
      this.error = null;
      try {
        await this.$store.dispatch('roles/fetchRoleById', {id: this.id});
      } catch (err) {
        this.error = err.response?.data?.message || 'Failed to fetch role. Please try again.';
      } finally {
        this.loading = false;
      }
    }
  },
  created() {
    this.fetchRole();
  }
};
</script>

<style scoped>
.no-bullet {
  list-style-type: none;
  padding: 0;
  margin: 0;
}
.role-card{
  width: 200px;
}
</style>
