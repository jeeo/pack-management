"use client"

import React, { useState } from 'react';
import TextField from '@mui/material/TextField';
import Button from '@mui/material/Button';
import { Dialog, DialogActions, DialogContent, DialogTitle, Grid } from '@mui/material';
import { calculateOrder, OrderPack } from '@/service/order';

const regexp = new RegExp('[a-zA-Z]');


const CalculateOrder = () => {
  const [orderAmount, setOrderAmount] = useState(0);
  const [orderPacks, setOrderPacks] = useState<OrderPack[] | null>(null);
  const [isDialogOpen, setIsDialogOpen] = useState(false);


  const onChangeAmount = (e: React.ChangeEvent<HTMLInputElement>) => {
    if (e.target.value === '') {
      setOrderAmount(0);
      return
    }

    if (regexp.test(e.target.value)) {
      return
    }
  
    setOrderAmount(parseInt(e.target.value, 10));
  }

  const handleCalculateClick = () => {
    calculateOrder(orderAmount).then(result => {
      console.log(result)
      setOrderPacks(result)
      setIsDialogOpen(true);
    })
  };

  const handleDialogClose = () => {
    setOrderPacks([]);
    setIsDialogOpen(false);
  };


  return (
    <>
    <Grid item container spacing={2} alignItems="center">
      <Grid item>
        <TextField
          label="Order amount"
          variant="outlined"
          value={orderAmount}
          onChange={onChangeAmount}
        />
      </Grid>
      <Grid item>
        <Button variant="contained" onClick={handleCalculateClick}>
          Calculate
        </Button>
      </Grid>
    </Grid>
    <Dialog open={isDialogOpen} onClose={handleDialogClose}>
        <DialogTitle>Order Packs</DialogTitle>
        <DialogContent>
          {orderPacks && (
            <ul>
              {orderPacks.map((orderPack, index) => (
                <li key={index}>
                  {`Pack: ${orderPack.pack.amount}, Quantity: ${orderPack.quantity}`}
                </li>
              ))}
            </ul>
          )}
        </DialogContent>
        <DialogActions>
          <Button onClick={handleDialogClose} color="primary">
            Close
          </Button>
        </DialogActions>
      </Dialog>
    </>
  );
};

export default CalculateOrder;