import Plugin from '@strana/plugin'
import HelloWorld from './components/HelloWorld.vue'

export default (options) => {
  const plugin = new Plugin({
    name: 'example',
    summary: 'an example plugin',
    components: {
      HelloWorld: HelloWorld,
    },
    routes: [
      {
        path: '/hello',
        name: 'hello',
        component: 'HelloWorld',
      },
    ],
  })

  return plugin
}
