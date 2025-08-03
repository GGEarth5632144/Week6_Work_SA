import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client' // Ensure you have the correct import path for createRoot
import './index.css'
import App from './App.tsx'

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <App />
  </StrictMode>,
)
