<template>
  <div class="haircut">
    <!--{{ $route }}-->
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

    <div class="item" v-if="haircut">
      <h2>{{ haircut.title }} ({{ haircut.price }})</h2>
      <div class="img">
        <img :src="haircut.photo">
        <button v-show="auth" @click="removeHaircut(haircut.id)" title="Удалить"><i class="fa fa-close fa-2x"></i></button>
      </div>
      <p>{{ haircut.description }}</p>
    </div>

    <hr>

    <div class="item" v-if="haircuts">
      <h2>Случайное из категории Стрижки</h2>
      <div class="haircuts_item" v-for="haircut in haircuts">
        <router-link :to="{ name: 'Haircut', params: { id: haircut.id } }" @click="getHaircut">
          <img :src="haircut.photo" class="photo">
          <p class="description">{{ haircut.description }}</p>
        </router-link>
      </div>
    </div>
  </div>

</template>

<style lang="sass">
.haircut
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
    .haircuts_item
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
  name: 'haircut',
    data () {
      return {
        haircut: null,
        haircuts: null,
        auth: false,
        status: null,
        error: null,
        loading: true
      }
    },
    created () {
      this.getHaircut();
      this.getRandomHaircuts();
      this.checkAuth();
    },
    watch: {
      '$route' (to, from) {
        this.getHaircut();
        this.getRandomHaircuts();
        this.checkAuth();
      }
    },
    methods: {
      getHaircut() {
        soso.request('haircuts', 'get', {id: +this.$route.params.id}).then(resp => {
          if (resp.data) {
            this.haircut = resp.data.object;
            this.loading = false
          } else {
            this.error = resp.log.user_msg;
            this.loading = false
          }
        })
      },
      getRandomHaircuts(){
        soso.request('haircuts', 'random', {id: +this.$route.params.id}).then(resp => {
          this.haircuts = resp.data.list
        })
      },
      removeHaircut(id){
        soso.delete('haircuts', {'id': +id}).then(resp => {
          if (resp.log.code_key == 200) {
            this.haircut = null;
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
