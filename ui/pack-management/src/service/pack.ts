import { getClientInstance } from './http_client';

export type Pack = {
  id: string
  amount: number
}

export const listPacks = async (): Promise<Pack[]> => {
  const client = getClientInstance();
  const response = await client.get('/package')
  if (response.status !== 200) {
    throw new Error('Failed to fetch packs', response.data);
  }

  const data = response.data

  return data.packs;
}

// still wrong
export const createPack = async (amount: number): Promise<Pack> => {
  const client = getClientInstance();
  const response = await client.post('/package', { amount })
  if (response.status !== 201) {
    throw new Error('Failed to create pack', response.data);
  }

  const data = response.data

  return data.pack;
}

export const updatePack = async (id: string, amount: number): Promise<Pack> => {
  const client = getClientInstance();
  const response = await client.put(`/package/${id}`, { amount })
  if (response.status !== 200) {
    throw new Error('Failed to update pack', response.data);
  }

  
  const data = response.data

  return data.pack;
}

export const deletePack = async (id: string): Promise<void> => {
  const client = getClientInstance();
  const response = await client.delete(`/package/${id}`)
  if (response.status !== 200) {
    throw new Error('Failed to delete pack', response.data);
  }
}