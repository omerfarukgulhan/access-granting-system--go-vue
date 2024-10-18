<template>
  <div class="login-container">
    <div class="box has-background-white">
      <h1 class="title has-text-centered has-text-black">Login</h1>
      <form @submit.prevent="handleLogin">
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
            />
          </div>
        </div>
        <div class="field">
          <div class="control">
            <button type="submit" class="button is-primary is-fullwidth">Login</button>
          </div>
        </div>
        <p v-if="errorMessage" class="has-text-danger">{{ errorMessage }}</p>
      </form>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      email: 'omer@omer.com',
      password: 'P4ssword',
      errorMessage: '',
    };
  },
  methods: {
    async handleLogin() {
      try {
        await this.$store.dispatch('auth/login', {email: this.email, password: this.password});
        this.$router.push('/');
      } catch (error) {
        this.errorMessage = error.message;
      }
    },
  },
};
</script>

<style scoped>
.login-container {
  display: flex;
  justify-content: center;
  align-items: center;
  margin: 2rem;
 }

.box {
  width: 300px;
}
</style>
