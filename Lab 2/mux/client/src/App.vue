<template>
  <div id="app">
    <post/>
    <get/>
    <put/>
    <delete/>
  </div>
</template>

<script>
import Post from './components/Post.vue'
import Get from './components/Get.vue'
import Put from './components/Put.vue'
import Delete from './components/Delete.vue'

export default {
  name: 'app',
  components: {
    Post, Get, Put, Delete
  }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
