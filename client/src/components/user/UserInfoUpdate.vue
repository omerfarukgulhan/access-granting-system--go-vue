<template>
  <div class="box has-background-white container user-info">
    <h2 class="subtitle has-text-grey-dark">Update Information</h2>
    <form @submit.prevent="updateUser">
      <div class="field">
        <label class="label has-text-black">Username</label>
        <div class="control">
          <input class="input" type="text" v-model="username" required :disabled="isUpdatingUser"/>
        </div>
      </div>
      <div class="field">
        <label class="label has-text-black">Profile Image</label>
        <div class="control">
          <input class="input" type="file" @change="handleFileUpload" :disabled="isUpdatingUser"/>
        </div>
      </div>
      <div class="field">
        <div class="control">
          <button type="submit" class="button is-primary is-fullwidth" :class="{ 'is-loading': isUpdatingUser }"
                  :disabled="isUpdatingUser">
            Update Info
          </button>
        </div>
      </div>
      <p v-if="userError" class="has-text-danger">{{ userError }}</p>
    </form>
  </div>
</template>

<script>
export default {
  props: ['user'],
  data() {
    return {
      username: this.user.username,
      profileImage: '',
      isUpdatingUser: false,
      userError: null,
    };
  },
  methods: {
    handleFileUpload(event) {
      this.profileImage = event.target.files[0];
    },
    async updateUser() {
      this.isUpdatingUser = true;
      this.userError = null;
      const formData = new FormData();
      formData.append('username', this.username);
      if (this.profileImage) {
        formData.append('profileImage', this.profileImage);
      }
      try {
        this.$emit('user-updated', formData);
      } catch (error) {
        this.userError = error.response?.data?.message || 'Failed to update user information';
      } finally {
        this.isUpdatingUser = false;
      }
    }
  }
};
</script>

<style>
.user-info{
  width: 600px;
}
</style>
