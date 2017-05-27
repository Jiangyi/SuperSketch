<template lang="pug">
  li.colour(@click="setColour")
    button(v-bind:style="{ backgroundColor: stateColour }")
</template>

<script>
  import colours from '../../colours';

  export default {
    name: 'ss-colour-button',
    props: ['colour'],
    computed: {
      stateColour() {
        return colours.selected[this.colour];
      },
    },
    methods: {
      setColour() {
        this.$store.commit('setColour', this.colour);
      },
    },
  };
</script>

<style lang="stylus">
  @import "../../css/mixins.styl"
  @import "../../css/colours.styl"

  .colour
    position: relative
    top: 0.5rem
    margin: 0 1px
    _trans(0.1s)

    button
      _mat-shadows(2)
      _trans(0.4s)
      &:hover
        width: 2.5rem
        height: 2.5rem
      position: relative
      background-color: #85d5f7
      width: 2.5rem
      height: 2.5rem
      border: 0.25rem solid #ffffff
      border-radius: 50%
      &:focus
        outline: 0

  @media(min-width: 768px)
    .colour
      &:active
        padding-bottom: 1rem
      top: -2em
      button
        width: 2.5rem
        height: 2.5rem
        &:hover
          width: 3rem
          height: 3rem
        border: 0.25rem solid #ffffff
</style>
