import { createClient } from "redis";

const client = createClient({ url: "redis://redis:6379" });
client.connect();

export async function getCached(key: string) {
  return client.get(key);
}

export async function setCached(key: string, value: string) {
  await client.set(key, value, { EX: 30 });
}

// import { createClient } from "redis";

// export const redisClient = createClient({
//   url: "redis://redis:6379",
// });

// redisClient.connect();