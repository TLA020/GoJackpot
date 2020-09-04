<template>
  <v-form ref="form">
    <v-row>
      <v-slider
        v-model="amount"
        append-icon="mdi-currency-usd"
        prepend-icon="mdi-currency-usd-off"
        class="align-center col-8 p-4"
        :max="100"
        thumb-label="always"
        :min="1"
        hide-details
      >
      </v-slider>
    </v-row>

    <v-btn color="warning" class="mr-4">
      Reset
    </v-btn>
    <v-btn :disabled="false" color="success">
      Place bet
    </v-btn>
  </v-form>
</template>

<script>
export default {
  name: "Bet",
  props: {
    value: {
      type: Number,
      required: false,
      default: 0
    }
  },

  computed: {
    amount: {
      get() {
        return this.value;
      },
      set(val) {
        this.$emit("input", val)
      }
    }
  },

  methods: {
    placeBet() {
      this.$store.dispatch("sendSocket", {
        name: "place-bet",
        data: { amount: Math.floor(Math.random() * 50) + 1 }
      });
    },
  }
};
</script>

<style scoped></style>
