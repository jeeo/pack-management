import React, { FunctionComponent, useState } from 'react';
import Grid from '@mui/material/Grid';
import TextField from '@mui/material/TextField';
import IconButton from '@mui/material/IconButton';
import EditIcon from '@mui/icons-material/Edit';
import DeleteIcon from '@mui/icons-material/Delete';
import { Pack } from '@/service/pack';
import CheckIcon from '@mui/icons-material/Check';
import CloseIcon from '@mui/icons-material/Close';
import { Button, Dialog, DialogActions, DialogContent, DialogTitle } from '@mui/material';

const regexp = new RegExp('[a-zA-Z]');

type EditablePackProps = {
  pack: Pack;
  actions: {
    updateItem: (id: string, amount: number) => void;
    deleteItem: (id: string) => void;
  };
};

const EditablePack: FunctionComponent<EditablePackProps> = ({ pack, actions }) => {
  const [isEditing, setIsEditing] = useState(false);
  const [amount, setAmount] = useState(pack.amount);
  const [isDeleteDialogOpen, setIsDeleteDialogOpen] = useState(false);

  const onChangeAmount = (e: React.ChangeEvent<HTMLInputElement>) => {
    if (e.target.value === '') {
      setAmount(0);
      return
    }

    if (regexp.test(e.target.value)) {
      return
    }
  
    setAmount(parseInt(e.target.value));
  }

  const handleEditClick = () => {
    setIsEditing(!isEditing);
  };

  const handleDeleteClick = () => {
    setIsDeleteDialogOpen(true);
  };

  const handleConfirmClick = () => {
    if (amount != pack.amount) {
      actions.updateItem(pack.id, amount);
    }
    setIsEditing(false);
  };

  const handleCancelClick = () => {
    setIsEditing(false);
  };

  const handleDialogClose = () => {
    setIsDeleteDialogOpen(false);
  };

  const handleDeleteConfirm = () => {
    actions.deleteItem(pack.id);
    setIsDeleteDialogOpen(false);
  };


  return (
    <Grid container spacing={2} justifyContent={'center'}>
      {!isEditing ? (
        <>
          <Grid item xs={12} sm={6} md={4}>
            <TextField
              variant="standard"
              fullWidth
              value={pack.amount}
              disabled
            />
          </Grid>
          <Grid item xs={6} sm={3} md={2}>
            <IconButton
              aria-label="Edit"
              color="primary"
              onClick={handleEditClick}
            >
              <EditIcon />
            </IconButton>
          </Grid>
          <Grid item xs={6} sm={3} md={2}>
            <IconButton
              aria-label="Delete"
              color="secondary"
              onClick={handleDeleteClick}
            >
              <DeleteIcon />
            </IconButton>
          </Grid>
        </>
      ) : (
        <>
          <Grid item xs={12} sm={6} md={4}>
            <TextField
              variant='standard'
              fullWidth
              value={amount}
              onChange={onChangeAmount}
              onKeyDown={(e) => {
                if (e.key === 'Enter') {
                  handleConfirmClick();
                }
              }
            }
            />
          </Grid>
          <Grid item xs={6} sm={3} md={2}>
            <IconButton
              aria-label="Confirm"
              color="primary"
              onClick={handleConfirmClick}
            >
              <CheckIcon />
            </IconButton>
          </Grid>
          <Grid item xs={6} sm={3} md={2}>
            <IconButton
              aria-label="Cancel"
              color="secondary"
              onClick={handleCancelClick}
            >
              <CloseIcon />
            </IconButton>
          </Grid>
        </>
      )}
      <Dialog open={isDeleteDialogOpen} onClose={handleDialogClose}>
        <DialogTitle>Delete Confirmation</DialogTitle>
        <DialogContent>
          Are you sure you want to delete this entry?
        </DialogContent>
        <DialogActions>
          <Button onClick={handleDialogClose} color="primary">
            Cancel
          </Button>
          <Button onClick={handleDeleteConfirm} color="secondary">
            Delete
          </Button>
        </DialogActions>
      </Dialog>
    </Grid>
  );
};

export default EditablePack;
