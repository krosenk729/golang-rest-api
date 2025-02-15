import Faker from 'faker';
import { Entry, Verbs } from './models';

/* eslint-disable */
export const mockData: Entry[] = [
    { notes: Faker.lorem.sentence(), tags: Array(5).fill('').map(Faker.vehicle.color), verb: 'created', date: '12-12-2020' },
    // { notes: Faker.lorem.sentence(), tags: Array(5).fill('').map(Faker.vehicle.color), verb: Verbs.created, date: '12-12-2020' },
    // { notes: Faker.lorem.sentence(), tags: Array(5).fill('').map(Faker.vehicle.color), verb: Verbs.led, date: '12-12-2020' },
    // { notes: Faker.lorem.sentence(), tags: Array(5).fill('').map(Faker.vehicle.color), verb: Verbs.participated, date: '12-12-2020' },
    // { notes: Faker.lorem.sentence(), tags: Array(5).fill('').map(Faker.vehicle.color), date: '12-12-2020' }
];
/* eslint-enable */
