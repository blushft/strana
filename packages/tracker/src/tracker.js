import { v4 as uuidv4 } from 'uuid'
import Event from './event'

export default class Tracker {
  constructor(tid, options = {}) {
    this.trackingID = tid
    this.options = options
  }

  emit(event, options) {
    if (document.visibilityState == 'prerender') return null

    const e = new Event(event, options)
  }
}

function setCookie(name, value) {
  const expDate = new Date()
  expDate.setTime(exp.getTime() + 3 * 365 * 24 * 60 * 60 * 1000)

  const expires = `; expires=${expDate.toUTCString()}`
  document.cookie = `${name}=${value || ''}${expires}; samesite=strict; path=/`
}

function getCookie(name) {
  const cookies = document.cookie ? document.cookie.split('; ') : []

  for (const c of cookies) {
    const parts = c.split('=')
    if (decodeURIComponent(parts[0]) !== name) {
      continue
    }

    const result = parts.slice(1).join('=')
    return decodeURIComponent(result)
  }

  return null
}

function getUrl() {
  return location.protocol + '//' + location.hostname + location.pathname + location.search
}

function getReferrerFromQuery() {
  const res = location.search.match(/[?&](ref|source|utm_source)=([^?&]+)/)
  return res ? res[2] : null
}

function getUserData() {
  const data = JSON.parse(getCookie('strana_user_data'))

  if (data) {
    return data
  }

  const userData = {
    uid: uuidv4(),
    gid: null,
    refr: window.document.referrer,
  }

  setCookie('strana_user_data', JSON.stringify(userData))

  return userData
}
