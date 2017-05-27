<template lang="pug">
  header#header
    span.logo
      a(href="/")
        img(src="../assets/super-sketch-v4.svg")
    span.options
</template>

<script>
  export default {
    name: 'ss-header',
  };
</script>

<style lang="stylus">
  @import "../css/mixins.styl"
  @import "../css/colours.styl"

  #header
    width: 100%
    _mat-shadows(2)
    height: 3rem
    background-color: $c-header

    .logo img
      user-select: none
      height: 4rem
      filter: drop-shadow(0 2px 4px rgba(0,0,0,0.23))
</style>
