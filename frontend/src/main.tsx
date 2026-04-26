import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import { Toaster } from 'sonner'
import './index.css'
import { BoardPage } from './app/boardPage'

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <BoardPage />
    <Toaster richColors position="top-right" />
  </StrictMode>,
)
