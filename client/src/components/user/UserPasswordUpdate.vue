<template>
  <div class="box has-background-white container user-delete">
    <h2 class="subtitle has-text-grey-dark">Change Password</h2>
    <form @submit.prevent="updatePassword">
      <div class="field">
        <label class="label has-text-black">Old Password</label>
        <div class="control">
          <input class="input" type="password" v-model="oldPassword" required :disabled="loading"/>
        </div>
      </div>
      <div class="field">
        <label class="label has-text-black">New Password</label>
        <div class="control">
          <input class="input" type="password" v-model="newPassword" required :disabled="loading"/>
        </div>
      </div>
      <div class="field">
        <div class="control">
          <button type="submit" class="button is-primary is-fullwidth" :class="{ 'is-loading': loading }"
                  :disabled="loading">
            Change Password
          </button>
        </div>
      </div>
      <p v-if="error" class="has-text-danger">{{ error }}</p>
    </form>
  </div>
</template>

<script>
export default {
  data() {
    return {
      oldPassword: '',
      newPassword: '',
      loading: false,
      error: null,
    };
  },
  methods: {
    async updatePassword() {
      this.loading = true;
      this.error = null;
      try {
        this.$emit('password-updated', {oldPassword: this.oldPassword, newPassword: this.newPassword});
      } catch (err) {
        this.error = err.response?.data?.message || 'Failed to update password';
      } finally {
        this.loading = false;
      }
    }
  }
};
</script>

<style>
.user-delete{
  width: 600px;
}
</style>