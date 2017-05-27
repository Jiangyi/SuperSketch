<template lang="pug">
  main.container.index
    img(src="../../assets/super-sketch-v4.svg").big-logo
    .buttons
      a(href="/login") Login
      a(href="/register") Register
      a(href="/play") Play
</template>

<script>
  export default {
    name: 'index',
  };
</script>

<style lang="stylus">
  @import "../../css/mixins.styl"
  @import "../../css/colours.styl"

  .big-logo
    padding: 1rem
    display: block
    margin: 3rem auto
    width: 15rem
    filter: drop-shadow(0 2px 4px rgba(0,0,0,0.5))

  .buttons
    text-align: center
    a
      line-height: 3rem
      text-decoration: none
      padding: 1px 6px
      display: block
      margin: 1.5rem auto
      color: #fff
      background-color: #475669
      border: none
      &:hover
        background-color: #59697d
      _trans(0.2s)
      _mat-shadows(2, 3)
      width: 7rem
      height: 3rem
      font-size: 1.25rem
      border-radius: 3px

  @media(min-width: 768px)
    .big-logo
      margin: 3rem auto
      width: 30rem
    .buttons
      button
        display: inline-block
        margin: 0 1rem

  @media(min-width: 2000px)
    .big-logo
      width: 40rem
    .buttons
      a
        display: inline-block
        margin: 1.5rem 1rem

</style>
