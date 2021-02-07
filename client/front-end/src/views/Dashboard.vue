<template>
  <div class="dashboard page">
    <nav class="menu">
    <ul>
      <li
      v-for="tab in tabs"
      :key="tab"
      :class="['menu__link', { 'menu__link--active': currentTab === tab }]"
      >
        <router-link :to="tab.path">{{ tab.label }}</router-link>
      </li>
      <li>
        <button @click="changeWeek(-1)">-</button>
        {{ formattedDate }}
        <button @click="changeWeek(1)">+</button>
      </li>
    </ul>
  </nav>
  <router-view/>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';

type Direction = 1 | -1;

export default defineComponent({
  name: 'Dashboard',
  data() {
    return {
      currentTab: 'Review',
      tabs: [
        {
          label: 'Calendar',
          path: '/calendar',
        },
        {
          label: 'Track',
          path: '/track',
        },
        {
          label: 'Review',
          path: '/review',
        },
      ],
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
