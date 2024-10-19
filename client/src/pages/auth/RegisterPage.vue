<template>
  <div class="register-container">
    <div class="box has-background-white">
      <h1 class="title has-text-centered has-text-black">Register</h1>
      <form @submit.prevent="handleRegister">
        <div class="field">
          <label class="label has-text-black" for="email">Email</label>
          <div class="control">
            <input
                v-model="email"
                id="email"
                class="input"
                placeholder="Email"
                type="email"
                required
                :disabled="loading"
            />
          </div>
        </div>
        <div class="field">
          <label class="label has-text-black" for="username">Username</label>
          <div class="control">
            <input
                v-model="username"
                id="username"
                class="input"
                placeholder="Username"
                type="text"
                required
                :disabled="loading"
            />
          </div>
        </div>
        <div class="field">
          <label class="label has-text-black" for="password">Password</label>
          <div class="control">
            <input
                v-model="password"
                id="password"
                class="input"
                placeholder="Password"
                type="password"
                required
                :disabled="loading"
            />
          </div>
        </div>
        <div class="field">
          <div class="control">
            <button type="submit" class="button is-primary is-fullwidth" :disabled="loading">
              <span v-if="!loading">Register</span>
              <span v-else>Loading...</span>
            </button>
          </div>
        </div>
        <p v-if="error" class="has-text-danger">{{ error }}</p>
      </form>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      email: 'faruk@faruk.com',
      username: 'username2',
      password: 'P4ssword',
      loading: false,
      error: '',
    };
  },
  methods: {
    async handleRegister() {
      this.loading = true;
      this.error = '';
      try {
        await this.$store.dispatch('auth/register', {
          email: this.email,
          username: this.username,
          password: this.password,
        });
        alert('Please activate your email');
        this.$router.push('/login');
      } catch (err) {
        this.error = err.response?.data?.message || 'Registration failed. Please try again.';
      } finally {
        this.loading = false;
      }
    },
  },
};
</script>

<style scoped>
.register-container {
  display: flex;
  justify-content: center;
  align-items: center;
  margin: 2rem;
}

.box {
  width: 300px;
}
</style>
