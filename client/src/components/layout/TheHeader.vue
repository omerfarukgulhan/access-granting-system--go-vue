<template>
  <nav class="navbar is-primary">
    <div class="navbar-brand">
      <router-link class="navbar-item" to="/" exact>
        <strong>Access Granting System</strong>
      </router-link>
    </div>
    <div class="navbar-menu">
      <div class="navbar-start">
        <router-link class="navbar-item" to="/" exact>Users</router-link>
        <router-link v-if="isAdmin" class="navbar-item" to="/roles">Roles</router-link>
        <router-link v-if="isAdmin" class="navbar-item" to="/user-roles">User Roles</router-link>
      </div>
      <div class="navbar-end">
        <div v-if="!isLoggedIn" class="navbar-item">
          <button class="button is-light">
            <router-link class="has-text-black" to="/login">Login</router-link>
          </button>
        </div>
        <div v-if="!isLoggedIn" class="navbar-item">
          <button class="button is-light">
            <router-link class="has-text-black" to="/register">Register</router-link>
          </button>
        </div>
        <div v-if="isLoggedIn" class="navbar-item">
          <button class="button is-light" @click="logout">Logout</button>
        </div>
      </div>
    </div>
  </nav>
</template>
<script>
export default {
  computed: {
    isLoggedIn() {
      return true
    },
    isAdmin() {
      return true
    }
  },
  methods: {
    logout() {
    }
  }
}
</script>
