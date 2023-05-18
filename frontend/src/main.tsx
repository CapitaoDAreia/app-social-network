import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './App.tsx'

ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
  <React.StrictMode >
    <div style={{height: '100vh'}} >
      <App />
    </div>
  </React.StrictMode>,
)
