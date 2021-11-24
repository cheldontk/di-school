<template>
  <div v-if="coffe">
    {{ coffe }}
  </div>
</template>

<script lang="js">
import axios from 'axios';

    export default {
        data() {
            this.coffe = undefined;
        },
        async created () {
            this.coffe = await this.getData().file;
        },
        methods: {
            async getData() {
                return await axios.get('https://coffee.alexflipnote.dev/random.json');
            }
        },
    }
</script>