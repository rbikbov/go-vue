<template>
  <div class="diploma">
    <div v-if="loading">
      <h1>Загрузка...</h1>
      <i class="fa fa-spinner fa-pulse fa-5x"></i>
    </div>

    <div class="error" v-if="error">
      <h1>Error: {{ error }}</h1>
      <i class="fa fa-warning fa-5x"></i>
    </div>

    <div class="status" v-if="status">
      <h1>{{ status }}</h1>
      <i class="fa fa-warning fa-5x"></i>
    </div>

    <div class="item" v-if="diploma">
      <h2>{{ diploma.title }}</h2>
      <div class="img">
        <img :src="diploma.photo">
        <button v-show="auth" @click="removeDiploma(diploma.id)" title="Удалить"><i class="fa fa-close fa-2x"></i></button>
      </div>
      <p>{{ diploma.description }}</p>
    </div>

    <hr>

    <div class="item" v-if="diplomas">
      <h2>Случайное из категории Дипломы</h2>
      <div class="diplomas_item" v-for="diploma in diplomas">
        <router-link :to="{ name: 'Diploma', params: { id: diploma.id } }" @click="getDiploma">
          <img :src="diploma.photo" class="photo">
          <p class="description">{{ diploma.description }}</p>
        </router-link>
      </div>
    </div>
  </div>

</template>

<style lang="sass">
.diploma
  text-align: center
  color: black
  div.error
    color: red
  div.status
    color: green
  div.item
    h2, p
      margin: 0
      margin-bottom: 15px
    div.img
      display: inline-block
      position: relative
      /*@media (min-width: 1200px)*/
        /*width: 1200px*/
      img
        height: 100vh
        max-width: 100%
      button
        position: absolute
        top: 5px
        right: 5px
        cursor: pointer
        &:hover
          color: red
    .diplomas_item
      @media (min-width: 576px)
        width: 100%
      @media (min-width: 768px)
        width: 50%
      @media (min-width: 992px)
        width: 33.333333%
      @media (min-width: 1200px)
        width: 25%
      vertical-align: bottom
      padding-left: 1%
      padding-right: 1%
      margin-bottom: 15px
      display: inline-block
      box-sizing: border-box
      color: black
      a
        .photo
          width: 100%
          border-radius: 15px
          transition: 0.5s
          @media (max-width: 576px)
            border: 2px solid rgba(0,0,0,0)
          &:hover
            @media (max-width: 576px)
              border-color: black
            @media (min-width: 576px)
              box-shadow: 3px 3px 3px 3px rgba(0, 0, 0, 0.5)
              transform: scale(1.05)
              transition: 0.5s
              z-index: 2
        .description
          margin: 0
</style>

<script>
  import soso from '../service'

  export default {
  name: 'diploma',
    data () {
      return {
        diploma: null,
        diplomas: null,
        auth: false,
        status: null,
        error: null,
        loading: true
      }
    },
    created () {
      this.getDiploma();
      this.getRandomDiplomas();
      this.checkAuth();
    },
    watch: {
      '$route.params.id' () {
          this.getDiploma();
          this.getRandomDiplomas();
          this.checkAuth();
//        this.$router.go({
//          path: this.$route.name,
//          params: {id: +this.$route.params.id}
//        })
      }
    },
    methods: {
      getDiploma() {
        soso.request('diplomas', 'get', {id: +this.$route.params.id}).then(resp => {
          if (resp.data) {
            this.diploma = resp.data.object;
            this.loading = false
          } else {
            this.error = resp.log.user_msg;
            this.loading = false
          }
        })
      },
      getRandomDiplomas(){
        soso.request('diplomas', 'random', {id: +this.$route.params.id}).then(resp => {
          this.diplomas = resp.data.list
        })
      },
      removeDiploma(id){
        soso.delete('diplomas', {'id': +id}).then(resp => {
          if (resp.log.code_key === 200) {
            this.diploma = null;
            this.status = 'Успешно удалено'
          }
        });
      },
      checkAuth() {
        soso.search('user').then(resp => {
          this.auth = resp.data.auth;
        })
      }
    },
  }
</script>
