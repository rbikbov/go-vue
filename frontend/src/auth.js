import ls from 'local-storage'
import jwt_decode from 'jwt-decode';

class Auth {
  constructor() {
    this.isAuth = false
    this.isAnonymous = true
    this.userID = null
    this.userName = null
    this.tokenData = null
    this.token = null

    this.loadFromLS()
  }

  loadData( token ) {
    
    try {

      const td = jwt_decode( token )

      this.tokenData = td
      this.isAuth = true
      this.isAnonymous = td.isAnonymous
      this.userID = td.uid
      this.userName = td.username
      this.token = token

      return td

    } catch ( err ) {
      console.warn( 'incorrect token' )

      this.tokenData = null
      this.isAuth = false
      this.userName = null
      this.isAnonymous = true
      this.userID = null
      this.token = null

      return false
    }

  }

  loadFromLS() {
      
    var token = ls.get( 'authToken' )
    this.loadData( token )

  }

  setToken( token ) {
    if ( this.loadData( token ) ) {

      ls.set( 'authToken', token )

    }
  }

}

var auth = new Auth()

window.auth = auth

export default auth