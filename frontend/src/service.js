import SoSo from 'soso-client'
import auth from './auth'

var soso = new SoSo('ws://localhost:8000/soso')

soso.log = true
soso.logData = true

soso.middleware.beforeSend( resp => {
  resp.other["authToken"] = auth.token
} )

soso.handle("auth", "SUCCESS", function(resp) {

  if ( resp.data.token ) {
    auth.setToken( resp.data.token )
  }

  if ( window.authWindow !== undefined ) {
    window.authWindow.close()
    window.authWindow = undefined
  }

})

soso.handle("auth", "ERROR", function(resp) {

  if ( window.authWindow !== undefined ) {
    window.authWindow.close()
    window.authWindow = undefined
  }
  
})

export default soso