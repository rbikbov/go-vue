<template>
  <header>
    <router-link to="/" exact>
      <h1>KULIKOVA</h1>
      <h6>The godlike hairdresser</h6>
    </router-link>
    <nav class="container">
      <router-link :to="{ name: 'Hairstyles' }" exact>Прически</router-link>
      <router-link :to="{ name: 'Haircuts' }" exact>Стрижки</router-link>
      <router-link :to="{ name: 'Diplomas' }" exact>Дипломы</router-link>
      <router-link :to="{ name: 'Contacts' }">Контакты</router-link>
    </nav>
  </header>
</template>

<style lang="sass">
header
  margin-bottom: 15px
  color: grey
  text-align: center
  > a
    display: inline-block
    text-align: right
    margin: 40px auto
    color: black
    h1
      font-size: 80px
      font-weight: 100
      @media (max-width: 576px)
        font-size: 70px
    h6
      padding: 0
      margin: 0
      display: inline-block
      font-size: 14px
      font-weight: 400
  nav
    margin: 0 auto
    font-size: 20px
    font-weight: 300
    a
      display: inline-block
      color: grey
      padding: 0 10px
      border-bottom: 3px solid rgba(0,0,0,0)
      &.router-link-active
        border-bottom-color: black
        color: black
        font-weight: 400
      &:hover, &.focus
        color: black
</style>
