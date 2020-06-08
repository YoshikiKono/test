<template>
  <section class="container">
    <h1>Todo List</h1>
    <input type="text" name="content" v-model="content"/>
    <div>
      <button @click="addTodo">登録</button>
    </div>
    <ul>
      <li v-for="(todo, index) in todos" :key="index">
        <span>{{ todo.content }}</span><span>({{ todo.created }})</span><button @click="deleteTodo(todo)">×</button>
      </li>
    </ul>
  </section>
</template>

<script>
import {mapState} from 'vuex';

export default {
  data: function() {
    return {
      content: '',
    }
  },
  computed: {
    ...mapState(['todos']),
  },
  methods: {
    addTodo: function() {
      this.$store.commit('addTodo', {content: this.content});
      this.content = '';
    },
    deleteTodo: function(todo) {
      this.$store.commit('deleteTodo', todo)
    }
  }
}
</script>