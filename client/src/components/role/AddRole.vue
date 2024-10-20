<template>
  <div>
    <!-- Button to open modal -->
    <button class="button is-primary" @click="openModal">Create Role</button>

    <!-- Modal for adding a new role -->
    <div class="modal" :class="{ 'is-active': showModal }">
      <div class="modal-background" @click="closeModal"></div>
      <div class="modal-card">
        <header class="modal-card-head">
          <p class="modal-card-title">Add New Role</p>
          <button class="delete" aria-label="close" @click="closeModal"></button>
        </header>
        <section class="modal-card-body">
          <div class="field">
            <label class="label">Role Name</label>
            <div class="control">
              <input class="input" type="text" v-model="newRoleName" placeholder="Enter role name">
            </div>
          </div>
        </section>
        <footer class="modal-card-foot">
          <button class="button is-success" @click="createRole">Save Role</button>
          <button class="button ml-2" @click="closeModal">Cancel</button>
          <p v-if="error" class="has-text-danger ml-3">{{ error }}</p>
        </footer>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";

const serverUrl = import.meta.env.VITE_SERVER_URL + "/roles";

export default {
  data() {
    return {
      showModal: false,
      newRoleName: '',
      error: '',
    };
  },
  methods: {
    openModal() {
      this.showModal = true;
    },
    closeModal() {
      this.showModal = false;
      this.newRoleName = '';
      this.error = '';
    },
    async createRole() {
      if (this.newRoleName.trim() === '') {
        this.error = 'Role name cannot be empty';
        return;
      }
      try {
        await axios.post(serverUrl, {name: this.newRoleName}, {
          headers: {
            Authorization: `${this.$store.getters['auth/getPrefix']} ${this.$store.getters['auth/getToken']}`,
          },
        });
        this.$emit('roleAdded');
        this.closeModal();
      } catch (err) {
        this.error = err.response?.data?.message || 'Failed to create role. Please try again.';
      }
    }
  }
};
</script>

<style scoped>
.modal-card {
  width: 400px;
  margin: auto;
}

.has-text-danger {
  color: red;
}
</style>
