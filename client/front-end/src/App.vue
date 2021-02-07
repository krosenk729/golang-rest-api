<template>
  <div id="nav">
    <router-link to="/">Home</router-link> |
    <router-link to="/about">About</router-link>
  </div>
  <!-- <router-view/> -->
</template>

<script lang="ts">
import { defineComponent } from 'vue';

type Direction = 1 | -1;

export default defineComponent({
  name: 'App',
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

<style lang="scss">
:root {
  --accent-lt-grey: #efefef;
  --accent-1: #b4d455;
  --text-blk: #666666;
  --text-blk-lt: #737373;
}
$base-margin: 8px;
$border-width: .5px;

* {
  font-family: -apple-system,BlinkMacSystemFont,Segoe UI,Helvetica,Arial,
  sans-serif,Apple Color Emoji,Segoe UI Emoji;
 color: var(--text-blk);
  &:focus {
    border: none !important;
    outline: none !important;
    box-shadow: inset 0 1px 1px rgba(0,0,0,.075), 0 0 8px rgba(102,175,233,.6) !important;
  }
}

.page {
  .menu {
    ul {
      display: inline-flex;
      list-style-type: none;
      padding-inline-start: 0;
      & > li {
        margin-right: .7em;
      }
    }
    &__link {

      &--active {
        font-weight: bold;
      }
    }
  }
}

.calendar {
  width: 100%;
  border-collapse: collapse;

  &__wrapper {
    width: 100%;
  }

  &__row {
    display: flex;
    flex-wrap: no-wrap;

    &--days .calendar__cell {
      height: 2.5 * $base-margin;
      background-color: var(--accent-lt-grey);
    }
  }
  &__cell {
    flex: 1;
    border: $border-width solid var(--accent-lt-grey);
    height: 8 * $base-margin;
    padding: 1em;
  }
}

.track {
  fieldset {
    border: 0;
    border-bottom: $border-width solid var(--accent-lt-grey);
  }
  input,
  textarea,
  select,
  option {
    @extend .w-100;
  }
  &--entry{
    display: grid;
    grid-template-columns: 30px 1fr;
    .lhs {}
    .rhs {
      display: grid;
      grid-template-columns: minmax(125px, 15%) 1fr;
      > :first-child {
        grid-row: span 2;
      }
    }
  }

  &--tags {
    display: flex;
    flex-direction: row;
    flex-wrap: wrap;
    &__new {
      @extend .w-100;
      flex: 1 0 100%;
      display: flex;
      > * {
        min-width: 145px;
      }
    }
  }
}

.review {
    .chevron-list {
      padding-inline-start: 0;
      list-style-type: none;
      li {
        margin-top: $base-margin;
        margin-bottom: $base-margin;
        border-bottom: $border-width solid var(--accent-lt-grey);
      }
      li > button {
        @extend .w-100;
        background: none;
        border: none;
        display: flex;
        justify-content: flex-start;
        align-content: baseline;
        padding: $base-margin;
      }
  }

  $chevron-skew: 48deg;
  .chevron {
    top: 100px;
    text-align: center;
    padding: .1em;
    margin-right: $base-margin;
    margin-bottom: $base-margin;
    height: .5em;
    width: 1.2em;
    transform: rotate(-90deg);

    &.active {
      transform: rotate(0deg) !important;
    }

    &:before {
      content: "";
      position: absolute;
      top: 0;
      left: 0;
      height: 100%;
      width: 51%;
      background: var(--accent-lt-grey);
      transform: skew(0deg, $chevron-skew);
    }
    &:after {
      content: "";
      position: absolute;
      top: 0;
      right: 0;
      height: 100%;
      width: 50%;
      background: var(--accent-lt-grey);
      transform: skew(0deg, -$chevron-skew);
    }
  }
}

.badge {
  margin: $base-margin;
  padding: $base-margin/2;
  border-radius: $base-margin;

  button {
    border: none;
    background: none;
  }

  $badge-colors:
    1 'cyan' #d0f0fd #04283f,
    /* 2 'raspberry' #ffdce5 #4c0c1c,*/
    2 'pink' #ff9eb7 #4c0c1c,
    3 'green' #d1f7c4 #0b1d05,
    4 'dark-orange' #ffa981 #6b2613,
    5 'purple' #cdb0ff #280b42,
    6 'apricot' #ffeab6 #3b2501;
  @each $i, $name, $bg-col, $tx-col in $badge-colors {
    &:nth-of-type(6n + #{$i}) {
      background-color: $bg-col;
      color: $tx-col;
    }
  }
}
.w-100 {
  width: 100%;
}
.light-grey {
  color: var(--text-blk-lt);
}
.sr-only {
    position: absolute;
    margin: -1px 0 0 -1px;
    padding: 0;
    display: block;
    width: 1px;
    height: 1px;
    font-size: 1px;
    line-height: 1px;
    overflow: hidden;
    clip: rect(0,0,0,0);
    border: 0;
    outline: 0;
}
</style>
