import axios from 'axios'

export const API = axios.create({
  baseURL: process.env.NEXT_PUBLIC_API_URL,
})

export const getSalaryFromApiByUuid = (uuid, exchange = false) =>
  new Promise((resolve, reject) => {
    API.get(`/salary/${uuid}${exchange ? '/exchange' : ''}`)
      .then((response) => resolve(response.data))
      .catch(reject)
  })

export const createSalaryOnApi = (request) =>
  new Promise((resolve, reject) => {
    API.post('/salary', request)
      .then((response) => resolve(response.data))
      .catch(reject)
  })
