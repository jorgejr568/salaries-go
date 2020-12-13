import Head from 'next/head'
import { Navbar } from './Navbar'

export function Layout({ title, children }) {
  return (
    <>
      <Head>
        <title>{title} - SalariesAPI</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <Navbar />
      <main className="mt-4">{children}</main>
    </>
  )
}
