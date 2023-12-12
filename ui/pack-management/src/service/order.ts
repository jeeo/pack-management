import { getClientInstance } from './http_client'
import { Pack } from './pack'


export type OrderPack = {
  pack: Pack
  quantity: number
}

export const calculateOrder = async (amount: number): Promise<OrderPack[]> => {
  const client = getClientInstance();
  const response = await client.post(`/order/calculate`, {
    amount,
  })
  if (response.status !== 200) {
    throw new Error('Failed to calculate order', response.data);
  }

  const data = response.data
  
  return data['order_packs'];
}