export default class Plugin {
  constructor({ name, summary, components, routes, modules }) {
    this.name = name
    this.summary = summary

    this.components = components
    this.routes = routes
    this.modules = modules
  }

  mountComponents(Vue) {
    for (const c in this.components) {
      Vue.component(c)
    }
  }
}
