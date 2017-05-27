<template lang="pug">
  main.container
    .centered-box
      form(method="post")#ss-register.auth-box
        h1 Register
        input(type="email", name="email", placeholder="Email", v-model="email" required)
        input(type="text", name="username", placeholder="Username", v-model="username" required)
        input(type="text", name="firstName", placeholder="First Name", required)
        input(type="text", name="lastName", placeholder="Last Name", required)
        input(type="password", name="password", placeholder="Password", v-model="password" ref="password" required)
        input(type="password", name="password2", placeholder="Confirm Password", v-model="confirmPassword" ref="confirmPassword" required)
        .tac
          p Terms and Conditions
        .checkbox
          input(type="checkbox", name="agree", v-model="agree" required)
          label(for="agree") I Agree
        button(type="submit", @click="register").std-button Register

</template>

<script>
  export default {
    name: 'login',
    data() {
      return {
        username: '',
        password: '',
        confirmPassword: '',
        agree: false,
      };
    },
    methods: {
      register() {
        // confirm password validation
        if (this.password !== this.confirmPassword) {
          this.$refs.confirmPassword.setCustomValidity("Passwords Don't Match");
        } else {
          this.$refs.confirmPassword.setCustomValidity('');
        }
      },
    },
  };
</script>

<style lang="stylus">
  @import "../../css/mixins.styl"
  @import "../../css/colours.styl"
  .tac
    margin: 1rem 0
    padding: 1rem
    background-color: white
    overflow-x: hidden
    overflow-y: hidden
    width: 200px
    text-align: center
    p
      color: black
  .checkbox
    label
      margin: 0 0.5rem
</style>
