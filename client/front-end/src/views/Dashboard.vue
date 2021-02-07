<template>
  <div class="dashboard">
  <nav class="menu">
  <ul>
    <li
     v-for="tab in tabs"
     :key="tab"
     :class="['menu__link', { 'menu__link--active': currentTab === tab }]"
     >
      <a href="#" @click.prevent="currentTab = tab">{{ tab }}</a>
      </li>
    <li>
      <button @click="changeWeek(-1)">-</button>
      {{ formattedDate }}
      <button @click="changeWeek(1)">+</button>
    </li>
  </ul>
</nav>

  <component :is="currentTabComponent" class="tab" :user-date="userDate"></component>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import Calendar from '@/components/Calendar.vue';
import Review from '@/components/Review.vue';
import Track from '@/components/Track.vue';

type Direction = 1 | -1;

export default defineComponent({
  name: 'Dashboard',
  components: {
    Calendar,
    Review,
    Track,
  },
  data() {
    return {
      currentTab: 'Review',
      tabs: ['Calendar', 'Track', 'Review'],
      userDate: new Date(),
    };
  },
  methods: {
    changeWeek(dir: Direction = 1): void {
      const n = this.userDate.setDate(this.userDate.getDate() + (dir * 7));
      this.userDate = new Date(n);
    },
  },
  computed: {
    currentTabComponent(): string {
      return `tab-${this.currentTab.toLowerCase()}`;
    },
    formattedDate(): string {
      return this.userDate.toDateString();
    },
  },
});
</script>
