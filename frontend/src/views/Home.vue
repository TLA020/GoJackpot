<template>
    <div>
        <pre>{{user}}</pre>
        <v-btn @click="placeBet">place bet</v-btn>
    </div>
</template>

<script>
  export default {
    name:"Home",
    data: () => ({
    }),
    computed: {
        user() {
          return this.$store.state.$auth.user;
        }
    },
    methods: {
      placeBet() {
        this.$store.dispatch("sendSocket", { name: "place-bet", data: { amount: Math.floor(Math.random() * 50) + 1 }});
      }
    }
  }
</script>

<style scoped>

</style>