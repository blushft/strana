import { v4 as uuidv4 } from 'uuid'

const validEvents = {
  identify: 'identify',
  group: 'group',
  session: 'session',
  pageview: 'pageview',
  screenview: 'screenview',
  action: 'action',
}

export default class Event {
  /**
   *
   * @param {string} event - Event type to create
   * @param {Object} config - Event config
   *
   */
  constructor(type, config = {}) {
    this.eid = uuidv4()
    this.tid = config.trackingId || store.get(tid)

    this.sid = ''

    const storedSID = store.get(sid)
    if (storedSID) {
      this.sid = storedSID
      this.news = false
    }

    if (config.sessionId && config.sessionId !== this.sid) {
      this.sid = config.sessionId
      this.news = true
      store.set(sid, this.sid)
    }

    if (this.sid === '') {
      this.sid = uuidv4()
      this.news = true
    }

    this.uid = config.userId || store.get(uid) || ''
    this.gid = config.groupId || store.get(gid) || ''
    this.cid = config.deviceId || store.get(cid) || ''

    this.e = type
  }

  setTrackingId(id) {
    this.tid = id
  }

  setSessionId(id) {
    this.sid = id
  }

  setUserId(id) {
    this.uid = id
  }

  setGroupId(id) {
    this.gid = id
  }

  setDeviceId(id) {
    this.cid = id
  }
}
