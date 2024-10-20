<template>
  <div class="container">
    <h2 class="title is-4">Add Role to User</h2>
    <h4 class="subtitle">Select a role to add to the user</h4>
    <h4>User Name: <strong>{{ user.username }}</strong></h4>
    <h4>User ID: <strong>{{ user.id }}</strong></h4>
    <div class="notification is-info" v-if="loading">
      Loading roles...
    </div>
    <ul class="list">
      <li v-for="role in roles" :key="role.id" class="m-4">
        <button class="button is-primary box" @click="addRoleToUser(role.id)">
          {{ role.name }} (ID: {{ role.id }})
        </button>
      </li>
    </ul>
    <div class="notification is-danger notification-width" v-if="error">{{ error }}</div>
    <div class="notification is-success notification-width" v-if="successMessage">{{ successMessage }}</div>
  </div>
</template>

<script>
import axios from "axios";

const serverUrl = import.meta.env.VITE_SERVER_URL + "/user-roles";

export default {
  props: ['userId'],
  data() {
    return {
      loading: false,
      error: null,
      successMessage: null
    };
  },
  computed: {
    roles() {
      return this.$store.getters['roles/allRoles'];
    },
    user() {
      return this.$store.getters['users/currentUser'];
    }
  },
  methods: {
    async fetchData() {
      this.loading = true;
      this.error = null;
      try {
        await this.$store.dispatch('roles/fetchRoles');
        await this.$store.dispatch('users/fetchUserById', {id: this.userId});
      } catch (err) {
        this.error = 'Failed to load roles or user information.';
      } finally {
        this.loading = false;
      }
    },
    async addRoleToUser(roleId) {
      this.loading = true;
      this.successMessage = null;
      this.error = null;
      try {
        await axios.post(
            serverUrl,
            {
              userId: parseInt(this.userId),
              roleId
            },
            {
              headers: {
                Authorization: `${this.$store.getters['auth/getPrefix']} ${this.$store.getters['auth/getToken']}`,
              },
            }
        );
        const role = this.roles.find(r => r.id === roleId);
        if (role) {
          this.successMessage = `User ${this.user.username} added to role ${role.name} successfully!`;
        } else {
          this.successMessage = 'Role added successfully, but role name could not be found.';
        }
      } catch (err) {
        this.error =  'Failed to add role to user.';
      } finally {
        this.loading = false;
      }
    }
  },
  created() {
    this.fetchData();
  }
};
</script>

<style scoped>
.list {
  list-style-type: none;
  padding: 0;
}

.box {
  width: 200px;
}

.notification-width{
  width: 300px;
}
</style>
