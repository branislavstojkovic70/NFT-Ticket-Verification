import { createRoot } from 'react-dom/client'
import './index.css'
import Navbar from './components/navbar.tsx';
import { createBrowserRouter, RouterProvider } from 'react-router-dom';
import Login from './pages/auth/login.tsx';
import { CssBaseline, ThemeProvider, createTheme } from '@mui/material';
import { Toaster } from 'react-hot-toast';
import { LocalizationProvider } from '@mui/x-date-pickers';
import { AdapterDateFns } from '@mui/x-date-pickers/AdapterDateFns';
import Home from './pages/home/home.tsx';
import Registration from './pages/auth/registration.tsx';

const theme = createTheme({
  palette: {
      primary: {
          main: "#092147",
          light: "#174384",
          contrastText: "#F5F5F5"
      },
      secondary: {
          main: "#F5F5F5",
          contrastText: '#092147'
      },
      error: {
          main: "#DD3636"
      },
      success: {
          main: "#66FF77"
      },
      background: {
          default: "#F5F5F5",
      }
  },
  typography: {
      allVariants: {
          color: "#092147",
          fontFamily: "Poppins"
      },

  },
});

const router = createBrowserRouter([
  	{
		path:"/",
		element: <Navbar/>,
		children: [
			{
				path:"/home",
				element:
					<Home/>
			},
		],
	},
	{
        path:"/login", 
		element:
            <Login/>
    },
	{
        path:"/registration", 
		element:
            <Registration/>
    },
])

createRoot(document.getElementById('root')!).render(
	<ThemeProvider theme={theme} >
	<LocalizationProvider dateAdapter={AdapterDateFns}>

		<CssBaseline/>
		<RouterProvider router={router}/>
		<Toaster position="bottom-center"
			 toastOptions={{
				 success: {
					 style: {
						 background: theme.palette.success.main,
					 },
				 },
				 error: {
					 style: {
						 background: theme.palette.error.main,
						 color: "#F5F5F5"
					 },
				 },
			 }}
		/>
	</LocalizationProvider>
</ThemeProvider>
)
