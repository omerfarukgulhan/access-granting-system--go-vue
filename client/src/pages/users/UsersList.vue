<template>
  <div>
    <h1>User List</h1>
    <table v-if="users.length">
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
      <tr v-for="user in users" :key="user.id">
        <td>{{ user.id }}</td>
        <td>{{ user.username }}</td>
        <td>{{ user.email }}</td>
        <td><img :src="user.profileImage" alt="Profile Image" class="profile-img"/></td>
        <td>
          <router-link :to="'/users/' + user.id">View</router-link>
        </td>
      </tr>
      </tbody>
    </table>
    <div v-else>
      <p>No users available.</p>
    </div>
  </div>
</template>

<script>
export default {
  computed: {
    users() {
      return this.$store.getters['users/allUsers'];
    }
  },
  methods: {
    fetchUsers() {
      this.$store.dispatch('users/fetchUsers');
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
}

th, td {
  padding: 10px;
  text-align: left;
  border-bottom: 1px solid #ddd;
}

.profile-img {
  width: 50px;
  height: 50px;
  border-radius: 50%;
}
</style>
