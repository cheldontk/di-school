<template>
<div>
  <Tutorial :message="welcome"/>
  <Header />
</div>
</template>

<script lang="ts">
export default {

  data() {
    return { welcome: "DI-School" }
  }
}

</script>
