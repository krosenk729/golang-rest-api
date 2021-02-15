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
        <button @click="decreaseWeek()">-</button>
        {{ formattedDate }}
        <button @click="increaseWeek()">+</button>
      </li>
    </ul>
  </nav>
  <router-view/>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { mapState, mapGetters } from 'vuex';

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
    };
  },
  methods: {
    increaseWeek() {
      this.$store.dispatch('increaseWeek');
    },
    decreaseWeek() {
      this.$store.dispatch('decreaseWeek');
    },
  },
  computed: {
    ...mapState({
      userDate: 'date',
    }),
    ...mapGetters([
      'formattedDate',
    ]),
  },
});
</script>
