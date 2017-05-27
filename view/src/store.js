/* eslint-disable no-param-reassign */
import Vue from 'vue';
import Vuex from 'vuex';
import colours from './colours';

Vue.use(Vuex);

const store = new Vuex.Store({
  state: {
    colourIndex: null,
    playerMap: {},
  },
  mutations: {
    setColour(state, colIndex) {
      state.colourIndex = colIndex;
    },
    addPlayer(state, player) {
      state.playerMap[player.id] = { firstName: player.fName, lastName: player.lName };
      console.log(state.playerMap);
    },
    removePlayer(state, id) {
      console.log('deleting player');
      delete state.playerMap[id];
    },
  },
  getters: {
    playerName: state => id => state.playerMap[id],
    colour: state => colours.selected[state.colourIndex],
    wordBorder: state => colours.wordBorder[state.colourIndex],
    wordInner: state => colours.wordInner[state.colourIndex],
    wordColour: state => colours.wordColour[state.colourIndex],
  },
});

export default store;
