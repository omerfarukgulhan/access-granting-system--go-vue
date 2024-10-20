<template>
  <div class="container">
    <h2 class="title is-4">Add User to Role</h2>
    <h4 class="subtitle">Select a user to add to the role</h4>
    <h4>Role Name: {{ role.name }}</h4>
    <h4>Role Id: {{ role.id }}</h4>
    <div class="notification is-info" v-if="loading">
      Loading users...
    </div>

    <ul class="list">
      <li v-for="user in users" :key="user.id" class="m-4">
        <button class="button is-primary box" @click="addUserToRole(user.id)">
          {{ user.username }} (ID: {{ user.id }})
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
  props: ['roleId'],
  data() {
    return {
      loading: false,
      error: null,
      successMessage: null
    };
  },
  computed: {
    users() {
      return this.$store.getters['users/allUsers'];
    },
    role() {
      return this.$store.getters['roles/currentRole'];
    }
  },
  methods: {
    async fetchData() {
      this.loading = true;
      this.error = null;
      try {
        await this.$store.dispatch('users/fetchUsers');
        await this.$store.dispatch('roles/fetchRoleById', {id: this.roleId});
      } catch (err) {
        this.error = 'Failed to load users';
      } finally {
        this.loading = false;
      }
    },
    async addUserToRole(userId) {
      this.loading = true;
      this.successMessage = null;
      this.error = null;
      try {
        await axios.post(
            serverUrl,
            {
              userId,
              roleId: parseInt(this.roleId),
            },
            {
              headers: {
                Authorization: `${this.$store.getters['auth/getPrefix']} ${this.$store.getters['auth/getToken']}`,
              },
            }
        );

        const user = this.users.find((u) => u.id === userId);
        if (user) {
          this.successMessage = `Role ${this.role.name} added to user ${user.username} successfully!`;
        } else {
          this.successMessage = 'Role added successfully, but user name could not be found.';
        }
      } catch (err) {
        this.error = 'Failed to add user to role.';
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
  width: 300px;
}

.notification-width{
  width: 300px;
}
</style>
