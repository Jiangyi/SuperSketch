<template lang="pug">
  #ss-canvas
    #word-guess(v-bind:style="{ borderColor: wordBorder, backgroundColor: wordInner, color: wordColour }") Sidewinder
    #game-canvas(v-bind:style="{ borderColor: colour }")
      .stretchy-wrapper(ref="wrapper")
        ss-drawing
        ul#colour-palette
          each colIndex in [0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
            ss-colour-button(colour=colIndex)
</template>

<script>
  import SsColourButton from './ColourButton';
  import SsDrawing from './Drawing';

  export default {
    components: {
      SsColourButton,
      SsDrawing,
    },
    computed: {
      colour() {
        return this.$store.getters.colour;
      },
      wordInner() {
        return this.$store.getters.wordInner;
      },
      wordBorder() {
        return this.$store.getters.wordBorder;
      },
      wordColour() {
        return this.$store.getters.wordColour;
      },
    },
    name: 'ss-canvas',
  };
</script>

<style lang="stylus">
  @import "../../css/mixins.styl"
  @import "../../css/colours.styl"

  .stretchy-wrapper
    position: absolute
    top: 0
    bottom: 0
    left: 0
    right: 0

  #ss-canvas
    margin: 0 0 2em 0
    width: 95%
    #game-canvas
      _mat-shadows(2)
      _trans(0.4s)
      display: flex
      text-align: center
      justify-content: center
      align-items: center
      position: relative
      border: 6px solid $g-border2
      padding-bottom: 62.5%
      border-radius: 10px
      background-color: white

    #word-guess
      _mat-shadows(2, 3)
      _trans()
      background-color: $g-inner
      border: 6px solid $g-border
      text-align: center
      padding: 0 5%
      margin: 0 auto
      width: 6em
      font-size: 2em
      line-height: 1.25em
      color: #5c6873
      border-radius: 2em
      user-select: none

    #colour-palette
      display: flex
      flex-direction: row
      flex-wrap: wrap
      align-items: center
      justify-content: center
      height: 0
      margin: 0
      padding: 0
      list-style: none

  @media(min-width: 900px)
    #ss-canvas
      margin: 2rem
      width: 50%
      #word-guess
        top: -0.75em
        font-size: 3em
</style>
