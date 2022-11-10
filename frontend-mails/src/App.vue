<template>
  <div id="wrapper">
    <nav id=navbar class="navbar navbar-dark bg-dark">
      <div class="container-fluid">
        <a class="navbar-brand">Mamuro Email</a>
        <button class="btn btn-outline-success" type="button" @click="disminuir">Anterior</button>
        <button class="btn btn-outline-success" type="button" @click="aumentar">Siguiente</button>
        <p id="rango">mostrando datos [{{from}}, {{to}}]</p>
        <form class="d-flex" role="search" @submit.prevent="sendRequestReset">
          
          <input class="form-control me-2" type="search" placeholder="Buscar" aria-label="Buscar" v-model="message">
          <button class="btn btn-outline-success" type="submit">Buscar</button>
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
          <tr v-for="todo in todos" :key="todo" @click="getMessage(todo._source)">
          <td>{{todo._source.Subject}}</td>
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
        from: 0,
        to: 20
      }
    },
    mounted(){
      this.getTodos();
    },
    methods:{
      getTodos(){
        console.log('codigo get TODOS.')
        axios.post('http://localhost:8081/api/getZincSearch?term=manipulated&from=0&to=20')
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
      sendRequestReset() {
        this.from=0
        this.to=20
        this.sendRequest(this.from, this.to)

    },
    sendRequest(from, to) {
        axios.post("http://localhost:8081/api/getZincSearch?term="+this.message+"&from="+from+"&to="+to)
            .then(response => {
              console.log(response)
              this.todos = response.data
      })
      .catch(e => console.log(e))

    },
      aumentar(){
        this.from += 20
        this.to +=20
        console.log("from: "+this.from+" to: "+this.to)
        this.sendRequest(this.from, this.to)
      },
      disminuir(){
        if(this.from != 0){
          this.from -= 20
          this.to -=20
        }else{
          this.from = 0
          this.to =20
        }
        
        console.log("from: "+this.from+" to: "+this.to)
        this.sendRequest(this.from, this.to)
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
  height: 10%;
  background-color: #e3f2fd;

}
#rango{
  color: #fff;
}
table, th, td {
  border: 1px solid;
}

#tablediv tr:hover{
  background-color: #e3f2fd;
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