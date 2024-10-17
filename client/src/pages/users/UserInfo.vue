<template>
  <div v-if="user">
    <h1>User Information</h1>
    <div class="user-info">
      <img :src="user.profileImage" alt="Profile Image"/>
      <p><strong>Username:</strong> {{ user.username }}</p>
      <p><strong>Email:</strong> {{ user.email }}</p>
      <p><strong>Roles:</strong></p>
      <ul>
        <li v-for="role in user.roles" :key="role.id">{{ role.name }}</li>
      </ul>
    </div>
  </div>
  <div v-else>
    <p>Loading user information...</p>
  </div>
</template>

<script>
export default {
  props: ['id'],
  computed: {
    user() {
      return this.$store.getters['users/currentUser'];
    }
  },
  methods: {
    fetchUser() {
      this.$store.dispatch('users/fetchUserById', {id: this.id});
    }
  },
  created() {
    this.fetchUser();
  }
};
</script>

<style scoped>
.user-info {
  border: 1px solid #ccc;
  padding: 20px;
  border-radius: 10px;
  max-width: 400px;
  margin: 0 auto;
  text-align: center;
}

.user-info img {
  width: 100px;
  height: 100px;
  border-radius: 50%;
}
</style>
