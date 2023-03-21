import '@/styles/globals.scss'
import type { AppProps } from 'next/app'
import ProgressBar from '@/components/ProgressBar'
import { config } from '@fortawesome/fontawesome-svg-core'
import '@fortawesome/fontawesome-svg-core/styles.css'

config.autoAddCss = false

export default function App({ Component, pageProps }: AppProps) {
  return (
    <>
      <ProgressBar />
      <Component {...pageProps} />
    </>
  )
}