import { Layout } from '../components/Layout/Layout'
import { YourSalariesPage } from '../components/Pages/YourSalariesPage'

export default function Home() {
  return (
    <div>
      <Layout title="Your salaries">
        <YourSalariesPage />
      </Layout>
    </div>
  )
}
