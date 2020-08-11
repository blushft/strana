import { v4 as uuidv4 } from 'uuid'

const validEvents = {
  identify: 'identify',
  pageview: 'pageview',
  screenview: 'screenview',
  structured: 'structured',
}

export default class Event {
  /**
   *
   * @param {string} event - Event type to create
   * @param {Object} config - Event config
   *
   */
  constructor(event, config = {}) {
    this.tid = config.trackingId || ''

    this.sid = config.sessionId || ''
    this.uid = config.userId || ''
    this.gid = config.groupId || ''
    this.eid = config.eventId || uuidv4()

    this.news = this.sid !== '' ? true : false
    this.newu = this.uid !== '' ? true : false

    this.e = event

    this.collect()
  }

  collect() {
    switch (this.e) {
      case 'pageview':
        return this.collectPageview()
    }
  }

  collectPageview() {}
}
