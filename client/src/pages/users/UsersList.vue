<template>
  <div>
    <h1 class="title has-text-centered m-4">User List</h1>
    <div v-if="loading" class="has-text-centered">
      <p>Loading users...</p>
    </div>
    <div v-if="error" class="notification is-danger">
      <p>{{ error }}</p>
    </div>
    <table v-if="!loading && users.length">
      <thead>
      <tr>
        <th>ID</th>
        <th>Username</th>
        <th>Email</th>
        <th>Profile</th>
        <th>Actions</th>
      </tr>
      </thead>
      <tbody>
      <tr v-for="(user, index) in users" :key="user.id" :class="{ 'striped': index % 2 === 0 }">
        <td>{{ user.id }}</td>
        <td>{{ user.username }}</td>
        <td>{{ user.email }}</td>
        <td>
          <img :src="user.profileImage" alt="Profile Image" class="profile-img"/>
        </td>
        <td>
          <button class="button is-link">
            <router-link :to="'/users/' + user.id" class="has-text-white">View</router-link>
          </button>
        </td>
      </tr>
      </tbody>
    </table>
    <div v-else-if="!loading && !users.length">
      <p>No users available.</p>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      loading: true,
      error: '',
    };
  },
  computed: {
    users() {
      return this.$store.getters['users/allUsers'];
    }
  },
  methods: {
    async fetchUsers() {
      this.loading = true;
      this.error = '';

      try {
        await this.$store.dispatch('users/fetchUsers');
      } catch (err) {
        this.error = err.response?.data?.message || 'Failed to fetch users. Please try again.';
      } finally {
        this.loading = false;
      }
    }
  },
  created() {
    this.fetchUsers();
  }
};
</script>

<style scoped>
table {
  width: 100%;
  border-collapse: collapse;
  margin-top: 20px;
  background-color: #fff;
}

thead {
  background-color: #f4f4f4;
}

th, td {
  padding: 10px;
  text-align: left;
  border-bottom: 1px solid #ddd;
  color: #333;
}

.profile-img {
  width: 50px;
  height: 50px;
  border-radius: 50%;
}

tbody tr.striped {
  background-color: #f9f9f9;
}

tbody tr {
  background-color: #fff;
}

tbody tr:hover {
  background-color: #e0e0e0;
}

.notification {
  margin-top: 20px;
}
</style>
