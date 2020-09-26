<template>
  <v-card>
    <v-row align="center" justify="start">
      <v-col class="ml-4">
        <div class="caption text--secondary pb-2">{{ label }}</div>

        <div>
          <h3>
            {{ value }}
            <v-icon :color="trendVals.color">{{ trendVals.icon }}</v-icon>
          </h3>
        </div>
      </v-col>
    </v-row>
  </v-card>
</template>

<script>
export default {
  name: 'TrendWidget',
  props: {
    label: {
      type: String,
      required: true,
    },
    value: {
      type: [String, Number],
      required: true,
    },
    trend: {
      type: String,
      default: 'even',
    },
    upIcon: {
      type: String,
      default: 'mdi-arrow-up-bold',
    },
    upColor: {
      type: String,
      default: 'primary',
    },
    downIcon: {
      type: String,
      default: 'mdi-arrow-down-bold',
    },
    downColor: {
      type: String,
      default: 'warning',
    },
    evenIcon: {
      type: String,
      default: '',
    },
  },
  computed: {
    trendVals() {
      switch (this.trend) {
        case 'up':
          return {
            icon: this.upIcon,
            color: this.upColor,
          }
        case 'down':
          return {
            icon: this.downIcon,
            color: this.downColor,
          }
        case 'even':
          return {
            icon: this.evenIcon,
            color: this.evenColor,
          }
        default:
          return { icon: '', color: '' }
      }
    },
  },
}
</script>