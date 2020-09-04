<template>
    <v-list subheader>
        <v-list-item
                v-for="user in currentUsers"
                :key="user.Email"
                class="mt-1"
        >
            <v-list-item-avatar>
                <v-img :src="getAvatar(user.Email)"></v-img>
            </v-list-item-avatar>

            <v-list-item-content>
                <v-list-item-title v-text="user.Email"></v-list-item-title>
            </v-list-item-content>
        </v-list-item>
    </v-list>
</template>

<script>
  export default {
    name: "UsersOnline",
    computed: {
      currentUsers() {
        return this.$store.state.$game.currentUsers;
      }
    },

    methods: {
      getAvatar(x) {
        return `https://api.adorable.io/avatars/${x}`;
      }
    }
  }
</script>

<style scoped>

</style>