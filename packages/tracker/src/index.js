import Tracker from './tracker'

if (typeof define !== 'undefined' && define.amd) {
  define(function() {
    'use strict'
    return Tracker
  })
} else if (typeof module !== 'undefined' && module.exports) {
  module.exports = Tracker
} else {
  window.strana = Tracker
}
