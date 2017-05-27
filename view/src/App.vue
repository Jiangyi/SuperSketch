<template lang="pug">
  #app
    ss-header
    router-view
</template>

<script>
  import SsHeader from './components/Header';
  import getCookie from './flash';

  export default {
    name: 'app',
    components: {
      SsHeader,
    },
    mounted() {
      const flashMessage = getCookie('Flash-Msg');
      if (flashMessage !== null) {
        console.log(flashMessage);
      }
    },
  };
</script>

<style lang="stylus" src="./css/main.styl"></style>
