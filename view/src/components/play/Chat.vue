<template lang="pug">
  #ss-chat.sidebox
    .chat-box
      input(placeholder="Chat Here!", v-model="inputText", @keydown.enter="submitChat")
      button.game-btn(type="submit", @click="submitChat")
        i.material-icons.md-18 send
    .chat-buttons
      button.game-btn Report
      button.game-btn Invite
    .message-box
      ul.messages
        li(v-for="message in messages") {{ message }}
</template>

<script>
  import EventBus from '../../eventbus';

  export default {
    name: 'ss-chat',
    data() {
      return {
        inputText: '',
        messages: [],
      };
    },
    created() {
      EventBus.$on('chat.received', (event) => {
        const id = event.id;
        const playerName = this.$store.getters.playerName(id);
        if (event.state !== 0) {
          this.messages.push(`System: ${playerName.firstName} ${playerName.lastName} has joined the room.`);
        } else {
          this.messages.push(`${playerName.firstName} ${playerName.lastName}: ${event.msg}`);
        }
      });
    },
    methods: {
      submitChat() {
        if (this.inputText !== '') {
          EventBus.$emit('chat.send', this.inputText);
        }
        this.clearChat();
      },
      clearChat() {
        this.inputText = '';
      },
    },
  };
</script>

<style lang="stylus">
  @import "../../css/mixins.styl"
  @import "../../css/colours.styl"

  #ss-chat
    margin-top: 4em
    height: 14em
    color: $chat-text
    display: flex
    flex-direction: column
    position: relative
    font-family: 'Lato', sans-serif;

    input
      font-family: inherit
      &:focus
        outline: 0

    button
      _trans(0.2s)
      border-radius: 2rem

    .chat-buttons
      text-align: center
      button
        user-select: none
        box-sizing: inherit
        display: inline-block
        height: 1.5rem
        width: 4rem
        margin: 0.5rem 1rem
        font-size: 1rem

    .message-box
      overflow-y: auto
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
          word-wrap: break-word

    .chat-box
      position: relative
      width: 100%
      height: 3rem
      padding: 0.5rem
      box-sizing: border-box
      input
        _trans(0.1s)
        color: $chat-text
        &::placeholder
          color: $g-border
          user-select: none
        &:focus
          _trans(0.1s)
          border: 2px solid #C0CCDA
        border-radius: 1rem
        padding: 0 2.1rem 0 0.5rem
        border: 2px solid $g-border2
        box-sizing: border-box
        width: 100%
        font-size: 1rem
        height: 2rem
      button
        height: 2rem
        width: 2rem
        margin-right: 0.5rem
        position: absolute
        right: 0
        line-height: 2.25rem

  @media (min-width: 900px)
    #ss-chat
      height: 70%
      margin-top: 0
      .message-box
        margin-bottom: 3rem
      .chat-box
        position: absolute
        bottom: 0

</style>
