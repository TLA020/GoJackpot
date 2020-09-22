<template>
  <v-container fluid>
    <GameStatsBar></GameStatsBar>
    <v-col cols="6"> <Game></Game> </v-col>
    <v-col cols="6"> <Bet></Bet> </v-col>
  </v-container>
</template>

<script>
import GameStatsBar from "@/components/GameStatsBar";
import Game from "@/components/Game";
import Bet from "@/components/Bet";

export default {
  name: "jackpot",
  components: { GameStatsBar, Game, Bet },
  data() {
    return {
      bet: {
        amount: 0
      }
    };
  },

  methods: {},

  computed: {
    currentUsers() {
      return this.$store.state.$game.currentUsers;
    }
  }
};
</script>

<style scoped></style>
