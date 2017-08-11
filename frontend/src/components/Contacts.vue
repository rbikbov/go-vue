<template>
  <div class="contacts">
    <div class="contact" v-for="contact in contacts">
      <i :class="contact.icon"></i>
      {{ contact.title }}:
      <a :href="contact.prefix + contact.description">{{ contact.description }}</a>
      <button v-show="auth" @click="removeContact(contact.id)" title="Удалить" ><i class="fa fa-close"></i></button>
    </div>
    <a class="add" v-if="auth" @click="form = !form">Добавить контакт</a>
    <form v-if="form" @submit.prevent="createNewContact">

      <table>
        <tbody>
        <tr>
          <th>Выберите иконку:</th>
          <td>
            <i class="fa fa-facebook" @click="iconClass"></i>
            <i class="fa fa-facebook-official" @click="iconClass"></i>
            <i class="fa fa-google" @click="iconClass"></i>
            <i class="fa fa-odnoklassniki" @click="iconClass"></i>
            <i class="fa fa-skype" @click="iconClass"></i>
            <i class="fa fa-telegram" @click="iconClass"></i>
            <i class="fa fa-twitter" @click="iconClass"></i>
            <i class="fa fa-vk" @click="iconClass"></i>
            <i class="fa fa-instagram" @click="iconClass"></i>
            <i class="fa fa-whatsapp" @click="iconClass"></i>
            <i class="fa fa-phone" @click="iconClass"></i>
            <i class="fa fa-mobile" @click="iconClass"></i>
          </td>
        </tr>
        <tr>
          <th>Выберите размер:</th>
          <td>
            <i :class="newContact.iconClass" @click="iconSize('')"></i>
            <i :class="newContact.iconClass + ' fa-2x'" @click="iconSize('fa-2x')"></i>
            <i :class="newContact.iconClass + ' fa-3x'" @click="iconSize('fa-3x')"></i>
            <i :class="newContact.iconClass + ' fa-4x'" @click="iconSize('fa-4x')"></i>
            <i :class="newContact.iconClass + ' fa-5x'" @click="iconSize('fa-5x')"></i>
          </td>
        </tr>

        <tr>
          <th>Тип контакта:</th>
          <td><input type="text" v-model="newContact.title" placeholder=" Тип контакта" required /></td>
        </tr>
        <tr>
          <th>Классы иконки:</th>
          <td><input type="text" v-model="newContact.icon" placeholder=" Классы иконки" required /></td>
        </tr>
        <tr>
          <th>Префикс контакта:</th>
          <td><input type="text" v-model="newContact.prefix" placeholder=" Например: 'skype:' или 'tel:' (без кавычек)" required /></td>
        </tr>
        <tr>
          <th>Значение контакта:</th>
          <td><input type="text" v-model="newContact.description" placeholder=" 'http://vk.com/id1' или '+79373616295' (без кавычек)" required /></td>
        </tr>

        <tr>
          <th>Конечный вид:</th>
          <td class="final">
            <i :class="newContact.icon"></i>
            {{ newContact.title }}:
            <a :href="newContact.prefix + newContact.description">{{ newContact.description }}</a>
          </td>
        </tr>

        <tr>
          <th colspan="2"><input type="submit" value="Добавить контакт" /></th>
        </tr>
        </tbody>
      </table>


      <br />

      <div>
      </div>

    </form>
    <div v-if="error" style="color: red">{{ error }}</div>
  </div>
</template>

<style lang="sass">
.contacts
  text-align: center
  .contact
    height: 40px
    a
      transition: 0.3s
      &:hover
        text-decoration: underline
        font-size: 150%
        color: royalblue
        /*transform: scale(1.1)*/
        transition: 0.3s
        z-index: 2
    button
      cursor: pointer
      &:hover
        color: red
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
            width: 40%
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
            i
              cursor: pointer
              margin: 5px
            &.final
              i
                margin: 0

</style>

<script>
  import soso from '../service'

  export default {
  name: 'contacts',
    data () {
      return {
        contacts: null,
        newContact: {
          'title': '',
          'description': '',
          'icon': '',
          'iconClass': '',
          'iconSize': '',
          'prefix': ''
        },
        auth: false,
        email: null,
        form: false,
        error: null
      }
    },
    created () {
      this.getContacts();
      this.checkAuth()
    },
    methods: {
      getContacts(){
        soso.search('contacts').then(resp => {
          this.contacts = resp.data.list;
        })
      },
      removeContact(delContactId){
        soso.delete('contacts', {'id': +delContactId}).then(resp => {
        });
        this.getContacts()
      },
      createNewContact(){
        soso.create('contacts', {
          'title': this.newContact.title,
          'description': this.newContact.description,
          'icon': this.newContact.icon,
          'prefix': this.newContact.prefix
        }).then(resp => {
          this.newContact.title = null;
          this.newContact.description = null;
          this.newContact.icon = null;
          this.newContact.prefix = null;
          this.getContacts()
        })
      },
      checkAuth() {
        soso.search('user').then(resp => {
          this.auth = resp.data.auth;
          this.email = resp.data.email;
        })
      },
      iconClass(iconClass) {
        this.newContact.iconClass = iconClass.target.className;
        this.newContact.icon = this.newContact.iconClass + ' ' + this.newContact.iconSize
      },
      iconSize(iconSize) {
        this.newContact.iconSize = iconSize;
        this.newContact.icon = this.newContact.iconClass + ' ' + this.newContact.iconSize
      }
    },
  }
</script>