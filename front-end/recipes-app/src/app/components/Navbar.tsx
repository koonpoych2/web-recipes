import Link from "next/link"
import { AppBar, Toolbar, Typography, Box, TextField} from "@mui/material"

const Navbar = () => {
  return (
    <AppBar position="static" color="default" elevation={1}>
        <Toolbar sx={{ display: "flex", justifyContent: "space-between", px: 2 }}>
            <Box>
                <Typography variant="h6">
                    <Link href="/" style={{ textDecoration: "none", color: "inherit" }}>Logo</Link>
                </Typography>
            </Box>
            <Box sx={{ flexGrow: 1, mx: 4, maxWidth: "700px" }}>
                <TextField 
                    fullWidth
                    placeholder="Search"
                    variant="outlined"
                    size="small"
                    sx={{
                        borderRadius: "999px",
                        "& .MuiOutlinedInput-root": {
                          borderRadius: "999px",
                        },
                    }}
                />
            </Box>
            <Box>
                <Typography variant="h6">Profile</Typography>
            </Box>
        </Toolbar>
    </AppBar>
  )
}
export default Navbar