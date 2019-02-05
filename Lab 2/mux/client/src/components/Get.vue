<template>
  <div>
    <h1>GET</h1>
    <input type="text" placeholder="ID" v-model="id">
    <input type="button" value="Get" @click="getItem" :disabled="id.length === 0">
    <input type="button" value="Get All" @click="getItems">

    <p>{{ result }}</p>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'get',
  components: {
  },
  data() {
    return {
      id: '',
      result: ''
    }
  },
  methods: {
    getItem() {
      const self = this
      axios({
        method: 'get',
        url: `http://localhost:8080/item/${this.id}`,
      }).then(function (res) {
        console.log(res)
        self.result = res.data
      });
    },
    getItems() {
      const self = this
      axios({
        method: 'get',
        url: `http://localhost:8080/item/`,
      }).then(function (res) {
        console.log(res)
        self.result = res.data
      });
    }
  }
}
</script>

<style>
</style>