var app = new Vue({
  el: '#app',
  vuetify: new Vuetify(),
  data: {
    addText: "",
    list: [],
    absolute: 0,
  },
  methods: {
    addToDo: function(e) {
      //var added = alert('予定の登録をします');
      if (this.addText) {
        this.list.unshift({
            'text': this.addText,
            'id': this.absolute + 1
        });
        this.addText = '';
        this.absolute++;
      }
    },
    deleteToDo(id) {
      var deleteIndex = '';
      var check = confirm('予定を削除しますか？');
      if (check === true) { 
          this.list.some(function (value, index) {
              if (value.id === id) {
                  deleteIndex = index;
              }
          });
          this.list.splice(deleteIndex, 1);
      }
   },
  }
}).$mount('#app')