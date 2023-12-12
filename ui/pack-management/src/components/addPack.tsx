import React, { useState } from 'react';
import Grid from '@mui/material/Grid';
import TextField from '@mui/material/TextField';
import Button from '@mui/material/Button';
import AddIcon from '@mui/icons-material/Add';
import CheckIcon from '@mui/icons-material/Check';
import CloseIcon from '@mui/icons-material/Close';
import { IconButton } from '@mui/material';
import { createPack } from '@/service/pack';

const regexp = new RegExp('[a-zA-Z]');

interface AddPackButtonProps {
  onAddPack: () => void;
}

const AddPackButton: React.FC<AddPackButtonProps> = ({ onAddPack }) => {
  const [newAmount, setNewAmount] = useState(0);
  const [isAdding, setIsAdding] = useState(false);

  const handleAddClick = () => {
    setIsAdding(true);
  };

  const onChangeAmount = (e: React.ChangeEvent<HTMLInputElement>) => {
    if (e.target.value === '') {
      setNewAmount(0);
      return;
    }

    if (regexp.test(e.target.value)) {
      return;
    }

    setNewAmount(parseInt(e.target.value));
  };

  const handleConfirmClick = () => {
    createPack(newAmount).then(() => {
      onAddPack();
    })
    setNewAmount(0);
    setIsAdding(false);
  };

  const handleCancelClick = () => {
    setNewAmount(0);
    setIsAdding(false);
  };

  return (
    <Grid container spacing={1} alignItems="center" justifyContent={'center'}>
      {isAdding ? (
        <>
          <Grid item xs={8}>
            <TextField
              label="New Pack Amount"
              variant="outlined"
              fullWidth
              value={newAmount}
              onChange={onChangeAmount}
            />
          </Grid>
          <Grid item>
            <IconButton color="primary" onClick={handleConfirmClick}>
              <CheckIcon />
            </IconButton>
          </Grid>
          <Grid item>
            <IconButton color="secondary" onClick={handleCancelClick}>
              <CloseIcon />
            </IconButton>
          </Grid>
        </>
      ) : (
        <Grid item>
          <Button variant='contained' endIcon={<AddIcon />} onClick={handleAddClick}>
            Add new
          </Button>
        </Grid>
      )}
    </Grid>
  );
};

export default AddPackButton;
