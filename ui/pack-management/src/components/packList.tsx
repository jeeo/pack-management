"use client"

import { Pack, deletePack, listPacks, updatePack } from '@/service/pack'
import { Grid } from '@mui/material'
import { useEffect, useState } from 'react'
import EditablePack from './editablePack'
import AddPackButton from './addPack'


export const PackList = () => {
  const [packs, setPacks] = useState<Pack[]>([])

  const loadList = () => {
    listPacks().then((packs) => {
      setPacks(packs)
    })
  }

  useEffect(() => {
    loadList()
  }, [])

  const updateItem = (id: string, amount: number) => {
    updatePack(id, amount).then(() => {
      setPacks(packs.map((pack) => {
        if (pack.id === id) {
          return { ...pack, amount }
        }
        return pack
      }))
    })
  }

  const deleteItem = (id: string) => {
    deletePack(id).then(() => {
      setPacks(packs.filter((pack) => pack.id !== id))
    })
  }


  return (
    <Grid container direction={'column'} alignContent={'center'}>
    {packs?.map((pack) => {
      return (
        <EditablePack key={pack.id} pack={pack} actions={{updateItem, deleteItem}}/>
      )
    })}
      <Grid item container justifyContent="flex-end">
        <AddPackButton onAddPack={loadList}/>
      </Grid>
    </Grid>
  )

}