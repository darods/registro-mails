<template>
  <div id="wrapper">
  <nav id=navbar class="navbar navbar-expand-lg navbar-light bg-light">
    <a class="navbar-brand">Prueba técnica Truora</a>
    
    <div class="collapse navbar-collapse" id="navbarSupportedContent">
    <form class="form-inline my-2 my-lg-0" @submit.prevent="sendRequest">
      <div>
        <input class="form-control mr-sm-2" type="search" placeholder="Search" aria-label="Search" v-model="message">
        <button class="btn btn-outline-success my-2 my-sm-0" type="submit">Buscar</button>
      </div>
    </form>
    </div>
  </nav>



  <div id=tablediv class="table-wrapper-scroll-y my-custom-scrollbar">

    <table id="tableResponses" class="table table-striped table-bordered table-sm">
      <thead>
      <td>Subject</td>
      <td>From</td>
      <td>To</td>
    </thead>
      <tbody>
          <tr v-for="todo in todos" :key="todo">
          <td @click="getMessage(todo._source)">{{todo._source.Subject}}</td>
          <td>{{todo._source.From}}</td>
          <td>{{todo._source.To}}</td>


         
        </tr>

      </tbody>
    </table>
  </div>
  <div id= msgdiv class="overflow-auto p-3 bg-light">{{msg}}</div>

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
  height: 100%;
  background-color: #fff;
}
#navbar {
  position: fixed;
  width: 100%;

}
table, th, td {
  border: 1px solid;
}

#msgdiv{
  margin-left: 60%;
  margin-top: 5%;
  max-width: fit-content; 
  max-height: 800px;
  position: fixed;
  
}
#tableResponses{
  overflow-y: auto;
  width: fit-content;
  height: fit-content;
}

#tablediv{
  margin-top: 5%;
  width: 60%; 
  float: left;
  height: 800px;
}

.my-custom-scrollbar {
position: relative;
height: 200px;
overflow: auto;
}
.table-wrapper-scroll-y {
display: block;
}

</style>