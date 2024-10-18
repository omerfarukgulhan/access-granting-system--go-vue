<template>
  <section class="section">
    <div class="container has-text-centered is-flex is-justify-content-center is-align-items-center full-height">
      <div class="card has-background-white-bis" style="max-width: 400px; margin: 0 auto;">
        <!-- Use Bulma class for white background -->
        <div class="card-image">
          <figure class="image is-4by3">
            <img
                :src="user.profileImage"
                alt="User Image"
            />
          </figure>
        </div>
        <div class="card-content">
          <div class="media">
            <div class="media-content">
              <p class="title is-4 has-text-black">{{ user.username }}</p>
              <p class="subtitle is-6 has-text-grey">@{{ user.username }}</p>
            </div>
          </div>

          <div class="content">
            <p class="has-text-black">Email:<strong class="has-text-dark">{{ user.email }}</strong></p>
            <p class="has-text-black">Roles</p>
            <ul class="no-bullet has-text-left">
              <li v-for="role in user.roles" :key="role.id" class="has-text-black">{{ role.name }}</li>
            </ul>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script>
export default {
  props: ['id'],
  computed: {
    user() {
      return this.$store.getters['users/currentUser'];
    }
  },
  methods: {
    fetchUser() {
      this.$store.dispatch('users/fetchUserById', {id: this.id});
    }
  },
  created() {
    this.fetchUser();
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
