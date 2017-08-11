<template>
  <footer>
    <div class="social">
      <a href="https://vk.com/kulikovamary">
        <i class="fa fa-vk fa-2x"></i>
        <br />
        <span class="text">Личный</span>
      </a>
      <a href="https://www.instagram.com/">
        <i class="fa fa-instagram fa-2x"></i>
        <br />
        <span class="text">Личный</span>
      </a>
      <a href="https://www.instagram.com/">
        <i class="fa fa-instagram fa-2x"></i>
        <br />
        <span class="text">Рабочий</span>
      </a>
    </div>
    <div class="phone">
      <a href="tel:79639639063">+7 (963) 963-90-63</a>
      <a href="tel:79388717900">+7 (938) 871-79-00</a>
      <p>Салон Чародейка г. Сочи, ул. Параллельная 9</p>
    </div>
    <i class="fa fa-google" @click='authByOauth("google")'></i>
  </footer>
</template>

<style lang="sass">
footer
  position: absolute
  bottom: 0
  left: 0
  right: 0
  text-align: center
  color: #454545
  .social
    margin-bottom: 10px
    a
      display: inline-block
      margin: auto 10px
      &:hover
        color: royalblue
      .text
        font-size: 12px
  .phone
    margin-top: 20px
    margin-bottom: 40px
    a, p
      margin: 10px
  i
    cursor: pointer
</style>

<script>
import soso from '../../service'

const popupCenter = ( url, title, w, h ) => {
  // Fixes dual-screen position             Most browsers    Firefox
  var dualScreenLeft = window.screenLeft != undefined ? window.screenLeft : screen.left;
  var dualScreenTop = window.screenTop != undefined ? window.screenTop : screen.top;

  var width = window.innerWidth ? window.innerWidth : document.documentElement.clientWidth ? document.documentElement.clientWidth : screen.width;
  var height = window.innerHeight ? window.innerHeight : document.documentElement.clientHeight ? document.documentElement.clientHeight : screen.height;

  var left = ((width / 2) - (w / 2)) + dualScreenLeft;
  var top = ((height / 2) - (h / 2)) + dualScreenTop;
  var newWindow = window.open(url, title, 'scrollbars=yes, width=' + w + ', height=' + h + ', top=' + top + ', left=' + left);

  // Puts focus on the newWindow
  if ( window.focus && newWindow !== undefined && newWindow.focus !== null && newWindow !== undefined ) {
    newWindow.focus();
  }

  return newWindow
};

export default {
name: 'footer',
data () {
    return {
      authURLs: {
        github: null,
        google: null,
      }
    }
  },

  mounted() {
    soso.request('auth', 'google').then( resp => {
      this.authURLs.google = resp.data.url
    })
  },

  methods: {
    authByOauth( type ) {
      if ( type !== 'github' && type !== 'google') {
        return
      }

      let url = this.authURLs[type];

      if ( url === null ) {
        alert("Try again.\n\nSomething went wrong.");
        return
      }

      window.authWindow = popupCenter(
          this.authURLs[type], "Authorization", 800, 600
      )

    }
  },

}
</script>
