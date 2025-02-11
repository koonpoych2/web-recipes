import Link from "next/link";
import Navbar from "./components/Navbar";
import Card from "@mui/material/Card";
import { Button, CardContent, Grid, Typography } from "@mui/material";
import Footer from "./components/Footer";
import Image from 'next/image';

const home_page = () => {
  return (
    <div>
      <Navbar />

      <Card sx={{ maxWidth: 700, m: 2 }}>
      <Grid container>
        <Grid item xs={12} sm={5}>
          <Image
            src="/"
            alt="Picture"
            width={240}
            height={240}
            style={{ width: '100%', height: '100%', objectFit: 'cover' }}
          />
        </Grid>
        <Grid item xs={12} sm={7}>
          <CardContent>
            <Typography variant="subtitle2" color="text.secondary" gutterBottom>
              Name
            </Typography>
            <Typography variant="h5" component="div" gutterBottom>
              Manu Name
            </Typography>

            <Typography variant="body2" color="text.secondary" paragraph>
              A perfectly grilled salmon with zesty lemon flavor, creating a healthy and delicious meal.
            </Typography>

            <Typography 
              variant="body2" 
              color="text.secondary"
              sx={{ display: 'flex', gap: 1, flexWrap: 'wrap', mt: 2 }}
            >
              <span>144 votes</span>
              <span>•</span>
              <span>5 minutes</span>
              <span>•</span>
              <span>20 comments</span>
              <span>•</span>
              <span>share</span>
            </Typography>
          </CardContent>

          <Button href="/"
            fullWidth 
            variant="contained" 
            sx={{ 
              bgcolor: 'primary.main',
              '&:hover': { bgcolor: 'primary.dark' },
              borderRadius: 0,
              py: 1.5
            }}
          >
            Read Full Recipe
          </Button>
        </Grid>
      </Grid>
    </Card>
    
    </div>
  );
};
export default home_page;
