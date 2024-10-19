<template>
  <div class="box has-background-white container user-delete">
    <h2 class="subtitle has-text-grey-dark">Change Password</h2>
    <form @submit.prevent="updatePassword">
      <div class="field">
        <label class="label has-text-black">Old Password</label>
        <div class="control">
          <input class="input" type="password" v-model="oldPassword" required :disabled="isUpdatingPassword"/>
        </div>
      </div>
      <div class="field">
        <label class="label has-text-black">New Password</label>
        <div class="control">
          <input class="input" type="password" v-model="newPassword" required :disabled="isUpdatingPassword"/>
        </div>
      </div>
      <div class="field">
        <div class="control">
          <button type="submit" class="button is-primary is-fullwidth" :class="{ 'is-loading': isUpdatingPassword }"
                  :disabled="isUpdatingPassword">
            Change Password
          </button>
        </div>
      </div>
      <p v-if="passwordError" class="has-text-danger">{{ passwordError }}</p>
    </form>
  </div>
</template>

<script>
export default {
  data() {
    return {
      oldPassword: '',
      newPassword: '',
      isUpdatingPassword: false,
      passwordError: null,
    };
  },
  methods: {
    async updatePassword() {
      this.isUpdatingPassword = true;
      this.passwordError = null;
      try {
        this.$emit('password-updated', {oldPassword: this.oldPassword, newPassword: this.newPassword});
      } catch (error) {
        this.passwordError = error.response?.data?.message || 'Failed to update password';
      } finally {
        this.isUpdatingPassword = false;
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