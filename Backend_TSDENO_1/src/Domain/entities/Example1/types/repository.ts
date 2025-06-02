import { Example1 } from "@example1";

interface IExample1Repository {
  collection: Example1[]
  findByEmail(Email: string): Promise<Example1>
  save(model: Example1): Promise<Example1>
}

export type { IExample1Repository }