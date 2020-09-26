<template>
  <v-card>
    <v-card-text dense>
      <v-row align="center" justify="start">
        <v-col>
          <div>
            <span class="display-1 mr-1">{{ value }}</span>
            <span class="caption">
              <slot name="value-stat"></slot>
            </span>
          </div>
          <div class="subtitle text--secondary">{{ label }}</div>
          <div class="caption text--secondary pt-2">
            <v-icon size="14" class="mr-1">mdi-clock</v-icon>Updated:
            <span class="text--white text--primary">1 hr ago</span>
          </div>
        </v-col>
      </v-row>
    </v-card-text>
  </v-card>
</template>

<script>
export default {
  name: 'StatCard',
  props: {
    value: {
      type: [String, Number],
      required: true,
    },
    label: {
      type: String,
      required: true,
    },
    trend: {
      type: String,
      default: 'even',
    },
    updated: {
      type: Date,
      default: null,
    },
  },
}
</script>
