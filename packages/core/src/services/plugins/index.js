import layouts from '../../layouts'

export class PluginLoader {
  constructor() {
    this.plugins = {}
  }

  register(plugin) {
    this.plugins[plugin.name] = plugin
  }

  routes() {
    const rts = []
    for (const plug in this.plugins) {
      rts.push(createRoute(this.plugins[plug].routes, this.plugins[plug].components))
    }

    return rts
  }
}

function createRoute(def, comps) {
  const rt = {
    path: `/${def.name.toLower()}`,
    name: def.name,
    component: layouts[def.layout],
  }

  if (def.children.length > 0) {
    rt.children = []
    for (const child in def.children) {
      rt.children.push({
        path: `/${def.name.toLower()}/${child.path}`,
        name: child.path,
        component: comps[child.component],
      })
    }
  }

  return rt
}
