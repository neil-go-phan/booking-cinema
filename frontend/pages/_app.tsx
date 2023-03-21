import '@/styles/globals.scss'
import type { AppProps } from 'next/app'
import ProgressBar from '@/components/ProgressBar'
import { config } from '@fortawesome/fontawesome-svg-core'
import '@fortawesome/fontawesome-svg-core/styles.css'
import { useState } from 'react'
import { IsLoggedContext } from '@/helpers/isLoggedContext';
config.autoAddCss = false

export default function App({ Component, pageProps }: AppProps) {
  const [isLogged, setIsLogged] = useState<boolean>(false)
  return (
    <IsLoggedContext.Provider value={{isLogged, setIsLogged}}>
      <ProgressBar />
      <Component {...pageProps} />
    </IsLoggedContext.Provider>
  )
}
