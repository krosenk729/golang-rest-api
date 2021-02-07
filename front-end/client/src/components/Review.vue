<template>
  <div class="review">
    <h1>Review Your Year</h1>
    <br>
    <h2>Top Tags</h2>
    <ul class="chevron-list">
      <li v-if="!tagsList.length">loading...</li>
        <li v-for="{ tag, count } of tagsList" v-bind:key="tag">
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
      <li v-if="!verbList.length">loading...</li>
        <li v-for="{ verb, count } of verbList" v-bind:key="verb">
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
import * as Util from '../shared/utils';
import { mockData } from '../shared/mocks';

export default defineComponent({
  name: 'Review',
  props: [''],
  data() {
    return {
      tagsList: [],
      verbList: [],
      selectedTag: undefined,
      selectedVerb: undefined,
    }
  },
  async created() {
      interface CountHash {
        tag?: string; 
        verb?: string; 
        count: number;
    };
    this.data = await Promise.resolve(mockData);
    const {
      tagsHash,
      verbHash
    } = this.data.reduce(({ tagsHash, verbHash }, { tags, verb }) => {

      tags.forEach((t) => {
        tagsHash[t] = tagsHash[t] || { tag: t, count: 0 };
        tagsHash[t].count += 1;
      });
      verbHash[verb] = verbHash[verb] || { verb, count: 0 };
      verbHash[verb].count += 1;
      return { tagsHash, verbHash };
    }, { tagsHash: {}, verbHash: {} });
    this.tagsList = Object.values(tagsHash).sort((a: CountHash, b: CountHash) => b.count - a.count);
    this.verbList = Object.values(verbHash).sort((a: CountHash, b: CountHash) => b.count - a.count);
  },
  methods: {
    expandTag(tag: string?): void {
      this.selectedTag = (!tag || tag === this.selectedTag) ? undefined : tag;
    },
    expandVerb(verb: string?): void {
      this.selectedVerb = (!verb || verb === this.selectedVerb) ? undefined : verb;
    },
  },
  computed: {
    selectedTagEntries(): string[] {
      return this.selectedTag ? this.data.filter((e) => e.tags.includes(this.selectedTag)) : [];
    },
    selectedVerbEntries(): string[] {
      return this.selectedVerb ? this.data.filter((e) => e.verb === this.selectedVerb) : [];
    },
  }
});
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
</style>
