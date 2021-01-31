'use strict';
// const faker = { lorem: {}, random: {} };

const getDayNum = (m = 1): number => m * 24 * 60 * 60 * 1000;
const getDayOfWeek = (orig = new Date(), i = 0) => {
	const d = new Date(+orig);
	const day = d.getDay();
	if (day === i) return d;
	d.setHours(-24 * (day - i)); 
	return d;
}

type Direction = 1 | -1;
const app = Vue.createApp({
  data() {
    return {
      currentTab: 'Review',
      tabs: ['Calendar', 'Track', 'Review'],
      userDate: new Date()
    }
  },
  methods: {
    changeWeek(dir: Direction = 1): void {
      const n = this.userDate.setDate(this.userDate.getDate() + (dir * 7));
      this.userDate = new Date(n);
    }
  },
  computed: {
    currentTabComponent(): string {
      return 'tab-' + this.currentTab.toLowerCase()
    },
    formattedDate(): string {
      return this.userDate.toDateString();
    }
  }
})

app.component('tab-calendar', {
  props: ['userDate'],
  template: `
  <div class="calendar">
    <button class="">Change weeks -</button>
    <table class="calendar__wrapper">
      <tr class="calendar__row calendar__row--days">
        <th class="calendar__cell" v-for="d in daysOTW">{{ d }}</th>
      </tr>
      <tr class="calendar__row" v-for="w in weeks">
        <td class="calendar__cell" v-for="(d, dayNum) in w">
          {{ d.getDate() }}
          <div v-if="dayNum == 0 || d.getDate() == 1">{{ getMonth(d) }}</div>
        </td>
        </tr>
    </table>
    <button class="">Change weeks +</button>
  </div>`,
  data() {
    return {
      daysOTW: ['Sunday','Monday','Tuesday','Wednesday','Thursday','Friday','Saturday'],
      monthsOTY: ['Jan','Feb','Mar','April','May','June','July','Aug', 'Sept', 'Oct', 'Nov', 'Dec'],
      numWeeksShown: 4,
    }
  },
  computed: {
    weeks() {
      const d = new Date(this.userDate);
      const base = d.setDate(d.getDate() - 14);
      const weeks = Array(this.numWeeksShown).fill('').map((_, i) => {
        const sunday = base + getDayNum(7 * i);
        return Array(7).fill('').map((__, j) => new Date(sunday + getDayNum(j)));
      });
      return weeks;
    }
  },
  methods: {
    progressWeek(): void {
      // emit event
    },
    getMonth(d: Date): string {
      return this.monthsOTY[d.getMonth()];
    }
  }
})

enum Verbs {
  created = 'Created',
  completed = 'Completed',
  initiative = 'Took Initiative',
  led = 'Led',
  participated = 'Participated',
}
interface Entry {
  notes: string;
  verb: keyof typeof Verbs;
  tags: string[];
}
app.component('tab-track', {
  props: ['userDate'],
  template: `
  <div class="track">
    <h1>Track Week: {{ range }}</h1>
    <p>Add entries for each thing you accomplished this week</p>
    <br>
    <fieldset v-for="(e, i) in entries" class="track--entry">
        <div class="lhs">
            <span>{{ i + 1 }}</span>
        </div>
        <div class="rhs">
            <label>
                <span class="sr-only">Verb</span>
                <select  v-model="e.verb">
                    <option></option>
                    <option v-for="(v, val) in verbs" :value="val">{{ v }}</option>
                </select>
             </label>
             <label>
                <span class="sr-only">Notes:</span>
                <textarea v-model="e.notes" placeholder="Did some amazing things I'll describe here..."
                    rows="3" style="resize:none"></textarea>
              </label>
            <div class="track--tags">
                <div v-for="(t, j) in e.tags" class="badge">
                    {{ t }} <button @click="deleteTag(e.tags, j)" aria-label="delete tag" title="Remove tag">✖</button>
                </div>
                <label for="tag" class="track--tags__new">
                    <span class="sr-only">Tag:</span>
                    <input id="tag" v-model="tag" placeholder="Add new tag..." @keyup.enter="addTag(e, tag)" />
                    <button @click="addTag(e, tag)">Add Tag</button>
                </label>
              </div>
        </div><!-- rhs -->
    </fieldset>
    <button @click="addEntry">Add Entry</button>
    <button @click="alert('TODO')">Save</button>
</div>
  `,
  data() {
    return {
      verbs: Verbs,
      entries: [{
        notes: 'Lorem ipsum',
        tags: ['a', 'b', 'c', 'd', 'e', 'f']
      }]
    }
  },
  computed: {
    range() {
      return getDayOfWeek(this.userDate, 0).toDateString() + ' - ' + getDayOfWeek(this.userDate, 6).toDateString()
    }
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
      this.entries.push({ notes: '', tags: [] })
    },
    deleteEntry(i): void {
      this.entries.splice(i, 1);
    }
  }
})

