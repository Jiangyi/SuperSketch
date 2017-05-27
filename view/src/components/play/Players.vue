<template lang="pug">
  #ss-players.sidebox
    .status-box
      h2 Bill Nye's room
      h3 {{ players.length }} /8 Players
    .player-box
      ul.players
        li(v-for="player in players", v-bind:class="{ highlighted: player.highlight }")
          div
            img(:src="getProfileImage(player.image)")
            h3 {{ player.name }}
            p {{ player.points }} PTS
          h4(v-if="player.name === 'Malt Barley'", class="you-banner") You
          h4(v-if="player.highlight", class="drawing-banner") Drawing

</template>

<script>
  const req = require.context('../../assets/', false, /^.*\.(jpg|png)$/);
  export default {
    name: 'ss-players',
    data() {
      return {
        round: 2,
        players: [
          { name: 'Joe Blow', points: 4, image: 'default', highlight: false },
          { name: 'Marie Pooppins', points: 1, image: 'default', highlight: true },
          { name: 'Larry Lobster', points: 14, image: 'default', highlight: false },
          { name: 'Ben Dover', points: 7, image: 'default', highlight: false },
          { name: 'Bill Nye', points: 6, image: 'default', highlight: false },
          { name: 'Malt Barley', points: 7, image: 'default', highlight: false },
          { name: 'Gary Paulson', points: 12, image: 'default', highlight: false },
          { name: 'Lee Ho Suk', points: 9, image: 'default', highlight: false },
        ],
      };
    },
    methods: {
      getProfileImage(path) {
        return req(`./${path}.jpg`);
      },
    },
  };
</script>

<style lang="stylus">
  @import "../../css/mixins.styl"
  @import "../../css/colours.styl"

  #ss-players
    color: $chat-text;
    display: flex
    flex-direction: column
    position: relative
    font-family: 'Lato', sans-serif;

    .status-box
      text-align: center
      padding: 0 0.5rem
      margin: 0
      h2
        margin: 0.5rem 0.25rem 0
      h3
        margin: 0.5rem 0.25rem 0
    .player-box
      overflow-y: auto
      overflow-x: hidden
      flex: 1
      background-color: $g-inner2
      margin: 0.5rem
      border-radius: 6px
      border: 2px solid $g-border2
      ul
        margin: 0
        list-style: none
        padding: 0.5rem
        line-height: 1.25rem
        li
          padding: 0 0.5rem
          margin: 0
          border-radius: 6px
          border: 1px solid $g-border2
          position: relative
          div
            img
              margin: 0 0.5rem 0 0
              height: 50px
              float: left
              border-radius: 6px
            h3
              padding: 0.5rem 0 0
              margin: 0.5rem 0 0
            p
              padding: 0
              margin: 0 0 0.5rem
          h4
            position: absolute
            padding: 0.2rem 0.5rem
            border-radius: 1px
          .you-banner
            left: -0.75rem
            top: -1.5rem
            background-color: #F54033
            color: $g-inner2
            _rotate(-7deg)
          .drawing-banner
            right: -0.75rem
            top: -1.8rem
            background-color: #FF9900
            color: $g-inner2
            _rotate(7deg)
        .highlighted
          border-radius: 6px
          border: 1px solid $g-border2
</style>
