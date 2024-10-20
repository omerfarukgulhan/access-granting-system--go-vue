<template>
  <div>
    <h1 class="title has-text-centered m-4">Role List</h1>
    <div class="has-text-centered">
      <add-role @roleAdded="fetchRoles"></add-role>
    </div>
    <div v-if="loading" class="has-text-centered">
      <p>Loading roles...</p>
    </div>
    <div v-if="error" class="notification is-danger">
      <p>{{ error }}</p>
    </div>
    <table v-if="!loading && roles.length">
      <thead>
      <tr>
        <th>ID</th>
        <th>Role name</th>
        <th class="actions-column">Actions</th>
      </tr>
      </thead>
      <tbody>
      <tr v-for="(role, index) in roles" :key="role.id" :class="{ 'striped': index % 2 === 0 }">
        <td>{{ role.id }}</td>
        <td>{{ role.name }}</td>
        <td class="actions-column">
          <!-- View Role Button -->
          <button class="button is-link">
            <router-link :to="'/roles/' + role.id" class="has-text-white">View</router-link>
          </button>
          <!-- Delete Role Button -->
          <button class="button is-danger ml-2" @click="deleteRole(role.id)">Delete</button>
        </td>
      </tr>
      </tbody>
    </table>
    <div v-else-if="!loading && !roles.length">
      <p>No roles available.</p>
    </div>
  </div>
</template>

<script>
import AddRole from '../../components/role/AddRole.vue';
import axios from 'axios';

const serverUrl = import.meta.env.VITE_SERVER_URL + '/roles';

export default {
  components: {
    AddRole,
  },
  data() {
    return {
      loading: true,
      error: '',
    };
  },
  computed: {
    roles() {
      return this.$store.getters['roles/allRoles'];
    }
  },
  methods: {
    async fetchRoles() {
      this.loading = true;
      this.error = '';
      try {
        await this.$store.dispatch('roles/fetchRoles');
      } catch (err) {
        this.error = 'Failed to fetch roles. Please try again.';
      } finally {
        this.loading = false;
      }
    },
    async deleteRole(roleId) {
      try {
        await axios.delete(`${serverUrl}/${roleId}`, {
          headers: {
            Authorization: `${this.$store.getters['auth/getPrefix']} ${this.$store.getters['auth/getToken']}`,
          }
        });
        this.fetchRoles(); // Refresh the role list after deletion
      } catch (err) {
        this.error = 'Failed to delete role. Please try again.';
      }
    }
  },
  created() {
    this.fetchRoles();
  }
};
</script>

<style scoped>
table {
  width: 600px;
  border-collapse: collapse;
  margin: 20px auto;
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

.actions-column {
  width: 200px;
  text-align: center;
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

.ml-2 {
  margin-left: 0.5rem;
}
</style>
