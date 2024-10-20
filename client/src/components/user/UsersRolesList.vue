<template>
  <div>
    <div class="roles-header is-flex is-align-items-center is-justify-content-space-between">
      <h6 class="has-text-black">Roles:</h6>
      <button v-if="isAdmin" @click="addRole" class="button is-small is-link">Add Role</button>
    </div>
    <div v-if="loading" class="has-text-centered">
      <p>Loading...</p>
    </div>
    <div v-if="error" class="has-text-centered has-text-danger">
      <p>{{ error }}</p>
    </div>
    <ul v-if="!loading && roles && roles.length > 0" class="no-bullet">
      <li v-for="role in roles" :key="role.id"
          class="role-item is-flex is-align-items-center is-justify-content-space-between">
        <span class="has-text-black">{{ role.name }}</span>
        <button v-if="isAdmin" @click="removeRole(role.id)" class="button is-small is-danger" :disabled="loading">
          Remove
        </button>
      </li>
    </ul>
    <p v-else-if="!loading && roles.length === 0" class="has-text-grey">User has no roles.</p>
  </div>
</template>

<script>
export default {
  props: {
    roles: {
      type: Array,
      required: true,
    },
    userId: {
      type: Number,
      required: true,
    },
    isAdmin: {
      type: Boolean,
      required: true,
    }
  },
  data() {
    return {
      loading: false,
      error: null,
    };
  },
  methods: {
    addRole() {
      this.$router.push(`/user-roles/add-role/${this.userId}`);

    },
    async removeRole(roleId) {
      console.log(roleId)
      this.loading = true;
      this.error = null;
      try {
        await this.$emit('remove-role', roleId);
      } catch (err) {
        this.error = 'Failed to remove role. Please try again.';
      } finally {
        this.loading = false;
      }
    }
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
