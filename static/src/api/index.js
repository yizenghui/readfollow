
import axios from 'axios'

axios.defaults.baseURL = process.env.API_ROOT

export default {
  get(data, callback){
    axios.get(data).then(
        (response) => {
        callback(null, response.data)
    })
  },  
  post(url, callback){
    axios.post("/post","url="+encodeURIComponent(url)).then(
        (response) => {
        callback(null, response.data)
    })
  },
}

