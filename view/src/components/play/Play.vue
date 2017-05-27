<template lang="pug">
  main.game-container
    ss-players
    ss-canvas
    ss-chat
</template>

<script>
  import SsCanvas from './Canvas';
  import SsChat from './Chat';
  import SsPlayers from './Players';
  import EventBus from '../../eventbus';

  export default {
    name: 'play',
    data() {
      return {
        conn: null,
      };
    },
    components: {
      SsCanvas,
      SsChat,
      SsPlayers,
    },
    methods: {
      onClose(event) {
        console.log(event);
      },

      onMessage(event) {
        console.log(event);
        const msg = JSON.parse(event.data);
        const state = msg.state;
        switch (state) {
          case 1:
            console.log(msg);
            this.$store.commit('addPlayer', { id: msg.id, fName: msg.fName, lName: msg.lName });
            if (msg.players) {
              Object.keys(msg.players).forEach((key) => {
                const name = msg.players[key];
                console.log(name);
                this.$store.commit('addPlayer', { id: key, fName: name, lName: name });
              });
            }
            EventBus.$emit('chat.received', msg);
            break;
          case 2:
            this.$store.commit('removePlayer', msg.id);
            EventBus.$emit('chat.received', msg);
            break;
          default:
            EventBus.$emit('chat.received', msg);
            break;
        }
      },
    },
    created() {
      EventBus.$on('chat.send', (message) => {
        console.log('Sending Message!');
        this.conn.send(JSON.stringify({
          id: 1,
          sendType: 0,
          msg: message,
          state: 0,
        }));
      });

      if (window.WebSocket) {
        this.conn = new WebSocket(`ws://${document.location.host}/ws${document.location.search}`);
        this.conn.onclose = this.onClose;
        this.conn.onmessage = this.onMessage;
      } else {
        console.log('Your browser does not support websockets.');
      }
    },
  };
</script>

<style lang="stylus">
  @import "../../css/mixins.styl"
  @import "../../css/colours.styl"

  .game-container
    display: flex
    align-items: center
    flex-direction: column

  .sidebox
    flex: 0 0 auto
    margin: 2rem 0
    _mat-shadows(2)
    width: 90%
    height: 10em
    background-color: $g-inner
    border: 6px solid $g-border
    border-radius: 10px

  @media(min-width: 900px)
    .game-container
      flex-direction: row
      justify-content: center
      height: calc(100% - 3rem)

    .sidebox
      width: 20%
      height: 70%
</style>
