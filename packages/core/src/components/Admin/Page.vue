<template>
  <v-container>
    <slot name="header">
      <v-row>
        <v-col>{{ title }}</v-col>
      </v-row>
    </slot>
    <slot></slot>
  </v-container>
</template>

<script>
export default {
  name: 'AdminPage',
  props: {
    title: {
      type: String,
      default: 'Admin',
    },
  },
}
</script>
