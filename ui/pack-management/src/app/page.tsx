import CalculateOrder from '@/components/calculateOrder';
import { PackList } from '@/components/packList';
import { Button, Grid, Paper, TextField, Typography } from '@mui/material';

export default function Home() {


  return (
    <Grid container sx={{height: '100vh'}} justifyContent={'space-between'}>
      <Grid item md={3}>
        <Grid container>
          <Typography variant='h4'> Available Packages </Typography>
          <PackList />
        </Grid>
      </Grid>
      <Grid item md={8}>
        <Grid container>
          <Typography variant='h4'> Calculate Packages </Typography>
          <CalculateOrder />
        </Grid>
      </Grid>
    </Grid>
  );
}
