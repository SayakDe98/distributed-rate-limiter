const { createClient } = require("redis");

export const redisClient = createClient({ url: "redis://redis:6379" });
redisClient.connect();

export async function getCached(key: string) {
  return redisClient.get(key);
}

export async function setCached(key: string, value: string) {
  await redisClient.set(key, value, { EX: 30 });
}