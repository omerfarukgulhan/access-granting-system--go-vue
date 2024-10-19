<template>
  <section class="section">
    <div class="container ">
      <div class="box has-text-centered has-background-white">
        <h2 class="title is-4 has-text-black">
          <span v-if="loading">Activating your account...</span>
          <span v-else>Account Activation</span>
        </h2>
        <div v-if="activationMessage" class="notification" :class="messageClass">
          {{ activationMessage }}
        </div>
        <div v-if="activationMessage && !loading">
          <button class="button is-primary mt-4" @click="goToLogin">
            Go to Login
          </button>
        </div>
      </div>
    </div>
  </section>
</template>

<script>
import axios from 'axios';

const serverUrl = import.meta.env.VITE_SERVER_URL + "/users";

export default {
  props: ['token'],
  data() {
    return {
      activationMessage: '',
      messageClass: '',
      loading: true,
    };
  },
  created() {
    this.activateUser();
  },
  methods: {
    async activateUser() {
      this.loading = true;
      try {
        const response = await axios.put(`${serverUrl}/activate/${this.token}`);
        if (response.data.success) {
          this.activationMessage = 'Your account has been activated successfully!';
          this.messageClass = 'is-success';
        } else {
          this.activationMessage = 'Account activation failed: ' + response.data.message;
          this.messageClass = 'is-danger';
        }
      } catch (error) {
        this.activationMessage = 'Error activating account. Please try again later.';
        this.messageClass = 'is-danger';
      } finally {
        this.loading = false;
      }
    },
    goToLogin() {
      this.$router.push('/login');
    }
  }
};
</script>

<style scoped>
.section {
  margin-top: 50px;
}

.box {
  max-width: 400px;
  margin: 0 auto;
}

.mt-4 {
  margin-top: 1.5rem;
}
</style>
