<template>
  <div class="calendar">
    <button class="">Change weeks -</button>
    <table class="calendar__wrapper">
      <tr class="calendar__row calendar__row--days">
        <th class="calendar__cell" v-for="d in daysOTW" v-bind:key="d">{{ d }}</th>
      </tr>
      <tr class="calendar__row" v-for="w in weeks" v-bind:key="w">
        <td class="calendar__cell" v-for="(d, dayNum) in w" v-bind:key="d">
          {{ d.getDate() }}
          <div v-if="dayNum == 0 || d.getDate() == 1">{{ getMonth(d) }}</div>
        </td>
        </tr>
    </table>
    <button class="">Change weeks +</button>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import * as Util from '../shared/utils';

export default defineComponent({
  name: 'Caelndar',
  props: ['userDate'],
  data() {
    return {
      daysOTW: ['Sunday', 'Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday', 'Saturday'],
      monthsOTY: ['Jan', 'Feb', 'Mar', 'April', 'May', 'June', 'July', 'Aug', 'Sept', 'Oct', 'Nov', 'Dec'],
      numWeeksShown: 4,
    };
  },
  computed: {
    weeks(): Date[][] {
      const d = new Date(this.userDate);
      const base = d.setDate(d.getDate() - 14);
      const weeks = Array(this.numWeeksShown).fill('').map((_, i) => {
        const sunday = base + Util.getDayNum(7 * i);
        return Array(7).fill('').map((__, j) => new Date(sunday + Util.getDayNum(j)));
      });
      return weeks;
    },
  },
  methods: {
    progressWeek(): void {
      // emit event
    },
    getMonth(d: Date): string {
      return this.monthsOTY[d.getMonth()];
    },
  },
});
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
</style>
