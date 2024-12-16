import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './index.css'
import { LoginPage } from './pages/LoginPage.tsx'
import { Provider } from 'react-redux'
import store from './store/store.ts'

createRoot(document.getElementById('root')!).render(
  <Provider store={store}>
    <LoginPage />
  </Provider>,
)