const mockData = [
  { notes: faker.lorem.sentence(), tags: Array(5).fill('').map(faker.vehicle.color), verb: Verbs.created, date: '12-12-2020' },
  { notes: faker.lorem.sentence(), tags: Array(5).fill('').map(faker.vehicle.color), verb: Verbs.created, date: '12-12-2020' },
  { notes: faker.lorem.sentence(), tags: Array(5).fill('').map(faker.vehicle.color), verb: Verbs.led, date: '12-12-2020' },
  { notes: faker.lorem.sentence(), tags: Array(5).fill('').map(faker.vehicle.color), verb: Verbs.participated, date: '12-12-2020' },
  { notes: faker.lorem.sentence(), tags: Array(5).fill('').map(faker.vehicle.color), date: '12-12-2020' }
]
app.component('tab-review', {
  props: [''],
  template: `
  <div class="review">
    <h1>Review Your Year</h1>
    <br>
    <h2>Top Tags</h2>
    <ul class="chevron-list">
      <li v-if="!tagsList.length">loading...</li>
        <li v-for="{ tag, count } of tagsList">
        <button 
          class="transparent"
          :aria-pressed="selectedTag === tag"
          @click="expandTag(tag)">
          <div :class="['chevron', { 'active': selectedTag === tag }]" aria-hidden="true"></div>
          {{ tag }} <span class="light-grey"> &emsp;({{ count }})</span>
          </button>
          <ul v-if="selectedTag === tag">
            <li v-for="e of selectedTagEntries">{{ e.date }}: {{ e.notes }}</li>
          </ul>
        </li>
    </ul>
    <h2>Your Verbs</h2>
    <ul class="chevron-list">
      <li v-if="!verbList.length">loading...</li>
        <li v-for="{ verb, count } of verbList">
        <button 
          class="transparent"
          :aria-pressed="selectedVerb === verb"
          @click="expandVerb(verb)">
          <div :class="['chevron', { 'active': selectedVerb === verb }]" aria-hidden="true"></div>
          {{ verb }} <span class="light-grey"> &emsp;({{ count }})</span>
          </button>
          <ul v-if="selectedVerb === verb">
            <li v-for="e of selectedVerbEntries">{{ e.date }}: {{ e.notes }}</li>
          </ul>
        </li>
    </ul>
</div>
  `,
  data() {
    return {
      tagsList: [],
      verbList: [],
      selectedTag: undefined,
      selectedVerb: undefined,
    }
  },
  async created() {
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
    this.tagsList = Object.values(tagsHash).sort((a, b) => b.count - a.count);
    this.verbList = Object.values(verbHash).sort((a, b) => b.count - a.count);
  },
  methods: {
    expandTag(tag: string?): void {
      this.selectedTag = (!tag || tag === this.selectedTag) ? undefined : tag;
    }
    expandVerb(verb: string?): void {
      this.selectedVerb = (!verb || verb === this.selectedVerb) ? undefined : verb;
    }
  },
  computed: {
    selectedTagEntries(): string[] {
      return this.selectedTag ? this.data.filter((e) => e.tags.includes(this.selectedTag)) : [];
    },
    selectedVerbEntries(): string[] {
      return this.selectedVerb ? this.data.filter((e) => e.verb === this.selectedVerb) : [];
    },
  }
})

app.mount('#app')