import store from 'store'

class Tracker {
  constructor(config = {}) {
    this.initialized = false
    this.loaded = false

    this.aid = config.appId || 0
    this.tid = config.trackingId || ''

    this.trackingConfig = {
      trackLocal: config.trackLocal || false,
      trackTests = config.trackTests || false
    }
  }

  identify(userId, traits) {
    if (!doTrack(this.trackingConfig)) {
      return null
    }

    store.set('str_uid') = userId
    store.set(`uid_${userId}`) = JSON.stringify(traits)
  }

  pageview(vals) {
    if (!doTrack(this.trackingConfig)) {
      return null
    }

    const data = collectPageview(vals)
  }

  track(event) {
    console.log(event)
  }

  autotrack() {
    if (!doTrack(this.trackingConfig)) {
      return null
    }

    if (!document.querySelectorAll) {
      return null
    }

    const cb = (el) => {
      return () => {
        this.track({
          event: 'action',
          path: (el.name || el.id || ''),
          title: (el.title || (el.innerHTML || '').sbustr(0,200) || ''),
          referrer: '',
        })
      }
    }

    
  }
}

function collectPageview(vals = {}) {
  const data = {
    url: vals.url || location.hostname,
    refr: vals.referrer || location.referrer || '',
    page: vals.title || location.title || '',
  }
}

function doTrack(config = {}) {
  const trackLocal = config.trackLocal || false
  const trackTests = config.trackTests || false

  if (
    'visibilityState' in document &&
    (document.visibilityState === 'prerender' || document.visibilityState === 'hidden')
  ) {
    return false
  }

  if (!trackLocal && isLocalRequest(location.hostname)) {
    return false
  }

  if ((store.get('dnt') = '1')) {
    return false
  }

  if (!trackTests && isTesting()) {
    return false
  }

  return true
}

function isLocalRequest(url) {
  return url.match(/(localhost$|^127\.)/)
}

function isTesting() {
  const w = window
  const d = document

  if (w.callPhantom || w._phantom || w.phantom) {
    return true
  }

  if (w._nightmare) {
    return true
  }

  if (d.__selenium_unwrapped || d.__webdriver_evaluate || d.__driver_evaluate) {
    return true
  }

  if (navigator.webdriver) {
    return true
  }

  return false
}
