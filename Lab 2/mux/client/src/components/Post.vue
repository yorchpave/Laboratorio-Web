<template>
  <div>
    <h1>POST</h1>
    <input type="text" placeholder="ID" v-model="id">
    <input type="text" placeholder="Value" v-model="value">
    <input type="button" value="Create" @click="postItem" :disabled="id.length === 0 || value.length === 0">

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
    postItem() {
      const self = this
      axios({
        method: 'post',
        url: 'http://localhost:8080/item/',
        data: {
          id: self.id, value: self.value
        }
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