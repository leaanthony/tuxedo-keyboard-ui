<template>
  <div id="app">
    <Sketch v-model="colors" style="width:200px; height:300px"/>
  </div>
</template>

<script>
import "./assets/css/main.css";

import { Sketch } from 'vue-color';

export default {
  name: "app",
  components: {
    Sketch
  },
  data() {
    return {
      colors: "#ffffff"
    }
  },
  watch: {
    colors: function (val) {
      let rgba = val.rgba;
      window.backend.updateKeyboardLights(rgba.r, rgba.g, rgba.b, rgba.a);
    }
  }
};
</script>
