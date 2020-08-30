<template>
    <div>
        <pre>{{user}}</pre>
    </div>
</template>

<script>
  export default {
    name:"Home",
    data: () => ({
    }),
    computed: {
        user() {
          return this.$store.state.auth.user;
        }
    },
  }
</script>

<style scoped>

</style>