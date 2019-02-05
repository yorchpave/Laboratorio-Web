<template>
  <div>
    <h1>DELETE</h1>
    <input type="text" placeholder="ID" v-model="id">
    <input type="button" value="Delete" @click="deleteItem" :disabled="id.length === 0">
    <input type="button" value="Delete All" @click="deleteItems">

    <p>{{ result }}</p>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'post',
  components: {
  },
  data() {
    return {
      id: '',
      value: '',
      result: ''
    }
  },
  methods: {
    deleteItem() {
      const self = this
      axios({
        method: 'delete',
        url: `http://localhost:8080/item/${this.id}`,
      }).then(function (res) {
        console.log(res)
        self.result = res.data
      });
    },
    deleteItems() {
      const self = this
      axios({
        method: 'delete',
        url: 'http://localhost:8080/item/',
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