<template>
  <v-app>
    <AppNav />
    <AppBar />
    <v-main>
      <router-view />
    </v-main>
  </v-app>
</template>

<script>
import AppNav from '@/components/App/Nav.vue'
import AppBar from '@/components/App/Bar.vue'
export default {
  name: 'Admin',
  components: {
    AppBar,
    AppNav,
  },
}
</script>
