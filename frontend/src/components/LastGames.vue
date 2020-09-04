<template>
    <div>
    <v-tabs>
      <v-tab>Last Games</v-tab>
      <v-tab>My Games</v-tab>
      <v-tab>24H Ranking</v-tab>
    </v-tabs>

    <v-simple-table>
        <template v-slot:default>
            <thead>
            <tr>
                <th class="text-left">Winner</th>
                <th class="text-left">Total pot</th>
                <th class="text-left">Total profit</th>
            </tr>
            </thead>
            <tbody>
            <tr>
                <td>Karel Camo</td>
                <td>€ 23.40</td>
                <td>€ 8.23</td>
            </tr>
            <tr>
                <td>Tjabo Camo</td>
                <td>€ 13.40</td>
                <td>€ 10.28</td>
            </tr>
            <tr>
                <td>Karel Camo</td>
                <td>€ 144.40</td>
                <td>€ 45.30</td>
            </tr>
            </tbody>
        </template>
    </v-simple-table>
    </div>
</template>

<script>
export default {
  name: "LastGames"
};
</script>

<style scoped></style>
