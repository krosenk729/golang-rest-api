<template>
  <div class="track">
    <h1>Track Week: {{ range() }}</h1>
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
                          v-model="e.newTag"
                          placeholder="Add new tag..."
                          @keyup.enter="addTag(e, e.newTag)" />
                    <button @click="addTag(e, e.newTag)">Add Tag</button>
                </label>
              </div>
        </div><!-- rhs -->
    </fieldset>
    <button @click="addEntry">Add Entry</button>
    <button @click="saveForm">Save</button>
</div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { mapGetters, mapActions } from 'vuex';
import { Entry, Verbs } from '../shared/models';

interface EntryInputs extends Entry {
  newTag?: string|undefined;
}

export default defineComponent({
  name: 'Track',
  data() {
    return {
      verbs: Verbs,
      entries: [{
        notes: 'Lorem ipsum',
        tags: ['a', 'b', 'c', 'd', 'e', 'f'],
      }],
      tag: '',
    };
  },
  computed: {
    ...mapGetters([
      'userDate',
      'formattedDate',
      'weekStart',
      'weekEnd',
    ]),
  },
  methods: {
    ...mapActions([
      'addEntries',
    ]),
    range(): string {
      const start = this.weekStart.toDateString();
      const end = this.weekEnd.toDateString();
      return `${start} - ${end}`;
    },
    addTag(e: EntryInputs, tag: string): void {
      if (!tag.trim()) return;
      e.tags.push(tag.trim());
      e.newTag = '';
    },
    deleteTag(arr: string[], i: number): void {
      arr.splice(i, 1);
    },
    addEntry(): void {
      this.entries.push({ notes: '', tags: [] });
    },
    deleteEntry(i: number): void {
      this.entries.splice(i, 1);
    },
    saveForm(): void {
      this.addEntries({ date: this.userDate, newEntries: this.entries })
        .then(() => {
          this.$router.push('calendar');
        });
    },
  },
});
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
</style>
