<template>
  <AdminPage>
    <v-row>
      <v-col v-for="n in 4" :key="n">
        <TrendWidget label="Views" value="489" trend="up"></TrendWidget>
      </v-col>
    </v-row>
    <v-row>
      <v-col v-for="n in 3" :key="n">
        <StatCard label="Total Visitors" value="143">
          <template slot="value-stat"> <v-icon size="12" color="accent">mdi-triangle</v-icon>+2% </template>
        </StatCard>
      </v-col>
    </v-row>
    <v-row>
      <v-col v-for="z in 2" :key="z">
        <ChartCard v-model="chartVals" unit="sec" icon="mdi-heart-pulse" caption="info" />
      </v-col>
    </v-row>
  </AdminPage>
</template>

<script>
import AdminPage from '@/components/Admin/Page.vue'
import TrendWidget from '@/components/Widgets/Trend.vue'
import StatCard from '@/components/Widgets/Cards/Stat.vue'
import ChartCard from '@/components/Widgets/Cards/Chart.vue'
export default {
  name: 'Home',
  components: {
    AdminPage,
    TrendWidget,
    StatCard,
    ChartCard,
  },
  data: () => ({
    chartVals: [22, 16, 42, 69, 74, 66, 77, 89],
  }),
}
</script>
