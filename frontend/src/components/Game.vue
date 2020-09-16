<template>
  <div>
    GAMESTATE: {{ gameState }}
    <div v-if="gameState === 'winnerPicked'">
      <h2>WINNER: {{ winner.user.email }}</h2>
      <h3>WON: €{{ winner.amount }}</h3>
      <p>New game starting soon.</p>
    </div>

    <div v-if="gameState === 'idle'">
      <v-chip color="orange" v-if="gameState === 'idle'"
        >Starting when 2 players place bets</v-chip
      >
    </div>

    <div v-if="gameState === 'inProgress'">
      <v-row>
        <h2 class="mr-4" v-if="totalPot">Total: €{{ totalPot }}</h2>
        <h2 v-if="timeLeft">Time Left: {{ timeLeft }}</h2>
        <v-expansion-panels popout>
          <v-expansion-panel
            v-for="(userBet, i) in game.userBets"
            :key="i"
            hide-actions
          >
            <v-expansion-panel-header>
              <v-row align="center" class="spacer" no-gutters>
                <v-col cols="4" sm="2" md="1">
                  <v-avatar size="36px">
                    <img alt="Avatar" :src="getAvatar(userBet.player.email)" />
                  </v-avatar>
                </v-col>

                <v-col class="hidden-xs-only" sm="5" md="3">
                  <strong v-html="userBet.player.email"></strong>
                </v-col>

                <v-col class="text-no-wrap" cols="5" sm="3">
                  <v-chip
                    :color="`lighten-4`"
                    class="ml-0 mr-2 black--text"
                    label
                  >
                    <strong class="mr-1">Total bet: </strong> €{{
                      userBet.bets
                        .map(o => o.amount)
                        .reduce((a, c) => {
                          return a + c;
                        })
                    }},-
                  </v-chip>
                </v-col>

                <v-col>
                  <span class="float-right"
                    >({{ userBet.bets.length }})
                  </span></v-col
                >
              </v-row>
            </v-expansion-panel-header>

            <v-expansion-panel-content>
              <strong>Bets placed</strong>
              <v-divider></v-divider>
              <v-card-text>
                <ul v-for="bet in userBet.bets" :key="bet.Created">
                  <li>
                    <span>€{{ bet.amount }},-</span>
                  </li>
                </ul>
              </v-card-text>
            </v-expansion-panel-content>
          </v-expansion-panel>
        </v-expansion-panels>
      </v-row>
    </div>
  </div>
</template>

<script>
export default {
  name: "Game",
  methods: {
    getAvatar(x) {
      return `https://api.adorable.io/avatars/${x}`;
    }
  },

  computed: {
    game() {
      return this.$store.state.$game.game;
    },

    userBets() {
      return this.game.userBets;
    },

    timeLeft() {
      return this.$store.state.$game.timeLeft;
    },

    winner() {
      return this.$store.state.$game.winner
    },

    gameState() {
      const game_states = ["idle", "inProgress", "ended", "winnerPicked"];
      let currentState = this.$store.state.$game.game.state || 0;
      return game_states[currentState];
    },

    totalPot() {
      let total = 0;
      if (!this.game.userBets) {
        return 0;
      }
      this.game.userBets.forEach(f => {
        total += f.bets
          .map(o => o.amount)
          .reduce((a, c) => {
            return a + c;
          });
      });
      return total;
    }
  }
};
</script>

<style scoped></style>
