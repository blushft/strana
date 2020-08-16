import store from 'store'
import AES from 'crypto-js/aes'

export default class Storage {
  constructor(options = {}) {
    this.valuePrefix = options.prefix || 'strana_enc_aes'
    this.key = options.key || 'STRANADATA'
  }

  encrypt(val) {
    return `${this.valuePrefix}${AES.encrypt(val, this.key).toString()}`
  }
}
