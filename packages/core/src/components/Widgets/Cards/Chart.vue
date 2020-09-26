<template>
  <v-card>
    <v-card-title>
      <v-icon color="primary" class="mr-12" size="64">{{ icon }}</v-icon>
      <v-row align="start">
        <div class="caption grey--text text-uppercase">{{ caption }}</div>
        <div>
          <span class="display-2 font-weight-black" v-text="avg || '-'"> </span>
          <strong v-if="avg">{{ unit }}</strong>
        </div>
      </v-row>
    </v-card-title>
    <v-sheet color="transparent">
      <v-sparkline :key="String(avg)" :smooth="16" :line-width="3" :value="value" auto-draw></v-sparkline>
    </v-sheet>
  </v-card>
</template>

<script>
export default {
  name: 'ChartCard',
  props: {
    value: {
      type: Array,
      required: true,
    },
    unit: {
      type: String,
      required: true,
    },
    icon: {
      type: String,
      default: '',
    },
    caption: {
      type: String,
      default: '',
    },
  },
  data: () => ({
    chartVals: [],
  }),
  computed: {
    avg() {
      const sum = this.value.reduce((acc, cur) => acc + cur, 0)
      const len = this.value.length
      if (!sum && !len) return 0

      return Math.ceil(sum / len)
    },
  },
}
</script>

<style></style>
