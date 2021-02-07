<template>
  <div class="track">
    <h1>Track Week: {{ range }}</h1>
    <p>Add entries for each thing you accomplished this week</p>
    <br>
    <fieldset v-for="(e, i) in entries" class="track--entry" v-bind:key="i">
        <div class="lhs">
            <span>{{ i + 1 }}</span>
        </div>
        <div class="rhs">
            <label>
                <span class="sr-only">Verb</span>
                <select  v-model="e.verb">
                    <option></option>
                    <option v-for="(v, val) in verbs" :value="val" v-bind:key="v">{{ v }}</option>
                </select>
             </label>
             <label>
                <span class="sr-only">Notes:</span>
                <textarea v-model="e.notes"
                    placeholder="Did some amazing things I'll describe here..."
                    rows="3"
                    style="resize:none"></textarea>
              </label>
            <div class="track--tags">
                <div v-for="(t, j) in e.tags" class="badge" v-bind:key="j">
                    {{ t }} <button @click="deleteTag(e.tags, j)"
                                    aria-label="delete tag"
                                    title="Remove tag">✖</button>
                </div>
                <label for="tag" class="track--tags__new">
                    <span class="sr-only">Tag:</span>
                    <input id="tag"
                          v-model="tag"
                          placeholder="Add new tag..."
                          @keyup.enter="addTag(e, tag)" />
                    <button @click="addTag(e, tag)">Add Tag</button>
                </label>
              </div>
        </div><!-- rhs -->
    </fieldset>
    <button @click="addEntry">Add Entry</button>
    <button @click="alert('TODO')">Save</button>
</div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import * as Util from '../shared/utils';
import { Verbs } from '../shared/models';

export default defineComponent({
  name: 'Track',
  props: ['userDate'],
  data() {
    return {
      verbs: Verbs,
      entries: [{
        notes: 'Lorem ipsum',
        tags: ['a', 'b', 'c', 'd', 'e', 'f'],
      }],
    };
  },
  computed: {
    range() {
      const start = Util.getDayOfWeek(this.userDate, 0).toDateString();
      const end = Util.getDayOfWeek(this.userDate, 6).toDateString();
      return `${start} - ${end}`;
    },
  },
  methods: {
    addTag(e, tag): void {
      if (!tag.trim()) return;
      e.tags.push(tag.trim());
      this.tag = '';
    },
    deleteTag(arr, i): void {
      arr.splice(i, 1);
    },
    addEntry(): void {
      this.entries.push({ notes: '', tags: [] });
    },
    deleteEntry(i): void {
      this.entries.splice(i, 1);
    },
  },
});
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
</style>
