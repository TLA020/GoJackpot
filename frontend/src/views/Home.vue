<template>
    <div>
        <h2>Home</h2>
        <pre>Welcome {{user.email}}</pre>
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

    mounted() {
      this.$socketEvent("authorized", () => {
        console.log("authorized from sokcethandler")
      });
    },

  }
</script>

<style scoped>

</style>