import axios, { AxiosInstance } from 'axios';

export let clientSingleton: AxiosInstance;

export const getClientInstance = (): AxiosInstance => {
  if (clientSingleton) {
    return clientSingleton;
  }

  const baseUrl = process.env.NEXT_PUBLIC_API_URL;

  clientSingleton = axios.create({
    baseURL: baseUrl,
  });

  return clientSingleton;
}

