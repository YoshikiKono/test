import Vuex from 'vuex';

const createStore = () => {
    return new Vuex.Store({
        state: () => ({
            todos: []
        }),
        mutations: {
            addTodo: function(state, obj) {
                var d = new Date();
                 state.todos.unshift({
                    content: obj.content,
                })
            },
            deleteTodo: function(state, obj) {
                for(let i = 0; i < state.todos.length; i++) {
                    const ob = state.todos[i];
                    if(ob.content == obj.content && ob.created == obj.created) {
                        alert('削除します ' + '"' + ob.content + '"');
                        state.todos.splice(i, 1);
                        return;
                    }
                }
            }
        }
    })
}

export default createStore;