<template>
  <div id="wrapper">
  <div id="top-nav">
    <form @submit.prevent="sendRequest">
      <div>
        <label for="userId">Buscar </label>
        <input for="text" id="userId" v-model="message">
        <button>Enviar</button>  
      </div>
    </form>

  </div>

  <div style="width: 600px; float: left;">
    <table>
      <td>Subject</td>
      <td>From</td>
      <td>To</td>
    </table>
    <table id="tableResponses">
      <tbody>
          <tr v-for="todo in todos" :key="todo">
          <td @click="getMessage(todo._source)">{{todo._source.Subject}}</td>
          <td>{{todo._source.From}}</td>
          <td>{{todo._source.To}}</td>


         
        </tr>

      </tbody>
    </table>
  </div>
  <div id="divMsg">{{msg}}</div>
</div>
</template>

<script>
  import axios from 'axios'
  export default{
    data(){
      return {
        todos:null,
        msg:null,
        message:null,
        userId: '',
        title: '',
        name: '',
      }
    },
    mounted(){
      this.getTodos();
    },
    methods:{
      getTodos(){
        console.log('codigo get TODOS.')
        //axios.get('http://jsonplaceholder.typicode.com/posts')
        axios.post('http://localhost:8081/api/getZincSearch?term=manipulated')
        .then(response => {
          console.log(response)
          this.todos = response.data
        })
        .catch(e => console.log(e))
      },

      getMessage(item){
        console.log('me hicieron clic '+item.Message)
        this.msg = item.Message
      },
      sendRequest() {
        axios.post("http://localhost:8081/api/getZincSearch?term="+this.message)
            .then(response => {
              console.log(response)
              this.todos = response.data
      })
      .catch(e => console.log(e))

    },
    createPost(){
      axios.post("http://localhost:8081/nameExample?name="+this.name)
      .then(response => {
        console.log(response);
      })
      .catch(e => console.log(e))
    }
    }
  }
</script>

<style>
#wrapper {
  width: 100%;
  background-color: #fff;
}
#top-nav {
  position: fixed;
  left: 0;
  right: 0;
  top: 0;
  height: 60px;
  width: 100%;
  background-color: green;
}
table, th, td {
  border: 1px solid;
}
#entryBox{
  padding-bottom: 3%;
  position: fixed;
}
#divMsg{
  margin-left: 1020px;
  position:fixed;
  overflow-y: auto;
  
}
#tableResponses{
  overflow-y: auto;
}

</style>