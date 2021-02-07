export enum Verbs {
    created = 'Created',
    completed = 'Completed',
    initiative = 'Took Initiative',
    led = 'Led',
    participated = 'Participated',
}

export interface Entry {
    notes: string;
    verb: keyof typeof Verbs;
    tags: string[];
    date?: string|Date;
}
