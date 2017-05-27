<template lang="pug">
  main.container
    .centered-box
      form(method="post")#ss-login.auth-box
        h1 Login
        input(type="text", name="username", placeholder="Username", v-model="username" required)
        input(type="password", name="password", placeholder="Password", v-model="password" required)
        button(type="submit", @click="submitCredentials").std-button Login

</template>

<script>

  export default {
    name: 'login',
    data() {
      return {
        username: '',
        password: '',
        message: 'Login not implemented',
      };
    },
    methods: {
      submitCredentials() {

      },
    },
  };
</script>

<style lang="stylus">
  @import "../../css/mixins.styl"
  @import "../../css/colours.styl"

</style>

