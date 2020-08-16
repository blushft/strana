import Queue from '@segment/localstorage-retry'
import axios from 'axios'

const options = {
  maxRetryDelay: 360000,
  minRetryDelay: 1000,
  backoffMult: 2,
  maxRetryAttempts: 10,
  maxItems: 100,
}

class EventQueue {
  constructor(config) {
    this.eventBuffer = []
    this.trackingId = config.trackingId
    this.url = config.url

    this.api = axios.create({
      baseURL: this.url,
    })

    this.queue = new Queue('strana', options, (item, done) => {
      this.processEvent()
    })

    this.queue.start()
  }

  _setHeaders() {}

  processEvent() {}

  queue(event) {
    btoa
  }
}
