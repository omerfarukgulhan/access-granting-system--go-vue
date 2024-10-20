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
            <div class="roles-header is-flex is-align-items-center is-justify-content-space-between">
              <h6 class="has-text-black">Roles:</h6>
              <button @click="addUser" class="button is-small is-link">Add User</button>
            </div>
            <ul v-if="role && role.users && role.users.length > 0" class="no-bullet">
              <li v-for="user in role.users" :key="user.id"
                  class="role-item is-flex is-align-items-center is-justify-content-space-between">
                <span class="has-text-black">{{ user.username }}</span>
                <button @click="deleteUserRole(user.id)" class="button is-small is-danger"
                        :disabled="loading">
                  Remove
                </button>
              </li>
            </ul>
            <p v-else class="has-text-grey">role has no users.</p>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script>
import axios from "axios";

const serverUrl = import.meta.env.VITE_SERVER_URL + "/user-roles";

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
    addUser() {
      this.$router.push(`/user-roles/add-user/${this.id}`);
    },
    async deleteUserRole(userId) {
      this.loading = true;
      this.error = null;
      try {
        await axios.delete(`${serverUrl}/${userId}/${this.id}`, {
          headers: {
            Authorization: `${this.$store.getters['auth/getPrefix']} ${this.$store.getters['auth/getToken']}`,
          },
        });
        await this.fetchRole();
      } catch (err) {
        this.error = err.response?.data?.message || 'Failed to delete user.';
      } finally {
        this.loading = false;
      }
    },
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

.role-card {
  width: 200px;
}
</style>
