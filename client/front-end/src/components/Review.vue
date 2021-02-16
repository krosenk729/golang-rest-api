<template>
  <div class="review">
    <h1>Review Your Year</h1>
    <br>
    <h2>Top Tags</h2>
    <ul class="chevron-list">
      <li v-if="!lists?.tagsList.length">loading...</li>
        <li v-for="{ tag, count } of lists?.tagsList" v-bind:key="tag">
        <button
          class="transparent"
          :aria-pressed="selectedTag === tag"
          @click="expandTag(tag)">
          <div :class="['chevron', { 'active': selectedTag === tag }]" aria-hidden="true"></div>
          {{ tag }} <span class="light-grey"> &emsp;({{ count }})</span>
          </button>
          <ul v-if="selectedTag === tag">
            <li v-for="e of selectedTagEntries" v-bind:key="e">{{ e.date }}: {{ e.notes }}</li>
          </ul>
        </li>
    </ul>
    <h2>Your Verbs</h2>
    <ul class="chevron-list">
      <li v-if="!lists?.verbList.length">loading...</li>
        <li v-for="{ verb, count } of lists?.verbList" v-bind:key="verb">
        <button
          class="transparent"
          :aria-pressed="selectedVerb === verb"
          @click="expandVerb(verb)">
          <div :class="['chevron', { 'active': selectedVerb === verb }]" aria-hidden="true"></div>
          {{ verb }} <span class="light-grey"> &emsp;({{ count }})</span>
          </button>
          <ul v-if="selectedVerb === verb">
            <li v-for="e of selectedVerbEntries" v-bind:key="e">{{ e.date }}: {{ e.notes }}</li>
          </ul>
        </li>
    </ul>
</div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { mapState, mapGetters } from 'vuex';
import { Entry } from '../shared/models';

export default defineComponent({
  name: 'Review',
  data() {
    return {
      tagsList: [],
      verbList: [],
      selectedTag: undefined as undefined|string,
      selectedVerb: undefined as undefined|string,
    };
  },
  async created() {
    await this.$store.dispatch('refreshData');
  },
  methods: {
    expandTag(tag?: string): void {
      this.selectedTag = (!tag || tag === this.selectedTag) ? undefined : tag;
    },
    expandVerb(verb?: string): void {
      this.selectedVerb = (!verb || verb === this.selectedVerb) ? undefined : verb;
    },
  },
  computed: {
    ...mapState({
      entries: 'entries',
    }),
    ...mapGetters([
      'lists',
    ]),
    selectedTagEntries(): Entry[] {
      return this.selectedTag
        ? this.entries.filter((e: Entry) => e.tags.includes(this.selectedTag as string))
        : [];
    },
    selectedVerbEntries(): Entry[] {
      return this.selectedVerb
        ? this.entries.filter((e: Entry) => e.verb === this.selectedVerb)
        : [];
    },
  },
});
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
</style>
