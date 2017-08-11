<template>
  <div class="haircuts">
    <div class="haircuts_item" v-for="haircut in haircuts">
      <router-link :to="{ name: 'Haircut', params: { id: haircut.id } }">
        <img :src="haircut.photo" class="photo">
        <p class="description">{{ haircut.description }}</p>
      </router-link>
    </div>
    <br />
    <a v-if="auth" @click="form = !form" class="add">Добавить фото</a>
    <form v-if="form" id="file-form" enctype="multipart/form-data" action="http://127.0.0.1:8000/create_haircut" method="post" @submit.prevent="createNewHaircut">
      <input type="text" name="email" v-model="email" style="display: none" />
      <table>
        <tbody>
        <tr>
          <th>Название стрижки:</th>
          <td><input type="text" name="title" placeholder="Введите название стрижки" required /></td>
        </tr>
        <tr>
          <th>Описание стрижки:</th>
          <td><input type="text" name="description" placeholder="Введите описание стрижки" required /></td>
        </tr>
        <tr>
          <th>Цена стрижки:</th>
          <td><input type="text" name="price" placeholder="Введите цену стрижки" required /></td>
        </tr>
        <tr>
          <th>Фотография:</th>
          <td><input type="file" id="file-select" name="uploadfile" required /></td>
        </tr>
        <tr>
          <th colspan="2"><input id="upload-button" type="submit" value="Добавить" /></th>
        </tr>
        </tbody>
      </table>
    </form>
    <div v-if="error" style="color: red">{{ error }}</div>

  </div>
</template>

<style lang="sass">
.haircuts
  padding: 2%
  text-align: center
  .haircuts_item
    width: 100%
    /*@media (min-width: 576px)*/
      /*width: 100%*/
    @media (min-width: 768px)
      width: 50%
    @media (min-width: 992px)
      width: 33.333333%
    vertical-align: middle
    padding-left: 1%
    padding-right: 1%
    margin-bottom: 15px
    display: inline-block
    box-sizing: border-box
    a
      .photo
        width: 100%
        height: 500px
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
        margin-top: 10px
        margin-bottom: 20px
  a.add
    display: inline-block
    padding: 15px 30px
    border: 2px solid grey
    border-radius: 25px
    background-color: white
    color: grey
    cursor: pointer
    &:hover
      background-color: grey
      color: white
      font-weight: bold
      border-color: white
  form
    width: 100%
    table
      border: 2px solid black
      padding: 5px
      margin: 5px auto
      @media (min-width: 992px)
        width: 50%
      tbody
        tr
          th
            width: 50%
            text-align: right
            input
              width: 100%
              cursor: pointer
              &:hover
                color: blue
          td
            input
              width: 100%
              border: none
              padding: 3px 0

</style>

<script>
  import soso from '../service'

  export default {
  name: 'haircuts',
    data () {
      return {
        haircuts: null,
        auth: false,
        email: null,
        form: false,
        error: null
      }
    },
    created () {
      this.getHaircuts();
      this.checkAuth()
    },
    methods: {
      getHaircuts(){
        soso.search('haircuts').then(resp => {
          this.haircuts = resp.data.list;
        })
      },
      createNewHaircut(e){
        let formData = new FormData(e.target);
        this.$http.post(e.target.action, formData)
            .then(function (response) {
              if (response.status == 200) {
                console.log(response);
                e.target.reset();
                this.getHaircuts();
              }

            }, function (response) {
              // console.log(response)
            });
      },
      checkAuth() {
        soso.search('user').then(resp => {
          this.auth = resp.data.auth;
          this.email = resp.data.email;
        })
      }
    },
  }
</script>
